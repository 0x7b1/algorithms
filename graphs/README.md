# Graphs

* `G = (V, E)`; `m = |V|`; `n = |E|`
* Adjacency matrix is an efficient way to encode a dense graph but is wasteful for a sparse graph
* Adjacency matrix takes O(n²) space, whereaas adjacency list takes O(m + n)
* Adjacency lists are perfect for graph exploration
* Why do we want to search a graph:
  * To check connectivity
  * Shortest path
  * Planning
  * Connected components

### Strongly Connected Components (SCC)

* A directed graph `G = (V, E)` is strongly connected if for all `v`,`w` in `V`: There's a path from `v` to `w` and from `w` to `v`
* We can decompose a graph into SCC. These tell you about "groups"
* ### Algorithm [View](/graphs/toposort.go)
  * Do DFS to create a DFS forest
    * Choose starting vertices in any order
    * Keep track of finishing times
  * Reverse all the edges in the graph
  * Do DFS again to create another DFS forest
    * This time, order the nodes in the reverse order of the finishing times that they had from the first DFS run
  * The SCCs are the different trees in **the second DFS forest**
* ### Complexity

| Time       |
| ---------- |
| `O(m + n)` |

### Breadth-first search (BFS)

* It's a good way to find connected components
* Useful for testing if a graph is bipartite or not
* Only works with unweighted graphs
* ### Algorithm [View](graphs/bfs.go)
  * Explore the vertices of a graph in "layers". Explores the neighbor nodes first, before moving to the next level neighbors
* ### Complexity

| Time       |
| ---------- |
| `O(m + n)` |

### Depth-first search (DFS)

DFS is a graph technique for searching exhaustively all the neighbors of each vertex that are connected. It goes deeply.

* ### Algorithm [View](graphs/dfs.go)
  * Traverse deep into the graph by visiting the children before sibling/neighbor nodes
  * Walk though a path, backtrack until we found a new path
* ### Complexity

| Time       |
| ---------- |
| `O(m + n)` |

### Topological Sort

Is a linear ordering of a DAG' nodes such that from every node `u` to node `v`, `u` comes before `v` in the ordering

* ### Algorithm [View](graphs/toposort.go)
  * Same as DFS but maintain an array updading the visited elements in the first position as in a stack
* ### Complexity

| Time       |
| ---------- |
| `O(m + n)` |

### Dijkstra's Algorithm

* Dijkstra is an algotihm that finds shortest paths in weightet graphs with non-negative edge weights
* The **cost** of a graph is the sum of the weights along that path
* The **shortest path** is the one with minimum cost
* Is very centralized: need to keep track of all the vertices to know which to update
* ### Algorithm [View](/graphs/dijkstra.go)
  * Init the shortest distance to MAX except for the initial nodhttps://aur.archlinux.org/qrcp.gite
  * Init a priority queue where the comparator will be on the total distance so far
  * Init a set to store all visited nodes
  * Add initial vertex to the priority queue
  * While queue is not empty: Mark vertex as visited, check the total distance to each neighbor, update shortest and previous arrays if smaller. If destination was unvisited, adds it to the queue
* ### Complexity

| Time             |
| ---------------- |
| `O(m + n log n)` |

### Bellman-Ford Algorithm

* Algorithm that solves the single source shortest paths problem on graphs with edges with potentially negative weights.
* Can be done in a distributed fashion, every vertex using only information from its neighbors
* ### Algorithm [View](/graphs/bellman-ford.go)
  * Given a directed graph and edges with weights
  * Detects a negative cycle if it exists and is reachable from s, or
  * Computes the shortest path distances
* ### Complexity

| Time    |
| ------- |
| `O(mn)` |

### Todo

* [ ] Floyd-Warshall Algorithm
* [ ] Prim's Algorithm
* [ ] Kruskal's Algorithm