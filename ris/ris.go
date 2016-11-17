package ris

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mae-global/rigo2/ris/args"
)

var (
	ErrRMANTREE = errors.New("is RMANTREE set")
)

var internal struct {
	sync.RWMutex
	Cache map[string]*cacheFile
}

func init() {
	internal.Cache = make(map[string]*cacheFile, 0)
}

type cacheFile struct {
	Filepath string
	Contents []byte
	Loaded   time.Time

	Info *args.Information
	/* TODO */
}

type Information struct {
	*args.Information
}

func (info *Information) ListOfParams() []string {
	out := make([]string, 0)
	for _, param := range info.Params {
		out = append(out, string(param.Name))
	}

	for _, page := range info.Pages {
		for _, param := range page.Params {
			out = append(out, string(param.Name))
		}
	}
	return out
}

func (info *Information) ListOfPages() []string {
	out := make([]string, 0)
	for _, page := range info.Pages {
		out = append(out, string(page.Name))
	}
	return out
}

func (info *Information) ListOfOutputs() []string {
	out := make([]string, 0)
	for _, output := range info.Outputs {
		out = append(out, string(output.Name))
	}
	return out
}

func (info *Information) LookupParam(name string) *args.InfoParam {
	for _, param := range info.Params {
		if string(param.Name) == name {
			return &param
		}
	}

	for _, page := range info.Pages {
		for _, param := range page.Params {
			if string(param.Name) == name {
				return &param
			}
		}
	}

	return nil
}

func (info *Information) LookupOutput(name string) *args.InfoOutput {
	for _, output := range info.Outputs {
		if string(output.Name) == name {
			return &output
		}
	}
	return nil
}

func read(name, filepath string) (*Information, error) {

	internal.RLock()
	if cache, exists := internal.Cache[filepath]; exists {
		internal.RUnlock()
		return &Information{cache.Info}, nil
	}
	internal.RUnlock()

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	info, err := args.Parse(name, file)
	if err != nil {
		return nil, err
	}

	internal.Lock()
	defer internal.Unlock()

	cache := &cacheFile{Filepath: filepath, Info: info, Loaded: time.Now()}
	cache.Contents = make([]byte, len(file))
	copy(cache.Contents, file)

	internal.Cache[filepath] = cache

	return &Information{info}, nil
}

func tofilepath(name, path string) (string, error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return "", ErrRMANTREE
	}

	debug := os.Getenv("RIGO2_DEBUG")
	if debug == "testing" {
		name = strings.Replace(name, "Pxr", "Test", -1)
	}

	return rmantree + path + name + ".args", nil
}

func PrintStats() {
	internal.RLock()
	defer internal.RUnlock()
	now := time.Now()

	fmt.Printf("%d in cache...\n", len(internal.Cache))
	for _, cache := range internal.Cache {
		fmt.Printf("\tloaded %s ago [%s] = %s\n", now.Sub(cache.Loaded), cache.Info.Name, cache.Info.ShaderType)
	}
	fmt.Printf("\n")
}

func Load(name string) (*Information,error) {

	filepath, err := tofilepath(name,"/lib/plugins/Args/")
	if err != nil {
		return nil,err
	}
	return read(name,filepath)
}

func Generic(name, filepath string) (*Information, error) {
	return read(name, filepath)
}
