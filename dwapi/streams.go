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

func (s *StreamService) Append(owner, id, streamid string, body io.Reader) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s", owner, id, streamid)
	headers := s.client.buildHeaders(POST, endpoint)
	headers.ContentType = "application/json-l"

	r, err := s.client.rawRequest(headers, body)
	if err != nil {
		return
	}

	err = s.client.unmarshal(r, &response)
	return
}

func (s *StreamService) Delete(owner, id, streamid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/records", owner, id, streamid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *StreamService) RetrieveSchema(owner, id, streamid string) (response StreamSchema, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/schema", owner, id, streamid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *StreamService) SetOrUpdateSchema(owner, id, streamid string, body *StreamSchemaUpdateRequest) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/streams/%s/%s/%s/schema", owner, id, streamid)
	headers := s.client.buildHeaders(PATCH, endpoint)
	err = s.client.request(headers, body, &response)
	return
}
