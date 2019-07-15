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

type QueryService struct {
	client *Client
}

// CreateSavedQueryInDataset creates a saved query in the specified dataset.
func (s *QueryService) CreateSavedQueryInDataset(owner, datasetid string, body *QueryCreateRequest) (
	response QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries", owner, datasetid)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// CreateSavedQueryInProject creates a saved query in the specified project.
func (s *QueryService) CreateSavedQueryInProject(owner, projectid string, body *QueryCreateRequest) (
	response QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/queries", owner, projectid)
	headers := s.client.buildHeaders(POST, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// DeleteSavedQueryInDataset deletes a saved query in the specified dataset.
func (s *QueryService) DeleteSavedQueryInDataset(owner, datasetid, queryid string) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries/%s", owner, datasetid, queryid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// DeleteSavedQueryInProject deletes a saved query in the specified project.
func (s *QueryService) DeleteSavedQueryInProject(owner, projectid, queryid string) (
	response SuccessResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/queries/%s", owner, projectid, queryid)
	headers := s.client.buildHeaders(DELETE, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// ExecuteSavedQuery runs a saved query against a dataset or data project.
//
// SPARQL results are available in a variety of formats. See https://apidocs.data.world/api/queries/executequery
// for the full list of return types.
func (s *QueryService) ExecuteSavedQuery(queryid, acceptType string, body *SavedQueryExecutionRequest) (
	response io.ReadCloser, err error) {
	endpoint := fmt.Sprintf("/queries/%s/results", queryid)
	headers := s.client.buildHeaders(POST, endpoint)
	headers.AcceptType = acceptType

	b, err := s.client.encodeBody(body)
	if err != nil {
		return
	}
	return s.client.rawRequest(headers, b)
}

// ExecuteSavedQueryAndSave runs a saved query against a dataset or data project and saves the results
// to a file.
//
// SPARQL results are available in a variety of formats. See https://apidocs.data.world/api/queries/executequery
// for the full list of return types.
func (s *QueryService) ExecuteSavedQueryAndSave(queryid, acceptType, path string, body *SavedQueryExecutionRequest) (
	response SuccessResponse, err error) {
	r, err := s.ExecuteSavedQuery(queryid, acceptType, body)
	if err != nil {
		return
	}

	if err = s.client.saveToFile(path, r); err != nil {
		return
	}

	return SuccessResponse{
		Message: fmt.Sprintf("Results saved to %s", path),
	}, nil
}

// ExecuteSPARQL runs a SPARQL query against a dataset or data project.
//
// SPARQL results are available in a variety of formats. See https://apidocs.data.world/api/queries/sparqlpost
// for the full list of return types.
func (s *QueryService) ExecuteSPARQL(owner, id, acceptType string, body *SPARQLQueryRequest) (
	response io.ReadCloser, err error) {
	endpoint := fmt.Sprintf("/sparql/%s/%s", owner, id)
	headers := s.client.buildHeaders(POST, endpoint)
	headers.AcceptType = acceptType

	b, err := s.client.encodeBody(body)
	if err != nil {
		return
	}
	return s.client.rawRequest(headers, b)
}

// ExecuteSPARQLAndSave runs a SPARQL query against a dataset or data project and saves the results
// to a file.
//
// SPARQL results are available in a variety of formats. See https://apidocs.data.world/api/queries/sparqlpost
// for the full list of return types.
func (s *QueryService) ExecuteSPARQLAndSave(owner, id, acceptType, path string, body *SPARQLQueryRequest) (
	response SuccessResponse, err error) {
	r, err := s.ExecuteSPARQL(owner, id, acceptType, body)
	if err != nil {
		return
	}

	if err = s.client.saveToFile(path, r); err != nil {
		return
	}

	return SuccessResponse{
		Message: fmt.Sprintf("Results saved to %s", path),
	}, nil
}

// ExecuteSQL runs a SQL query against a dataset or data project.
//
// SQL results are available in a variety of formats. See https://apidocs.data.world/api/queries/sqlpost
// for the full list of return types.
func (s *QueryService) ExecuteSQL(owner, id, acceptType string, body *SQLQueryRequest) (
	response io.ReadCloser, err error) {
	endpoint := fmt.Sprintf("/sql/%s/%s", owner, id)
	headers := s.client.buildHeaders(POST, endpoint)
	headers.AcceptType = acceptType

	b, err := s.client.encodeBody(body)
	if err != nil {
		return
	}
	return s.client.rawRequest(headers, b)
}

// ExecuteSQLAndSave runs a SQL query against a dataset or data project and saves the results to a file.
//
// SQL results are available in a variety of formats. See https://apidocs.data.world/api/queries/sqlpost
// for the full list of return types.
func (s *QueryService) ExecuteSQLAndSave(owner, id, acceptType, path string, body *SQLQueryRequest) (
	response SuccessResponse, err error) {
	r, err := s.ExecuteSQL(owner, id, acceptType, body)
	if err != nil {
		return
	}

	if err = s.client.saveToFile(path, r); err != nil {
		return
	}

	return SuccessResponse{
		Message: fmt.Sprintf("Results saved to %s", path),
	}, nil
}

// ListQueries lists the saved queries associated with a dataset.
//
// Query definitions will be returned, not the query results. To retrieve the query results
// use `Query.ExecuteSavedQuery`.
func (s *QueryService) ListQueriesAssociatedWithDataset(owner, datasetid string) (
	response []QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries", owner, datasetid)
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// ListQueries lists the saved queries associated with a project.
//
// Query definitions will be returned, not the query results. To retrieve the query results
// use `Query.ExecuteSavedQuery`.
func (s *QueryService) ListQueriesAssociatedWithProject(owner, projectid string) (
	response []QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/queries", owner, projectid)
	if err = s.client.requestMultiplePages(endpoint, &response); err != nil {
		return nil, err
	}
	return
}

// Retrieve fetches a saved query.
//
// Query definitions will be returned, not the query results. To retrieve the query results
// use `Query.ExecuteSavedQuery`.
func (s *QueryService) Retrieve(queryid string) (response QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/queries/%s", queryid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// RetrieveVersion fetches a version of a saved query.
//
// Query definitions will be returned, not the query results. To retrieve the query results
// use `Query.ExecuteSavedQuery`.
func (s *QueryService) RetrieveVersion(queryid, versionid string) (response QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/queries/%s/v/%s", queryid, versionid)
	headers := s.client.buildHeaders(GET, endpoint)
	err = s.client.request(headers, nil, &response)
	return
}

// UpdateSavedQueryInDataset updates a saved query in the specified dataset.
func (s *QueryService) UpdateSavedQueryInDataset(owner, datasetid, queryid string, body *QueryUpdateRequest) (
	response QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/datasets/%s/%s/queries/%s", owner, datasetid, queryid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}

// UpdateSavedQueryInProject updates a saved query in the specified project.
func (s *QueryService) UpdateSavedQueryInProject(owner, datasetid, queryid string, body *QueryUpdateRequest) (
	response QuerySummaryResponse, err error) {
	endpoint := fmt.Sprintf("/projects/%s/%s/queries/%s", owner, datasetid, queryid)
	headers := s.client.buildHeaders(PUT, endpoint)
	err = s.client.request(headers, body, &response)
	return
}
