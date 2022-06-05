package main

import (
	"flag"
	"log"

	accountcredential "github.com/yinyin/go-accountcredential"
)

func main() {
	var userName, groupName string
	flag.StringVar(&userName, "user", "", "User name of account.")
	flag.StringVar(&groupName, "group", "", "Group name of account.")
	flag.Parse()
	log.Printf("Lookup user=[%s], group=[%s].", userName, groupName)
	c, err := accountcredential.Lookup(userName, groupName)
	if nil != err {
		log.Fatalf("cannot lookup account credential: user=[%s], group=[%s], err=%v", userName, groupName, err)
		return
	}
	if c == nil {
		log.Print("empty result credential.")
	} else {
		log.Printf("result UID=%d, GID=%d.", c.Uid, c.Gid)
	}
}
