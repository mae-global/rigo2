package catribdriver

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"

	. "github.com/mae-global/rigo2/ri"
	. "github.com/mae-global/rigo2/ri/core"
	. "github.com/mae-global/rigo2/drivers"
)



type catribDriver struct {
	sync.RWMutex

	last string
	cmd  *exec.Cmd
	in   io.WriteCloser
	out  io.ReadCloser
	err  io.ReadCloser
}

func (d *catribDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* do nothing */
}

func (d *catribDriver) GetProgress() RtInt {
	return RtInt(100)
}

func (d *catribDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {
	d.last = RIBStream(name, args, tokens, values)
	fmt.Fprintf(d.in, "%s\n", d.last)
	return nil
}

func (d *catribDriver) Close() *RtError {

	d.out.Close()
	//d.cmd.Wait()
	return nil
}

func (d *catribDriver) GetLastRIB() string {
	return d.last
}

func BuildDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	//logger.Printf("CatribDriver options = %s\n",options)

	cmd := exec.Command("catrib", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	d := &catribDriver{}
	d.cmd = cmd
	d.in = stdin
	d.out = stdout
	d.err = stderr

	return d, d.cmd.Start()
}

