package wiz

import (
	time "time"
)

// Now returns unix timestamp of now.
func Now() uint64 {
	utc := time.Now()         //gets time from system
	return uint64(utc.Unix()) //Converts unix time from int64 to uint64
}

// Sleep sleeps for n seconds(int)
func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
