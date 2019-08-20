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

type DoiService struct {
	client *Client
}

// Associate a DOI (Digital Object Identifier) with a dataset.
func (s *DoiService) Associate(owner, datasetid, doi string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/dois/%s", owner, datasetid, doi)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// AssociateWithVersion associates a DOI (Digital Object Identifier) with a version of dataset.
func (s *DoiService) AssociateWithVersion(owner, datasetid, versionid, doi string) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s/dois/%s", owner, datasetid, versionid, doi)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Delete a DOI (Digital Object Identifier) associated with a dataset.
func (s *DoiService) Delete(owner, datasetid, doi string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/dois/%s", owner, datasetid, doi)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// DeleteAssociatedWithVersion deletes a DOI (Digital Object Identifier) associated with a dataset.
func (s *DoiService) DeleteAssociatedWithVersion(owner, datasetid, versionid, doi string) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s/dois/%s", owner, datasetid, versionid, doi)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}
