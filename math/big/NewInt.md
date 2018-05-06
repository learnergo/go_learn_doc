这一节主要介绍big.NewInt源码实现

1、使用举例

```
bigNum := big.NewInt(58)
```

2、进入NewInt 实现


```
func NewInt(x int64) *Int {
	return new(Int).SetInt64(x)
}
```

这里有两个新“事物”

```
- Int
- SetInt64
```
3、Int定义

此Int是big int定义，并非我们常用int类型


```
type Int struct {
	neg bool // 是否负数
	abs nat  // 绝对值
}
```

再看nat定义

```
type nat []Word
```
再看Word定义

```
type Word uint
```

看到这里也许明白了golang实现big int的方式

```
- neg保存正负判断
- abs（[]uint） 以uint字节保存绝对值
```

4、SetInt64内部实现


```
func (z *Int) SetInt64(x int64) *Int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	z.abs = z.abs.setUint64(uint64(x))
	z.neg = neg
	return z
}
```

这段比较简单，主要看abs赋值这段：


```
z.abs = z.abs.setUint64(uint64(x))
```
5、setUint64 内部实现


```
func (z nat) setUint64(x uint64) nat {
	// single-word value
	if w := Word(x); uint64(w) == x {
		return z.setWord(w)
	}
	// 2-word value
	z = z.make(2)
	z[1] = Word(x >> 32)
	z[0] = Word(x)
	return z
}
```

以uint64做判断，如果传入值小于uint32用一个字节保存，反之用两个字节保存




