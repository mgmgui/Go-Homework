module github.com/mgmgui/go-homework/hello

go 1.20

require (
	github.com/mgmgui/go-homework/hello/calculator v0.0.0-00010101000000-000000000000
	rsc.io/quote/v3 v3.1.0
)

require (
	golang.org/x/text v0.8.0 // indirect
	rsc.io/sampler v1.3.1 // indirect
)

replace github.com/mgmgui/go-homework/hello/calculator => ./calculator
