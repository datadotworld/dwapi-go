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

type UserService struct {
	client *Client
}

// DatasetsContributing lists the datasets that the currently authenticated user has access to
// because they are a contributor.
func (s *UserService) DatasetsContributing() (response []DatasetSummaryResponse, err error) {
	endpoint := "/user/datasets/contributing"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// DatasetsLiked lists the datasets that the currently authenticated user has liked (bookmarked).
func (s *UserService) DatasetsLiked() (response []DatasetSummaryResponse, err error) {
	endpoint := "/user/datasets/liked"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// DatasetsOwned lists the datasets that the currently authenticated user has access to
// because they are the owner.
func (s *UserService) DatasetsOwned() (response []DatasetSummaryResponse, err error) {
	endpoint := "/user/datasets/own"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// ProjectsContributing lists the projects that the currently authenticated user has access to
// because they are a contributor.
func (s *UserService) ProjectsContributing() (response []ProjectSummaryResponse, err error) {
	endpoint := "/user/projects/contributing"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// ProjectsLiked lists the projects that the currently authenticated user has liked (bookmarked).
func (s *UserService) ProjectsLiked() (response []ProjectSummaryResponse, err error) {
	endpoint := "/user/projects/liked"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// ProjectsOwned lists the datasets that the currently authenticated user has access to
// because they are the owner.
func (s *UserService) ProjectsOwned() (response []ProjectSummaryResponse, err error) {
	endpoint := "/user/projects/own"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// Retrieve the user profile information for the specified account.
func (s *UserService) Retrieve(agentid string) (response UserInfoResponse, err error) {
	endpoint := fmt.Sprintf("/users/%s", agentid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Self retrieves the user profile information of the currently authenticated user.
func (s *UserService) Self() (response UserInfoResponse, err error) {
	endpoint := "/user"
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}
