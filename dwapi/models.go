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

type DatasetCreateRequest struct {
	Description string              `json:"description,omitempty"`
	Files       []FileCreateRequest `json:"files,omitempty"`
	License     string              `json:"license,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Title       string              `json:"title"`
	Visibility  string              `json:"visibility"`
}

type DatasetCreateResponse struct {
	Message string `json:"message,omitempty"`
	URI     string `json:"uri"`
}

type DatasetOrProjectIdentifier struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
}

type DatasetReplaceRequest struct {
	Description string              `json:"description,omitempty"`
	Files       []FileCreateRequest `json:"files,omitempty"`
	License     string              `json:"license,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Title       string              `json:"title"`
	Visibility  string              `json:"visibility"`
}

type DatasetSummaryResponse struct {
	AccessLevel string                    `json:"accessLevel"`
	Created     string                    `json:"created"`
	Description string                    `json:"description,omitempty"`
	Dois        []DigitalObjectIdentifier `json:"dois,omitempty"`
	Files       []FileSummaryResponse     `json:"files,omitempty"`
	ID          string                    `json:"id"`
	IsProject   bool                      `json:"isProject"`
	License     string                    `json:"license,omitempty"`
	Owner       string                    `json:"owner"`
	Status      string                    `json:"status"`
	Summary     string                    `json:"summary,omitempty"`
	Tags        []string                  `json:"tags,omitempty"`
	Title       string                    `json:"title"`
	Updated     string                    `json:"updated"`
	Version     string                    `json:"version"`
	VersionDois []DigitalObjectIdentifier `json:"versionDois,omitempty"`
	Visibility  string                    `json:"visibility"`
}

type DatasetUpdateRequest struct {
	Description string              `json:"description,omitempty"`
	Files       []FileCreateRequest `json:"files,omitempty"`
	License     string              `json:"license,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Title       string              `json:"title,omitempty"`
	Visibility  string              `json:"visibility,omitempty"`
}

type DigitalObjectIdentifier struct {
	Created string `json:"created"`
	Doi     string `json:"doi"`
}

type FileCreateRequest struct {
	Description string                          `json:"description,omitempty"`
	Labels      []string                        `json:"labels,omitempty"`
	Name        string                          `json:"name"`
	Source      FileSourceCreateOrUpdateRequest `json:"source"`
}

type FileSourceCreateOrUpdateRequest struct {
	Authorization  WebAuthorization    `json:"authorization,omitempty"`
	Credentials    WebCredentials      `json:"credentials,omitempty"`
	ExpandArchive  bool                `json:"expandArchive,omitempty"`
	Method         string              `json:"method,omitempty"`
	OauthToken     OauthTokenReference `json:"oauthToken,omitempty"`
	RequestEntity  string              `json:"requestEntity,omitempty"`
	RequestHeaders interface{}         `json:"requestHeaders,omitempty"`
	URL            string              `json:"url,omitempty"`
}

type FileSourceResponse struct {
	Authorization   WebAuthorization    `json:"authorization,omitempty"`
	Credentials     WebCredentials      `json:"credentials,omitempty"`
	ExpandArchive   bool                `json:"expandArchive,omitempty"`
	Method          string              `json:"method,omitempty"`
	SyncStatus      string              `json:"syncStatus"`
	SyncSummary     []string            `json:"labels,omitempty"`
	LastSyncFailure string              `json:"lastSyncFailure,omitempty"`
	LastSyncStart   string              `json:"lastSyncStart,omitempty"`
	LastSyncSuccess string              `json:"lastSyncSuccess,omitempty"`
	OauthToken      OauthTokenReference `json:"oauthToken,omitempty"`
	RequestEntity   string              `json:"requestEntity,omitempty"`
	RequestHeaders  interface{}         `json:"requestHeaders,omitempty"`
	URL             string              `json:"url,omitempty"`
}

type FileSummaryResponse struct {
	Created     string             `json:"created"`
	Description string             `json:"description,omitempty"`
	Labels      []string           `json:"labels,omitempty"`
	Name        string             `json:"name"`
	SizeInBytes string             `json:"sizeInBytes,omitempty"`
	Source      FileSourceResponse `json:"source,omitempty"`
	Updated     string             `json:"updated"`
}

type InsightBody struct {
	EmbedURL     string `json:"embedUrl,omitempty"`
	ImageURL     string `json:"imageUrl,omitempty"`
	MarkdownBody string `json:"markdownBody,omitempty"`
}

type InsightCreateRequest struct {
	Body            InsightBody `json:"body"`
	DataSourceLinks []string    `json:"dataSourceLinks,omitempty"`
	Description     string      `json:"description,omitempty"`
	SourceLink      string      `json:"sourceLink,omitempty"`
	Thumbnail       string      `json:"thumbnail,omitempty"`
	Title           string      `json:"title"`
}

type InsightCreateResponse struct {
	Message string `json:"message,omitempty"`
	URI     string `json:"uri"`
}

type InsightReplaceRequest struct {
	Body            InsightBody `json:"body"`
	DataSourceLinks []string    `json:"dataSourceLinks,omitempty"`
	Description     string      `json:"description,omitempty"`
	SourceLink      string      `json:"sourceLink,omitempty"`
	Thumbnail       string      `json:"thumbnail,omitempty"`
	Title           string      `json:"title"`
}

type InsightSummaryResponse struct {
	Author          string      `json:"author"`
	Body            InsightBody `json:"body"`
	Created         string      `json:"created"`
	DataSourceLinks []string    `json:"dataSourceLinks,omitempty"`
	Description     string      `json:"description,omitempty"`
	ID              string      `json:"id"`
	SourceLink      string      `json:"sourceLink,omitempty"`
	Thumbnail       string      `json:"thumbnail,omitempty"`
	Title           string      `json:"title"`
	Updated         string      `json:"updated"`
	Version         string      `json:"version"`
}

type InsightUpdateRequest struct {
	Body            InsightBody `json:"body,omitempty"`
	DataSourceLinks []string    `json:"dataSourceLinks,omitempty"`
	Description     string      `json:"description,omitempty"`
	SourceLink      string      `json:"sourceLink,omitempty"`
	Thumbnail       string      `json:"thumbnail,omitempty"`
	Title           string      `json:"title,omitempty"`
}

type LinkedDatasetCreateOrUpdateRequest struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
}

type LinkedDatasetSummaryResponse struct {
	AccessLevel string   `json:"accessLevel"`
	Created     string   `json:"created"`
	Description string   `json:"description,omitempty"`
	ID          string   `json:"id"`
	License     string   `json:"license,omitempty"`
	Owner       string   `json:"owner"`
	Summary     string   `json:"summary,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Title       string   `json:"title"`
	Updated     string   `json:"updated"`
	Version     string   `json:"version"`
	Visibility  string   `json:"visibility"`
}

