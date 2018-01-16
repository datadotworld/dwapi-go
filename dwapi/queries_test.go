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

func TestQueryService_ListQueriesAssociatedWithDataset(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponses

	owner := client.Owner
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
	got, err := client.Query.ListQueriesAssociatedWithDataset(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_ListQueriesAssociatedWithProject(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponses

	owner := client.Owner
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
	got, err := client.Query.ListQueriesAssociatedWithProject(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	owner := client.Owner
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
	got, err := client.Query.Retrieve(queryid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestQueryService_RetrieveVersion(t *testing.T) {
	setup()
	defer teardown()

	want := querySummaryResponse

	owner := client.Owner
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
	got, err := client.Query.RetrieveVersion(queryid, versionid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
