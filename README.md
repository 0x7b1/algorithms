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
- #### Algorithm [View](sorting/insertion.go)
    - From i to 0..n, insert a[i] to its correct position to the left (0..i)

- #### Complexity
    | Time | Space |
    |---|---|
    | `O(n²)` | `O(1)` |

### Merge Sort
- #### Algorithm [View](sorting/mergesort.go)
    - Splits a collection into two halves
    - Sort the two halves (recursive call)
    - Merge them together to form one sorted collection

- #### Complexity
    | Time | Space |
    |---|---|
    | `Θ(n log n)` | `O(n)` |

### Quick Sort
- #### Algorithm [View](sorting/quicksort.go)
    - Choose a pivot element `p`
    - Partition array `A` around `p`
    - recursively sort first part of `A`
    - recursively sort second part of `A`

- #### Complexity
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

- #### Algorithm [View](datastructures/heap.go)
    - ##### Insert
        - Stick the new object at the end of the heap and increment the heap size
        - Repeatedly swap the new object with its parent until the heap property is restored
    - ##### ExtractMin
        - **Operation:** Given a heap `H`, remove and return from `H` an object with the smallest key
        - Overwrite the root with the last object `x` in the heap, and decrement the heap size
        - Repeatedly swap `x` with its smallest child until the heap property is restored

- #### Complexity
    |Operation|Running Time|
    |--- |---|
    |Insert|`O(log n)`|
    |ExtractMin|`O(log n)`|
    |FindMin|`O(1)`|
    |Heapify|`O(n)`|
    |Delete|`O(log n)`|
