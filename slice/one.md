1、声明变量，go自动初始化为nil，长度：0；容量：0；地址：0x0
例如：
    
```
var s []int
	if s == nil {
		fmt.Println("I am nil")
	}
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println(len(s))
	fmt.Println(cap(s))
```

输出：
   
```
I am nil
    0x0
    0
    0
```

	
2、切片的追加
例如：
   
```
var s []int
	s = append(s, 1)
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println(len(s))
	fmt.Println(cap(s))
```

输出：
    
```
0xc04204e080
	1
	1
```

会自动分配地址，容量和大小自动加一

3、容量和大小
在上面基础再加入1：
  
```
var s []int
	s = append(s, 1)
	fmt.Printf("%p", s)
	fmt.Println()
	s = append(s, 1)
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println(len(s))
	fmt.Println(cap(s))
```

输出:
	
```
0xc04204e080
	0xc04204e0b0
	2
	2
```

我们发现，两次地址不一致，是因为超出容量会重新分配地址，所以我们最好指定容量和大小

4、指定大小的陷阱
例如:
	
```
var s []int = make([]int, 2)
	fmt.Println(s)
	s = append(s, 1)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
```

输出：

```
[0 0]
	[0 0 1]
	3
	4
```

有没有感觉奇怪，append 1后不是我们想象的[1 0]，而是[0 0 1]，容量也加了一倍。
原来make会默认初始值，append会在末尾加入元素，如果满了就会扩容（现在容量的2倍）。
所以可以看出容量参数的重要性：make([]int,0,10),初始大小赋值为0

5、扩容问题

前面提到扩容，会重新分配地址，这在共享数据时会出现问题
解决方案：指针
如：

```
func main() {
    var osa = make ([]string,0);
    sa:=&osa;
    for i:=0;i<10;i++{
        *sa=append(*sa,fmt.Sprintf("%v",i))
        fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n",osa,sa,sa);
    }
    fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n",osa,sa,sa);
   
}
```

sa就是osa这个切片的指针，我们共享切片数据和操作切片的时候都使用这个切片地址就ok了，
其本质上是：append操作亦然会在需要的时候构造新的切片，不过是将地址都保存到了sa中，
因此我们通过该指针始终可以访问到真正的数据。