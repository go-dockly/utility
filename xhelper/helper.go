package xhelper

import (
	"github.com/go-dockly/utility/xlogger"
	"github.com/stretchr/testify/suite"
)

type Helper struct {
	suite  *suite.Suite
	logger *xlogger.Logger
}

func NewHelper(s *suite.Suite, logger *xlogger.Logger) *Helper {
	return &Helper{suite: s, logger: logger}
}
