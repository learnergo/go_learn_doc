最基本的场景

##### 方式一 使用http.Newrequest

先生成http.client -> 再生成 http.request -> 之后提交请求：client.Do(request) -> 处理返回结果。
每一步的过程都可以设置一些具体的参数，下面是一个最朴素最基本的例子：

```
//生成client 参数为默认
	client := &http.Client{}
	
	//生成要访问的url
	url := "http://www.baidu.com"
	    
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		panic(err)
	}
	
	//处理返回结果
	response, _ := client.Do(reqest)
   
   //将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
   
   //返回的状态码
	status := response.StatusCode

	fmt.Println(status)
```

	
	
##### 方式二 先生成client，之后用client.get/post..
client结构自己也有一些发送api的方法，比如client.get,client.post,client.postform..等等。基本上涵盖了主要的http请求的类型，通常不进行什么特殊的配置的话，这样就可以了，其实client的get或者post方法，也是对http.Newerequest方法的封装，里面还额外添加了req.Header.Set("Content-Type", bodyType)一般用的话，也是ok的

##### 方式三 http. Get/Post..

具体实现的时候，还是采用的先前提到的模式，先生成一个默认的client，之后调用http.Newrequest方法。