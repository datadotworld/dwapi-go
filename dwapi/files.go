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
	"os"
)

type FileService struct {
	client *Client
}

func (s *FileService) AddFilesFromURLs(owner, id string, body *FileCreateRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/files", owner, id)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *FileService) Delete(owner, id, filename string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/files/%s", owner, id, filename)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *FileService) Download(owner, id, filename string) (response io.Reader, err error) {
	endpoint := fmt.Sprintf("/file_download/%s/%s/%s", owner, id, filename)
	headers := s.client.buildHeaders(GET, endpoint)
	return s.client.rawRequest(headers, nil)
}

func (s *FileService) DownloadAndSave(owner, id, filename, path string) (response SuccessResponse, err error) {
	r, err := s.Download(owner, id, filename)
	if err != nil {
		return
	}

	if err = s.client.saveToFile(path, r); err != nil {
		return
	}

	return SuccessResponse{
		Message: fmt.Sprintf("File saved to %s", path),
	}, nil
}

func (s *FileService) DownloadDataset(owner, id string) (response io.Reader, err error) {
	endpoint := fmt.Sprintf("/download/%s/%s", owner, id)
	headers := s.client.buildHeaders(GET, endpoint)
	return s.client.rawRequest(headers, nil)
}

func (s *FileService) DownloadAndSaveDataset(owner, id, path string) (response SuccessResponse, err error) {
	r, err := s.DownloadDataset(owner, id)
	if err != nil {
		return
	}

	if err = s.client.saveToFile(path, r); err != nil {
		return
	}

	return SuccessResponse{
		Message: fmt.Sprintf("ZIP file saved to %s", path),
	}, nil
}

func (s *FileService) Sync(owner, id string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/sync", owner, id)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *FileService) UploadStream(owner, id, filename string, body io.Reader, expandArchive bool) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/uploads/%s/%s/files/%s", owner, id, filename)

	if expandArchive {
		endpoint = endpoint + "?expandArchive=true"
	}

	headers := s.client.buildHeaders(PUT, endpoint)
	headers.ContentType = "application/octet-stream"

	r, err := s.client.rawRequest(headers, body)
	if err != nil {
		return
	}

	err = s.client.unmarshal(r, &response)
	return
}

func (s *FileService) Upload(owner, id, filename, path string, expandArchive bool) (response SuccessResponse, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	response, err = s.UploadStream(owner, id, filename, f, expandArchive)
	if err != nil {
		return
	}

	err = f.Close()
	return
}
