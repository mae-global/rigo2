package main

import (
	"net/http"
	"flag"
	"log"
	"html/template"
	"sync"
	"fmt"
	"strings"

	"github.com/mae-global/rigo2/rma"	

	"github.com/gorilla/mux"
)

/* Concept :- load a prebuilt gzip archive, with renders and navigations already done */


func main() {

	host := flag.String("host","localhost:8080","host interface")

	flag.Parse()

	log.Printf("RiGO2/RenderManAssetLibrary/Browser\n")
	log.Printf("Version 1\n")
	log.Printf("Running on %s\n",*host)

	ctx := &Context{}
	ctx.Libraries = make(map[string]*rma.RenderManAssetLibrary,0)

	library,err := rma.LoadLibrary() /* load the default library */
	if err != nil {
		log.Fatal(err)
	}

	ctx.Libraries[library.Target] = library

	t := template.New("index").Funcs(template.FuncMap{})
	t,err = t.Parse(html_template)
	if err != nil {
		log.Fatal(err)
	}

	ctx.T = t
	ctx.def = library.Target

	log.Printf("loaded default %s library\n",library.Target)

	r := mux.NewRouter()

	r.HandleFunc("/",wrap(ctx,IndexHandler)).Methods("GET")
	r.HandleFunc("/library/{library}/{nav}",wrap(ctx,NavHandler)).Methods("GET")
	r.HandleFunc("/library/{library}/",wrap(ctx,NavHandler)).Methods("GET")
	r.HandleFunc("/image/{name}",wrap(ctx,ImageHandler)).Methods("GET")

	http.Handle("/",r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


type Context struct {
	mux sync.RWMutex
	def string

	Libraries map[string]*rma.RenderManAssetLibrary
	
	
	T *template.Template
}

func (ctx *Context) Image(w http.ResponseWriter,req *http.Request,name string) error {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()

	node := ctx.Libraries[ctx.def].Find(name)
	if node == nil {
		return fmt.Errorf("Not Found")
	}

	node.PrettyPrint()

	//log.Printf("Image %s -> %s\n",name,node.Path() + "/asset_100.png")
	http.ServeFile(w,req, node.Path() + "/asset_100.png")
	return nil
} 

func (ctx *Context) Render(w http.ResponseWriter,fragment string,val interface{}) error {
	ctx.mux.RLock()
	defer ctx.mux.RUnlock()

	if err := ctx.T.ExecuteTemplate(w,fragment,val); err != nil {
		return fmt.Errorf("template: \"%s\" execute fragment error -- %v",fragment,err)
	}

	return nil
}


func wrap(ctx *Context,f func(*Context,http.ResponseWriter,*http.Request)) func(http.ResponseWriter,*http.Request) {

	return func(w http.ResponseWriter,req *http.Request) { f(ctx,w,req) }
}


type T struct {
	Header interface{}
	Service interface{}
	Info interface{}
	Body interface{}
}


func IndexHandler(ctx *Context,w http.ResponseWriter,req *http.Request) {

	/* build the library menus */



	if err := ctx.Render(w,"index",T{}); err != nil {
		log.Printf(err.Error())
		http.Error(w,err.Error(),500)
	}	
}

type TSectionPart struct {
	Label string	
	Selected bool
	Href string
}

type TSection struct {
	Navigation []TSectionPart
	Assets []TSectionPart
	Filepath string
}


func NavHandler(ctx *Context,w http.ResponseWriter,req *http.Request) {

	vars := mux.Vars(req)
	library := vars["library"]
	nav := vars["nav"]

	parts := strings.Split(nav,",")
	
	/* TODO: build menu */
	
	log.Printf("NavHandler \"%s\" --> %v\n",library,parts)

	sections := make([]TSection,0)

	root := ctx.Libraries[ctx.def].Root

	common := ""
	
	for _,part := range parts {

		section := TSection{Filepath:root.Filepath}
		section.Navigation = make([]TSectionPart,0)
		section.Assets = make([]TSectionPart,0)



		/* go through all the nodes and render */
		for _,node := range root.Nodes {
			if node.Dir {
				selected := false
				if node.Name == part {
					selected = true
				}

				href := ""
				if len(common) > 0 {
					href = fmt.Sprintf("/library/RenderMan/%s,%s",common,node.Name)
				} else {
					href = fmt.Sprintf("/library/RenderMan/%s",node.Name)
				}

				section.Navigation = append(section.Navigation,TSectionPart{Label:node.Name,Selected:selected,Href:href}) 

			} else {

				href := fmt.Sprintf("/image/%s",strings.Replace(node.Name," ","_",-1))

				section.Assets = append(section.Assets,TSectionPart{Label:node.Name,Href:href})	
			}
		}	
		
		sections = append(sections,section)
	
		found := false
		/* move the root node onwards */
		for _,node := range root.Nodes {
			if node.Dir && node.Name == part {
				root = node
				found = true
				break
			}
		}

		if !found {
			break
		}

		if len(common) > 0 {
			common = fmt.Sprintf("%s,%s",common,part)
		} else {
			common = part
		}
	}

	if root.Name != "" { /* not root */
		/* do the tail */
		section := TSection{}
		section.Navigation = make([]TSectionPart,0)
		section.Assets = make([]TSectionPart,0)

		for _,node := range root.Nodes {
			if node.Dir {
				href := ""
				if len(common) > 0 {
					href = fmt.Sprintf("/library/RenderMan/%s,%s",common,node.Name)
				} else {
					href = fmt.Sprintf("/library/RenderMan/%s",node.Name)
				}
	
				section.Navigation = append(section.Navigation,TSectionPart{Label:node.Name,Href:href}) 
			} else {
	
				href := fmt.Sprintf("/image/%s",strings.Replace(node.Name," ","_",-1))
				section.Assets = append(section.Assets,TSectionPart{Label:node.Name,Href:href})	
			}		
		}

		sections = append(sections,section)
	}


	if err := ctx.Render(w,"library",T{Body:sections}); err != nil {
		log.Printf(err.Error())
		http.Error(w,err.Error(),500)
	}
}



func ImageHandler(ctx *Context,w http.ResponseWriter,req *http.Request) {
	
	vars := mux.Vars(req)
	name := strings.Replace(vars["name"],"_"," ",-1)

	log.Printf("ImageHandler \"%s\"\n",name)


	if err := ctx.Image(w,req,name); err != nil {
		log.Printf(err.Error())
		http.Error(w,err.Error(),404)
	}
}

























