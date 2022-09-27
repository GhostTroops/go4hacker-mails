package pkg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r02 = regexp.MustCompile(`^-+=`)

func SaveFile(s string, data []byte) {
	// lists.txt 20
	//if 3 < len(os.Args) {
	//	ioutil.WriteFile(s, data, 0644)
	//}
}

// BuildFileName builds a file name for a MIME part, using information extracted from
// the part itself, as well as a radix and an index given as parameters.
func BuildFileName(part *multipart.Part, radix string, index int) (filename string) {
	// 1st try to get the true file name if there is one in Content-Disposition
	filename = part.FileName()
	if strings.HasSuffix(filename, "=?") {
		filename = DecodeTitle(filename)
	}
	if len(filename) > 0 {
		return
	}
	// If no defaut filename defined, try to build one of the following format :
	// "radix-index.ext" where extension is comuputed from the Content-Type of the part
	mediaType, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
	if err == nil {
		mime_type, e := mime.ExtensionsByType(mediaType)
		if e == nil {
			//if strings.HasSuffix(radix, "=?") {
			//	radix = DecodeTitle(radix)
			//} else {
			radix = r02.ReplaceAllString(radix, "")
			//}
			var s09 string
			if 0 == len(mime_type) {
				s09 = ""
			} else {
				s09 = mime_type[len(mime_type)-1]
			}
			return fmt.Sprintf("%s-%d%s", radix, index, s09)
		}
	}
	return

}

func DoFileName(s string) string {
	if strings.HasPrefix(s, "=?") {
		s = DecodeTitle(s)
	} else if strings.HasSuffix(s, "?=") {
		s = DecodeTitle("=?UTF-8?B?" + s)
	}
	return s
}

// 追加到文件中
func AppendFile(szFile, szOut string) {
	f, err := os.OpenFile(szFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(szOut + "\n"); err != nil {
		log.Println(err)
	}
}

func SearchText(data []byte) {
	if 2 < len(os.Args) {
		s := string(data)
		s1 := os.Args[2]
		j := strings.Index(s, s1)
		if -1 < j {
			x01 := 200
			if 3 < len(os.Args) {
				if m, err := strconv.Atoi(os.Args[3]); nil == err {
					x01 = m
				}
			}
			n := j + len(s1) + x01
			if n > len(s) {
				n = len(s)
			}

			AppendFile("SMResults.txt", s[j:n]+"\n===================\n")
			fmt.Printf("found in: %s\n==========\n", s[j:n])
		}
	}
}

// WitePart decodes the data of MIME part and writes it to the file filename.
func WritePart(part *multipart.Part, filename string, path string) {
	// Read the data for this MIME part
	part_data, err := ioutil.ReadAll(part)
	if err != nil || 0 == len(part_data) {
		//log.Println("Error reading MIME part data -", err)
		return
	}
	content_transfer_encoding := strings.ToUpper(part.Header.Get("Content-Transfer-Encoding"))
	switch {
	case strings.Compare(content_transfer_encoding, "BASE64") == 0:
		decoded_content, err := base64.StdEncoding.DecodeString(string(part_data))
		if err != nil {
			//log.Println("Error decoding base64 -", err)
		} else {
			SearchText(decoded_content)
			SaveFile(path+"/"+DoFileName(filename), decoded_content)
		}

	case strings.Compare(content_transfer_encoding, "QUOTED-PRINTABLE") == 0:
		decoded_content, err := ioutil.ReadAll(quotedprintable.NewReader(bytes.NewReader(part_data)))
		if err != nil {
			//log.Println("Error decoding quoted-printable -", err)
		} else {
			SearchText(decoded_content)
			SaveFile(path+"/"+DoFileName(filename), decoded_content)
		}
	default:
		SearchText(part_data)
		SaveFile(path+"/"+DoFileName(filename), part_data)
	}
}

// ParsePart parses the MIME part from mime_data, each part being separated by
// boundary. If one of the part read is itself a multipart MIME part, the
// function calls itself to recursively parse all the parts. The parts read
// are decoded and written to separate files, named uppon their Content-Descrption
// (or boundary if no Content-Description available) with the appropriate
// file extension. Index is incremented at each recursive level and is used in
// building the filename where the part is written, as to ensure all filenames
// are distinct.
func (r *PopMail) ParsePart(mime_data io.Reader, boundary string, index int, path string) {
	// Instantiate a new io.Reader dedicated to MIME multipart parsing
	// using multipart.NewReader()
	reader := multipart.NewReader(mime_data, boundary)
	if reader == nil {
		return
	}
	//fmt.Println(strings.Repeat("  ", 2*(index-1)), ">>>>>>>>>>>>> ", boundary)
	// Go through each of the MIME part of the message Body with NextPart(),
	// and read the content of the MIME part with ioutil.ReadAll()
	for {
		new_part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		//for key, value := range new_part.Header {
		//	fmt.Printf("%s Key: (%+v) - %d Value: (%#v)\n", strings.Repeat("  ", 2*(index-1)), key, len(value), value)
		//}
		//fmt.Println(strings.Repeat("  ", 2*(index-1)), "------------")
		mediaType, params, err := mime.ParseMediaType(new_part.Header.Get("Content-Type"))
		if err == nil && strings.HasPrefix(mediaType, "multipart/") {
			r.ParsePart(new_part, params["boundary"], index+1, path)
		} else {
			filename := BuildFileName(new_part, boundary, 1)
			WritePart(new_part, filename, path)
		}
	}
}
