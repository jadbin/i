package blog

import (
	"io/ioutil"
	"log"
)

func GetArchivePath(meta *Meta) string {
	if meta.Series == "" {
		return ARCHIVE_FOLDER + "/" + meta.Id + ".html"
	}
	return ARCHIVE_FOLDER + "/" + meta.Series + "/" + meta.Id + ".html"
}

func ReadContent(file string) string {
	text := "<p><code>- EOF -</code></p>"
	if bytes, err := ioutil.ReadFile(file); err != nil {
		log.Println(err)
	} else {
		text = string(bytes)
	}
	return text
}

func ReadAbstract(file string) string {
	text := "<p><code>- EOF -</code><p>"
	if bytes, err := ioutil.ReadFile(file); err != nil {
		log.Println(err)
	} else {
		n := len(bytes)
		p, s, cnt := 0, 0, 0
		tar := []byte{'<', '/', 'p', '>'}
		for ; p < n; p++ {
			if s >= 4 {
				cnt++
				s = 0
				if cnt >= PARA_NUM_IN_ABSTRACT {
					break
				}
			}
			if tar[s] == bytes[p] {
				s++
			} else {
				s = 0
			}
		}
		text = string(bytes[0:p])
	}
	return text
}
