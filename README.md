# Algorithms & Data Structures

## I. Time Complexity

### Asymptotic Notation
Idea: Suppress constant factors and lower-order terms

|Big O - O(…)|Big Omega - Ω(…)|Big Theta - Θ(…)|
|----------|-------------|------|
|<ul><li>Upper bound</li><li>T(n) is O(f(n)) if f(n) grows at least as fast as T(n) as n gets large</li></ul> | <ul><li>Lower bound</li>T(n) is Ω(f(n)) if f(n) grows at most as fast as T(n) as n gets large</li></ul> | <ul><li>Lower and Upper bound</li><li>T(n) = O(f(n)) AND T(n) = Ω(f(n))</li></ul> |
|![](images/big_o.png)|![](images/big_omega.png)|![](images/big_theta.png)|

### Master Theorem
- **Example** : Binary Search

Recurses on either the left half of the input array or the right half. One recursive call, a=1
The recursive call is on half of the input array, b=2
Outside the recursive call, a single comparison between the middle element and the searched one. O(1); d=0
Since a=1=20=bd
O(ndlog n) = O(log n)

## II. Sorting
### Insertion Sort
#### Algorithm [View](sorting/mergesort.go)
- From i to 0..n, insert a[i] to its correct position to the left (0..i)

#### Complexity
| Time | Space |
|---|---|
| O(n²) | O(1) |

### Merge Sort
#### Algorithm
- Splits a collection into two halves
- Sort the two halves (recursive call)
- Merge them together to form one sorted collection

#### Complexity
| Time | Space |
|---|---|
| O(n²) | O(1) |
