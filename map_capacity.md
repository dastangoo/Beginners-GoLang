# Map Capacity
```go
make(map[keytype]valuetype, cap)

map2 := make(map[string]float32, 100)

noteFrequency := map[string]float32{
  "C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
  "G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440
}

mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)

_, ok := map1[key1]; ok {
  ...
}

delete(map1, key1)
```