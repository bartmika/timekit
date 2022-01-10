# timekit
Additional helpful functions for Golang [`time.Time`](https://pkg.go.dev/time) package to make your life easier.

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

startOfYearDate := timekit.FirstDayOfThisYear(time.Now())
fmt.Println(startOfYearDate)
```

## Documentation

All [documentation](https://pkg.go.dev/github.com/bartmika/timekit) can be found here.

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [Unlicense](LICENSE).

Resource used:

* [Stubbing Time.Now() in golang](https://labs.yulrizka.com/en/stubbing-time-dot-now-in-golang/) was a tremendous help in getting me to understand how to unit test with the [`time.Time`](https://pkg.go.dev/time) package.
