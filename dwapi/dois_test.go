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

func TestDoiService_Associate(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	doi := "10.1109/5.771073"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/dois/%s", owner, datasetid, doi)
	mux.HandleFunc(endpoint, handler)
	got, err := client.DOI.Associate(owner, datasetid, doi)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDoiService_AssociateWithVersion(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	versionid := "a.dataset.version"
	doi := "10.1109/5.771073"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PUT, "Expected method 'PUT', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s/dois/%s", owner, datasetid, versionid, doi)
	mux.HandleFunc(endpoint, handler)
	got, err := client.DOI.AssociateWithVersion(owner, datasetid, versionid, doi)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDoiService_DeleteAssociatedWithVersion(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	versionid := "a.dataset.version"
	doi := "10.1109/5.771073"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/v/%s/dois/%s", owner, datasetid, versionid, doi)
	mux.HandleFunc(endpoint, handler)
	got, err := client.DOI.DeleteAssociatedWithVersion(owner, datasetid, versionid, doi)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestDoiService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	datasetid := "my-awesome-dataset"
	doi := "10.1109/5.771073"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/datasets/%s/%s/dois/%s", owner, datasetid, doi)
	mux.HandleFunc(endpoint, handler)
	got, err := client.DOI.Delete(owner, datasetid, doi)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
