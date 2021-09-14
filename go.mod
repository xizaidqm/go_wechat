module main

go 1.15

replace (
	entity => ./entity
	proxy => ./proxy
	task => ./task
)

require (
	github.com/PuerkitoBio/goquery v1.7.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	task v0.0.1
)
