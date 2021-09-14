module task

go 1.15

replace (
	entity => ../entity
	proxy => ../proxy
)

require (
	entity v0.0.0
	github.com/antchfx/htmlquery v1.2.3
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	proxy v0.0.0
)
