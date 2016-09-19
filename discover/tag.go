package discover

import (
	"sort"

	"github.com/jadbin/i/blog"
)

func GetTagList(metalist []*blog.Meta) []*Sticker {
	m := make(map[string]int)
	for _, meta := range metalist {
		for _, tag := range meta.Tags {
			m[tag] = m[tag] + 1
		}
	}
	list := []*Sticker{}
	for name, count := range m {
		list = append(list, &Sticker{Name: name, Count: count})
	}
	sort.Sort(StickerList(list))
	return list
}
