package main

import "fmt"

const (
	Unvisited  = "UNVISITED"
	InProgress = "IN_PROGRESS"
	Done       = "DONE"
)

type Vertex struct {
	inNeighbors  []*Vertex
	outNeighbors []*Vertex
	value        string
	// useful for DFS
	inTime  int
	outTime int
	status  string
}

func (v *Vertex) hasOutNeighbor(u *Vertex) bool {
	for _, w := range v.outNeighbors {
		if w == u {
			return true
		}
	}

	return false
}

func (v *Vertex) hasInNeighbor(u *Vertex) bool {
	for _, w := range v.inNeighbors {
		if u == w {
			return true
		}
	}

	return false
}

func (v *Vertex) hasNeighbor(u *Vertex) bool {
	return v.hasOutNeighbor(u) || v.hasInNeighbor(u)
}

func (v *Vertex) addOutNeighbor(u *Vertex) {
	v.outNeighbors = append(v.outNeighbors, u)
}

func (v *Vertex) addInNeighbor(u *Vertex) {
	v.inNeighbors = append(v.inNeighbors, u)
}

func (v *Vertex) String() string {
	return fmt.Sprint(v.value)
}

// Directed graph
type Graph struct {
	vertices []*Vertex
}

func (g *Graph) addVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) addDiEdge(u, v *Vertex) {
	u.addOutNeighbor(v)
	v.addInNeighbor(u)
}

func (g *Graph) addBiEdge(u, v *Vertex) {
	g.addDiEdge(u, v)
	g.addDiEdge(v, u)
}

func (g *Graph) getDirEdges() [][]*Vertex {
	var ret [][]*Vertex
	for _, v := range g.vertices {
		for _, u := range v.outNeighbors {
			ret = append(ret, []*Vertex{v, u})
		}
	}

	return ret
}

func (g *Graph) String() string {
	ret := "Vertices:\n\t"
	for _, v := range g.vertices {
		ret += fmt.Sprint(v) + ","
	}

	ret += "\n"
	ret += "Edges:\n\t"

	for _, de := range g.getDirEdges() {
		ret += fmt.Sprintf("(%s,%s) ", de[0], de[1])
	}

	ret += "\n"

	return ret
}

func DFS_helper(w *Vertex, currentTime int, verbose bool) int {
	if verbose {
		fmt.Println("Time", currentTime, "entering", w.value)
	}

	w.inTime = currentTime
	currentTime++
	w.status = InProgress

	for _, v := range w.outNeighbors {
		if v.status == Unvisited {
			currentTime= DFS_helper(v, currentTime, verbose)
			currentTime++
		}
	}

	w.outTime = currentTime
	w.status = Done
	if verbose {
		fmt.Println("Time", currentTime, "leaving", w.value)
	}

	return currentTime
}

func DFS(w *Vertex, g *Graph, verbose bool)  {
	for _, v := range g.vertices {
		v.status = Unvisited
		v.inTime = 0
		v.outTime = 0
	}

	DFS_helper(w, 0, verbose)
}

func topoSort_helper(w *Vertex, currentTime int, ordering *[]*Vertex, verbose bool) int {
	if verbose {
		fmt.Println("Time", currentTime, "entering", w.value)
	}

	w.inTime = currentTime
	currentTime++
	w.status = InProgress

	for _, v := range w.outNeighbors {
		if v.status == Unvisited {
			currentTime = topoSort_helper(v, currentTime, ordering, verbose)
			currentTime++
		}
	}

	*ordering = append([]*Vertex{w}, *ordering...)

	w.outTime = currentTime
	w.status = Done

	if verbose {
		fmt.Println("Time", currentTime, "leaving", w.value)
	}

	return currentTime
}

func topoSort(w *Vertex, g *Graph, verbose bool) []*Vertex {
	var ordering []*Vertex

	for _, v := range g.vertices {
		v.status = Unvisited
		v.inTime = 0
		v.outTime = 0
	}

	topoSort_helper(w, 0, &ordering, verbose)

	return ordering
}

func main() {
	G := new(Graph)
	v_E := &Vertex{value: "E", status: Unvisited}
	v_F := &Vertex{value: "F", status: Unvisited}
	v_G := &Vertex{value: "G", status: Unvisited}
	v_H := &Vertex{value: "H", status: Unvisited}
	v_I := &Vertex{value: "I", status: Unvisited}
	v_J := &Vertex{value: "J", status: Unvisited}

	G.addVertex(v_E)
	G.addVertex(v_F)
	G.addVertex(v_G)
	G.addVertex(v_H)
	G.addVertex(v_I)
	G.addVertex(v_J)

	G.addDiEdge(v_E, v_F)
	G.addDiEdge(v_F, v_G)
	G.addDiEdge(v_G, v_E)
	G.addDiEdge(v_G, v_J)
	G.addDiEdge(v_H, v_G)
	G.addDiEdge(v_H, v_I)

	fmt.Println(G)

	DFS(v_E, G, true)

	fmt.Println("\n---------------\n")

	sorted := topoSort(v_H, G, true)
	for _, v := range sorted {
		fmt.Print(" -> ", v)
	}
}
