package renderdriver

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


type renderDriver struct {
	sync.RWMutex

	last string
	cmd  *exec.Cmd
	in   io.WriteCloser
	out  io.ReadCloser
	err  io.ReadCloser

	progress *ProtectedInteger
}

func (d *renderDriver) Flush(marker RtString, synchronous RtBoolean, flushmode RtToken) {
	/* TODO: talk to renderer via Ric */
}

func (d *renderDriver) GetProgress() RtInt {
	d.progress.RLock()
	defer d.progress.RUnlock()
	return RtInt(d.progress.Value)
}

func (d *renderDriver) Handle(name RtString, args []RtPointer, tokens []RtPointer, values []RtPointer) *RtError {
	d.last = RIBStream(name, args, tokens, values)
	fmt.Fprintf(d.in, "%s\n", d.last)
	return nil
}

func (d *renderDriver) Close() *RtError {

	d.out.Close()

	/* wait on the process to finish */
	//d.cmd.Wait()

	return nil
}

func (d *renderDriver) GetLastRIB() string {
	return d.last
}

func BuildDriver(logger *log.Logger, options []RtPointer, args ...string) (Driver, error) {

	//logger.Printf("RenderDriver options = %s\n",options)

	cmd := exec.Command("render", args...)

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

	d := &renderDriver{}
	d.cmd = cmd
	d.in = stdin
	d.out = stdout
	d.err = stderr
	d.progress = new(ProtectedInteger)

	go func(pi *ProtectedInteger) {
		buffer := make([]byte, 256)
		for {
			n, err := stderr.Read(buffer)
			if err != nil {
				if err == io.EOF {
					return
				}

				logger.Fatal(err)
			}

			str := string(buffer[:n])
			/* basic progress parser */
			end := -1
			for i, c := range str {
				if c == '%' {
					end = i
					break
				}
			}

			if end != -1 && end > 0 {
				p, err := strconv.Atoi(strings.TrimSpace(str[:end]))
				if err != nil {
					logger.Printf("unable to parse progress -- %v (%s)\n", err, str)
					continue
				}

				pi.Lock()
				pi.Value = p
				pi.Unlock()
			}
		}
	}(d.progress)

	return d, d.cmd.Start()
}





