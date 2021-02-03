# Localizer: convenient localization for Go

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/razor-1/localizer/CI?style=flat-square)](https://github.com/razor-1/localizer/actions?query=workflow%3ABuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/razor-1/localizer?style=flat-square)](https://goreportcard.com/report/github.com/razor-1/localizer)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.15-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/razor-1/localizer)](https://pkg.go.dev/github.com/razor-1/localizer)

Localizer intends to make it easy for you to work with locales in Go. It was inspired by 
many good tools that came before it, such as:
* [Babel](http://babel.pocoo.org/) for Python
* [go-i18n](https://github.com/nicksnyder/go-i18n) 
* [cldr](https://github.com/theplant/cldr)

I couldn't find one Go package that did everything, so I took the best of what was out there, made some
improvements & tweaks, and put them together into localizer. I also wanted to leverage golang.org/x/text/language 
as much as possible, since it takes care of some difficult tasks. 

## Getting Started: Example
```go
import (
    "fmt"
    "time"

    "github.com/razor-1/cldr/resources/currency"
    "golang.org/x/text/language"

    "github.com/razor-1/localizer"
)

func main() {
    l, err := localizer.NewLocale(language.Spanish)
    if err != nil {
        panic(err)
    }

    jan2020 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
    month, err := l.Calendar.Format(jan2020, "MMMM")
    if err != nil {
        panic(err)
    }

    fmt.Println(month) // "enero"
    fmt.Println(l.Calendar.FormatNames.Months.Wide.Jan) // "enero"
    fmt.Println(l.FmtNumber(10000.12)) // "10.000,12"
    cur, err := l.FmtCurrency(currency.USD, 10000.12)
    if err != nil {
        panic(err)
    }

    fmt.Println(cur) // "10.000,12 US$"
}
```

As you can see, the locale object makes it easy to interact with CLDR data. What about translated strings?
```go
gotl := gotext.NewLocale("translations", "es")
gotl.AddDomain("messages")

gtstore := gotextstore.GetTranslationStore(gotl)
l, err := localizer.NewLocaleWithStore(language.Spanish, gtstore)
if err != nil {
    panic(err)
}

fmt.Println(l.Get("form.button.login")) //"Iniciar sesión"
fmt.Println(l.GetPlural("%d hours", 1, 1)) //"1 hora"
fmt.Println(l.GetPlural("%d hours", 2, 2)) //"2 horas"
```

You can use any package which implements the `localizer/store.TranslationStore` interface to load translations. The above
example uses a helper package, localizer/gotextstore, which integrates with [gotext](https://github.com/leonelquinteros/gotext) to
provide gettext po/mo support for localizer. If other packages implement this interface, support for xliff, xmb, and 
other common formats can be added. You can also easily implement it for your own custom store.

localizer also exposes a [message](https://godoc.org/golang.org/x/text/message) printer if you'd like to use it. 
Just call `l.NewPrinter()` and you can then call `Printf()` and other methods on it.
