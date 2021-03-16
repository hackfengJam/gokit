package isafe

import (
	"runtime"

	"github.com/hackfengJam/gokit/ilog"
)

// Recover recover
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if r := recover(); r != nil {
		buf := make([]byte, 64<<10)
		buf = buf[:runtime.Stack(buf, false)]
		ilog.Errorf("[Recovery] panic recovered: %s\n%s", r, buf)
	}
	return
}
