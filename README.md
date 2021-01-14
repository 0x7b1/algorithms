# Algorithms & Data Structures

## Index
* ### [I. Time Complexity](complexity.md)
* ### [II. Sorting](./sorting/README.md)
* ### [III. Data Structures](./datastructures/README.md)
* ### [IV. Graphs](./graphs/README.md)
* ### [V. Concurrency](./concurrency/README.md)

## Heuristics
- Divide-and-Conquer
- Dynamic Programming
- Greedy Programming

## Algorithm Toolbox
- Lists, arrays, stack
- Trees
- Sorting and searching
- Priority queues
- Pattern matching and parsing
- Hashing
- Disjoint sets
- Graph algorithms
- Minimum spanning tree
- Shortest path
- State space search algorithms

| Algorithm or Data Structure | Runtime                                                                                             | Use Cases                                                                                                                                                                                                                                                                                                                                                                                                            |
| --------------------------- | --------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Hash Tables                 | `O(1)` insertion, lookup, and deletion                                                              | - When you only need to store and lookup objects.<br/>- When you need to partition a list of objects distinct groups by some property (basically what a group by in SQL does)<br/>- You need to count the number of distinct items in a list                                                                                                                                                                         |
| Linked Lists                | `O(1)` insertion, lookup, and deletion at the ends or next to a node you already have a pointer to. | The main use cases of a linked list revolve around the fact that a linked list maintains relative ordering of the elements. In programming interviews, a linked list is primarily used as either a stack or queue.                                                                                                                                                                                                   |
| Binary Trees                | `O(log n)` insertion, lookup, and deletion.                                                         | Used when you need to keep your data in sorted order. This lets you quickly answer questions like how many elements fall into a given range or what is the Nth highest element in the tree.                                                                                                                                                                                                                          |
| Binary Search               | `O(log n)`                                                                                          | You need to find the number in a sorted array closest to another number.<br/>You need to find the smallest number in a sorted array that is larger than another number.<br/>You need to find the largest number in a sorted array that is smaller than another number.<br/>If you aren’t able to use a hash table for whatever reason, you can use a binary search to check if a given element is in a sorted array. |
| Depth-first Search          | `O(n)`                                                                                              | Traverse an entire graph.<br/>Find a specific element in a graph.<br/>Find the connected components of a graph.                                                                                                                                                                                                                                                                                                      |
| Sorting                     | `O(n log n)`                                                                                        | Can be used if you need to process elements in a specific order. First sort by that order, then iterate through the elements.<br/>Can be used to sort an array that you will later perform binary search on.                                                                                                                                                                                                         |

## References
- CS 161 - Design and Analysis of Algorithms
- Algorithms Illuminated Book series
- Algorithms 4th Edition
