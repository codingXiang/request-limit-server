package limit

import (
	"github.com/spf13/viper"
	"time"
)


/*
Range	時間限制物件
- param
	- unit : 單位
	- per  : 數量
- method
	- NewRange : 建立物件 instance
	- Get      : 取得實際時間範圍
*/
type Range struct {
	unit time.Duration
	per  int
}

func NewRange(conf *viper.Viper) *Range {
	u := NewUnit(conf.GetString("limit.range.unit"))
	p := conf.GetInt("limit.range.per")
	return &Range{
		unit: u.Duration(),
		per:  p,
	}
}

func (t *Range) Get() time.Duration {
	return time.Duration(t.per) * t.unit
}
