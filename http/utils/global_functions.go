package utils

import(
	"flag"
)

func GetAddrFromFlag()string{
	var addr string
	flag.StringVar(&addr,"addr","0.0.0.0:8888","请输入监听端口")
	flag.Parse()
	return addr
}