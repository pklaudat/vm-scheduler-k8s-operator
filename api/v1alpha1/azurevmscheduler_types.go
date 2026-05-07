/*
Copyright 2026 Paulo Klaudat.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureVmSchedulerSpec defines the desired state of AzureVmScheduler.
type AzureVmSchedulerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AzureVmScheduler. Edit azurevmscheduler_types.go to remove/update
	Foo string `json:"foo,omitempty"`
	SubscriptionId string `json:"subsriptionid"`
	ResourceGroup string `json:"resourceGroup"`
	VmNames string[] `json:"vmNames"`
	startUpAt string `json:"startUpAt"`
	shutdownAt string `json:"shutdownAt"`
	timezone string `json:"timezone,omitempty"`
}

// AzureVmSchedulerStatus defines the observed state of AzureVmScheduler.
type AzureVmSchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureVmScheduler is the Schema for the azurevmschedulers API.
type AzureVmScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureVmSchedulerSpec   `json:"spec,omitempty"`
	Status AzureVmSchedulerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureVmSchedulerList contains a list of AzureVmScheduler.
type AzureVmSchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVmScheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureVmScheduler{}, &AzureVmSchedulerList{})
}
