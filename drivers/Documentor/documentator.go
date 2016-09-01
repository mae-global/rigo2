package documentordriver

/* For documenting purposes, this driver will generate html/css code for both the minimal go program and the RIB output */ 

import (
	"sync"
	"log"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)


type documentorDriver struct {
	sync.RWMutex

	last string
	

	progress *ProtectedInteger
}

func (d *documentorDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {

}

func (d *documentorDriver) GetProgress() RtInt {
	d.progress.RLock()
	defer d.progress.RUnlock()
	return RtInt(d.progress.Value)
}

func (d *documentorDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {
	d.last = RIBStream(name, args, tokens, values)

	return nil
}

func (d *documentorDriver) Close() *RtError {

	return nil
}

func (d *documentorDriver) GetLastRIB() string {
	return d.last
}

func BuildDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {


	d := &documentorDriver{}
	d.progress = new(ProtectedInteger)
	d.progress.Value = 100

	return d, nil
}





