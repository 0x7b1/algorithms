# Algorithms & Data Structures

## I. Time Complexity

### Asymptotic Notation
Idea: Suppress constant factors and lower-order terms

|Big O - O(…)|Big Omega - Ω(…)|Big Theta - Θ(…)|
|----------|-------------|------|
|<ul><li>Upper bound</li><li>T(n) is O(f(n)) if f(n) grows at least as fast as T(n) as n gets large</li></ul> | <ul><li>Lower bound</li>T(n) is Ω(f(n)) if f(n) grows at most as fast as T(n) as n gets large</li></ul> | <ul><li>Lower and Upper bound</li><li>T(n) = O(f(n)) AND T(n) = Ω(f(n))</li></ul> |
|![](images/big_o.png)|![](images/big_omega.png)|![](images/big_theta.png)|

### Master Theorem
![](images/master_theorem.png)

**Example** : Binary Search
- Recurses on either the left half of the input array or the right half.
- One recursive call, `a=1`
- The recursive call is on half of the input array, `b=2`
- Outside the recursive call, a single comparison between the middle element and the searched one. `O(1); d=0`
- Since `a = 1 = 20 = b^d` -> `O(n^d log n) = O(log n)`

## II. Sorting
### Insertion Sort
- ### Algorithm [View](sorting/insertion.go)
    - From i to 0..n, insert a[i] to its correct position to the left (0..i)

- ### Complexity
    | Time | Space |
    |---|---|
    | `O(n²)` | `O(1)` |

### Merge Sort
- ### Algorithm [View](sorting/mergesort.go)
    - Splits a collection into two halves
    - Sort the two halves (recursive call)
    - Merge them together to form one sorted collection

- ### Complexity
    | Time | Space |
    |---|---|
    | `Θ(n log n)` | `O(n)` |

### Quick Sort
- ### Algorithm [View](sorting/quicksort.go)
    - Choose a pivot element `p`
    - Partition array `A` around `p`
    - recursively sort first part of `A`
    - recursively sort second part of `A`

- ### Complexity
    | Time | Space |
    |---|---|
    | <ul><li>Best and average: `O(n log n)`</li><li>Worst: `O(n²)`</li> | `O(n log n)` |

## III. Data Structures
- They exist for organizing and manipulating data
- Different data structures support different set of operations
- The more operations a data structure supports, the slower the operations and greater the space overhead
- __Rule:__ Choose the simplest data structure that supports all the operations required by the problem

### Heap

- Data structure that keeps track of an evolving set of objects with keys and can quickly identify the object with the mininum of maximum key.
- In a max heap, the keys of parent nodes are always greater than or equal to those of the children and the highest key is in the root node.
- In a min heap, the keys of parent nodes are less than or equal to those of the children and the lowest key is in the root node

![](images/heap.png)

- ### Algorithm [View](datastructures/heap.go)
    - #### Insert
        - Stick the new object at the end of the heap and increment the heap size
        - Repeatedly swap the new object with its parent until the heap property is restored
    - #### ExtractMin
        - **Operation:** Given a heap `H`, remove and return from `H` an object with the smallest key
        - Overwrite the root with the last object `x` in the heap, and decrement the heap size
        - Repeatedly swap `x` with its smallest child until the heap property is restored

- ### Complexity
    |Operation|Running Time|
    |---|---|
    |Insert|`O(log n)`|
    |ExtractMin|`O(log n)`|
    |FindMin|`O(1)`|
    |Heapify|`O(n)`|
    |Delete|`O(log n)`|
    
### Binary Search Tree

- Data structure optimized for -wait for it- search. Compared to heap that optimizes fast minimum computations.
- Type of binary tree which maintains the property that the value in each node must be greater than or equal to any value stored in the left sub-tree, and less than or equal to any value stored in the right sub-tree

- ### Algorithm [View](datastructures/bst.go)
    - #### Insert
        - Start at the root node
        - Repeatedly traverses left and right child pointers, until a null pointer is encountered
        - Replace the null pointer with one to the new object. Set the new node's parent pointer to its parent, and child pointers to null
    - #### Delete
        - Use Search to locate an object `x` with key `k`. (If no such object exists, halt)
        - If `x` has no children, delete `x` by setting the appropriate child pointer of `x`'s parent to null. (If `x` was the root, the new tree is empty)
        - If `x` has one child, splice `x` out by rewiring the appropriate child pointer of `x`'s child to `x`'s parent. (If `x` was the root, its child becomes the new root)
        - Otherwise, swap `x` with the object in its left subtree that has the biggest key, and delete `x` from its new position (where it has at most one child).
    - #### Search
        - Start at the root node
        - Repeatedly traverses the left and right child pointers. Left if `k` is less than current node's key, right otherwise
        - Return a pointer to an object with key `k` or none
    - #### Min (Max)
        - Start at the root node
        - Traverse left (or right) child pointers until reaching a null pointer
        - Return the pointer of the last visited object
    - #### Predecessor
        - If `x`'s left subtree is non-empty, return the result of Max applied to this subtree.
        - Otherwise, traverse parent pointers upward toward the root. If the traversal visits consecutive nodes `y` and `z` with `y` a right child of `z`, return a pointer to `z`.
        - Otherwise, return none.
    
- ### Complexity
    |Operation|Running Time|
    |---|---|
    |Insert|`O(log n)`|
    |Delete|`O(log n)`|    
    |Min, Max|`O(log n)`| 
    |Predecessor|`O(log n)`|

### Black-Red Tree

- ### Algorithm

- ### Complexity

### B-Tree

- ### Algorithm

- ### Complexity

### Hash Table

- ### Algorithm

- ### Complexity

### Bloom Filter

- ### Algorithm

- ### Complexity

