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

func TestUserService_DatasetsOwned(t *testing.T) {
	setup()
	defer teardown()

	want := datasetSummaryResponses

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
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
			}]
		}`, owner, datasetid)
	}
	endpoint := "/user/datasets/own"
	mux.HandleFunc(endpoint, handler)
	got, err := client.User.DatasetsOwned()
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestUserService_ProjectsOwned(t *testing.T) {
	setup()
	defer teardown()

	want := projectSummaryResponses

	owner := testClientOwner
	datasetid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
				"owner": "%s",
				"id": "%s",
				"title": "My Awesome Project",
				"visibility": "OPEN",
				"status": "LOADED",
				"created": "2016-07-13T23:38:44.026Z",
				"updated": "2018-08-03T14:56:41.777Z",
				"accessLevel": "READ",
				"version": "some.version.identifier"
			}]
		}`, owner, datasetid)
	}
	endpoint := "/user/projects/own"
	mux.HandleFunc(endpoint, handler)
	got, err := client.User.ProjectsOwned()
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestUserService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	want := UserInfoResponse{
		DisplayName: "Adam Warlock",
		ID:          "adam",
		Created:     "2017-07-13T23:38:44.026Z",
		Updated:     "2018-09-03T14:56:41.777Z",
	}

	agentid := "adam"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"displayName": "Adam Warlock",
			"id": "%s",
			"created": "2017-07-13T23:38:44.026Z",
			"updated": "2018-09-03T14:56:41.777Z"
		}`, agentid)
	}
	endpoint := fmt.Sprintf("/users/%s", agentid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.User.Retrieve(agentid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestUserService_Self(t *testing.T) {
	setup()
	defer teardown()

	want := UserInfoResponse{
		DisplayName: "Tim Notes",
		ID:          testClientOwner,
		Created:     "2016-07-13T23:38:44.026Z",
		Updated:     "2018-08-03T14:56:41.777Z",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"displayName": "Tim Notes",
			"id": "tim-notes",
			"created": "2016-07-13T23:38:44.026Z",
			"updated": "2018-08-03T14:56:41.777Z"
		}`)
	}
	mux.HandleFunc("/user", handler)
	got, err := client.User.Self()
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
