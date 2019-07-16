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

func TestWebhookService_List(t *testing.T) {
	setup()
	defer teardown()

	want := subscriptionSummaryResponses

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"count": 1,
			"records": [{
				"events": ["ALL"]
			}]
		}`)
	}
	endpoint := "/user/webhooks"
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.List()
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_RetrieveAccountSubscription(t *testing.T) {
	setup()
	defer teardown()

	want := subscriptionSummaryResponse

	user := "arnold"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"events": ["ALL"]
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/users/%s", user)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.RetrieveAccountSubscription(user)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_RetrieveDatasetSubscription(t *testing.T) {
	setup()
	defer teardown()

	want := subscriptionSummaryResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"events": ["ALL"]
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.RetrieveDatasetSubscription(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_RetrieveProjectSubscription(t *testing.T) {
	setup()
	defer teardown()

	want := subscriptionSummaryResponse

	owner := testClientOwner
	projectid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"events": ["ALL"]
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.RetrieveProjectSubscription(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_SubscribeToAccount(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := SubscriptionCreateRequest{
		Events: []string{"ALL"},
	}
	user := "arnold"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/users/%s", user)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.SubscribeToAccount(user, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_SubscribeToDataset(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := SubscriptionCreateRequest{
		Events: []string{"ALL"},
	}
	owner := "arnold"
	datasetid := "my-awesome-dataset"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.SubscribeToDataset(owner, datasetid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_SubscribeToProject(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := SubscriptionCreateRequest{
		Events: []string{"ALL"},
	}
	owner := "arnold"
	projectid := "my-awesome-project"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.SubscribeToProject(owner, projectid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_UnsubscribeFromAccount(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	user := "arnold"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/users/%s", user)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.UnsubscribeFromAccount(user)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_UnsubscribeFromDataset(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := "arnold"
	datasetid := "my-awesome-dataset"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/datasets/%s/%s", owner, datasetid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.UnsubscribeFromDataset(owner, datasetid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestWebhookService_UnsubscribeFromProject(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := "arnold"
	projectid := "my-awesome-project"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/user/webhooks/projects/%s/%s", owner, projectid)
	mux.HandleFunc(endpoint, handler)
	got, err := dw.Webhook.UnsubscribeFromProject(owner, projectid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
