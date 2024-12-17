```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var data []int
        var mu sync.Mutex

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock()
                        data = append(data, i)
                        mu.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println(data)
}
```