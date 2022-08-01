/*
Copyright 2022 The BMCGO Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AuthenticationMode
// +kubebuilder:validation:Enum=AuthNone;BasicAuth;OemAuth;RedfishSessionAuth
type AuthenticationMode string

const (
	AuthNone               AuthenticationMode = "AuthNone"
	AuthBasic              AuthenticationMode = "BasicAuth"
	AuthOemAuth            AuthenticationMode = "OemAuth"
	AuthRedfishSessionAuth AuthenticationMode = "RedfishSessionAuth"
)

// RedfishEndpointSpec defines the desired state of RedfishEndpoint
type RedfishEndpointSpec struct {
	AuthenticationMode AuthenticationMode `json:"authenticationMode"`
	BasicAuth          BasicAuth          `json:"basicAuth"`
	RedfishSessionAuth RedfishSessionAuth `json:"redfishSessionAuth"`
	//OemAuth TODO
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RedfishSessionAuth struct {
	//TODO
}

// RedfishEndpointStatus defines the observed state of RedfishEndpoint
type RedfishEndpointStatus struct {
	ErrorMessage string      `json:"errorMessage"`
	LastUpdated  metav1.Time `json:"lastUpdated"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RedfishEndpoint is the Schema for the redfishendpoints API
type RedfishEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedfishEndpointSpec   `json:"spec,omitempty"`
	Status RedfishEndpointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RedfishEndpointList contains a list of RedfishEndpoint
type RedfishEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedfishEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedfishEndpoint{}, &RedfishEndpointList{})
}
