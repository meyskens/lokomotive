// Copyright 2020 The Lokomotive Authors
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

// Package gangway has code related to deployment of gangway component.
package gangway

const chartValuesTmpl = `
clusterName: {{ .ClusterName }}

ingressHost: {{ .IngressHost }}

sessionKey: {{ .SessionKey }}

apiServerURL: {{ .APIServerURL }}

authorizeURL: {{ .AuthorizeURL }}

tokenURL: {{ .TokenURL }}

clientID: {{ .ClientID }}

clientSecret: {{ .ClientSecret }}

redirectURL: {{ .RedirectURL }}

certManagerClusterIssuer: {{ .CertManagerClusterIssuer }}
`
