# go-walker

## Installation

```bash
go get gopkg.in/minodisk/go-walker.v1
```

## Usage

```go
import "github.com/minodisk/go-walker"

func main() {
  filenames, err := walker.FindFiles("fixtures")
  if err != nil {
    panic(err)
  }
  // ...
}
```

## Reference
