# UnixTime

[![Test](https://github.com/eosswedenorg-go/unixtime/actions/workflows/test.yml/badge.svg)](https://github.com/eosswedenorg-go/unixtime/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/eosswedenorg-go/unixtime.svg)](https://pkg.go.dev/github.com/eosswedenorg-go/unixtime)

Simple module to handle unix timestamp json encoding/decoding.

## Example

```go

package main

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/eosswedenorg-go/unixtime"
)

type MyStruct struct {
    Timestamp unixtime.Time
}

func main() {
    var s MyStruct

    input := `{"timestamp": 1205647965800}`

    if err := json.Unmarshal([]byte(input), &s); err != nil {
        panic(err)
    }

    // Output: Decoded: 2008-03-16 06:12:45.800 +0000 UTC
    fmt.Println("Decoded:", s.Timestamp.Time().Format("2006-01-02 15:04:05.000 -0700 MST"))

    s.Timestamp.FromTime(time.Date(2014, 6, 11, 14, 3, 55, 0, time.UTC).Add(time.Millisecond * 625))

    data, err := json.Marshal(s)
    if err != nil {
        panic(err)
    }

    // Output: Encoded: {"Timestamp":1402495435625}
    fmt.Println("Encoded:", string(data))
}

```

## Author

Henrik Hautakoski - [Sw/eden](https://eossweden.org/) - [henrik@eossweden.org](mailto:henrik@eossweden.org)