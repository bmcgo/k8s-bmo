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
// +kubebuilder:validation:Enum=Idle;Inspecting;Provisioning;PoweredOff
type ActualState string

const (
	StateNotManaged  DesiredState = "NotManaged"
	StateInspected   DesiredState = "Inspected"
	StateProvisioned DesiredState = "Provisioned"
	StatePowerOff    DesiredState = "PowerOff"

	StateIdle         ActualState = "Idle"
	StateInspecting   ActualState = "Inspecting"
	StateProvisioning ActualState = "Provisioning"
	StatePoweredOff   ActualState = "PoweredOff"
)

// SystemSpec defines the desired state of System
type SystemSpec struct {
	State DesiredState `json:"foo,omitempty"`
}

// SystemStatus defines the observed state of System
type SystemStatus struct {
	Id    string      `json:"id"`
	State ActualState `json:"state"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

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
