package pkg

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"mime"
	"regexp"
	"strings"
	"sync"
)

type PopMail struct {
	User   string
	Pswd   string
	Server string
	Wg     *sync.WaitGroup
}

func GetStr(r *io.Reader) string {
	buf := new(strings.Builder)
	n, err := io.Copy(buf, *r)
	if err != nil || 0 == n {
		//log.Println("GetStr ", err)
		return ""
	}
	return buf.String()
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

var rg01 = regexp.MustCompile(`=\?UTF-8\?B\?(.*?)\?=`)

func DecodeBase64(s string) string {
	if s2, err := base64.StdEncoding.DecodeString(s); nil == err {
		return string(s2)
	} else {
		log.Println(err)
	}
	return s
}

func DecodeTitle(s string) string {
	a := rg01.FindAllStringSubmatch(s, -1)
	var a1 []string
	for i, _ := range a {
		if 2 == len(a[i]) {
			a1 = append(a1, DecodeBase64(a[i][1]))
		}
	}
	//log.Println(a)
	return strings.Join(a1, " ")
}
func GetPopMail(u, p, s string, Wg *sync.WaitGroup) *PopMail {
	//os.MkdirAll(s, os.ModePerm)
	if -1 == strings.Index(u, "@") {
		a := strings.Split(s, ".")[1:]
		u = u + "@" + strings.Join(a, ".")
	}

	return &PopMail{User: u, Pswd: p, Server: s, Wg: Wg}
}

func (r *PopMail) WriteFile(id int, args ...string) {
	s := fmt.Sprintf("%d %s", id, strings.Join(args, "\n"))
	szFileName := fmt.Sprintf("%s/%d.txt", r.Server, hash(s))
	SaveFile(szFileName, []byte(s))
}
func (r *PopMail) PopAllMails() {
	defer r.Wg.Done()
	p := New(Opt{
		Host:       r.Server,
		Port:       995,
		TLSEnabled: true,
		//Port:       110,
		//TLSEnabled: false,
	})

	// Create a new connection. POP3 connections are stateful and should end
	// with a Quit() once the opreations are done.
	c, err := p.NewConn()
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer c.Quit()

	// Authenticate.
	if err := c.Auth(r.User, r.Pswd); err != nil {
		log.Printf("%v", err)
		return
	}
	// Print the total number of messages and their size.
	count, size, _ := c.Stat()
	fmt.Println("total messages=", count, "size=", size)

	// Pull the list of all message IDs and their sizes.
	msgs, _ := c.List(0)
	for _, m := range msgs {
		fmt.Sprintln("id=", m.ID, "size=", m.Size)
	}

	// Pull all messages on the server. Message IDs go from 1 to N.
	for id := 1; id <= count; id++ {
		m, _ := c.Retr(id)
		//s1 := m.Header.Get("subject")
		dec := new(mime.WordDecoder)
		from, _ := dec.DecodeHeader(m.Header.Get("From"))
		to, _ := dec.DecodeHeader(m.Header.Get("To"))
		subject, _ := dec.DecodeHeader(m.Header.Get("Subject"))

		mediaType, params, err := mime.ParseMediaType(m.Header.Get("Content-Type"))
		if err != nil {
			log.Printf("%v", err)
			return
		}
		if strings.HasPrefix(mediaType, "multipart/") {
			r.ParsePart(m.Body, params["boundary"], 1, r.Server)
		}
		// DecodeTitle(s1)
		r.WriteFile(id, from, to, subject)
	}

}
