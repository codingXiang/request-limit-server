package limit

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

//Suite struct
type Suite struct {
	suite.Suite
	config *viper.Viper
}

//初始化 Suite
func (s *Suite) SetupSuite() {
	s.config = viper.New()
	s.config.Set("limit.range.unit", "minute")
	s.config.Set("limit.range.per", "1")

}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

var (
	testRangeObj = &Range{
		unit: time.Minute,
		per:  1,
	}
)

func (s *Suite) TestNewRange() {
	r := NewRange(s.config)
	assert.Equal(s.T(), r, testRangeObj)
}

func (s *Suite) TestGet() {
	r := NewRange(s.config)
	assert.Equal(s.T(), r.Get(), testRangeObj.Get())
}
