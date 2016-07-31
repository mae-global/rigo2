package rigo

import (
	"fmt"
	"sync"

	"github.com/mae-global/rigo2/drivers"
	"github.com/mae-global/rigo2/drivers/RIB"
	"github.com/mae-global/rigo2/drivers/Render"
	"github.com/mae-global/rigo2/drivers/Catrib"

)

/* Load all the default drivers */


/* TODO: MOVE to blockdriver */
const (
	DefaultBlockFile = "out.rib.block"
)


var internal struct {
	sync.RWMutex
	Drivers map[string]drivers.Builder
}

func init() {
	
	/* build the default drivers table */
	dd := make(map[string]drivers.Builder,0)
	//dd["block"] = BuildBlockDiagrammingDriver
	//dd["stdout"] = BuildRIBStdoutDriver
	//dd["catrib"] = BuildCatribDriver
	//dd["render"] = BuildRenderDriver
	
	dd["file"]   = ribdriver.BuildFileDriver
	dd["stdout"] = ribdriver.BuildStdoutDriver

	dd["render"] = renderdriver.BuildDriver

	dd["catrib"] = catribdriver.BuildDriver
	
	internal.Drivers = dd
}


func AddDriver(name string, builder drivers.Builder) error {
	internal.Lock()
	defer internal.Unlock()
	internal.Drivers[name] = builder
	return nil /* TODO */
}

func RemoveDriver(name string) error {
	internal.Lock()
	defer internal.Unlock()
	delete(internal.Drivers,name)
	return nil
}

func FindDriver(name string) (drivers.Builder,error) {
	internal.RLock()
	defer internal.RUnlock()

	if d,exists := internal.Drivers[name]; exists {
		return d,nil
	}

	return nil,fmt.Errorf("Driver \"%s\" Not Found",name)
}

