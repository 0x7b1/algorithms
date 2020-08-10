# Sorting

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
