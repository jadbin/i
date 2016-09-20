package blog

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/jadbin/cherry"
	"github.com/jadbin/i/config"
)

type SeriesB struct {
	resRoot string
	metaMap map[string]([]*Meta)
}

func (this *SeriesB) Init(resRoot string) {
	this.resRoot = resRoot
	metalist := ReadMetalist(resRoot + METALIST)
	this.metaMap = make(map[string]([]*Meta))
	for _, meta := range metalist {
		if meta.Series != "" {
			this.metaMap[meta.Series] = append(this.metaMap[meta.Series], meta)
		}
	}
}

func (this *SeriesB) Handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	seriesName := r.Form.Get(":series-name")
	pageId, err := this.getPageId(r.Form.Get(":page-id"))
	if err != nil {
		cherry.HttpErr(w, r, 400)
		return
	}
	pageId--
	begin := pageId * ARCHIVE_NUM_PER_PAGE
	end := begin + ARCHIVE_NUM_PER_PAGE
	n := len(this.metaMap[seriesName])
	if end > n {
		end = n
	}
	if begin < 0 || begin >= end {
		cherry.HttpErr(w, r, 404)
		return
	}
	tpl, err := template.ParseFiles(this.resRoot+"/blog/tmpl/page.html", this.resRoot+"/tmpl/header.html", this.resRoot+"/tmpl/footer.html")
	if err != nil {
		log.Println(err)
		cherry.HttpErr(w, r, 500)
		return
	}
	data := this.getArchiveListData(this.metaMap[seriesName], begin, end)
	pageCount := 0
	if n > 0 {
		pageCount = (n-1)/ARCHIVE_NUM_PER_PAGE + 1
	}
	data["PageNumber"] = pageId + 1
	data["PreviousPageNumber"] = pageId
	data["NextPageNumber"] = pageId + 2
	data["PageCount"] = pageCount
	data["PreviousPageHref"] = fmt.Sprintf("/blog/series/%s/page/%d", seriesName, pageId)
	data["NextPageHref"] = fmt.Sprintf("/blog/series/%s/page/%d", seriesName, pageId+2)
	data["PageTitle"] = fmt.Sprintf("%s - jadbin.com", seriesName)
	data["PageDescription"] = fmt.Sprintf("%s - jadbin.com", seriesName)
	data["JS"] = config.CDN_JS
	data["CSS"] = config.CDN_CSS
	if err := tpl.Execute(w, data); err != nil {
		log.Println(err)
		cherry.HttpErr(w, r, 500)
		return
	}
}

func (this *SeriesB) getPageId(idStr string) (int, error) {
	if idStr == "" {
		idStr = "1"
	}
	return strconv.Atoi(idStr)
}

func (this *SeriesB) getArchiveListData(list []*Meta, begin int, end int) map[string]interface{} {
	data := make(map[string]interface{})
	archiveList := [](map[string]interface{}){}
	for id := begin; id < end; id++ {
		archiveList = append(archiveList, this.getArchiveData(list[id]))
	}
	data["ArchiveList"] = archiveList
	return data
}

func (this *SeriesB) getArchiveData(meta *Meta) map[string]interface{} {
	abstract := ReadAbstract(this.resRoot + GetArchivePath(meta))
	data := make(map[string]interface{})
	data["ArchiveId"] = meta.Id
	data["ArchiveTitle"] = meta.Title
	data["ArchiveDate"] = meta.Date
	data["ArchiveTags"] = meta.Tags
	data["ArchiveSeries"] = meta.Series
	data["ArchiveAbstract"] = abstract
	return data
}
