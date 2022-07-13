# timekit
[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/bartmika/timekit)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartmika/timekit)](https://goreportcard.com/report/github.com/bartmika/timekit)
[![License](https://img.shields.io/github/license/bartmika/timekit)](https://github.com/bartmika/timekit/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/bartmika/timekit)

Convenience functions to make your life easier when using with Golang's [`time`](https://pkg.go.dev/time) package.

## Installation

In your Golang project, please run:

```
go get github.com/bartmika/timekit
```

## Documentation

All [documentation](https://pkg.go.dev/github.com/bartmika/timekit) can be found here.

## Example Usage
The following usage section will show a few interesting solutions that can be solved by this library.

### How do I get the first day of this year in Golang?

```go
import (
    "fmt"

    "github.com/bartmika/timekit"
)

startOfYearDate := timekit.FirstDayOfThisYear(time.Now)
fmt.Println(startOfYearDate)
```

If you have interest in finding out more functions for getting different date/times then checkout [timekit.go](timekit.go) file.

### How do parse JavaScript time into Golang time?

In your browser console, try writing this:
```javascript
// EXAMPLE JAVASCRIPT CODE
var dt = new Date()
console.log(dt.getTime()) // 1643082322380
console.log(dt) // Mon Jan 24 2022 22:45:22 GMT-0500 (Eastern Standard Time)
```

Then try this in your go source file:

```go
import (
    "fmt"

    "github.com/bartmika/timekit"
)

jsTime := int64(1643082322380)
goTime := ParseJavaScriptTime(jsTime)
fmt.Println(goTime)
```

### How get a range of date/times between two dates?

```go
import (
    "fmt"

    "github.com/bartmika/timekit"
)

start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
dts := RangeFromTimeStepper(start, end, 0, 0, 1, 0, 0, 0) // Step by day.
fmt.Println(dts) // Output:
                 // Jan 7th 2022
                 // Jan 8th 2022
                 // Jan 9th 2022
                 // Jan 10th 2022
```

If you have interest in finding out more range functions then checkout [range.go](range.go) and [timestepper.go](timestepper.go) files.


### How get iterate between two dates?

```go
import (
    "fmt"

    "github.com/bartmika/timekit"
)

loc := time.UTC                                 // closure can be used if necessary
start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
ts := NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)

var actual time.Time
running := true
for running {
    // Get the value we are on in the timestepper.
    actual = ts.Get()

    log.Println(actual) // For debugging purposes only.

    // Run our timestepper to get our next value.
    ts.Next()

    running = ts.Done() == false
}
```

If you have interest in finding out iterating between two date/times then checkout [timestepper.go](timestepper.go) file.

## Contributing

Found a bug? Want a feature to improve your developer experience when dealing with the [`time`](https://pkg.go.dev/time) package? Please create an [issue](https://github.com/bartmika/timekit/issues).

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [ISC License](LICENSE).

Resource used:

* [Stubbing Time.Now() in golang](https://labs.yulrizka.com/en/stubbing-time-dot-now-in-golang/) was a tremendous help in getting me to understand how to unit test with the [`time.Time`](https://pkg.go.dev/time) package.
* [dannav/hhmmss](https://github.com/dannav/hhmmss) package is used for the Golang `hh:mm:ss` string to duration conversation method.
