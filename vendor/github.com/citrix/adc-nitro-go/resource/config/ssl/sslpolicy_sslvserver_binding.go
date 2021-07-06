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

package ssl

/**
* Binding class showing the sslvserver that can be bound to sslpolicy.
*/
type Sslpolicysslvserverbinding struct {
	/**
	* The entity name to which policy is bound
	*/
	Boundto string `json:"boundto,omitempty"`
	Priority int `json:"priority,omitempty"`
	Activepolicy int `json:"activepolicy,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Type of policy label invocation.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to invoke if the current policy rule evaluates to TRUE.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Name of the SSL policy for which to display detailed information.
	*/
	Name string `json:"name,omitempty"`


}