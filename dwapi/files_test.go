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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleFileService_AddFilesFromURLs() {
	owner := "dataset-owner"
	datasetid := "my-awesome-dataset"
	file := FileCreateRequest{
		Name:        "my-file.csv",
		Description: "A description for my test file.",
		Source: FileSourceCreateOrUpdateRequest{
			URL: "http://mysite.com/my-file.csv",
		},
	}
	files := []FileCreateRequest{
		file,
	}

	_, err := dw.File.AddFilesFromURLs(owner, datasetid, &files)
	if err != nil {
		log.Fatal(err)
	}
}

func TestFileService_AddFilesFromURLs(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	addRequest := FileCreateRequest{
		Name:        "my-file.csv",
		Description: "My test file",
		Source: FileSourceCreateOrUpdateRequest{
			URL: "http://www.test.url/my-file.csv",
		},
	}
	body := []FileCreateRequest{
		addRequest,
	}
	owner := testClientOwner
	id := "my-test-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/files", owner, id)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.File.AddFilesFromURLs(owner, id, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestFileService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	id := "my-awesome-dataset"
	filename := "arbitrary.file"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/files/%s", owner, id, filename)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.File.Delete(owner, id, filename)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func ExampleFileService_DownloadAndSave() {
	owner := "dataset-owner"
	datasetid := "my-awesome-dataset"
	filename := "my-file.csv"
	savePath := "./my-file.csv"

	_, err := dw.File.DownloadAndSave(owner, datasetid, filename, savePath)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleFileService_DownloadAndSaveDataset() {
	owner := "dataset-owner"
	datasetid := "my-awesome-dataset"
	savePath := "./my-file.zip"

	_, err := dw.File.DownloadAndSaveDataset(owner, datasetid, savePath)
	if err != nil {
		log.Fatal(err)
	}
}

func TestFileService_Download(t *testing.T) {
	setup()
	defer teardown()

	want := "test content"

	owner := testClientOwner
	id := "my-awesome-dataset"
	filename := "test-file"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `test content`)
	}
	endpoint := fmt.Sprintf("/file_download/%s/%s/%s", owner, id, filename)
	mux.HandleFunc(endpoint, handler)
	r, err := dw.File.Download(owner, id, filename)
	if assert.NoError(t, err) {
		got, _ := ioutil.ReadAll(r)
		assert.Equal(t, want, string(got))
	}
	r.Close()
}

func TestFileService_DownloadAndSave(t *testing.T) {
	setup()
	defer teardown()

	filename := "test-file"
	path := filepath.Join(os.TempDir(), filename)
	want := SuccessResponse{
		fmt.Sprintf("File saved to %s", path),
	}

	owner := testClientOwner
	id := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `test content`)
	}
	endpoint := fmt.Sprintf("/file_download/%s/%s/%s", owner, id, filename)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.File.DownloadAndSave(owner, id, filename, path)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
		assert.FileExists(t, path)

		c, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, "test content", string(c))
	}
	_ = os.Remove(path)
}

func TestFileService_DownloadDataset(t *testing.T) {
	setup()
	defer teardown()

	want := "test content"

	owner := testClientOwner
	id := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `test content`)
	}
	endpoint := fmt.Sprintf("/download/%s/%s", owner, id)
	mux.HandleFunc(endpoint, handler)
	r, err := dw.File.DownloadDataset(owner, id)
	if assert.NoError(t, err) {
		got, _ := ioutil.ReadAll(r)
		assert.Equal(t, want, string(got))
	}
	r.Close()
}

func TestFileService_DownloadAndSaveDataset(t *testing.T) {
	setup()
	defer teardown()

	filename := "test-file"
	path := filepath.Join(os.TempDir(), filename)
	want := SuccessResponse{
		fmt.Sprintf("ZIP file saved to %s", path),
	}

	owner := testClientOwner
	id := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `test content`)
	}
	endpoint := fmt.Sprintf("/download/%s/%s", owner, id)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.File.DownloadAndSaveDataset(owner, id, path)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
		assert.FileExists(t, path)

		c, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, "test content", string(c))
	}
	_ = os.Remove(path)
}

func TestFileService_Sync(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	id := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/sync", owner, id)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.File.Sync(owner, id)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func ExampleFileService_Upload() {
	owner := "dataset-owner"
	datasetid := "my-awesome-dataset"
	filename := "my-file.csv"
	filepath := "./my-file.csv"

	_, err := dw.File.Upload(owner, datasetid, filename, filepath, false)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleFileService_UploadStream() {
	owner := "dataset-owner"
	datasetid := "my-awesome-dataset"
	filename := "my-file.csv"
	s := []string{"first_name,last_name", "Abe,Marcos", "Abby,Johnson"}
	sj := strings.Join(s, "\n")
	body := strings.NewReader(sj)

	_, err := dw.File.UploadStream(owner, datasetid, filename, body, false)
	if err != nil {
		log.Fatal(err)
	}
}
