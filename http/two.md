对每个步骤进行细节性的说明
生成client时候的参数配置

最常见的一个参数是使用https的方式发送信息时候client端的设置。如果生成client的时候，什么信息都不添加，就会使用默认的值。具体的信息包括：

	
```
Transport RoundTripper

CheckRedirect func(req *Request, via []*Request) error

Jar CookieJar

Timeout time.Duration
```

第一个参数是一个RoundTripper接口，里面包含了一个RoundTrip函数，指定了一些http请求的基本机制。
http.Transport中涉及到的参数较多，要是不指定的话，就会使用默认的DefaultTransport参数，里面包含一些默认的请求时间以及proxy机制之类的。
具体的细节参数涉及到好多，有的都没有使用到过比如那些我握手时间之类的，目前使用到的最多的就是https的相关参数：TLSClientConfig，
这是一个*tls.Config类型，其中涉及到的参数还是有很多，一个基本的是用案例如下，仅仅是在配置中制定了rooca以及客户度端使用的证书。

通常发送https请求的时候，前面的参数可以使用如下方式进行处理：


```
    pool := x509.NewCertPool()
	caCertPath := "certs/cert_server/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("certs/cert_server/client.crt", "certs/cert_server/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}
	
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
```

	
	
	生成request时候的参数配置

生成request的时候，主要的是几个基本的参数。NewRequest函数有三个基本的参数，NewRequest(method, urlStr string, body io.Reader)第一个是请求的类型，GET, POST, PUT, etc.要设成大写的形式。第二个参数是请求要访问的url，第三个参数是请求的body中的内容，需要是一个io.Reader的类型。

注意io.Reader的接口中是一个Read方法，实现了Read方法的类型应该都可以作为io.Reader来返回,Read(p []byte) (n int, err error)函数具体的功能就是读入len(p)长度的内容到p中，返回读入的长度以及错误信息。

通常是采用strings.NewReader函数，将一个string类型转化为io.Reader类型，或者bytes.NewBuffer函数，将[]byte类型转化为io.Reader类型。

此外还可以给request的header中添加一些额外的信息，比如下面例子中添加了请求的body的类型以及token的信息。


```
reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
reqest.Header.Set("Authorization", "qwertyuiopasdfghjklzxcvbnm1234567890")
```


还有比如模拟表单提交，可以把提交的类型设置为url.Values类型再进行Encode：

```
// use map as struct
	var clusterinfo = url.Values{}
	//var clusterinfo = map[string]string{}
	clusterinfo.Add("userName", user)
	clusterinfo.Add("password", pw)
	clusterinfo.Add("cloudName", clustername)
	clusterinfo.Add("masterIp", masterip)
	clusterinfo.Add("cacrt", string(caCrt))

	data := clusterinfo.Encode()
	
	url := "https://10.10.105.124:8443/user/checkAndUpdate"
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
```

	
	最常见的一种情况是发送一个json文件过去，可以把Header的类型设置成为：


```
"Content-Type", "application/json; charset=utf-8"
```

其余的部分按照先前同样的方式进行设置发送提交就好。

request的类型的属性还是比较多的，慢慢整理。

生成的response结果的处理

一般在client构建好之后，要采用client.Do(request)方法提交client请求，之后会返回一个*Response类型。
response中的参数一般也比较多，我们需要的最多的通常是Body参数，一般通过body, _ := ioutil.ReadAll(resp.Body)会把body转化为[]byte类型返回过来, 之后再进行其他的处理。