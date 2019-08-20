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
)

type InsightService struct {
	client *Client
}

// Create a new insight.
func (s *InsightService) Create(owner, projectid string, body *InsightCreateRequest) (
	response InsightCreateResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// Delete an insight.
func (s *InsightService) Delete(owner, projectid, insightid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// List insights associated with a project.
func (s *InsightService) List(owner, projectid string) (response []InsightSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s", owner, projectid)
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// Replace an insight.
func (s *InsightService) Replace(owner, projectid, insightid string, body *InsightReplaceRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// Retrieve fetches an insight.
func (s *InsightService) Retrieve(owner, projectid, insightid string) (response InsightSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// RetrieveVersion fetches a version of an insight.
func (s *InsightService) RetrieveVersion(owner, projectid, insightid, versionid string) (
	response InsightSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s/%s/v/%s", owner, projectid, insightid, versionid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Update an insight.
//
// Note that only elements included in the request will be updated. All omitted elements will
// remain untouched.
func (s *InsightService) Update(owner, projectid, insightid string, body *InsightUpdateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}
