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
