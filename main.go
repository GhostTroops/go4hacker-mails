package main

import (
	"fmt"
	"github.com/hktalent/gopop3/pkg"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("How use?\n./popmails lists.txt xxxUser 50 pop3.xxx.com\n")
	//os.Args = []string{"", "lists.txt", "voraporn", "50", "mail.lionairthai.com"}
	Wg := sync.WaitGroup{}
	if data, err := ioutil.ReadFile(os.Args[1]); nil != err {
		log.Printf("read %s is error: %v\n", os.Args[1], err)
		return
	} else {
		a := strings.Split(string(data), "\n")
		for _, x := range a {
			j := strings.Split(x, "\t")
			if 2 <= len(j) {
				x := pkg.GetPopMail(j[0], j[1], os.Args[4], &Wg)
				Wg.Add(1)
				go x.PopAllMails()
			}
		}
	}
	Wg.Wait()
}
