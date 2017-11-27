# Sync Packages
```go
import "sync"
type Info struct {
  mu sync.Mutex
  ...other fields
  Str string
}


func Update(info * Info)  {
  info.mu.Lock()
  info.Str = // new value
  info.mu.Unlock()
}
```