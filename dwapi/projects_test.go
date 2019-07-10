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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_Create(t *testing.T) {
	setup()
	defer teardown()

	want := ProjectCreateResponse{
		URI: "https://data.world/tim-notes/my-awesome-project",
	}

	body := ProjectCreateOrUpdateRequest{
		Title:      "My Awesome Project",
		Visibility: "OPEN",
		LinkedDatasets: []LinkedDatasetCreateOrUpdateRequest{
			{
				ID:    "my-awesome-dataset",
				Owner: testClientOwner,
			},
		},
	}
	owner := testClientOwner

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
		fmt.Fprintf(w, `{
			"URI": "https://data.world/%s/my-awesome-project"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/projects/%s", owner)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.Create(owner, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_CreateOrReplace(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := ProjectCreateOrUpdateRequest{
		Title:      "My Awesome Project",
		Visibility: "OPEN",
		LinkedDatasets: []LinkedDatasetCreateOrUpdateRequest{
			{
				ID:    "my-awesome-dataset",
				Owner: testClientOwner,
			},
		},
	}
	owner := testClientOwner
	projectid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.CreateOrReplace(owner, projectid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	projectid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.Delete(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_Link(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	projectid := "my-awesome-project"
	linkedDatasetOwner := testClientOwner
	linkedDatasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/linkedDatasets/%s/%s",
		owner, projectid, linkedDatasetOwner, linkedDatasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.LinkDataset(owner, projectid, linkedDatasetOwner, linkedDatasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	want := projectSummaryResponse

	owner := testClientOwner
	projectid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"owner": "%s",
			"id": "%s",
			"title": "My Awesome Project",
			"visibility": "OPEN",
			"status": "LOADED",
			"created": "2016-07-13T23:38:44.026Z",
			"updated": "2018-08-03T14:56:41.777Z",
			"accessLevel": "READ",
			"version": "some.version.identifier"
		}`, owner, projectid)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.Retrieve(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_RetrieveVersion(t *testing.T) {
	setup()
	defer teardown()

	want := projectSummaryResponse

	owner := testClientOwner
	projectid := "my-awesome-project"
	versionid := "some.version.identifier"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"owner": "%s",
			"id": "%s",
			"title": "My Awesome Project",
			"visibility": "OPEN",
			"status": "LOADED",
			"created": "2016-07-13T23:38:44.026Z",
			"updated": "2018-08-03T14:56:41.777Z",
			"accessLevel": "READ",
			"version": "%s"
		}`, owner, projectid, versionid)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/v/%s", owner, projectid, versionid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.RetrieveVersion(owner, projectid, versionid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_Sync(t *testing.T) {
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
	got, err := client.Project.Sync(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_Unlink(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	projectid := "my-awesome-project"
	linkedDatasetOwner := testClientOwner
	linkedDatasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/linkedDatasets/%s/%s",
		owner, projectid, linkedDatasetOwner, linkedDatasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.UnlinkDataset(owner, projectid, linkedDatasetOwner, linkedDatasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_Update(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := ProjectCreateOrUpdateRequest{
		Title:      "My Awesome Project 2.0",
		Visibility: "PRIVATE",
	}
	owner := testClientOwner
	projectid := "my-awesome-project"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PATCH, "Expected method 'PATCH', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Project.Update(owner, projectid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
