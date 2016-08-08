package documentatordriver

/* For documenting purposes, this driver will generate html/css code for both the minimal go program and the RIB output */ 

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)


type documentatorDriver struct {
	sync.RWMutex

	last string
	

	progress *ProtectedInteger
}

func (d *documentatorDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* TODO: talk to renderer via Ric */
}

func (d *documentatorDriver) GetProgress() RtInt {
	d.progress.RLock()
	defer d.progress.RUnlock()
	return RtInt(d.progress.Value)
}

func (d *documentatorDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {
	d.last = RIBStream(name, args, tokens, values)

	return nil
}

func (d *documentatorDriver) Close() *RtError {

	d.out.Close()

	/* wait on the process to finish */
	//d.cmd.Wait()

	return nil
}

func (d *documentatorDriver) GetLastRIB() string {
	return d.last
}

func BuildDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {


	d := &documentatorDriver{}
	d.progress = new(ProtectedInteger)
	d.progress.Value = 100

	return d, nil
}





