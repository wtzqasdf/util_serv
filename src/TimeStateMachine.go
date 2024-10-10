package src

import "time"

type TimeStatusMachine struct {
	lastSnmpTime          time.Time
	lastSqlBackupTime     time.Time
	delaySnmpSeconds      float64
	delaySqlBackupSeconds float64
}

func NewTimeStatusMachine(delaySnmpSeconds float64, delaySqlBackupSeconds float64) *TimeStatusMachine {
	tsm := new(TimeStatusMachine)
	tsm.lastSnmpTime = time.Now()
	tsm.lastSqlBackupTime = time.Now()
	tsm.delaySnmpSeconds = delaySnmpSeconds
	tsm.delaySqlBackupSeconds = delaySqlBackupSeconds
	return tsm
}

func (tsm *TimeStatusMachine) CanSnmpWork() bool {
	canWork := false
	if time.Since(tsm.lastSnmpTime).Seconds() > tsm.delaySnmpSeconds {
		tsm.lastSnmpTime = time.Now()
		canWork = true
	}
	return canWork
}

func (tsm *TimeStatusMachine) CanSqlBackup() bool {
	canWork := false
	if time.Since(tsm.lastSqlBackupTime).Seconds() > tsm.delaySqlBackupSeconds {
		tsm.lastSqlBackupTime = time.Now()
		canWork = true
	}
	return canWork
}
