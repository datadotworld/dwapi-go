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
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatasetService_Create(t *testing.T) {
	setup()
	defer teardown()

	want := DatasetCreateResponse{
		URI: "https://data.world/tim-notes/my-awesome-dataset",
	}

	body := DatasetCreateRequest{
		Title:      "My Awesome Dataset",
		Visibility: "OPEN",
	}
	owner := testClientOwner

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
		fmt.Fprintf(w, `{
			"URI": "https://data.world/%s/my-awesome-dataset"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/datasets/%s", owner)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.Create(owner, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func ExampleDatasetService_Create() {
	owner := "dataset-owner"
	request := DatasetCreateRequest{
		Title:       "My Awesome Dataset",
		Description: "A short description",
		Summary:     "A long description",
		Tags:        []string{"first", "puppies and kittens"},
		License:     "PDDL",
		Visibility:  "PRIVATE",
	}
	response, err := dw.Dataset.Create(owner, &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dataset URI:", response.URI)
}

func TestDatasetService_CreateOrReplace(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := DatasetReplaceRequest{
		Title:      "My Awesome Dataset",
		Visibility: "OPEN",
	}
	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.CreateOrReplace(owner, datasetid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.Delete(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	want := datasetSummaryResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
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
		}`, owner, datasetid)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.Retrieve(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_RetrieveVersion(t *testing.T) {
	setup()
	defer teardown()

	want := datasetSummaryResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	versionid := "some.version.identifier"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"owner": "%s",
			"id": "%s",
			"title": "My Awesome Dataset",
			"visibility": "OPEN",
			"status": "LOADED",
			"created": "2016-07-13T23:38:44.026Z",
			"updated": "2018-08-03T14:56:41.777Z",
			"isProject": false,
			"accessLevel": "READ",
			"version": "%s"
		}`, owner, datasetid, versionid)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s", owner, datasetid, versionid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.RetrieveVersion(owner, datasetid, versionid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_Sync(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/sync", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.Sync(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_Update(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := DatasetUpdateRequest{
		Title:      "My Awesome Dataset 2.0",
		Visibility: "PRIVATE",
	}
	owner := testClientOwner
	datasetid := "my-awesome-dataset"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PATCH, "Expected method 'PATCH', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Dataset.Update(owner, datasetid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
