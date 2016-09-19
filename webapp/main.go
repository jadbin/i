package main

import (
	"log"
	"net/http"

	"github.com/jadbin/cherry"
	"github.com/jadbin/i/blog"
	"github.com/jadbin/i/discover"
)

func main() {
	log.Println("Hello, World!")
	//go redirect_https()
	server := cherry.NewServer()
	server.RouteGet("/", &blog.BlogB{})
	server.RouteGet("/blog/page/:page-id", &blog.BlogB{})
	server.RouteGet("/blog/archive/:archive-id", &blog.ArchiveB{})
	server.RouteGet("/blog/tag/:tag-name", &blog.TagB{})
	server.RouteGet("/blog/tag/:tag-name/page/:page-id", &blog.TagB{})
	server.RouteGet("/blog/series/:series-name", &blog.SeriesB{})
	server.RouteGet("/blog/series/:series-name/page/:page-id", &blog.SeriesB{})
	server.RouteGet("/discover", &discover.DiscoverB{})
	server.WebRoot = "WebRoot"
	server.ResRoot = "ResRoot"
	server.Serve()
}

func redirect_https() {
	server := cherry.NewServer()
	server.RouteGet("/*", &RedirectHttpsB{})
	server.Serve()
}

type RedirectHttpsB struct {
}

func (this *RedirectHttpsB) Init(resRoot string) {
}

func (this *RedirectHttpsB) Handle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}
