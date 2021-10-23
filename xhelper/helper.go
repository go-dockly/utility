package xhelper

import (
	"github.com/go-dockly/utility/xlogger"
	"github.com/stretchr/testify/suite"
)

// Helper for testify suite
type Helper struct {
	suite  *suite.Suite
	logger *xlogger.Logger
}

// NewHelper constructs a helper class to ease the most mundane test tasks
func NewHelper(s *suite.Suite, logger *xlogger.Logger) *Helper {
	return &Helper{suite: s, logger: logger}
}
