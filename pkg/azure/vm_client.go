package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

func (c *Client) GetPowerState(
	ctx context.Context,
	subscriptionID string,
	resourceGroup string,
	vmName string,
) (string, error) {

	vmClient, err := armcompute.NewVirtualMachinesClient(
		subscriptionID,
		c.credential,
		nil,
	)

	if err != nil {
		return "", fmt.Errorf(
			"failed to create vm client: %w",
			err,
		)
	}

	resp, err := vmClient.InstanceView(
		ctx,
		resourceGroup,
		vmName,
		nil,
	)

	if err != nil {
		return "", fmt.Errorf(
			"failed to get vm instance view: %w",
			err,
		)
	}

	for _, status := range resp.Statuses {

		if status.Code == nil {
			continue
		}

		switch *status.Code {

		case "PowerState/running":
			return PowerStateRunning, nil

		case "PowerState/deallocated":
			return PowerStateStopped, nil

		case "PowerState/stopped":
			return PowerStateStopped, nil
		}
	}

	return PowerStateUnknown, nil
}

// StartVM starts a virtual machine.
//
// This operation is asynchronous in Azure.
// The reconciler should requeue and observe completion later.
func (c *Client) StartVM(
	ctx context.Context,
	subscriptionID string,
	resourceGroup string,
	vmName string,
) error {

	vmClient, err := armcompute.NewVirtualMachinesClient(
		subscriptionID,
		c.credential,
		nil,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create vm client: %w",
			err,
		)
	}

	_, err = vmClient.BeginStart(
		ctx,
		resourceGroup,
		vmName,
		nil,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to start vm %s: %w",
			vmName,
			err,
		)
	}

	return nil
}

// StopVM deallocates a virtual machine.
//
// Deallocate fully stops billing for compute resources.
// This operation is asynchronous in Azure.
func (c *Client) StopVM(
	ctx context.Context,
	subscriptionID string,
	resourceGroup string,
	vmName string,
) error {

	vmClient, err := armcompute.NewVirtualMachinesClient(
		subscriptionID,
		c.credential,
		nil,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create vm client: %w",
			err,
		)
	}

	_, err = vmClient.BeginDeallocate(
		ctx,
		resourceGroup,
		vmName,
		nil,
	)
	if err != nil {
		return fmt.Errorf(
			"failed to stop vm %s: %w",
			vmName,
			err,
		)
	}

	return nil
}

// ListVMsByTags lists all VMs matching the provided tags.
//
// Current implementation:
// - scans all VMs in the resource group
// - filters client-side
//
// Future optimization:
// - Azure Resource Graph
// - Azure Resource Manager filtering
func (c *Client) ListVMsByTags(
	ctx context.Context,
	subscriptionID string,
	resourceGroup string,
	tags map[string]string,
) ([]string, error) {

	vmClient, err := armcompute.NewVirtualMachinesClient(
		subscriptionID,
		c.credential,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create vm client: %w",
			err,
		)
	}

	pager := vmClient.NewListPager(
		resourceGroup,
		nil,
	)

	var matchedVMs []string

	for pager.More() {

		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to list vms: %w",
				err,
			)
		}

		for _, vm := range page.Value {

			if vm.Name == nil {
				continue
			}

			if vm.Tags == nil {
				continue
			}

			match := true

			for expectedKey, expectedValue := range tags {

				actualValue, exists := vm.Tags[expectedKey]

				if !exists {
					match = false
					break
				}

				if actualValue == nil {
					match = false
					break
				}

				if *actualValue != expectedValue {
					match = false
					break
				}
			}

			if match {
				matchedVMs = append(
					matchedVMs,
					*vm.Name,
				)
			}
		}
	}

	return matchedVMs, nil
}
