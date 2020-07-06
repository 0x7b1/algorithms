package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Unvisited  = "UNVISITED"
	InProgress = "IN_PROGRESS"
	Done       = "DONE"

	Visited = "VISITED"
)

const Infinity = int(^uint(0) >> 1)

type Vertex struct {
	inNeighbors  []*VertexWeight
	outNeighbors []*VertexWeight
	value        int
	status       string
	parent       *Vertex
	estD         int
}

type Edge struct {
	from   *Vertex
	to     *Vertex
	weight int
}

type VertexWeight struct {
	v      *Vertex
	weight int
}

func (v *Vertex) hasOutNeighbor(u *Vertex) bool {
	for _, w := range v.getOutNeighbors() {
		if w == u {
			return true
		}
	}
	return false
}

func (v *Vertex) hasInNeighbor(u *Vertex) bool {
	for _, w := range v.getInNeighbors() {
		if u == w {
			return true
		}
	}

	return false
}

func (v *Vertex) hasNeighbor(u *Vertex) bool {
	return v.hasOutNeighbor(u) || v.hasInNeighbor(u)
}

func (v *Vertex) getOutNeighbors() []*Vertex {
	var arr []*Vertex
	for _, vm := range v.outNeighbors {
		arr = append(arr, vm.v)
	}

	return arr
}

func (v *Vertex) getInNeighbors() []*Vertex {
	var arr []*Vertex
	for _, vm := range v.inNeighbors {
		arr = append(arr, vm.v)
	}

	return arr
}

func (v *Vertex) getOutNeighborsWithWeights() []*VertexWeight {
	return v.outNeighbors
}

func (v *Vertex) getInNeighborsWithWeights() []*VertexWeight {
	return v.inNeighbors
}

func (v *Vertex) addOutNeighbor(u *Vertex, weight int) {
	v.outNeighbors = append(v.outNeighbors, &VertexWeight{u, weight})
}

func (v *Vertex) addInNeighbor(u *Vertex, weight int) {
	v.inNeighbors = append(v.inNeighbors, &VertexWeight{u, weight})
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

func (g *Graph) addDiEdge(u, v *Vertex, weight int) {
	u.addOutNeighbor(v, weight)
	v.addInNeighbor(u, weight)
}

func (g *Graph) addBiEdge(u, v *Vertex, weight int) {
	g.addDiEdge(u, v, weight)
	g.addDiEdge(v, u, weight)
}

func (g *Graph) getDirEdges() []*Edge {
	var edges []*Edge

	for _, v := range g.vertices {
		for _, u := range v.outNeighbors {
			edges = append(edges, &Edge{v, u.v, u.weight})
		}
	}

	return edges
}

func (g *Graph) String() string {
	ret := "Vertices:\n\t"
	for _, v := range g.vertices {
		ret += fmt.Sprint(v) + ","
	}

	ret += "\n"
	ret += "Edges:\n\t"

	for _, de := range g.getDirEdges() {
		ret += fmt.Sprintf("(%s -> %s; %d) ", de.from, de.to, de.weight)
	}

	ret += "\n"

	return ret
}

func randomGraph(n int, p float32, weight []int) *Graph {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	G := new(Graph)
	V := make([]*Vertex, n)

	for i := range V {
		V[i] = &Vertex{
			inNeighbors:  []*VertexWeight{},
			outNeighbors: []*VertexWeight{},
			value:        i,
			status:       Unvisited,
			parent:       nil,
			estD:         Infinity,
		}

		G.addVertex(V[i])
	}

	for _, v := range V {
		for _, w := range V {
			if v != w {
				if rnd.Float32() < p {
					G.addDiEdge(v, w, weight[rnd.Intn(len(weight))])
				}
			}
		}
	}

	return G
}

func dijkstraDumb(w *Vertex, G *Graph) {
	for _, v := range G.vertices {
		v.estD = Infinity
	}
	w.estD = 0
	unsureVertices := G.vertices[:]

	for len(unsureVertices) > 0 {
		var u *Vertex
		minD := Infinity
		for _, x := range unsureVertices {
			if x.estD < minD {
				minD = x.estD
				u = x
			}
		}

		if u == nil {
			return
		}

		for _, vw := range u.getOutNeighborsWithWeights() {
			if u.estD+vw.weight < vw.v.estD {
				vw.v.estD = u.estD + vw.weight
				vw.v.parent = u
			}
		}

		var unsureVerticesTmp []*Vertex
		for _, v := range unsureVertices {
			if v != u {
				unsureVerticesTmp = append(unsureVerticesTmp, v)
			}
		}

		unsureVertices = unsureVerticesTmp
	}
}

func dijkstraDumb_shortestPaths(w *Vertex, G *Graph) {
	dijkstraDumb(w, G)
	for _, v := range G.vertices {
		if v.estD == Infinity {
			fmt.Println("Cannot reach", v)
			continue
		}

		var path []*Vertex
		current := v
		for current != w {
			path = append(path, current)
			current = current.parent
		}

		path = append(path, current)
		fmt.Println(path)
	}

}

func main() {
	G := randomGraph(5, .2, []int{1})
	fmt.Println(G)
	dijkstraDumb_shortestPaths(G.vertices[0], G)
}
