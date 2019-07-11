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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamService_Append(t *testing.T) {
	setup()
	defer teardown()

	want := SuccessResponse{
		Message: "Accepted",
	}

	type Language struct {
		Name string
		Year int
		URL  string
	}

	python := Language{
		Name: "Python",
		Year: 1991,
		URL:  "http://python.org",
	}

	golang := Language{
		Name: "Go",
		Year: 2009,
		URL:  "http://golang.org",
	}

	l1, _ := json.Marshal(python)
	l2, _ := json.Marshal(golang)

	ls := [][]byte{l1, l2}
	b := bytes.Join(ls, []byte("\n"))

	body := bytes.NewReader(b)
	owner := testClientOwner
	id := "my-awesome-dataset"
	streamid := "arbitrary.file"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, POST, "Expected method 'POST', got %s", r.Method)
	}
	endpoint := fmt.Sprintf("/streams/%s/%s/%s", owner, id, streamid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Stream.Append(owner, id, streamid, body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestStreamService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	owner := testClientOwner
	id := "my-awesome-dataset"
	streamid := "arbitrary.file"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, DELETE, "Expected method 'DELETE', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/records", owner, id, streamid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Stream.Delete(owner, id, streamid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestStreamService_RetrieveSchema(t *testing.T) {
	setup()
	defer teardown()

	want := StreamSchema{
		PrimaryKeyFields: []string{"id"},
		SequenceField:    "creation_time",
	}

	owner := testClientOwner
	id := "my-awesome-dataset"
	streamid := "arbitrary.file"
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, GET, "Expected method 'GET', got %s", r.Method)
		fmt.Fprintf(w, `{
			"primaryKeyFields": ["id"],
			"sequenceField": "creation_time"
		}`)
	}
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/schema", owner, id, streamid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Stream.RetrieveSchema(owner, id, streamid)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestStreamService_SetOrUpdateSchema(t *testing.T) {
	setup()
	defer teardown()

	want := successResponse

	body := StreamSchemaUpdateRequest{
		PrimaryKeyFields: []string{"id"},
		SequenceField:    "creation_time",
		UpdateMethod:     "TRUNCATE",
	}
	owner := testClientOwner
	id := "my-awesome-dataset"
	streamid := "arbitrary.file"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, PATCH, "Expected method 'PATCH', got %s", r.Method)
		fmt.Fprintf(w, `{
			"message": "test.message"
		}`)
	}
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/schema", owner, id, streamid)
	mux.HandleFunc(endpoint, handler)
	got, err := client.Stream.SetOrUpdateSchema(owner, id, streamid, &body)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
