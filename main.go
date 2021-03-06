package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	netp := flag.Int("netp", 5000, "logstash connection port")
	flag.Parse()
	fmt.Println(*netp)
	log := logrus.New()
	uri := fmt.Sprintf("%s%d", "localhost:", *netp)
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"appid": "appid"}))
	log.Hooks.Add(hook)
	// 필드도 같이
	entry := log.WithFields(logrus.Fields{
		"method":   "get or post",
		"uid":      "USERID",
		"uip":      "127.0.0.1",
		"protocol": "http",
	})
	logSwitch := true

	for logSwitch {
		logmsg := bufio.NewReader(os.Stdin)
		line, err := logmsg.ReadString('\n')
		if err != nil {
			log.Fatal("err input")
		}
		entry.Error(line)
	}
}
