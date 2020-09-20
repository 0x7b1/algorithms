# Sorting

There are three practical reasons why study sorting algorithms is important

- Analyzing sorting algorithms is a good way to compare performance between algorithms
- Similar techniques are used to address other problems
- Sorting algorithms is usually a starting point to solve other problems



## Table of Contents

- Insertion Sort
- Merge Sort
- Quick Sort
- Bucket Sort
- Radix Sort



### Insertion Sort

* ### Algorithm [View](sorting/insertion.go)
  * From i to 0..n, insert a[i] to its correct position to the left (0..i)
* ### Complexity

| Time | Space   |
| ---- | ------- |
|      | `O(n²)` |

### Merge Sort

* ### Algorithm [View](sorting/mergesort.go)
  * Splits a collection into two halves
  * Sort the two halves (recursive call)
  * Merge them together to form one sorted collection
* ### Complexity

| Time         | Space        |
| ------------ | ------------ |
| `Θ(n log n)` | `Θ(n log n)` |

### Quick Sort

* ### Algorithm [View](sorting/quicksort.go)
  * Choose a pivot element `p`
  * Partition array `A` around `p`
  * recursively sort first part of `A`
  * recursively sort second part of `A`
* ### Complexity

| Time         | Space        |
| ------------ | ------------ |
| `Θ(n log n)` | `Θ(n log n)` |



- We can achieve a sorting time better than `O(nlogn)` by leveraging the shape of the input data. With this we can enable techniques beyond comparison, such as bucketing.