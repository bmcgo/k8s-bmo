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

// IPMIEndpointSpec defines the desired state of IPMIEndpoint
type IPMIEndpointSpec struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`

	BootMacAddress string `json:"bootMacAddress,omitempty"`
}

// IPMIEndpointStatus defines the observed state of IPMIEndpoint
type IPMIEndpointStatus struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	BMCGUID      string `json:"bmcGUID,omitempty"`
}

//+kubebuilder:resource:shortName=ipmie
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="BareMetalNodeName",type="string",JSONPath=".status.bareMetalNodeName",description="Baare Metal Node Name",priority=0

// IPMIEndpoint is the Schema for the ipmiendpoints API
type IPMIEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IPMIEndpointSpec   `json:"spec,omitempty"`
	Status IPMIEndpointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IPMIEndpointList contains a list of IPMIEndpoint
type IPMIEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPMIEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IPMIEndpoint{}, &IPMIEndpointList{})
}
