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

// AddFilesFromURLs adds files from URLs to a project. This method allows files published on the web
// to be added to a data.world dataset via their URL. This method can also be used to retrieve data
// via web APIs, with advanced options for HTTP method, request payload and authentication.
//
// The source URL will be stored so you can easily update your file anytime it changes via the
// `Sync now` button on the dataset page or by calling `Project.Sync()`.
func (s *ProjectService) AddFilesFromURLs(owner, projectid string, body *[]FileCreateRequest) (
	response SuccessResponse, err error) {
	return s.client.File.AddFilesFromURLs(owner, projectid, body)
}

// Contributing lists the projects that the currently authenticated user has access to because
// they are a contributor.
func (s *ProjectService) Contributing() (response []ProjectSummaryResponse, err error) {
	return s.client.User.ProjectsContributing()
}

// Create a project and associated data.
func (s *ProjectService) Create(owner string, body *ProjectCreateOrUpdateRequest) (
	response ProjectCreateResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s", owner)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// CreateOrReplace attempts to create a project with the given id, and will reset the project if it
// already exists, redefining all of its attributes.
func (s *ProjectService) CreateOrReplace(owner, projectid string, body *ProjectCreateOrUpdateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// Delete a project and associated data.
func (s *ProjectService) Delete(owner, projectid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// DownloadFile downloads a file within the project as originally uploaded.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *ProjectService) DownloadFile(owner, projectid, filename string) (response io.Reader, err error) {
	return s.client.File.Download(owner, projectid, filename)
}

// DownloadAndSaveFile downloads a file within the project as originally uploaded, and saves the results
// to a file.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *ProjectService) DownloadAndSaveFile(owner, projectid, filename, path string) (
	response SuccessResponse, err error) {
	return s.client.File.DownloadAndSave(owner, projectid, filename, path)
}

// Download downloads a .zip file containing all files within a project as originally uploaded.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *ProjectService) Download(owner, projectid, filename string) (response io.Reader, err error) {
	return s.client.File.DownloadDataset(owner, projectid)
}

// DownloadAndSave downloads a .zip file containing all files within a project as originally
// uploaded, and saves the results to a file.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *ProjectService) DownloadAndSave(owner, projectid, path string) (response SuccessResponse, err error) {
	return s.client.File.DownloadAndSaveDataset(owner, projectid, path)
}

// Liked lists the projects that the currently authenticated user has liked (bookmarked).
func (s *ProjectService) Liked() (response []ProjectSummaryResponse, err error) {
	return s.client.User.ProjectsLiked()
}

// LinkDataset adds a linked dataset to a project.
func (s *ProjectService) LinkDataset(owner, projectid, linkedDatasetOwner, linkedDatasetid string) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/linkedDatasets/%s/%s",
		owner, projectid, linkedDatasetOwner, linkedDatasetid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// ListQueries lists the saved queries associated with a project.
//
// Query definitions will be returned, not the query results. To retrieve the query results
// use `Query.ExecuteSavedQuery`.
func (s *ProjectService) ListQueries(owner, projectid string) (response []QuerySummaryResponse, err error) {
	return s.client.Query.ListQueriesAssociatedWithProject(owner, projectid)
}

// Owned lists the projects that the currently authenticated user has access to because they are
// the owner.
func (s *ProjectService) Owned() (response []ProjectSummaryResponse, err error) {
	return s.client.User.ProjectsOwned()
}

// Retrieve fetches a project.
//
// The definition of will be returned, not the associated data. Use `Query.ExecuteSQL()`
// or `Query.ExecuteSPARQL()` to query the data. You can also download the original
// files with `Project.Download` or `Project.DownloadFile`.
func (s *ProjectService) Retrieve(owner, projectid string) (response ProjectSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Retrieve fetches a version of a project.
//
// The definition of will be returned, not the associated data. Use `Query.ExecuteSQL()`
// or `Query.ExecuteSPARQL()` to query the data. You can also download the original
// files with `Project.Download` or `Project.DownloadFile`.
func (s *ProjectService) RetrieveVersion(owner, projectid, versionid string) (
	response ProjectSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/v/%s", owner, projectid, versionid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Sync files within a project. This method will process the latest data available for files added
// from URLs or via streams.
func (s *ProjectService) Sync(owner, projectid string) (response SuccessResponse, err error) {
	return s.client.File.Sync(owner, projectid)
}

// UnlinkDataset removes a linked dataset from a project.
func (s *ProjectService) UnlinkDataset(owner, projectid, linkedDatasetOwner, linkedDatasetid string) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/linkedDatasets/%s/%s",
		owner, projectid, linkedDatasetOwner, linkedDatasetid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Update a project.
func (s *ProjectService) Update(owner, id string, body *ProjectCreateOrUpdateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, id)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// UploadFile uploads one file at a time to a project.
func (s *ProjectService) UploadFile(owner, id, filename, path string, expandArchive bool) (
	response SuccessResponse, err error) {
	return s.client.File.Upload(owner, id, filename, path, expandArchive)
}
