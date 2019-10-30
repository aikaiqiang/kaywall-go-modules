package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello，this is first go web program!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 解析form
		r.ParseForm()
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		// 自动解析: Request本身也提供了FormValue()函数来获取用户提交的参数
		// r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串
		fmt.Println("username = ", r.FormValue("username"))
		fmt.Println("password = ", r.FormValue("password"))

		// 验证表单的输入
		if len(r.Form["username"][0]) == 0 {
			//为空的处理
		}

		// 1- 数字验证
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			//数字转化出错了，那么可能就不是数字
		}

		//接下来就可以判断这个数字的大小范围了
		if getint > 100 {
			//太大了
		}

		// 正则表达式
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			return
		}

		// 2- 中文
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {
			return
		}

		// 3- 英文
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
			return
		}

		// 4- 电子邮件地址
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}

		// 5- 手机号码
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			return
		}

		// 6- 下拉菜单
		// 7- 单选按钮
		// 8- 复选框
		// 9- 日期和时间
		// 10- 身份证号码

		// Go 里面是怎么做这个有效防护的，Go 的 html/template（text/template）里面带有下面几个函数可以帮你转义
		// func HTMLEscape(w io.Writer, b []byte)  // 把 b 进行转义之后写到w
		// func HTMLEscapeString(s string) string  // 转义 s 之后返回结果字符串
		// func HTMLEscaper(args ...interface{}) string  // 支持多个参数一起转义，返回结果字符串

		t, _ := template.ParseFiles("template/home.gtpl")
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	http.HandleFunc("/upload", upload)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("template/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
