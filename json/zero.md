##### json库分类


```
encoding/json
官方提供的标准json, 需要预定义struct，通过reflection和interface来完成工作, 性能低
```



```
go-simplejson, gabs, jason等
此类工具包依赖于encoding/json，提高了易用性. 不再需要预定义struct, 性能降低了许多.
```





```
easyjson, ffjson
舍弃了官方(encoding/json), 另辟蹊径.
不使用影响性能的reflection， 通过预先定义好的struct，进行命令行配置, 生成相关的解码/编码方法.easyjson甚至使用了unsafe的包, 来减少内存的损耗. 
虽然开发成本变高了,需要额外定义struct文件和编码/解码方法文件, 但性能极速提高.有不少大公司都用easyjson替换了官方(encoding/json), 比如东.
```



```
sonparser
喜欢简单和扩展性, 依赖于bytes. json就是字符串, 那么用bytes不是又美好又简约吗? 易用性不但好, 而且性能是极高的.
```



##### 各自的优缺点


```
encoding/json, 官方自带的, 文档最多, 易用性差, 性能差
```


```
go-simplejson, gabs, jason等衍生包, 简单且易用, 易于阅读, 便于维护, 但性能最差
```


```
easyjson, ffjson此类包, 适合固定结构的json, 易用性一般, 维护成本高, 性能特别好
```


```
jsonparser 适合动态和固定结构的json, 简单且易用, 维护成本低, 性能极好
```

以性能的高低排序: jsonparser > easyjson > encoding/json > go-simplejson, gabs, jason