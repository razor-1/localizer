module github.com/razor-1/localizer/cmd/test

go 1.14

replace (
	github.com/razor-1/gotext => /Users/jon/gocode/src/github.com/razor-1/gotext
	github.com/razor-1/localizer => /Users/jon/gocode/src/github.com/razor-1/localizer
	github.com/razor-1/localizer/store => /Users/jon/gocode/src/github.com/razor-1/localizer/store
)

require (
	github.com/leonelquinteros/gotext v1.4.0 // indirect
	github.com/razor-1/gotext v0.0.0-00010101000000-000000000000
	github.com/razor-1/localizer v0.0.3
	golang.org/x/text v0.3.3
)
