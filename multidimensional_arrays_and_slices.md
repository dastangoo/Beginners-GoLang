```go
0 <= len(s) <= cap(s)

var identifier []type // no length is needed

var slice1 []type = arr1[start:end]

arr1[0:len(arr1)]
slice1 = &arr1 

slice1 = slice1[:len(slice1)-1]
s := [3]int{1, 2, 3}[:]
s := [...]int{1, 2, 3}[:]

s == s[:i] + s[i:]
len(s) <= cap(s)

var x = []int{2, 3, 5, 7, 11}
```
