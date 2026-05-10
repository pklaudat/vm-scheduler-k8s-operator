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

const (
	PhasePending     = "Pending"
	PhaseReconciling = "Reconciling"
	PhaseReady       = "Ready"
	PhaseError       = "Error"

	PowerStateRunning = "Running"
	PowerStateStopped = "Stopped"
)

type Schedule struct {
	// +kubebuilder:validation:Pattern=`^([01][0-9]|2[0-3]):([0-5][0-9])$`
	Start string `json:"start,omitempty"`

	// +kubebuilder:validation:Pattern=`^([01][0-9]|2[0-3]):([0-5][0-9])$`
	Stop string `json:"stop,omitempty"`
}

type VMSelector struct {
	Names []string `json:"names,omitempty"`

	// Dynamic VM selection using Azure tags
	//
	// Example:
	// tags:
	//   autoSchedule: "true"
	//   environment: "dev"
	//
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// AzureVmSchedulerSpec defines the desired state of AzureVmScheduler.
type AzureVmSchedulerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:default=true
	Enabled bool `json:"enabled,omitempty"`

	// +kubebuilder:validation:MinLength=1
	SubscriptionID string `json:"subscriptionId"`

	// +kubebuilder:validation:MinLength=1
	ResourceGroup string `json:"resourceGroup"`

	Selector VMSelector `json:"selector"`

	Schedule Schedule `json:"schedule"`

	Timezone string `json:"timezone,omitempty"`
}

type VMStatus struct {
	Name string `json:"name,omitempty"`

	DesiredPowerState string `json:"desiredPowerState,omitempty"`

	ActualPowerState string `json:"actualPowerState,omitempty"`

	// Last operation executed against the VM
	// Example:
	// - StartRequested
	// - StopRequested
	// - WaitingForCompletion
	// - Completed
	// - Failed
	LastOperation string `json:"lastOperation,omitempty"`

	// Last status update timestamp
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
}

// AzureVmSchedulerStatus defines the observed state of AzureVmScheduler.
type AzureVmSchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Current reconciliation phase
	//
	// Examples:
	// - Pending
	// - Reconciling
	// - Ready
	// - Error
	Phase string `json:"phase,omitempty"`

	// Current observed generation
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Last reconciliation execution
	LastScheduleExecution metav1.Time `json:"lastScheduleExecution,omitempty"`

	// Last successful reconciliation execution
	LastSuccessfulExecution metav1.Time `json:"lastSuccessfulExecution,omitempty"`

	// High-level last operation
	LastOperation string `json:"lastOperation,omitempty"`

	// Per-VM reconciliation status
	VMStatuses []VMStatus `json:"vmStatuses,omitempty"`

	// Standard Kubernetes conditions
	//
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:printcolumn:name="Enabled",type="boolean",JSONPath=".spec.enabled"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="LastRun",type="date",JSONPath=".status.lastScheduleExecution"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
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
