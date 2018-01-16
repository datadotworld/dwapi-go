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

func (s *DatasetService) AddFilesFromURLs(owner, datasetid string, body *FileCreateRequest) (response SuccessResponse, err error) {
	return s.client.File.AddFilesFromURLs(owner, datasetid, body)
}

func (s *DatasetService) AssociateDOI(owner, datasetid, doi string) (response SuccessResponse, err error) {
	return s.client.DOI.Associate(owner, datasetid, doi)
}

func (s *DatasetService) AssociateDOIWithVersion(owner, datasetid, versionid, doi string) (response SuccessResponse, err error) {
	return s.client.DOI.AssociateWithVersion(owner, datasetid, versionid, doi)
}

func (s *DatasetService) Contributing() (response []DatasetSummaryResponse, err error) {
	return s.client.User.DatasetsContributing()
}

func (s *DatasetService) Create(owner string, body *DatasetCreateRequest) (response DatasetCreateResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s", owner)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *DatasetService) CreateOrReplace(owner, id string, body *DatasetReplaceRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, id)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *DatasetService) Delete(owner, datasetid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *DatasetService) DeleteDOI(owner, datasetid, doi string) (response SuccessResponse, err error) {
	return s.client.DOI.Delete(owner, datasetid, doi)
}

func (s *DatasetService) DeleteDOIAssociatedWithVersion(owner, datasetid, versionid, doi string) (response SuccessResponse, err error) {
	return s.client.DOI.DeleteAssociatedWithVersion(owner, datasetid, versionid, doi)
}

func (s *DatasetService) DownloadFile(owner, datasetid, filename string) (response io.Reader, err error) {
	return s.client.File.Download(owner, datasetid, filename)
}

func (s *DatasetService) DownloadAndSaveFile(owner, datasetid, filename, path string) (response SuccessResponse, err error) {
	return s.client.File.DownloadAndSave(owner, datasetid, filename, path)
}

func (s *DatasetService) Download(owner, datasetid, filename string) (response io.Reader, err error) {
	return s.client.File.DownloadDataset(owner, datasetid)
}

func (s *DatasetService) DownloadAndSave(owner, datasetid, path string) (response SuccessResponse, err error) {
	return s.client.File.DownloadAndSaveDataset(owner, datasetid, path)
}

func (s *DatasetService) Liked() (response []DatasetSummaryResponse, err error) {
	return s.client.User.DatasetsLiked()
}

func (s *DatasetService) ListQueries(owner, datasetid string) (response []QuerySummaryResponse, err error) {
	return s.client.Query.ListQueriesAssociatedWithDataset(owner, datasetid)
}

func (s *DatasetService) Owned() (response []DatasetSummaryResponse, err error) {
	return s.client.User.DatasetsOwned()
}

func (s *DatasetService) Retrieve(owner, datasetid string) (response DatasetSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *DatasetService) RetrieveVersion(owner, datasetid, versionid string) (response DatasetSummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s", owner, datasetid, versionid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *DatasetService) Sync(owner, datasetid string) (response SuccessResponse, err error) {
	return s.client.File.Sync(owner, datasetid)
}

func (s *DatasetService) Update(owner, id string, body *DatasetUpdateRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, id)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *DatasetService) UploadFile(owner, id, filename, path string, expandArchive bool) (response SuccessResponse, err error) {
	return s.client.File.Upload(owner, id, filename, path, expandArchive)
}
