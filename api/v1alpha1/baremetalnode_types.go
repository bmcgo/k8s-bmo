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
	"context"
	"github.com/bmcgo/k8s-bmo/ipmitool"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	ActualStatePoweredOff   ActualState = "PowerOff"
)

// BareMetalNodeSpec defines the desired state of BareMetalNode
type BareMetalNodeSpec struct {
	State DesiredState `json:"state"`
}

// BareMetalNodeStatus defines the observed state of BareMetalNode
type BareMetalNodeStatus struct {
	Id         string      `json:"id,omitempty"`
	State      ActualState `json:"state"`
	BMCGUID    string      `json:"bmcGUID,omitempty"`
	LastUpdate metav1.Time `json:"lastUpdate,omitempty"`
}

//+kubebuilder:resource:shortName=bmn
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="DesiredState",type="string",JSONPath=".spec.state",description="Desired state",priority=0
//+kubebuilder:printcolumn:name="ActualState",type="string",JSONPath=".status.state",description="Actual state",priority=0

// BareMetalNode is the Schema for the systems API
type BareMetalNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BareMetalNodeSpec   `json:"spec,omitempty"`
	Status BareMetalNodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BareMetalNodeList contains a list of BareMetalNode
type BareMetalNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BareMetalNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BareMetalNode{}, &BareMetalNodeList{})
}

func (b BareMetalNode) GetIPMIClient(ctx context.Context, c client.Client) (*ipmitool.IpmiTool, error) {
	ep := IPMIEndpoint{}
	err := c.Get(ctx, client.ObjectKey{Name: b.Name, Namespace: b.Namespace}, &ep)
	if err != nil {
		return nil, err
	}
	it := ipmitool.New(ep.Spec.Host, int(ep.Spec.Port), ep.Spec.Username, ep.Spec.Password)
	return &it, nil
}
