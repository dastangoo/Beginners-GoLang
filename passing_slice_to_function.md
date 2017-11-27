```go
func sum(a []int) int  {
  s := 0
  for i := 0; i < len(a); i++ {
    s += a[i]
  }
  return s
}

func main()  {
  var arr = [5]int{0, 1, 2, 3, 4}
  sum(arr[:])
}

var slice1 []type = make([]type, len)
slice1 := make([]type, len)

s2 := make([]int, 10)
cap(s2) == len(s2) == 10


slice1 := make([]type, len, cap)
func make([]T, len, cap)

make([]int, 50, 100)
new([100]int)[0:50]


var buffer bytes.Buffer

var r * bytes.Buffer = new(bytes.Buffer)

func NewBuffer(buf []bytes) * Buffer 

var buffer bytes.Buffer
for {
  if s, ok := getNextString(); ok {
    buffer.WriteString(s)
  } else {
    break
  }
}
```