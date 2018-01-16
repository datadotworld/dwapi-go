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
	"io"
)

type ProjectService struct {
	client *Client
}

func (s *ProjectService) AddFilesFromURLs(owner, projectid string, body *FileCreateRequest) (response SuccessResponse, err error) {
	return s.client.File.AddFilesFromURLs(owner, projectid, body)
}

func (s *ProjectService) Contributing() (response []ProjectSummaryResponse, err error) {
	return s.client.User.ProjectsContributing()
}

func (s *ProjectService) Create(owner string, body *ProjectCreateOrUpdateRequest) (response ProjectCreateResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s", owner)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *ProjectService) CreateOrReplace(owner, projectid string, body *ProjectCreateOrUpdateRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *ProjectService) Delete(owner, projectid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *ProjectService) DownloadFile(owner, projectid, filename string) (response io.Reader, err error) {
	return s.client.File.Download(owner, projectid, filename)
}

func (s *ProjectService) DownloadAndSaveFile(owner, projectid, filename, path string) (response SuccessResponse, err error) {
	return s.client.File.DownloadAndSave(owner, projectid, filename, path)
}

func (s *ProjectService) Download(owner, projectid, filename string) (response io.Reader, err error) {
	return s.client.File.DownloadDataset(owner, projectid)
}

func (s *ProjectService) DownloadAndSave(owner, projectid, path string) (response SuccessResponse, err error) {
	return s.client.File.DownloadAndSaveDataset(owner, projectid, path)
}

func (s *ProjectService) Liked() (response []ProjectSummaryResponse, err error) {
	return s.client.User.ProjectsLiked()
}

func (s *ProjectService) LinkDataset(owner, projectid, linkedDatasetOwner, linkedDatasetid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/linkedDatasets/%s/%s",
		owner, projectid, linkedDatasetOwner, linkedDatasetid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *ProjectService) ListQueries(owner, projectid string) (response []QuerySummaryResponse, err error) {
	return s.client.Query.ListQueriesAssociatedWithProject(owner, projectid)
}

func (s *ProjectService) Owned() (response []ProjectSummaryResponse, err error) {
	return s.client.User.ProjectsOwned()
}

func (s *ProjectService) Retrieve(owner, projectid string) (response ProjectSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *ProjectService) RetrieveVersion(owner, projectid, versionid string) (response ProjectSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/v/%s", owner, projectid, versionid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *ProjectService) Sync(owner, projectid string) (response SuccessResponse, err error) {
	return s.client.File.Sync(owner, projectid)
}

func (s *ProjectService) UnlinkDataset(owner, projectid, linkedDatasetOwner, linkedDatasetid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/linkedDatasets/%s/%s",
		owner, projectid, linkedDatasetOwner, linkedDatasetid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *ProjectService) Update(owner, id string, body *ProjectCreateOrUpdateRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, id)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *ProjectService) UploadFile(owner, id, filename, path string, expandArchive bool) (response SuccessResponse, err error) {
	return s.client.File.Upload(owner, id, filename, path, expandArchive)
}
