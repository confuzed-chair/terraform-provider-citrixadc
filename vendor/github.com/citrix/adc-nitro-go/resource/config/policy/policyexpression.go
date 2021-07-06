/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package policy

/**
* Configuration for expression resource.
*/
type Policyexpression struct {
	/**
	* Unique name for the expression. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression string. For example: http.req.body(100).contains("this").
	*/
	Value string `json:"value,omitempty"`
	/**
	* Description for the expression.
	*/
	Description string `json:"description,omitempty"`
	/**
	* Any comments associated with the expression. Displayed upon viewing the policy expression.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Message to display if the expression fails. Allowed for classic end-point check expressions only.
	*/
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
	/**
	* Type of expression. Can be a classic or default syntax (advanced) expression.
	*/
	Type string `json:"type,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Pihits string `json:"pihits,omitempty"`
	Type1 string `json:"type1,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
