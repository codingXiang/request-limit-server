package limit

import "time"

/*
Unit 限制時間單位
- 目前支援以下單位
	- Second : 秒
	- Minute : 分
	- Hour   : 小時
*/
type Unit int

const (
	SECOND = "second"
	MINUTE = "minute"
	HOUR   = "hour"
)

const (
	Second Unit = iota
	Minute
	Hour
)

func NewUnit(unit string) Unit {
	switch unit {
	case SECOND:
		return Second
	case MINUTE:
		return Minute
	case HOUR:
		return Hour
	default:
		return Minute
	}
}

//轉換單位為 duration
func (u Unit) Duration() time.Duration {
	switch u {
	case Second:
		return time.Second
	case Minute:
		return time.Minute
	case Hour:
		return time.Hour
	default:
		return time.Minute
	}
}