package discover

type Sticker struct {
	Name  string
	Count int
}

type StickerList []*Sticker

func (this StickerList) Len() int {
	return len(this)
}
func (this StickerList) Less(i int, j int) bool {
	return this[i].Count > this[j].Count
}
func (this StickerList) Swap(i int, j int) {
	this[i], this[j] = this[j], this[i]
}
