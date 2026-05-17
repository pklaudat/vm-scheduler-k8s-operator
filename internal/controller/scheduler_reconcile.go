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
		return fmt.Errorf("failed to authenticate against azure - invalid controller app credentials: %w", err)
	}

	vms, err := azureClient.ListVMsByTags(
		ctx,
		scheduler.Spec.SubscriptionID,
		scheduler.Spec.ResourceGroups,
		scheduler.Spec.Selector.Tags,
	)

	if err != nil {
		return fmt.Errorf("failed to list VMs by tags: %w", err)
	}

	for _, vm := range vms {

		powerState, err := azureClient.GetPowerState(ctx, vm.Subscription, vm.ResourceGroup, vm.Name)

		if err != nil {
			fmt.Printf("Failed to get current power state for vm %s - %w", vm.Name, err)

			continue
		}

		desiredState := r.shouldVMBeRunning(vm.Name, scheduler)

		if powerState != azure.PowerStateRunning && desiredState {

		}

	}

	return nil
}
