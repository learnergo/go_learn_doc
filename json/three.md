## 如何解析一个未知JSON

我们知道interface{}可以用来存储任意数据类型的对象，这种数据结构正好用于存储解析的未知结构的json数据的结果。JSON包中采用map[string]interface{}和[]interface{}结构来存储任意的JSON对象和数组。Go类型和JSON类型的对应关系如下：


```
- bool 代表 JSON booleans,
- float64 代表 JSON numbers,
- string 代表 JSON strings,
- nil 代表 JSON null.
```
----------------------------------------------


现在我们假设有如下的JSON数据


```
b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
```

如果在我们不知道他的结构的情况下，我们把他解析到interface{}里面


```
var f interface{}
err := json.Unmarshal(b, &f)
这个时候f里面存储了一个map类型，他们的key是string，值存储在空的interface{}里

f = map[string]interface{}{
    "Name": "Wednesday",
    "Age":  6,
    "Parents": []interface{}{
        "Gomez",
        "Morticia",
    },
}
```

那么如何来访问这些数据呢？通过断言的方式：


```
m := f.(map[string]interface{})
```

通过断言之后，你就可以通过如下方式来访问里面的数据了


```
for k, v := range m {
    switch vv := v.(type) {
    case string:
        fmt.Println(k, "is string", vv)
    case int:
        fmt.Println(k, "is int", vv)
    case float64:
        fmt.Println(k,"is float64",vv)
    case []interface{}:
        fmt.Println(k, "is an array:")
        for i, u := range vv {
            fmt.Println(i, u)
        }
    default:
        fmt.Println(k, "is of a type I don't know how to handle")
    }
}
```

通过上面的示例可以看到，通过interface{}与type assert的配合，我们就可以解析未知结构的JSON数了。