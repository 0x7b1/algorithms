
# Asymptotic Notation
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
