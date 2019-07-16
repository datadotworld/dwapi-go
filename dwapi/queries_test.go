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

func TestQueryService_CreateSavedQueryInDataset(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	body := QueryCreateRequest{
		Name:      "Metadata",
		Content:   "SELECT * FROM Tables",
		Language:  "SQL",
		Published: true,
	}
	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
		fmt.Fprintf(w, `{
			"body": "SELECT * FROM Tables",
			"created": "2018-08-03T15:56:41.777Z",
			"id": "unique.id",
			"language": "SQL",
			"name": "Metadata",
			"owner": "%s",
			"updated": "2018-08-03T15:56:41.777Z",
			"version": "some.version.identified"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.CreateSavedQueryInDataset(owner, datasetid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_CreateSavedQueryInProject(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	body := QueryCreateRequest{
		Name:      "Metadata",
		Content:   "SELECT * FROM Tables",
		Language:  "SQL",
		Published: true,
	}
	owner := testClientOwner
	projectid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
		fmt.Fprintf(w, `{
			"body": "SELECT * FROM Tables",
			"created": "2018-08-03T15:56:41.777Z",
			"id": "unique.id",
			"language": "SQL",
			"name": "Metadata",
			"owner": "%s",
			"updated": "2018-08-03T15:56:41.777Z",
			"version": "some.version.identified"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/queries", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.CreateSavedQueryInProject(owner, projectid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_DeleteSavedQueryInDataset(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	queryid := "my-saved-query"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries/%s", owner, datasetid, queryid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.DeleteSavedQueryInDataset(owner, datasetid, queryid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDatasetService_DeleteSavedQueryInProject(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	projectid := "my-awesome-dataset"
	queryid := "my-saved-query"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/queries/%s", owner, projectid, queryid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.DeleteSavedQueryInProject(owner, projectid, queryid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_ListQueriesAssociatedWithDataset(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponses

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
				"body": "SELECT * FROM Tables",
				"created": "2018-08-03T15:56:41.777Z",
				"id": "unique.id",
				"language": "SQL",
				"name": "Metadata",
				"owner": "%s",
				"updated": "2018-08-03T15:56:41.777Z",
				"version": "some.version.identified"
			}]
		}`, owner)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.ListQueriesAssociatedWithDataset(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_ListQueriesAssociatedWithProject(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponses

	owner := testClientOwner
	projectid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
				"body": "SELECT * FROM Tables",
				"created": "2018-08-03T15:56:41.777Z",
				"id": "unique.id",
				"language": "SQL",
				"name": "Metadata",
				"owner": "%s",
				"updated": "2018-08-03T15:56:41.777Z",
				"version": "some.version.identified"
			}]
		}`, owner)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/queries", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.ListQueriesAssociatedWithProject(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	owner := testClientOwner
	queryid := "query.id"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"body": "SELECT * FROM Tables",
			"created": "2018-08-03T15:56:41.777Z",
			"id": "unique.id",
			"language": "SQL",
			"name": "Metadata",
			"owner": "%s",
			"updated": "2018-08-03T15:56:41.777Z",
			"version": "some.version.identified"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/queries/%s", queryid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.Retrieve(queryid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_RetrieveVersion(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	owner := testClientOwner
	queryid := "query.id"
	versionid := "some.version.identified"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"body": "SELECT * FROM Tables",
			"created": "2018-08-03T15:56:41.777Z",
			"id": "unique.id",
			"language": "SQL",
			"name": "Metadata",
			"owner": "%s",
			"updated": "2018-08-03T15:56:41.777Z",
			"version": "%s"
		}`, owner, versionid)
	}
	endpoint := fmt.Sprintf("/queries/%s/v/%s", queryid, versionid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.RetrieveVersion(queryid, versionid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_UpdateSavedQueryInDataset(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	body := QueryUpdateRequest{
		Name:    "Metadata",
		Content: "SELECT * FROM Table",
	}
	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	queryid := "unique.id"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"body": "SELECT * FROM Tables",
			"created": "2018-08-03T15:56:41.777Z",
			"id": "unique.id",
			"language": "SQL",
			"name": "Metadata",
			"owner": "%s",
			"updated": "2018-08-03T15:56:41.777Z",
			"version": "some.version.identified"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries/%s", owner, datasetid, queryid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.UpdateSavedQueryInDataset(owner, datasetid, queryid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_UpdateSavedQueryInProject(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	body := QueryUpdateRequest{
		Name:    "Metadata",
		Content: "SELECT * FROM Table",
	}
	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	queryid := "unique.id"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"body": "SELECT * FROM Tables",
			"created": "2018-08-03T15:56:41.777Z",
			"id": "unique.id",
			"language": "SQL",
			"name": "Metadata",
			"owner": "%s",
			"updated": "2018-08-03T15:56:41.777Z",
			"version": "some.version.identified"
		}`, owner)
	}
	endpoint := fmt.Sprintf("/projects/%s/%s/queries/%s", owner, datasetid, queryid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Query.UpdateSavedQueryInProject(owner, datasetid, queryid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
