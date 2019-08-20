// Copyright © 2018 data.world, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This product includes software developed at
// data.world, Inc.(http://data.world/).

/*
Package dwapi makes it easy to use data.world's REST API (https://apidocs.data.world/api) with Go.

Users can create and update datasets, projects, and metadata; query datasets using SQL and SPARQL;
and download and upload files.
*/
package dwapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	DELETE = "DELETE"
	GET    = "GET"
	PATCH  = "PATCH"
	POST   = "POST"
	PUT    = "PUT"

	defaultBaseURL = "https://api.data.world"
)

type Client struct {
	BaseURL string
	Token   string

	Dataset *DatasetService
	DOI     *DoiService
	File    *FileService
	Insight *InsightService
	Project *ProjectService
	Query   *QueryService
	Stream  *StreamService
	User    *UserService
	Webhook *WebhookService
}

type headers struct {
	Method      string
	Endpoint    string
	AcceptType  string
	ContentType string
}

type paginatedResponse struct {
	Count         int           `json:"count"`
	NextPageToken string        `json:"nextPageToken,omitempty"`
	Records       []interface{} `json:"records"`
}

func NewClient(token string) *Client {
	c := &Client{
		BaseURL: getBaseURL(),
		Token:   token,
	}
	c.Dataset = &DatasetService{c}
	c.DOI = &DoiService{c}
	c.File = &FileService{c}
	c.Insight = &InsightService{c}
	c.Project = &ProjectService{c}
	c.Query = &QueryService{c}
	c.Stream = &StreamService{c}
	c.User = &UserService{c}
	c.Webhook = &WebhookService{c}
	return c
}

func getBaseURL() string {
	baseURL := defaultBaseURL
	if host := os.Getenv("DW_API_HOST"); host != "" {
		baseURL = host
	} else if env := os.Getenv("DW_ENVIRONMENT"); env != "" {
		baseURL = fmt.Sprintf("https://api.%s.data.world", env)
	}
	return baseURL + "/v0"
}

func (c *Client) buildHeaders(method, endpoint string) *headers {
	return &headers{
		Method:   method,
		Endpoint: endpoint,
	}
}

func (c *Client) encodeBody(body interface{}) (io.Reader, error) {
	b := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(b).Encode(body); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (c *Client) rawRequest(headers *headers, body io.Reader) (io.ReadCloser, error) {
	url := c.BaseURL + headers.Endpoint

	r, err := http.NewRequest(headers.Method, url, body)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	if headers.ContentType == "" {
		headers.ContentType = "application/json"
	}
	r.Header.Add("Content-Type", headers.ContentType)

	if headers.AcceptType != "" {
		r.Header.Add("Accept", headers.AcceptType)
	}

	h := &http.Client{
		Timeout: 60 * time.Second,
	}
	response, err := h.Do(r)
	if err != nil {
		return nil, err
	}

	if string(response.Status[0]) != "2" {
		return nil, errors.New(response.Status)
	}
	return response.Body, nil
}

func (c *Client) request(headers *headers, body, response interface{}) (err error) {
	b, err := c.encodeBody(body)
	if err != nil {
		return
	}

	r, err := c.rawRequest(headers, b)
	if err != nil {
		return
	}

	err = c.unmarshal(r, response)
	r.Close()
	return
}

func (c *Client) requestMultiplePages(endpoint string, response interface{}) error {
	var records []interface{}
	nextPageToken := ""
	url := endpoint

	for {
		if nextPageToken != "" {
			url = fmt.Sprintf("%s?next=%s", endpoint, nextPageToken)
		}

		headers := c.buildHeaders(GET, url)
		page := paginatedResponse{}
		if err := c.request(headers, nil, &page); err != nil {
			return err
		}
		records = append(records, page.Records...)
		nextPageToken = page.NextPageToken
		if nextPageToken == "" {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	b, err := json.Marshal(records)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &response); err != nil {
		return err
	}
	return nil
}

func (c *Client) saveToFile(path string, contents io.Reader) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}

	if _, err = io.Copy(f, contents); err != nil {
		return
	}

	err = f.Close()
	return
}

func (c *Client) unmarshal(reader io.Reader, response interface{}) (err error) {
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	err = json.Unmarshal(d, response)
	return
}
