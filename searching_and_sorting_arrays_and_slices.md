# Searching and Sorting Arrays and Slices
```go
func SearchFloat64s(a []float64, x float64) int
func SearchStrings(a []string, x string) int

a = append(a, b..)

b = make([]T, len(a))
copy(b, a)

a = append(a[:i], a[i+1:]...)

a = append(a[:i], a[j:]...)

a = append(a, make([]T, j)...)

a = append(a[:i], append([]T{x}, ...)...)

a = append(a[:i], append(b, a[:i]...)...)

x, a = a[len(a) - 1], a[:len(a) - 1]
a = append(a, x)

var digitRegexp = regex.MustCompile("[0-9]+")

func FindDigits(filename string) []byte  {
  b, _ := ioutil.ReadFile(filename)
  return digitRegexp.Find(b)
}

func FindDigits(filename string) []byte {
  b, _ := ioutil.ReadFile(filename)
  b = digitRegexp.Find(b)
  c := make([]byte, len(b))
  copy(c, b)
  return c
}
```