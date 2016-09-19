package blog

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/jadbin/cherry"
	"github.com/jadbin/i/config"
)

type ArchiveB struct {
	resRoot string
	metaMap map[string]*Meta
}

func (this *ArchiveB) Init(resRoot string) {
	this.resRoot = resRoot
	metalist := ReadMetalist(resRoot + METALIST)
	this.metaMap = make(map[string]*Meta)
	for _, meta := range metalist {
		this.metaMap[meta.Id] = meta
	}
}

func (this *ArchiveB) Handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get(":archive-id")
	if this.metaMap[id] == nil {
		cherry.HttpErr(w, r, 404)
		return
	}
	tpl, err := template.ParseFiles(this.resRoot+"/blog/tmpl/archive.html", this.resRoot+"/tmpl/header.html", this.resRoot+"/tmpl/footer.html")
	if err != nil {
		log.Println(err)
		cherry.HttpErr(w, r, 500)
		return
	}
	meta := this.metaMap[id]
	data := this.getArchiveData(meta)
	data["PageTitle"] = fmt.Sprintf("%s - wangyb.net", meta.Title)
	data["PageDescription"] = fmt.Sprintf("%s - wangyb.net", meta.Title)
	data["JS"] = config.CDN_JS
	data["CSS"] = config.CDN_CSS
	if err := tpl.Execute(w, data); err != nil {
		log.Println(err)
		cherry.HttpErr(w, r, 500)
		return
	}
}

func (this *ArchiveB) getArchiveData(meta *Meta) map[string]interface{} {
	content := ReadContent(this.resRoot + GetArchivePath(meta))
	data := make(map[string]interface{})
	data["ArchiveId"] = meta.Id
	data["ArchiveTitle"] = meta.Title
	data["ArchiveDate"] = meta.Date
	data["ArchiveTags"] = meta.Tags
	data["ArchiveSeries"] = meta.Series
	data["ArchiveContent"] = content
	return data
}
