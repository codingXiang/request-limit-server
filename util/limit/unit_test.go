package limit

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

//Suite struct
type UnitSuite struct {
	suite.Suite
}

//TestStart 為測試程式進入點
func TestStartUnitSuite(t *testing.T) {
	suite.Run(t, new(UnitSuite))
}

func (s *UnitSuite) TestNewUnit() {
	var u Unit
	u = NewUnit(SECOND)
	assert.Equal(s.T(), u, Second)
	u = NewUnit(MINUTE)
	assert.Equal(s.T(), u, Minute)
	u = NewUnit(HOUR)
	assert.Equal(s.T(), u, Hour)
	u = NewUnit("other")
	assert.Equal(s.T(), u, Minute)
}

func (s *UnitSuite) TestDuration() {
	var u Unit
	u = NewUnit(SECOND)
	assert.Equal(s.T(), u.Duration(), time.Second)
	u = NewUnit(MINUTE)
	assert.Equal(s.T(), u.Duration(), time.Minute)
	u = NewUnit(HOUR)
	assert.Equal(s.T(), u.Duration(), time.Hour)
	u = NewUnit("other")
	assert.Equal(s.T(), u.Duration(), time.Minute)
}