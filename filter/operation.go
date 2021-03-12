package filter

import (
	"github.com/hackfengJam/filter/core"
)

func init() {
}

// GetOperation Get Operation
func GetOperation(name string) core.Operation {
	return core.GetOperationFactory().Get(name)
}
