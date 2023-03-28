package solution

import "runtime"

func Version() string {
	return runtime.Version()
}
