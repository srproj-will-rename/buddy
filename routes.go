package buddy

import "fmt"

const (
	routeDelimiter = "_"
	streamIdentifier = "StreamController"
	controllerIdentifier = "Controller"
)

// this code is kind of messy

type RoutingTable map[string]Route

func (rt RoutingTable) Add(path string, route Route) {
	rt[path] = route
}

// Internally represents the routing table for the router
// maybe a mutex here
var routingTable = RoutingTable{}

func Table() RoutingTable {
	fmt.Println(routingTable)
	return routingTable
}

type Route struct {
	controller ControllerType
	function string
	persist bool
}

func (r *Route) Controller() ControllerType {
	return r.controller
}

func (r *Route) Function() string {
	return r.function
}

func Lookup(route string) Route {
	return Table()[route]
}

func Add(path string, controller ControllerType, function string) {
	fmt.Println("adding")
	fmt.Println(path)
	fmt.Println(controller)
	routingTable.Add(path, Route{controller: controller, function: function})
}
