package src

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	//Private
	baseURL string
	//Public
	LineToken                 string
	SnmpServerIP              string
	SnmpServerCommunity       string
	SnmpServerDevice          string
	UpsBatteryPercentShutdown int
	SqlUsername               string
	SqlPassword               string
	SqlBackupDirectory        string
	SqlBackupDatabases        []string
}

func NewConfig() *Config {
	dir, _ := os.Getwd()

	config := new(Config)
	config.baseURL = dir
	config.SqlBackupDatabases = make([]string, 0)
	config.readConfig()
	return config
}

func (conf *Config) readConfig() {
	//Read file content
	bytes, _ := os.ReadFile(conf.baseURL + "/app.config")
	content := string(bytes)

	//Parse content
	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		pair := strings.Split(strings.Trim(lines[i], "\r"), "=")
		if len(pair) > 1 {
			pair[0] = strings.Trim(pair[0], " ")
			pair[1] = strings.Trim(pair[1], " ")
			conf.assignProperty(pair[0], pair[1])
		}
	}
}

func (conf *Config) assignProperty(key string, value string) {
	if key == "snmp_server_ip" {
		conf.SnmpServerIP = value
	}
	if key == "snmp_server_community" {
		conf.SnmpServerCommunity = value
	}
	if key == "snmp_server_device" {
		conf.SnmpServerDevice = value
	}
	if key == "ups_battery_percent_shutdown" {
		result, _ := strconv.ParseInt(value, 10, 0)
		conf.UpsBatteryPercentShutdown = int(result)
	}
	if key == "sql_username" {
		conf.SqlUsername = value
	}
	if key == "sql_password" {
		conf.SqlPassword = value
	}
	if key == "sqlbackup_directory" {
		conf.SqlBackupDirectory = value
	}
	if key == "sqlbackup_databases" {
		databases := strings.Split(value, ",")
		conf.SqlBackupDatabases = databases
	}
}
