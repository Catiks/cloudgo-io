package main

import (
	"os"
	//服务路径,放在同一个人文件夹，相对路径
	"./service"

	flag "github.com/spf13/pflag"
)

const (
	// 端口号 : 8080
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
