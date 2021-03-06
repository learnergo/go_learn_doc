### panic 
-------------------------

什么是 panic？
在 Go 语言中，程序中一般是使用错误来处理异常情况。对于程序中出现的大部分异常情况，错误就已经够用了。

但在有些情况，当程序发生异常时，无法继续运行。在这种情况下，我们会使用 panic 来终止程序。
当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，
接着打印出堆栈跟踪（Stack Trace），最后程序终止。在编写一个示例程序后，我们就能很好地理解这个概念了。

在本教程里，我们还会接着讨论，当程序发生 panic 时，使用 recover 可以重新获得对该程序的控制。

可以认为 panic 和 recover 与其他语言中的 try-catch-finally 语句类似，只不过一般我们很少使用 panic 和 recover。
而当我们使用了 panic 和 recover 时，也会比 try-catch-finally 更加优雅，代码更加整洁。

----------------

什么时候应该使用 panic？
需要注意的是，你应该尽可能地使用错误，而不是使用 panic 和 recover。
只有当程序不能继续运行的时候，才应该使用 panic 和 recover 机制。

panic 有两个合理的用例。


```
- 发生了一个不能恢复的错误，此时程序不能继续运行。 一个例子就是web服务器无法绑定所要求的端口。在这种情况下，就应该使用 panic，因为如果不能绑定端口，啥也做不了。

- 发生了一个编程上的错误。假如我们有一个接收指针参数的方法，而其他人使用nil作为参数调用了它。在这种情况下，我们可以使用 panic，因为这是一个编程错误：用 nil 参数调用了一个只能接收合法指针的方法。
```

	

	
###  recover

-------------------------------

recover 是一个内建函数，用于重新获得 panic 协程的控制。

recover 函数的标签如下所示：


```
func recover() interface{}
```

只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，并且停止 panic 续发事件（Panicking Sequence），程序运行恢复正常。如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。

我们来修改一下程序，在发生 panic 之后，使用 recover 来恢复正常的运行。


```
func recoverName() {  
    if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName *string, lastName *string) {  
    defer recoverName()
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Elon"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}
```


该程序会输出：


```
recovered from  runtime error: last name cannot be nil  
returned normally from main  
deferred call in main
```




### panic，recover 和 Go 协程

--------------------------------------

只有在相同的 Go 协程中调用 recover 才管用。recover 不能恢复一个不同协程的 panic。我们用一个例子来理解这一点。


```
func recovery() {  
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}

func a() {  
    defer recovery()
    fmt.Println("Inside A")
    go b()
    time.Sleep(1 * time.Second)
}

func b() {  
    fmt.Println("Inside B")
    panic("oh! B panicked")
}

func main() {  
    a()
    fmt.Println("normally returned from main")
}
```

在上面的程序中，函数 b() 在第 23 行发生 panic。函数 a() 调用了一个延迟函数 recovery()，用于恢复 panic。在第 17 行，函数 b() 作为一个不同的协程来调用。下一行的 Sleep 只是保证 a() 在 b() 运行结束之后才退出。

你认为程序会输出什么？panic 能够恢复吗？答案是否定的，panic 并不会恢复。因为调用 recovery 的协程和 b() 中发生 panic 的协程并不相同，因此不可能恢复 panic。

运行该程序会输出：


```
Inside A  
Inside B  
panic: oh! B panicked

goroutine 5 [running]:  
main.b()  
    /tmp/sandbox388039916/main.go:23 +0x80
created by main.a  
    /tmp/sandbox388039916/main.go:17 +0xc0
```

从输出可以看出，panic 没有恢复。

如果函数 b() 在相同的协程里调用，panic 就可以恢复。

如果程序的第 17 行由 go b() 修改为 b()，就可以恢复 panic 了，因为 panic 发生在与 recover 相同的协程里。




### 运行时panic

-----------------------------------

运行时错误（如数组越界）也会导致 panic。这等价于调用了内置函数 panic，其参数由接口类型 runtime.Error 给出。runtime.Error 接口的定义如下：


```
type Error interface {  
    error
    // RuntimeError is a no-op function but
    // serves to distinguish types that are run time
    // errors from ordinary errors: a type is a
    // run time error if it has a RuntimeError method.
    RuntimeError()
}
```

而 runtime.Error 接口满足内建接口类型 error。

我们来编写一个示例，创建一个运行时 panic。


```
func a() {  
    n := []int{5, 7, 4}
    fmt.Println(n[3])
    fmt.Println("normally returned from a")
}
func main() {  
    a()
    fmt.Println("normally returned from main")
}
```

在 playground 上运行

在上面的程序中，第 9 行我们试图访问 n[3]，这是一个对切片的错误引用。该程序会发生 panic，输出如下：


```
panic: runtime error: index out of range

goroutine 1 [running]:  
main.a()  
    /tmp/sandbox780439659/main.go:9 +0x40
main.main()  
    /tmp/sandbox780439659/main.go:13 +0x20
```

你也许想知道，是否可以恢复一个运行时 panic？当然可以！我们来修改一下上面的代码，恢复这个 panic。


```
func r() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
    }
}

func a() {  
    defer r()
    n := []int{5, 7, 4}
    fmt.Println(n[3])
    fmt.Println("normally returned from a")
}

func main() {  
    a()
    fmt.Println("normally returned from main")
}
```


运行上面程序会输出：


```
Recovered runtime error: index out of range  
normally returned from main
```

从输出可以知道，我们已经恢复了这个 panic。


### 恢复后获得堆栈跟踪

-------------------------------------------

当我们恢复 panic 时，我们就释放了它的堆栈跟踪。实际上，在上述程序里，恢复 panic 之后，我们就失去了堆栈跟踪。

有办法可以打印出堆栈跟踪，就是使用 Debug 包中的 PrintStack 函数。


```
import (  
    "fmt"
    "runtime/debug"
)

func r() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
        debug.PrintStack()
    }
}
```
