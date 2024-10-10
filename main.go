package main

import (
	"fmt"
	"time"
	"util_serv/src"
	src_os "util_serv/src/os"
	src_snmp "util_serv/src/snmp"
)

func run(
	tsm *src.TimeStatusMachine,
	conf *src.Config,
	snmp src_snmp.ISNMP,
	system src_os.ISystem) {
	for {
		if tsm.CanSnmpWork() {
			status, _ := snmp.GetUpsStatus()
			battery, _ := snmp.GetUpsBatteryPercentage()
			//if ups changed as battery and battery less than 50 percent, to shutdown
			if status == src_snmp.OnBattery && battery < conf.UpsBatteryPercentShutdown {
				system.Shutdown()
				break
			}
		}
		if tsm.CanSqlBackup() {
			system.SqlBackup(conf)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	timeStatusMachine := src.NewTimeStatusMachine(1, 60*24)
	conf := src.NewConfig()
	snmp := src_snmp.NewASUSTOR_SNMP(conf)
	system := src_os.NewUbuntu()
	go run(timeStatusMachine, conf, snmp, system)

	fmt.Println("Util_Serv is started")
	time.Sleep(time.Duration(24*365*50) * time.Hour)
}
