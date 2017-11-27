# Appending and Copying Slices 2
len([]int(s))

utf8.RunCountInStrings(s)


var b []byte
var s string
b = append(b, s...)


str[i] = 'D'
s := "hello"
c := []bytes(s)
c[0] = 'c'
s2 := string(c)


0 if a==b, -1 if a < b, 1 if a > b

func Compare(a, b []byte) int {
  for i := 0, i < len(a) && i < len(b), i++ {
    switch {
      case a[i] > b[i]:
        return 1
      case a[i] < b[i]:
        return -1
      
    }
  }
}