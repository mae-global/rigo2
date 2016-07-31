package drivers

import (
	"sync"
	"log"

	. "github.com/mae-global/rigo2/ri/core"
)



type Driver interface {
	Flush(marker RtString, synchronous RtBoolean, flushmode RtToken)
	GetProgress() RtInt
	Handle(RtString, []RtPointer, []RtPointer, []RtPointer) *RtError
	Close() *RtError
	GetLastRIB() string
}

/* construction function for creating a new driver */
type Builder func(logger *log.Logger, options []RtPointer, args ...string) (Driver,error)


type ProtectedInteger struct {
	sync.RWMutex
	Value int
}






