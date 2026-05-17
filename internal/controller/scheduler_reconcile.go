package controller

import (
	"context"
	"fmt"

	klaudatiov1alpha1 "github.com/pklaudat/azure-vm-scheduler-operator/api/v1alpha1"
	"github.com/pklaudat/azure-vm-scheduler-operator/pkg/azure"
)

func (r *AzureVmSchedulerReconciler) reconcileVmScheduler(
	ctx context.Context,
	scheduler *klaudatiov1alpha1.AzureVmScheduler,
) error {

	azureClient, err := azure.NewClient()

	if err != nil {
		fmt.Errorf("Failed to authenticate against azure - invalid controller app credentials")
	}

	vms, err := azureClient.ListVMsByTags(
		ctx,
		scheduler.Spec.SubscriptionID,
		scheduler.Spec.ResourceGroups,
		scheduler.Spec.Selector.Tags,
	)

	if err != nil {
		fmt.Errorf("Failed to list VMs by Tags")
	}

	for _, vm := range vms {

		powerState, err := azureClient.GetPowerState(ctx, vm.Subscription, vm.ResourceGroup, vm.Name)

		if err != nil {
			fmt.Errorf("Failed to get current power state for vm %s", vm.Name)

			continue
		}

		desiredState := r.shouldVMBeRunning(vm.Name, scheduler)

		if powerState != azure.PowerStateRunning && desiredState {

		}

	}

	return nil
}
