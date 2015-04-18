# go-cloudatcost

go-cloudatcost is a Go client library for accessing the [cloudatcost API](https://github.com/cloudatcost/api).

go-cloudatcost was heavily inspired by the go-github library.

**travis-ci:** [![Build Status](https://travis-ci.org/masayukioguni/go-cloudatcost.svg?branch=master)](https://travis-ci.org/masayukioguni/go-cloudatcost)

**drone.io:** [![Build Status](https://drone.io/github.com/masayukioguni/go-cloudatcost/status.png)](https://drone.io/github.com/masayukioguni/go-cloudatcost/latest)

**Test Coverage:** [![Coverage Status](https://coveralls.io/repos/masayukioguni/go-cloudatcost/badge.png?branch=master)](https://coveralls.io/r/masayukioguni/go-cloudatcost?branch=master)([gocov report](https://drone.io/github.com/masayukioguni/go-cloudatcost/files/coverage.html))

**GoDoc:** [![GoDoc](https://godoc.org/github.com/masayukioguni/go-cloudatcost/cloudatcost?status.svg)](https://godoc.org/github.com/masayukioguni/go-cloudatcost/cloudatcost)

API Version: `v1`

## References
https://github.com/cloudatcost/api

## Installation
```bash
$ go get github.com/masayukioguni/go-cloudatcost/cloudatcost
```

## Example

```go
package main

import (
  "fmt"
  "github.com/masayukioguni/go-cloudatcost/cloudatcost"
  "os"
)

func main() {
  Login := os.Getenv("CLOUDATCOST_API_LOGIN")
  Key := os.Getenv("CLOUDATCOST_API_KEY")

  // Initializes a new CloudAtCost client
  client, _ := cloudatcost.NewClient(&cloudatcost.Option{Login: Login, Key: Key})

  listservers, hr, err := client.ServersService.List()

  if err != nil {
    fmt.Printf("error: %v\n\n", err)
    return
  }

  if hr.StatusCode != 200 {
    fmt.Printf("http response error: %+v %+v \n\n", hr, err)
    return
  }

  fmt.Printf("%v,%v\n", listservers, err)

}
```