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

// DesiredState
// +kubebuilder:validation:Enum=NotManaged;Inspected;Provisioned;PowerOff
type DesiredState string

// ActualState
// +kubebuilder:validation:Enum=NotManaged;Inspecting;Inspected;Provisioning;Provisioned;PoweredOff
type ActualState string

const (
	DesiredStateNotManaged  DesiredState = "NotManaged"
	DesiredStateInspected   DesiredState = "Inspected"
	DesiredStateProvisioned DesiredState = "Provisioned"
	DesiredStatePowerOff    DesiredState = "PowerOff"

	ActualStateNotManaged   ActualState = "NotManaged"
	ActualStateInspecting   ActualState = "Inspecting"
	ActualStateInspected    ActualState = "Inspected"
	ActualStateProvisioning ActualState = "Provisioning"
	ActualStateProvisioned  ActualState = "Provisioned"
	ActualStatePoweredOff   ActualState = "PoweredOff"
)

// SystemSpec defines the desired state of System
type SystemSpec struct {
	State DesiredState `json:"state"`
}

// SystemStatus defines the observed state of System
type SystemStatus struct {
	Id    string      `json:"id"`
	State ActualState `json:"state"`
}

//+kubebuilder:resource:shortName=sys
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="DesiredState",type="string",JSONPath=".spec.state",description="Desired state",priority=0
//+kubebuilder:printcolumn:name="ActualState",type="string",JSONPath=".status.state",description="Actual state",priority=0

// System is the Schema for the systems API
type System struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemSpec   `json:"spec,omitempty"`
	Status SystemStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SystemList contains a list of System
type SystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []System `json:"items"`
}

func init() {
	SchemeBuilder.Register(&System{}, &SystemList{})
}
