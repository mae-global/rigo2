package parser


type Param struct {
	Default interface{} `json:"default"`
	Type string `json:"type"`
	Name string `json:"name"`
	Value interface{} `json:"value"`
	Vstructmember string `json:"vstructmember"`
}


type Node struct {
	Params map[string]Param `json:"params"`
	RmanNode string `json:"rmanNode"`
	Type string `json:"type"`
	NodeClass string `json:"nodeClass"`
}


type Compatibility struct {
	HostNodeTypes []string `json:"hostNodeTypes"`
	Host map[string]string `json:"host"`
	Renderer map[string]string `json:"renderer"`
}

type Connection struct {
	Node string `json:"node"`
	Param string `json:"param"`
}

type ConnectionBlock map[string]Connection

type NodeGraph struct {
	Dependencies []string `json:"dependecies"`
	NodeList map[string]Node `json:"nodeList"`
	Compatibility Compatibility `json:"compactibility"`
	ConnectionList []ConnectionBlock `json:"connectionList"`
	Metadata map[string]string `json:"metadata"`
}

type Asset struct {
	NodeGraph NodeGraph `json:"nodeGraph"`
	Label string `json:"label"`
}

type RenderManAsset struct {
	
	UsedNodeTypes []string `json:"usedNodeTypes"`
	Version float64 `json:"version"`
	Asset Asset `json:"asset"`
}
	
type Outer struct {	
	RenderManAsset *RenderManAsset
}
