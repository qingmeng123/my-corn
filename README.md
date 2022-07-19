使用示例
```
算法：循环链表
```

```
package main

import (
   "fmt"
   "my-corn/corn"
   "time"
)


func task() {
   fmt.Println("hello")
}


func main() {
   scheduler:=corn.NewScheduler()
   scheduler.Every(2).Second().Do(task)
   go scheduler.Start()
   time.Sleep(5*time.Second)
   scheduler.Pause()
   fmt.Println("暂停5s")
   time.Sleep(5*time.Second)
   fmt.Println("继续工作")
   scheduler.On()
   time.Sleep(2*time.Second)
   //scheduler.Clear()
   scheduler.Remove("main.task")
   time.Sleep(5*time.Second)
}
```
