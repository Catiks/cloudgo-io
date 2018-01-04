# golang实现的一个web小应用，展示静态文件服务、js请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力

[GitHub 项目地址](https://github.com/Catiks/cloudgo-io)

## 程序运行

> 找到main可执行文件，命令行中输入./main 即可启动服务器

## 依赖说明
>本程序可能对一下包有依赖
>	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
解决协议来问题，输入go get "——————"
其中"——————"替换为以上任意一个地址

##结果展示
>启动
catik@catik-Aspire-V3-471:~/文档/cloudgo-io$ go run main.go 
[negroni] listening on :8080
[negroni] 2018-01-05T01:55:53+08:00 | 200 | 	 344.785µs | localhost:8080 | GET / 

>post
[negroni] 2018-01-05T02:00:41+08:00 | 200 | 	 356.045µs | localhost:8080 | POST /

>模板输出后的页面
<html>
<head>
<title>Personal information</title>
</head>
<body style="text-align:center">
<h3>Person general infor</h3>
<hr>
<ul>
<li>Name: zhang:<p>  *
<li>Phone Number: 1419012394 <p>  *
<li>ID: 189128498998881<p>*
</ul>
<hr>
</body>
</html>
加*的为替换位置
>unkonwn的处理
[negroni] 2018-01-05T02:03:37+08:00 | 514 | 	 64.091µs | localhost:8080 | GET /unknown
>可以看出，错误代码为514
>更多的截图细节详见result文件夹
