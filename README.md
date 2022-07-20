使用示例

```

package main

import (
	"fmt"
	"my-corn/corn"
	"time"
)

func task1() {
	fmt.Println("hello")
}

func task2(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	scheduler := corn.NewScheduler()
	scheduler.Every(2).Second().Do(task1)
	scheduler.Every(3).Second().Do(task2, 1, "first")
	scheduler.Every(3).Days().Do(task2, 1, "first")
	go scheduler.Start()

	time.Sleep(5 * time.Second)
	scheduler.Pause()
	fmt.Println("暂停5s")
	time.Sleep(5 * time.Second)
	fmt.Println("继续工作")
	scheduler.On()
	time.Sleep(2 * time.Second)
	//scheduler.Clear()
	scheduler.Remove(task2)
	time.Sleep(5 * time.Second)

}
```