package bandwidth

import "time"

const path = "/sys/class/net/"

type bwData struct {
	rx, tx    float64
	timestamp time.Time
}

var Stats []bwData

type rxGraphData struct {
	rx        float64
	timestamp int64
}

type txGraphData struct {
	tx        float64
	timestamp int64
}
