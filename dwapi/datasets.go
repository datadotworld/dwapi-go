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

type DatasetService struct {
	client *Client
}

// AddFilesFromURLs adds files from URLs to a dataset. This method allows files published on the web
// to be added to a data.world dataset via their URL. This method can also be used to retrieve data
// via web APIs, with advanced options for HTTP method, request payload and authentication.
//
// The source URL will be stored so you can easily update your file anytime it changes via the
// `Sync now` button on the dataset page or by calling `Dataset.Sync()`.
func (s *DatasetService) AddFilesFromURLs(owner, datasetid string, body *[]FileCreateRequest) (
	response SuccessResponse, err error) {
	return s.client.File.AddFilesFromURLs(owner, datasetid, body)
}

// AssociateDOI associates a DOI (Digital Object Identifier) with a dataset.
func (s *DatasetService) AssociateDOI(owner, datasetid, doi string) (response SuccessResponse, err error) {
	return s.client.DOI.Associate(owner, datasetid, doi)
}

// AssociateDOIWithVersion associates a DOI (Digital Object Identifier) with a version of a dataset.
func (s *DatasetService) AssociateDOIWithVersion(owner, datasetid, versionid, doi string) (
	response SuccessResponse, err error) {
	return s.client.DOI.AssociateWithVersion(owner, datasetid, versionid, doi)
}

// Contributing lists the datasets that the currently authenticated user has access to because
// they are a contributor.
func (s *DatasetService) Contributing() (response []DatasetSummaryResponse, err error) {
	return s.client.User.DatasetsContributing()
}

// Create a dataset and associated data.
func (s *DatasetService) Create(owner string, body *DatasetCreateRequest) (response DatasetCreateResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s", owner)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// CreateOrReplace attempts to create a dataset with the given id, and will reset the dataset if it
// already exists, redefining all of its attributes.
func (s *DatasetService) CreateOrReplace(owner, id string, body *DatasetReplaceRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, id)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// Delete a dataset and associated data.
func (s *DatasetService) Delete(owner, datasetid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// DeleteDOI deletes a DOI (Digital Object Identifier) associated with a version of a dataset.
func (s *DatasetService) DeleteDOI(owner, datasetid, doi string) (response SuccessResponse, err error) {
	return s.client.DOI.Delete(owner, datasetid, doi)
}

// DeleteDOIAssociatedWithVersion deletes a DOI (Digital Object Identifier) associated with a version
// of a dataset.
func (s *DatasetService) DeleteDOIAssociatedWithVersion(owner, datasetid, versionid, doi string) (
	response SuccessResponse, err error) {
	return s.client.DOI.DeleteAssociatedWithVersion(owner, datasetid, versionid, doi)
}

// DownloadFile downloads a file within the dataset as originally uploaded.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *DatasetService) DownloadFile(owner, datasetid, filename string) (response io.Reader, err error) {
	return s.client.File.Download(owner, datasetid, filename)
}

// DownloadAndSaveFile downloads a file within the dataset as originally uploaded, and saves the results
// to a file.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *DatasetService) DownloadAndSaveFile(owner, datasetid, filename, path string) (
	response SuccessResponse, err error) {
	return s.client.File.DownloadAndSave(owner, datasetid, filename, path)
}

// Download downloads a .zip file containing all files within a dataset as originally uploaded.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *DatasetService) Download(owner, datasetid, filename string) (response io.Reader, err error) {
	return s.client.File.DownloadDataset(owner, datasetid)
}

// DownloadAndSave downloads a .zip file containing all files within a dataset as originally
// uploaded, and saves the results to a file.
//
// Prefer `Query.ExecuteSQL()` or `Query.ExecuteSPARQL()` for retrieving clean and structured data.
func (s *DatasetService) DownloadAndSave(owner, datasetid, path string) (response SuccessResponse, err error) {
	return s.client.File.DownloadAndSaveDataset(owner, datasetid, path)
}

// Liked lists the datasets that the currently authenticated user has liked (bookmarked).
func (s *DatasetService) Liked() (response []DatasetSummaryResponse, err error) {
	return s.client.User.DatasetsLiked()
}

// ListQueries lists the saved queries associated with a dataset.
//
// Query definitions will be returned, not the query results. To retrieve the query results
// use `Query.ExecuteSavedQuery`.
func (s *DatasetService) ListQueries(owner, datasetid string) (response []QuerySummaryResponse, err error) {
	return s.client.Query.ListQueriesAssociatedWithDataset(owner, datasetid)
}

// Owned lists the datasets that the currently authenticated user has access to because they are
// the owner.
func (s *DatasetService) Owned() (response []DatasetSummaryResponse, err error) {
	return s.client.User.DatasetsOwned()
}

// Retrieve fetches a dataset.
//
// The definition of will be returned, not the associated data. Use `Query.ExecuteSQL()`
// or `Query.ExecuteSPARQL()` to query the data. You can also download the original
// files with `Dataset.Download` or `Dataset.DownloadFile`.
func (s *DatasetService) Retrieve(owner, datasetid string) (response DatasetSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Retrieve fetches a version of a dataset.
//
// The definition of will be returned, not the associated data. Use `Query.ExecuteSQL()`
// or `Query.ExecuteSPARQL()` to query the data. You can also download the original
// files with `Dataset.Download` or `Dataset.DownloadFile`.
func (s *DatasetService) RetrieveVersion(owner, datasetid, versionid string) (
	response DatasetSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s", owner, datasetid, versionid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// Sync files within a dataset. This method will process the latest data available for files added
// from URLs or via streams.
func (s *DatasetService) Sync(owner, datasetid string) (response SuccessResponse, err error) {
	return s.client.File.Sync(owner, datasetid)
}

// Update a dataset.
func (s *DatasetService) Update(owner, id string, body *DatasetUpdateRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, id)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// UploadFile uploads one file at a time to a dataset.
func (s *DatasetService) UploadFile(owner, id, filename, path string, expandArchive bool) (
	response SuccessResponse, err error) {
	return s.client.File.Upload(owner, id, filename, path, expandArchive)
}
