package src_os

import "util_serv/src"

type ISystem interface {
	Shutdown()
	SqlBackup(conf *src.Config)
}
