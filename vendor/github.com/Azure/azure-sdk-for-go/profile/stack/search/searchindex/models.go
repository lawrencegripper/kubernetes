// +build go1.9

// Copyright 2017 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/lawrencegripper/azure-sdk-for-go/tools/profileBuilder

package searchindex

import original "github.com/lawrencegripper/azure-sdk-for-go/service/search/2016-09-01/searchindex"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DocumentsProxyClient = original.DocumentsProxyClient
type IndexActionType = original.IndexActionType

const (
	Delete		IndexActionType	= original.Delete
	Merge		IndexActionType	= original.Merge
	MergeOrUpload	IndexActionType	= original.MergeOrUpload
	Upload		IndexActionType	= original.Upload
)

type QueryType = original.QueryType

const (
	Full	QueryType	= original.Full
	Simple	QueryType	= original.Simple
)

type SearchMode = original.SearchMode

const (
	All	SearchMode	= original.All
	Any	SearchMode	= original.Any
)

type DocumentIndexResult = original.DocumentIndexResult
type IndexingResult = original.IndexingResult
type Int64 = original.Int64
type SearchParametersPayload = original.SearchParametersPayload
type SuggestParametersPayload = original.SuggestParametersPayload

func NewDocumentsProxyClient() DocumentsProxyClient {
	return original.NewDocumentsProxyClient()
}
func NewDocumentsProxyClientWithBaseURI(baseURI string) DocumentsProxyClient {
	return original.NewDocumentsProxyClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func New() ManagementClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) ManagementClient {
	return original.NewWithBaseURI(baseURI)
}
