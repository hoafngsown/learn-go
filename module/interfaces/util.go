package interfaces

import "log"

type Util struct {
	Log         LogUtil
	Logger      *log.Logger
	Environment Environment
}
