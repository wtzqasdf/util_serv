package src_snmp

import (
	"util_serv/src"

	"github.com/alouca/gosnmp"
)

type ASUSTOR_SNMP struct {
	ISNMP
	//Private
	conf *src.Config
}

func NewASUSTOR_SNMP(conf *src.Config) *ASUSTOR_SNMP {
	snmp := new(ASUSTOR_SNMP)
	snmp.conf = conf
	return snmp
}

func (snmp *ASUSTOR_SNMP) GetUpsStatus() (UpsStatusType, error) {
	remote, err := gosnmp.NewGoSNMP(snmp.conf.SnmpServerIP, snmp.conf.SnmpServerCommunity, gosnmp.Version2c, 5)
	if err != nil {
		return OnLine, err
	}
	packet, err := remote.Get("1.3.6.1.4.1.44738.6.6.0")
	if err != nil {
		return OnLine, err
	}
	var result UpsStatusType = OnLine
	for _, v := range packet.Variables {
		if v.Value == "OL" {
			result = OnLine
		} else if v.Value == "OB DISCHRG" {
			result = OnBattery
		} else {
			result = LowBattery
		}
	}
	return result, nil
}

func (snmp *ASUSTOR_SNMP) GetUpsBatteryPercentage() (int, error) {
	remote, err := gosnmp.NewGoSNMP(snmp.conf.SnmpServerIP, snmp.conf.SnmpServerCommunity, gosnmp.Version2c, 5)
	if err != nil {
		return 0, err
	}
	packet, err := remote.Get("1.3.6.1.4.1.44738.6.7.0")
	if err != nil {
		return 0, err
	}
	var result int = 0
	for _, v := range packet.Variables {
		result = v.Value.(int)
	}
	return result, nil
}
