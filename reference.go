package bpaygo

type Status int64

const (
	NewStatus          Status = 1000
	PaidStatus         Status = 1001
	CancelledStatus    Status = 1002
	PayingStaus        Status = 1003
	ProviderPaidStatus Status = 1004
	ErrorStatus        Status = 1005
)