type OauthTokenReference struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Site  string `json:"site"`
}

type ProjectCreateOrUpdateRequest struct {
	Files          []FileCreateRequest                  `json:"files,omitempty"`
	License        string                               `json:"license,omitempty"`
	LinkedDatasets []LinkedDatasetCreateOrUpdateRequest `json:"linkedDatasets,omitempty"`
	Objective      string                               `json:"objective,omitempty"`
	Summary        string                               `json:"summary,omitempty"`
	Tags           []string                             `json:"tags,omitempty"`
	Title          string                               `json:"title"`
	Visibility     string                               `json:"visibility"`
}

type ProjectCreateResponse struct {
	Message string `json:"message,omitempty"`
	URI     string `json:"uri"`
}

type ProjectSummaryResponse struct {
	AccessLevel    string                         `json:"accessLevel"`
	Created        string                         `json:"created"`
	Files          []FileSummaryResponse          `json:"files,omitempty"`
	ID             string                         `json:"id"`
	License        string                         `json:"license,omitempty"`
	LinkedDatasets []LinkedDatasetSummaryResponse `json:"linkedDatasets,omitempty"`
	Objective      string                         `json:"objective,omitempty"`
	Owner          string                         `json:"owner"`
	Status         string                         `json:"status"`
	Summary        string                         `json:"summary,omitempty"`
	Tags           []string                       `json:"tags,omitempty"`
	Title          string                         `json:"title"`
	Updated        string                         `json:"updated"`
	Version        string                         `json:"version"`
	Visibility     string                         `json:"visibility"`
}

type QueryCreateRequest struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Language  string `json:"language"`
	Published bool   `json:"published,omitempty"`
}

type QueryParameter struct {
	Type     string `json:"type,omitempty"`
	Datatype string `json:"datatype,omitempty"`
}

type QuerySummaryResponse struct {
	Body       string                    `json:"body,omitempty"`
	Created    string                    `json:"created,omitempty"`
	ID         string                    `json:"id,omitempty"`
	Language   string                    `json:"language,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Owner      string                    `json:"owner,omitempty"`
	Updated    string                    `json:"updated,omitempty"`
	Version    string                    `json:"version,omitempty"`
	Parameters map[string]QueryParameter `json:"parameters,omitempty"`
}

type QueryUpdateRequest struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Published bool   `json:"published,omitempty"`
}

type SavedQueryExecutionRequest struct {
	Parameters         map[string]string `json:"parameters,omitempty"`
	IncludeTableSchema bool              `json:"includeTableSchema,omitempty"`
	MaxRows            float64           `json:"maxRows,omitempty"`
}

type SPARQLQueryRequest struct {
	Query string `json:"query"`
}

type SQLQueryRequest struct {
	Query              string `json:"query"`
	IncludeTableSchema bool   `json:"includeTableSchema,omitempty"`
}

type StreamSchema struct {
	PrimaryKeyFields []string `json:"primaryKeyFields,omitempty"`
	SequenceField    string   `json:"sequenceField,omitempty"`
}

type StreamSchemaUpdateRequest struct {
	PrimaryKeyFields []string `json:"primaryKeyFields,omitempty"`
	SequenceField    string   `json:"sequenceField,omitempty"`
	UpdateMethod     string   `json:"updateMethod"`
}

type Subscription struct {
	Dataset DatasetOrProjectIdentifier `json:"dataset,omitempty"`
	Events  []string                   `json:"events"`
	Project DatasetOrProjectIdentifier `json:"project,omitempty"`
	User    UserIdentifier             `json:"user,omitempty"`
}

type SubscriptionCreateRequest struct {
	Events []string `json:"events"`
}

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

type UserIdentifier struct {
	ID string `json:"id"`
}

type UserInfoResponse struct {
	AvatarURL   string `json:"avatarUrl,omitempty"`
	Created     string `json:"created"`
	DisplayName string `json:"displayName,omitempty"`
	ID          string `json:"id"`
	Updated     string `json:"updated"`
}

type WebAuthorization struct {
	Credentials string `json:"credentials,omitempty"`
	Type        string `json:"type"`
}

type WebCredentials struct {
	Password string `json:"password,omitempty"`
	User     string `json:"user"`
}
