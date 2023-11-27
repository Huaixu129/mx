package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "", "用户名默认为空")
	//&user用于接受-u后面的指定参数，“u”即为-u，如果无u则返回默认值value，最后一个参数为解释说明
	flag.StringVar(&pwd, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号默认为3306")

	//一个非常重要的操作转换，必须有
	flag.Parse()
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n", user, pwd, host, port)
}
