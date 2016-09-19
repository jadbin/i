package blog

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Meta struct {
	Id     string
	Title  string
	Date   string
	Tags   []string
	Series string
	State  string
}

func NewMeta() *Meta {
	m := &Meta{}
	m.Tags = []string{}
	return m
}

type MetaReader struct {
	data []byte
	cur  int
	n    int
}

func NewMetaReader(data []byte) *MetaReader {
	r := &MetaReader{}
	r.data = data
	r.n = len(data)
	r.cur = 0
	return r
}

func (this *MetaReader) NextMeta() (*Meta, error) {
	this.moveBack()
	if this.cur >= this.n {
		return nil, io.EOF
	}
	meta := NewMeta()
	for {
		if this.cur >= this.n {
			break
		}
		if this.isEnd(this.data[this.cur]) {
			this.moveOver(this.cur)
			break
		}
		if this.cur >= this.n {
			return nil, io.EOF
		}
		key, err := this.nextKey()
		if err != nil {
			return nil, io.ErrUnexpectedEOF
		}
		value, err := this.nextValue()
		if err != nil {
			return nil, io.ErrUnexpectedEOF
		}
		switch key {
		case "id":
			meta.Id = value
		case "title":
			meta.Title = value
		case "date":
			meta.Date = value
		case "tags":
			meta.Tags = strings.Split(value, "|")
		case "series":
			meta.Series = value
		case "state":
			meta.State = value
		}
	}
	return meta, nil
}

func (this *MetaReader) nextKey() (string, error) {
	if this.cur >= this.n {
		return "", io.EOF
	}
	p := this.cur
	var s string
	for {
		if p >= this.n || this.isEnd(this.data[p]) {
			return "", io.ErrUnexpectedEOF
		}
		if this.data[p] == ':' {
			s = string(this.data[this.cur:p])
			this.cur = p + 1
			break
		}
		p++
	}
	return strings.Trim(s, " \t"), nil
}

func (this *MetaReader) nextValue() (string, error) {
	if this.cur >= this.n {
		return "", io.EOF
	}
	p := this.cur
	var s string
	for {
		if p >= this.n || this.isEnd(this.data[p]) {
			s = string(this.data[this.cur:p])
			this.moveOver(p)
			break
		}
		p++
	}
	return strings.Trim(s, " \t"), nil
}

func (this *MetaReader) moveBack() {
	for this.cur < this.n && this.isEmpty(this.data[this.cur]) {
		this.cur++
	}
}

func (this *MetaReader) moveOver(p int) {
	if p >= this.n {
		this.cur = p
		return
	}
	if this.data[p] == '\r' && p+1 < this.n && this.data[p+1] == '\n' {
		this.cur = p + 2
	} else {
		this.cur = p + 1
	}
}

func (this *MetaReader) isEmpty(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func (this *MetaReader) isEnd(b byte) bool {
	return b == '\n' || b == '\r'
}

func ReadMetalist(file string) []*Meta {
	metalist := []*Meta{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return metalist
	}
	r := NewMetaReader(bytes)
	for {
		meta, err := r.NextMeta()
		if err != nil {
			break
		}
		if meta.State == "" {
			metalist = append(metalist, meta)
		}
	}
	for i, j := 0, len(metalist)-1; i < j; i, j = i+1, j-1 {
		metalist[i], metalist[j] = metalist[j], metalist[i]
	}
	return metalist
}
