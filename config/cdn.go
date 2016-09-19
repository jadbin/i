package config

var (
	CDN_CSS = map[string]string{
		"bootstrap":   `<link href="//cdn.bootcss.com/bootstrap/4.0.0-alpha.2/css/bootstrap.min.css" rel="stylesheet" />`,
		"fontawesome": `<link href="//cdn.bootcss.com/font-awesome/4.5.0/css/font-awesome.min.css" rel="stylesheet" />`,
		"highlightjs": `<link href="//cdn.bootcss.com/highlight.js/9.1.0/styles/github.min.css" rel="stylesheet" />`,
	}
	CDN_JS = map[string]string{
		"bootstrap":   `<script src="//cdn.bootcss.com/bootstrap/4.0.0-alpha.2/js/bootstrap.min.js"></script>`,
		"jquery":      `<script src="//cdn.bootcss.com/jquery/2.2.0/jquery.min.js"></script>`,
		"mathjax":     `<script src="//cdn.bootcss.com/mathjax/2.6.1/MathJax.js?config=default"></script>`,
		"highlightjs": `<script src="//cdn.bootcss.com/highlight.js/9.1.0/highlight.min.js"></script>`,
	}
)
