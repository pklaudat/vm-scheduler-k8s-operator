package azure

const (
	PowerStateRunning = "Running"
	PowerStateStopped = "Stopped"
	PowerStateUnknown = "Unknown"
)

type ScheduleVM struct {
	Name          string
	ResourceGroup string
	Subscription  string
}
