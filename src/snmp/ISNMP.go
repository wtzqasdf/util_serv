package src_snmp

type ISNMP interface {
	GetUpsStatus() (UpsStatusType, error)
	GetUpsBatteryPercentage() (int, error)
}

type UpsStatusType int8

const (
	OnLine     UpsStatusType = 1
	OnBattery  UpsStatusType = 2
	LowBattery UpsStatusType = 3
)
