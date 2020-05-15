package xhelper

import (
	"github.com/go-dockly/utility/xlogger"
	"github.com/pelletier/go-toml"
)

// GetLogger returns the application default logger
func (h *Helper) GetLogger() *xlogger.Logger {
	l, err := xlogger.New(new(xlogger.Config))
	h.suite.Require().NoError(err)

	return l
}

// GetConfig returns the config.ServiceConfig
func (h *Helper) GetConfig(path string, cfg interface{}) {
	b := h.BytesFromFile(path)
	err := toml.Unmarshal(b, cfg)
	h.suite.Require().NoError(err)
}
