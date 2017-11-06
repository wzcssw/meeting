package lib

import "flag"

// Port 项目启动端口
var Port string

// Flag Flag
func Flag() {
	flag.StringVar(&Port, "p", "3000", "http listen port")
	flag.StringVar(&ENV, "e", "develop", "enviorment")
	flag.Parse()
}
