
目标： 实现错误处理机制

### panic

Go 语言中，比较常见的错误处理方法是返回 error，由调用者决定后续如何处理。但是如果是无法恢复的错误，可以手动触发 panic，当然如果在程序运行过程中出现了类似于数组越界的错误，panic 也会被触发。panic 会中止当前执行的程序，退出。

主动触发的例子：
```go
func main() {
	fmt.Println("before panic")
	panic("crash")
	fmt.Println("after panic")
}
```

### defer 

panic 会导致程序被中止，但是在退出前，会先处理完当前协程上已经defer 的任务，执行完成后再退出。效果类似于 Python 语言的 try...Exception。

```go
func main() {
	defer func() {
		fmt.Println("defer func")
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
}
```
可以 defer 多个任务，在同一个函数中 defer 多个任务，会逆序执行。即先执行最后 defer 的任务。

在这里，defer 的任务执行完成之后，panic 还会继续被抛出，导致程序非正常结束。

### recover
Go 语言还提供了 recover 函数，可以避免因为 panic 发生而导致整个程序终止，recover 函数只在 defer 中生效。

```go
func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")
}

func main() {
	test_recover()
	fmt.Println("after recover")
}
```

我们可以看到，recover 捕获了 panic，程序正常结束。test_recover() 中的 after panic 没有打印，这是正确的，当 panic 被触发时，控制权就被交给了 defer 。就像在 java 中，try代码块中发生了异常，控制权交给了 catch，接下来执行 catch 代码块中的代码。而在 main() 中打印了 after recover，说明程序已经恢复正常，继续往下执行直到结束。

