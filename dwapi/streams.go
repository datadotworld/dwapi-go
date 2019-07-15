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
	"io"
)

type StreamService struct {
	client *Client
}

// Append JSON data to a stream associated with a dataset.
//
// data.world streams are append-only by default. Alternatively, if a primary key is specified
// (see: `Stream.SetOrUpdateSchema()`), data.world will replace records with the same primary key value.
//
// Streams don’t need to be created before you can append data to them. They will be created on-demand.
// Either when the first record is appended or when the schema is defined.
//
// Multiple records can be appended at once by using `application/json-l` as the request content type.
//
// Data uploaded to a dataset via a stream is not immediately processed. Instead, it is processed
// automatically in accordance with the dataset settings (default: daily) or as a result of calling
// `File.Sync()`.
//
// Once processed, the contents of a stream will appear as a .jsonl file.
func (s *StreamService) Append(owner, id, streamid string, body io.Reader) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s", owner, id, streamid)
	headers := s.client.buildHeaders(POST, endpoint)
	headers.ContentType = "application/json-l"

	r, err := s.client.rawRequest(headers, body)
	if err != nil {
		return
	}
	r.Close()
	return SuccessResponse{
		Message: "Accepted",
	}, nil
}

// Delete all records previously appended to a stream.
func (s *StreamService) Delete(owner, id, streamid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/records", owner, id, streamid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// RetrieveSchema fetches a stream’s schema.
func (s *StreamService) RetrieveSchema(owner, id, streamid string) (response StreamSchema, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/schema", owner, id, streamid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// SetOrUpdateSchema sets or updates a stream’s schema. The schema of a stream defines its primary
// key(s) and sort/sequence field.
//
// data.world streams are append-only by default. Alternatively, if a primary key is specified,
// data.world will replace records with the same primary key value. data.world will sort records
// by sequence field value and will discard all but the last record appended for each given
// primary key value.
//
//The updateMethod parameter specifies how data.world should handle existing records when the
// schema is updated. Currently, the only updateMethod supported is TRUNCATED. data.world
// will discard all records when the schema is updated.
func (s *StreamService) SetOrUpdateSchema(owner, id, streamid string, body *StreamSchemaUpdateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/schema", owner, id, streamid)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}
