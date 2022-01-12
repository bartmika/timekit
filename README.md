# timekit
[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/bartmika/timekit)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartmika/timekit)](https://goreportcard.com/report/github.com/bartmika/timekit)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fbartmika%2Ftimekit.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fbartmika%2Ftimekit?ref=badge_shield)

Convenience functions to make your life easier when using with Golang's [`time`](https://pkg.go.dev/time) package.

## Installation

In your Golang project, please run:

```
go get github.com/bartmika/timekit
```

## Usage

```go
import (
    "fmt"

    "github.com/bartmika/timekit"
)

startOfYearDate := timekit.FirstDayOfThisYear(time.Now)
fmt.Println(startOfYearDate)
```

## Documentation

All [documentation](https://pkg.go.dev/github.com/bartmika/timekit) can be found here.

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [Unlicense](LICENSE).

Resource used:

* [Stubbing Time.Now() in golang](https://labs.yulrizka.com/en/stubbing-time-dot-now-in-golang/) was a tremendous help in getting me to understand how to unit test with the [`time.Time`](https://pkg.go.dev/time) package.
