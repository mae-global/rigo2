package blockdriver

import (
	"sync"
	"log"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)


type blockDriver struct {
	sync.RWMutex

	last string


	progress *ProtectedInteger
}

func (d *blockDriver) Flush(marker RtString,synchronous RtBoolean,flushmode RtToken) {

}

func (d *blockDriver) GetProgress() RtInt {
	d.progress.RLock()
	defer d.progress.RUnlock()
	return RtInt(d.progress.Value)
}

func (d *blockDriver) Handle(name RtString,args []RtPointer,tokens []RtPointer,values []RtPointer) *RtError {

	d.last = RIBStream(name,args,tokens,values)

	return nil
}

func (d *blockDriver) Close() *RtError {
	return nil
}

func (d *blockDriver) GetLastRIB() string {
	return d.last
}

func BuildDriver(logger *log.Logger,options []RtPointer,args ...string) (Driver,error) {

	d := &blockDriver{}
	d.progress = new(ProtectedInteger)
	d.progress.Value = 100

	return d,nil
}


