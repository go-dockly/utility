package xhelper

import (
	"encoding/json"
)

// FromJSON decodes json into its expected struct
func (h *Helper) FromJSON(b []byte, expected interface{}) {
	err := json.Unmarshal(b, expected)
	h.suite.Require().NoError(err)
}
