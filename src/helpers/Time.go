package src_helpers

import (
	"fmt"
	"strconv"
	"time"
)

func GetFullTimeWithoutSymbol() string {
	local := time.Now().Local()
	year := strconv.Itoa(local.Year())
	month := fmt.Sprintf("%02d", local.Month())
	day := fmt.Sprintf("%02d", local.Day())
	hour := fmt.Sprintf("%02d", local.Hour())
	min := fmt.Sprintf("%02d", local.Minute())
	second := fmt.Sprintf("%02d", local.Second())
	return year + month + day + hour + min + second
}
