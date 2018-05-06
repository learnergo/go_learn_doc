4代表读权限，2代表写权限，1代表执行权限

7=4+2+1,表示拥有可读可写可执行权限
5=4+1,表示拥有可读可执行权限，但是没有写权限
4 代表拥有可读权限
0 代表没有任何权限
以此类推


介绍四种写入文件的方式：

##### 方式一：使用 io.WriteString 写入文件
函数原型

```
func WriteString(w Writer, s string) (n int, err error) {}
```


例子：


```
var wireteString = "测试数据"
if 存在(filename) { //如果文件存在
	f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
} else {
	f, err1 = os.Create(filename) //创建文件
}
defer f.Close()
n, err := io.WriteString(f, wireteString) //写入文件(字符串)
```

	
	
	
	
	
##### 第二种方式: 使用 ioutil.WriteFile 写入文件
函数原型

```
func WriteFile(filename string, data []byte, perm os.FileMode) error {}
```

例子：
 
```
err2 := ioutil.WriteFile("./output2.txt", byteData, 0666)
```




##### 第三种方式:  使用 File (Write或WriteString) 写入文件
函数原型



```
func (f *File) Write(b []byte) (n int, err error) {}

func (f *File) WriteString(s string) (n int, err error) {}
```

例子：

```
f, err3 := os.Create("./output3.txt") //创建文件
defer f.Close()
n2, err3 := f.Write(byteData) //写入文件(字节数组)

n3, err3 := f.WriteString(strData) //写入文件(字节数组)

f.Sync()
```

	
	
	
##### 第四种方式:  使用 bufio.NewWriter 写入文件
例子:

```
w := bufio.NewWriter(f) //创建新的 Writer 对象
n4, err3 := w.WriteString("bufferedn")
w.Flush()
f.Close()
```
