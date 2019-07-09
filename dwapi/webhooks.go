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
)

type WebhookService struct {
	client *Client
}

func (s *WebhookService) List() (response []Subscription, err error) {
	endpoint := "/user/webhooks"
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

func (s *WebhookService) RetrieveAccountSubscription(user string) (response Subscription, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/users/%s", user)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *WebhookService) RetrieveDatasetSubscription(owner, datasetid string) (response Subscription, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *WebhookService) RetrieveProjectSubscription(owner, projectid string) (response Subscription, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *WebhookService) SubscribeToAccount(user string, body *SubscriptionCreateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/users/%s", user)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *WebhookService) SubscribeToDataset(owner, datasetid string, body *SubscriptionCreateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *WebhookService) SubscribeToProject(owner, projectid string, body *SubscriptionCreateRequest) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

func (s *WebhookService) UnsubscribeToAccount(user string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/users/%s", user)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *WebhookService) UnsubscribeToDataset(owner, datasetid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/datasets/%s/%s", owner, datasetid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

func (s *WebhookService) UnsubscribeToProject(owner, projectid string) (response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/user/webhooks/projects/%s/%s", owner, projectid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}
