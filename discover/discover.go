package discover

import (
	"log"
	"net/http"
	"text/template"

	"github.com/jadbin/cherry"
	"github.com/jadbin/i/blog"
	"github.com/jadbin/i/config"
)

type DiscoverB struct {
	resRoot    string
	seriesList []*Sticker
	tagList    []*Sticker
}

func (this *DiscoverB) Init(resRoot string) {
	this.resRoot = resRoot
	metalist := blog.ReadMetalist(resRoot + blog.METALIST)
	this.seriesList = GetSeriesList(metalist)
	this.tagList = GetTagList(metalist)
}

func (this *DiscoverB) Handle(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(this.resRoot+"/discover/tmpl/discover.html", this.resRoot+"/tmpl/header.html", this.resRoot+"/tmpl/footer.html")
	if err != nil {
		log.Println(err)
		cherry.HttpErr(w, r, 500)
		return
	}
	data := this.getArchiveListData()
	data["PageTitle"] = "Discover - wangyb.net"
	data["PageDescription"] = "Discover - wangyb.net"
	data["JS"] = config.CDN_JS
	data["CSS"] = config.CDN_CSS
	if err := tpl.Execute(w, data); err != nil {
		log.Println(err)
		cherry.HttpErr(w, r, 500)
		return
	}
}

func (this *DiscoverB) getArchiveListData() map[string]interface{} {
	data := make(map[string]interface{})
	data["SeriesList"] = this.seriesList
	data["TagList"] = this.tagList
	return data
}
