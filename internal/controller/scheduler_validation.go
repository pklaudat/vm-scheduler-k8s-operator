package controller

import (
	
	"fmt"
	"time"

	klaudatiov1alpha1 "github.com/pklaudat/azure-vm-scheduler-operator/api/v1alpha1"
)


func validateSchedulerSpec(scheduler *klaudatiov1alpha1.AzureVmScheduler) error {

	spec := scheduler.Spec

	if len(spec.Names) == 0 && len(spec.Tags) == 0 {
		return fmt.Errorf("selector must contain at least one tag or vm name")
	}

	if spec.Schedule.Start == "" || spec.Schedule.stop == "" {
		return fmt.Errorf("schedule start time or stop time must not be empty")
	}

	startTime, err :=  time.Parse("00:00", spec.Schedule.Start)
	
	if err != nil {
		return fmt.Errorf("invalid start time %w", err)
	}
	
	stopTime, err := time.Parse("00:00", spec.Schedule.Stop)

	if err != nil {
		return fmt.Errof("invalid stop time %w", err)
	}

	if startTime.Equal(stopTime) {
		return fmt.Error("schedule.start and schedule.stop must not be the same")
	}
	

	return nil
}