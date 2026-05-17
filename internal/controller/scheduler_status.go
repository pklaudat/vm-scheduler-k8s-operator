package controller

import (
	klaudatiov1alpha1 "github.com/pklaudat/azure-vm-scheduler-operator/api/v1alpha1"
)

func (r *AzureVmSchedulerReconciler) shouldVMBeRunning(
	vmName string,
	scheduler *klaudatiov1alpha1.AzureVmScheduler,
) bool {

	// TBD

	return true

}
