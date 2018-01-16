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

func TestInsightService_Create(t *testing.T) {
	setup()
	defer teardown()

	want := InsightCreateResponse{
		URI: "https://data.world/tim-notes/my-awesome-project/insightid",
	}

	body := InsightCreateRequest{
		Body: InsightBody{
			ImageURL: "https://www.link.to.image.com/img.jpg",
		},
		Title: "My Awesome Insight",
	}
	owner := client.Owner
	projectid := "my-awesome-project"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
		fmt.Fprintf(w, `{
			"URI": "https://data.world/%s/%s/insightid"
		}`, owner, projectid)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.Create(owner, projectid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestInsightService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := client.Owner
	projectid := "my-awesome-project"
	insightid := "my-awesome-insight"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.Delete(owner, projectid, insightid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestInsightService_List(t *testing.T) {
	setup()
	defer teardown()

	want := insightSummaryResponses

	owner := client.Owner
	projectid := "my-awesome-project"
	insightid := "my-awesome-insight"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
				"author": "%s",
				"id": "%s",
				"title": "My Awesome Insight",
				"body": {"imageUrl": "https://www.link.to.image.com/img.jpg"},
				"created": "2018-07-13T23:38:44.026Z",
				"updated": "2018-08-03T14:56:41.777Z",
				"version": "some.version.identifier"
			}]
		}`, owner, insightid)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.List(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_ReplaceInsight(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := InsightReplaceRequest{
		Body: InsightBody{
			ImageURL: "https://www.link.to.image.com/img.jpg",
		},
		Title: "My Awesome Insight",
	}
	owner := client.Owner
	projectid := "my-awesome-project"
	insightid := "my-awesome-insight"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.Replace(owner, projectid, insightid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestInsightService_Retrieve(t *testing.T) {
	setup()
	defer teardown()

	want := insightSummaryResponse

	owner := client.Owner
	projectid := "my-awesome-project"
	insightid := "my-awesome-insight"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"author": "%s",
			"id": "%s",
			"title": "My Awesome Insight",
			"body": {"imageUrl": "https://www.link.to.image.com/img.jpg"},
			"created": "2018-07-13T23:38:44.026Z",
			"updated": "2018-08-03T14:56:41.777Z",
			"version": "some.version.identifier"
		}`, owner, insightid)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.Retrieve(owner, projectid, insightid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestInsightService_RetrieveVersion(t *testing.T) {
	setup()
	defer teardown()

	want := InsightSummaryResponse{
		Author: "tim-notes",
		Body: InsightBody{
			ImageURL: "https://www.link.to.image.com/img.jpg",
		},
		Created: "2018-07-13T23:38:44.026Z",
		ID:      "my-awesome-insight",
		Title:   "My Awesome Insight",
		Updated: "2018-08-03T14:56:41.777Z",
		Version: "some.version.identifier",
	}

	owner := client.Owner
	projectid := "my-awesome-project"
	insightid := "my-awesome-insight"
	versionid := "some.version.identifier"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"author": "%s",
			"id": "%s",
			"title": "My Awesome Insight",
			"body": {"imageUrl": "https://www.link.to.image.com/img.jpg"},
			"created": "2018-07-13T23:38:44.026Z",
			"updated": "2018-08-03T14:56:41.777Z",
			"version": "%s"
		}`, owner, insightid, versionid)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s/%s/v/%s", owner, projectid, insightid, versionid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.RetrieveVersion(owner, projectid, insightid, versionid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestProjectService_UpdateInsight(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := InsightUpdateRequest{
		Title: "My Awesome Insight 2.0",
	}
	owner := client.Owner
	projectid := "my-awesome-project"
	insightid := "my-awesome-insight"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PATCH, "Expected method 'PATCH', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/insights/%s/%s/%s", owner, projectid, insightid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Insight.Update(owner, projectid, insightid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
