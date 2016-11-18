package rma

import (
	"os"
	"io/ioutil"
	"errors"
	"sync"
	"time"
	"path/filepath"
	"fmt"
	"strings"

	"github.com/mae-global/rigo2/rma/parser"
)

var (
	ErrRMANTREE = errors.New("is RMANTREE set")
)

var internal struct {
	sync.RWMutex
	Cache map[string]*cacheFile
}

func init() {
	internal.Cache = make(map[string]*cacheFile,0)
}

type cacheFile struct {
	Filepath string
	Contents []byte
	Loaded time.Time

	Asset *parser.RenderManAsset
}


type LibraryNode struct {

	Name string	
	Filepath string
	Dir bool
	Rank int
	Nodes []*LibraryNode
}

func depth(current int,root *LibraryNode) int {
	if root == nil {
		return current
	}

	if root.Rank > current {
		current = root.Rank
	}	

	for _,n := range root.Nodes {
		if d := depth(n.Rank,n); d > current {
			current = d
		}
	}

	return current
} 

func (node *LibraryNode) Depth() int {
	return depth(0,node)
}
	
	
func (node *LibraryNode) PrettyPrint() {
	fmt.Printf("node \"%s\"\n",node.Name)
	fmt.Printf(" @ %s\n",node.Filepath)
}


type RenderManAssetLibrary struct {
	Target string
		
	Root *LibraryNode
	names map[string]*LibraryNode
}

func (lib *RenderManAssetLibrary) Depth() int {
	return depth(0,lib.Root)
}

func (lib *RenderManAssetLibrary) Find(name string) *LibraryNode {
	if node,exists := lib.names[name]; exists {
		return node
	}

	return nil
}
	


func (lib *RenderManAssetLibrary) PrettyPrint() {
	writefmt(0,lib.Root)
	fmt.Printf("\n= library %s\n= %d assets\n",lib.Target,len(lib.names))
}



type RenderManAsset struct {
	*parser.RenderManAsset
}



/* TODO: add a heap of useful helper functions here */


func tofilepath(name,path string) (string,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return "",ErrRMANTREE
	}

	return rmantree + path + name + "/asset.json",nil
}

func read(name, filepath string) (*RenderManAsset,error) {

	internal.RLock()
	if cache,exists := internal.Cache[filepath]; exists {
		internal.RUnlock()
		return &RenderManAsset{cache.Asset},nil
	}
	internal.RUnlock()	


	file,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}

	asset,err := parser.Parse(name,file)
	if err != nil {
		return nil,err
	}
	
	internal.Lock()
	defer internal.Unlock()

	cache := &cacheFile{Filepath: filepath,Asset: asset,Loaded:time.Now()}
	cache.Contents = make([]byte,len(file))
	copy(cache.Contents,file)

	internal.Cache[filepath] = cache

	return &RenderManAsset{asset},nil
}

func Load(name string) (*RenderManAsset,error) {

	filepath,err := tofilepath(name,"/lib/RenderManAssetLibrary")
	if err != nil { 
		return nil,err
	}
	
	return read(name,filepath)
}

func writefmt(depth int,root *LibraryNode) {
	if root == nil {
		return
	}
	for i := 0; i < depth; i++ {
		fmt.Printf("   ")
	}
	if root.Dir {
		fmt.Printf("+(%d) %s\n",root.Rank,root.Name)	
	} else {
		fmt.Printf("*(%d) %s\n",root.Rank,root.Name)
	}

	for _,n := range root.Nodes {
		writefmt(depth + 1,n)	
	}
}

func LoadLibrary() (*RenderManAssetLibrary,error) {

	rmantree := os.Getenv("RMANTREE")
	if len(rmantree) == 0 {
		return nil,ErrRMANTREE
	}

	filep := rmantree + "/lib/RenderManAssetLibrary"

	files := make([]string,0)

	root := &LibraryNode{Dir:true}
	root.Nodes = make([]*LibraryNode,0)

	nodes := make(map[string]*LibraryNode,0)

	count := 0

	/* walk, finding all /asset.json files */
	walkf := func(path string,info os.FileInfo,err error) error {

			if err != nil {
				return err
			}

			if !info.IsDir() {
	
				if info.Name() == "asset.json" {
	
					f := strings.TrimPrefix(path,filep)
					files = append(files,f)

					/* strip the asset.json file, then break into path components */

					parts := strings.Split(f,"/")

					c := root
					
					for _,p := range parts {
						if len(p) > 0 && p != "asset.json" { 
														
							found := false
							/* check for first part in the root nodes */
							for _,n := range c.Nodes {
								if n.Name == p {
									found = true
									c = n
								}
							}

							if !found {

								if filepath.Ext(p) == ".rma" {
								
									count ++
											
									name := strings.Replace(strings.TrimSuffix(p,".rma"),"_"," ",-1)

									n := &LibraryNode{Dir:false,Name:name,Filepath:path,Rank:c.Rank + 1} 
									c.Nodes = append(c.Nodes,n)
									c = n	

									nodes[name] = n
			
								} else {

									n := &LibraryNode{Dir:true,Name:p,Rank:c.Rank + 1}
									n.Nodes = make([]*LibraryNode,0)
									c.Nodes = append(c.Nodes,n)
									c = n
								}
							}							
						}
					}				
				}
			} 
	
			return nil
	}	


	if err := filepath.Walk(filep,walkf); err != nil {
		return nil,err
	}



	library := &RenderManAssetLibrary{Target:filep}
	library.Root = root
	library.names = make(map[string]*LibraryNode,0)

	for key,value := range nodes {
		library.names[key] = value
	}	

	return library,nil
}







		









/* TODO: add a LoadLocal function instead of the specific filepath */
/* TODO: add a scan for rma files (.rma/asset.json) from a specific filepath root, returns a map useful for display */


