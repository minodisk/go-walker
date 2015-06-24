# go-walker [![Circle CI](https://circleci.com/gh/minodisk/go-walker/tree/master.svg?style=svg)](https://circleci.com/gh/minodisk/go-walker/tree/master)

Walk a directory recursively in Go.

## Installation

```bash
go get gopkg.in/minodisk/go-walker.v1
```

## Usage

### Walk

```go
import (
  "os"
  "gopkg.in/minodisk/go-walker.v1"
)

func main() {
  isFile = false
  err := walker.Walk(func (name string, fi os.FileInfo) (bool, error) {
    if !fi.IsDir() && name == "target/file" {
      isFile = true
      return false, nil
    }
    return true, nil
  })
  // do something
}
```

### Find files

```go
import "gopkg.in/minodisk/go-walker.v1"

func main() {
  filenames, err := walker.FindFiles("fixtures")
  if err != nil {
    panic(err)
  }
  // do something
}
```

## Reference

http://godoc.org/gopkg.in/minodisk/go-walker.v1
