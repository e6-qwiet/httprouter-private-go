package namespace

import(
	"github.com/kyawmyintthein/httprouter"
)

// Namespace is a struct that can hold muliple routes and sub namespace
type Namespace struct {
	Name string
	Routes []*Route
	ChildNamespaces []*Namespace
	IsRoot  bool
}

// Route is a struct that can hold route information such as HttpMethod, Path, and Handle func
// Example : &Route{Path: "/",HttpMethod: "POST", Handle: func(http.ResponseWriter, *http.Request, Params)}
type Route struct {
	Path    string
	HttpMethod  string
	Func 	httprouter.Handle
}

// New is a func to create a namespace as root
// It can have multiple sub-namespace
// No namespace on top of root
func New(name string) *Namespace{
	return &Namespace{
		Name: name,
		IsRoot: true,
	}
}

// RouteTo is to keep the route information in namespace
func (n *Namespace) RouteTo(httpMethod, path string, handle httprouter.Handle){
	n.Routes = append(n.Routes, &Route{Path: path,HttpMethod: httpMethod,Func: handle})
}

// Handle the route from namaspace
func (root *Namespace) Handle(r *httprouter.Router){
	if root.IsRoot == true{
		root.handleRoutesFromNamespace(root.Name, r)
		if root.hasChildNamespaces(){
			for _, namespace := range root.ChildNamespaces{
				namespace.handleRoutesFromNamespace(root.Name + namespace.Name, r)
				if namespace.hasChildNamespaces(){
					handleChildNamespaces(namespace.ChildNamespaces, r)
				}
			}
		}
	}else{
		panic("Namespace is not root. Use RootNamespace function to create as root.")
	}
}

// Check namespace have childs
func (n *Namespace) hasChildNamespaces() bool{
	return len(n.ChildNamespaces) > 0
}


// Handle routes for child namespace
func handleChildNamespaces(childs []*Namespace,r *httprouter.Router){
	for _, child := range childs{
		child.handleRoutesFromNamespace(child.Name, r)
		if child.hasChildNamespaces(){
			handleChildNamespaces(child.ChildNamespaces, r)
		}
	}
	
}

// put routes to handler
func (n *Namespace) handleRoutesFromNamespace(prefix string, r *httprouter.Router){
	for _, route := range n.Routes{
		switch route.HttpMethod{
			case "GET":
				r.GET(prefix + route.Path,route.Func)
			case "POST":
				r.POST(prefix + route.Path,route.Func)
			case "OPTIONS":
				r.OPTIONS(prefix + route.Path,route.Func)
			case "HEAD":
				r.HEAD(prefix + route.Path,route.Func)
			case "PUT":
				r.PUT(prefix + route.Path,route.Func)
			case "PATCH":
				r.PATCH(prefix + route.Path,route.Func)
			case "DELETE":
				r.DELETE(prefix + route.Path,route.Func)
			default:
				panic("Invalid http method in routes")
		}
	}
}

// Use is to add sub-namespace for a namesapce
func (n *Namespace) Use(name string) *Namespace{
	newPath := n.Name + name
	subNamespace := &Namespace{Name: newPath}
	n.ChildNamespaces = append(n.ChildNamespaces, subNamespace)
	return subNamespace
}
