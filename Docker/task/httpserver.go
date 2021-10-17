package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"os"
)

func httplog(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//后台日志打印
	log.Printf("===================Http Logs:============\n")
	log.Printf("Path: %s\n", r.URL.Path)
	log.Printf("RemoteAddr: %s\n", r.RemoteAddr)
	log.Printf("StatusCode: %d\n", http.StatusOK)
	//获取系统环境变量VERSION
	version := os.Getenv("VERSION")
	if version != "" {
		//如果变量存在，添加进请求头
		r.Header.Add("VERSION", version)
		log.Println("VERSION: ",version)
	} else {
		log.Printf("VERSION not set!\n")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	httplog(w, r)
	//打印请求头内容
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	//打印访问路径
	io.WriteString(w, "===================Request URL Path:============\n")
	fmt.Fprintf(w, "Path = %q\n", r.URL.Path)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	httplog(w, r)
	io.WriteString(w, "200\n")
}

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", handler)
	//当访问路径为healthz时，执行healthz的func
	http.HandleFunc("/healthz", healthz)
	//监听80端口,并启动服务器，处理器参数为nil
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
