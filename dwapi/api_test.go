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

package dwapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dw     *Client
	mux    *http.ServeMux
	server *httptest.Server

	testClientOwner = "tim-notes"
)

var datasetSummaryResponse = DatasetSummaryResponse{
	Owner:       testClientOwner,
	ID:          "my-awesome-dataset",
	Title:       "My Awesome Dataset",
	Visibility:  "OPEN",
	Status:      "LOADED",
	Created:     "2016-07-13T23:38:44.026Z",
	Updated:     "2018-08-03T14:56:41.777Z",
	IsProject:   false,
	AccessLevel: "READ",
	Version:     "some.version.identifier",
}

var datasetSummaryResponses = []DatasetSummaryResponse{
	datasetSummaryResponse,
}

var insightSummaryResponse = InsightSummaryResponse{
	Author: testClientOwner,
	Body: InsightBody{
		ImageURL: "https://www.link.to.image.com/img.jpg",
	},
	Created: "2018-07-13T23:38:44.026Z",
	ID:      "my-awesome-insight",
	Title:   "My Awesome Insight",
	Updated: "2018-08-03T14:56:41.777Z",
	Version: "some.version.identifier",
}

var insightSummaryResponses = []InsightSummaryResponse{
	insightSummaryResponse,
}

var projectSummaryResponse = ProjectSummaryResponse{
	Owner:       testClientOwner,
	ID:          "my-awesome-project",
	Title:       "My Awesome Project",
	Visibility:  "OPEN",
	Status:      "LOADED",
	Created:     "2016-07-13T23:38:44.026Z",
	Updated:     "2018-08-03T14:56:41.777Z",
	AccessLevel: "READ",
	Version:     "some.version.identifier",
}

var projectSummaryResponses = []ProjectSummaryResponse{
	projectSummaryResponse,
}

var querySummaryResponse = QuerySummaryResponse{
	Owner:    testClientOwner,
	ID:       "unique.id",
	Name:     "Metadata",
	Body:     "SELECT * FROM Tables",
	Language: "SQL",
	Created:  "2018-08-03T15:56:41.777Z",
	Updated:  "2018-08-03T15:56:41.777Z",
	Version:  "some.version.identified",
}

var querySummaryResponses = []QuerySummaryResponse{
	querySummaryResponse,
}

var successResponse = SuccessResponse{
	Message: "test.message",
}

var subscriptionSummaryResponse = Subscription{
	Events: []string{"ALL"},
}

var subscriptionSummaryResponses = []Subscription{
	subscriptionSummaryResponse,
}

func getTestClient() *Client {
	return NewClient("secret.token")
}

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	dw = getTestClient()
	dw.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	dw = getTestClient()
	assert.NotEmpty(t, dw.BaseURL)
	assert.Equal(t, dw.Token, "secret.token")
}

func TestGetBaseURL(t *testing.T) {
	dw := getTestClient()
	assert.Equal(t, dw.BaseURL, defaultBaseURL+"/v0")

	_ = os.Setenv("DW_ENVIRONMENT", "sparklesquad")
	dw = getTestClient()
	assert.Equal(t, dw.BaseURL, "https://api.sparklesquad.data.world/v0")

	_ = os.Setenv("DW_API_HOST", "http://localhost:1010")
	dw = getTestClient()
	assert.Equal(t, dw.BaseURL, "http://localhost:1010/v0")
}

func TestClient_RequestMultiplePages(t *testing.T) {
	setup()
	defer teardown()

	want := datasetSummaryResponses

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
				"owner": "%s",
				"id": "%s",
				"title": "My Awesome Dataset",
				"visibility": "OPEN",
				"status": "LOADED",
				"created": "2016-07-13T23:38:44.026Z",
				"updated": "2018-08-03T14:56:41.777Z",
				"isProject": false,
				"accessLevel": "READ",
				"version": "some.version.identifier"
			}]
		}`, owner, datasetid)
	}
	endpoint := "/user/datasets/own"
	var got []DatasetSummaryResponse
	mux.HandleFunc(endpoint, handler)
	err := dw.requestMultiplePages(endpoint, &got)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestClient_buildHeaders(t *testing.T) {
	setup()
	defer teardown()

	want := &headers{
		Method:   "GET",
		Endpoint: "/an/endpoint",
	}
	got := dw.buildHeaders(GET, "/an/endpoint")
	assert.Equal(t, want, got)
}
