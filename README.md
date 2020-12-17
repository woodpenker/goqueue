
This is an implement of FIFO queue for golang. Simple queue is the simplest implement using slice. I use a cycle slice to implement the queue. Below is the benchmark to compare with container/list, simple queue and channel implement. This implement may reduce the memory allocate cost and is fast then container/list which implement queue using linked list.

BTW: This implementation is not thread safe.

这个库实现了先入先出Queue队列, 使用了循环的切片数组来实现, 以此减少simple queue实现和使用链表实现中内存分配的开销.以下是其对比基准测试的结果.

BTW: 这个实现不是线程并发安全的.

Benchmark results:

```shell
Running tool: /home/wudy/bin/go test -benchmem -run=^$ -coverprofile=/tmp/vscode-go105nsY/go-code-cover -bench . test/goqueue

goos: linux
goarch: amd64
pkg: test/goqueue
BenchmarkPushQueue-8               	   30787	     37580 ns/op	   22608 B/op	     770 allocs/op
BenchmarkPushSimpleQueue-8         	   27054	     52191 ns/op	   38928 B/op	     780 allocs/op
BenchmarkPushList-8                	   10000	    114365 ns/op	   55344 B/op	    1793 allocs/op
BenchmarkPushChannel-8             	   14604	     85135 ns/op	   22624 B/op	     770 allocs/op
BenchmarkRandomQueue-8             	   24532	     49889 ns/op	   22248 B/op	     725 allocs/op
BenchmarkRandomSimpleQueue-8       	   19254	     72182 ns/op	   36968 B/op	     748 allocs/op
BenchmarkRandomList-8              	    8326	    125716 ns/op	   45432 B/op	    1549 allocs/op
BenchmarkGrowShrinkQueue-8         	   17838	     63756 ns/op	   40890 B/op	    1020 allocs/op
BenchmarkGrowShrinkSimpleQueue-8   	   19856	     57777 ns/op	   40870 B/op	    1022 allocs/op
BenchmarkGrowShrinkList-8          	   10000	    142118 ns/op	   57182 B/op	    2022 allocs/op
PASS
coverage: 57.9% of statements
ok  	test/goqueue	17.345s

╰─$ cat /proc/cpuinfo|grep "model name"
model name      : Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz

╰─$ uname -a
Linux 5.9.13-200.fc33.x86_64 #1 SMP Tue Dec 8 15:42:52 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux

```

Example:

```go
import(
    queue "github.com/woodpenker/goqueue"
)
func main() {
    q := queue.NewQueue(10)
    q.Push(12345)
    q.Push("hello")
    v1 := q.Pop()
    fmt.Println(v1) // 12345
    v2 := q.Top()
    fmt.Println(v2) // hello
}
```