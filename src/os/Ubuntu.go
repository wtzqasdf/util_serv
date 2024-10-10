package src_os

import (
	"os"
	"os/exec"
	"util_serv/src"
	src_helpers "util_serv/src/helpers"
)

type Ubuntu struct {
	ISystem
}

func NewUbuntu() *Ubuntu {
	s := new(Ubuntu)
	return s
}

func (system *Ubuntu) Shutdown() {
	cmd := exec.Command("sudo", "shutdown", "-h", "now")
	cmd.Run()
}

func (system *Ubuntu) SqlBackup(conf *src.Config) {
	time := src_helpers.GetFullTimeWithoutSymbol()
	for i := 0; i < len(conf.SqlBackupDatabases); i++ {
		dbname := conf.SqlBackupDatabases[i]
		cmd := exec.Command("mysqldump", "-u", conf.SqlUsername, "-p"+conf.SqlPassword, dbname)
		//Write to file
		filename := conf.SqlBackupDirectory + dbname + "_" + time + ".sql"
		outFile, _ := os.Create(filename)
		defer outFile.Close()
		cmd.Stdout = outFile
		cmd.Run()
	}
}
