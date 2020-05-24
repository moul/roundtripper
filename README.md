# roundtripper

:smile: roundtripper

![CI](https://github.com/moul/roundtripper/workflows/CI/badge.svg)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/roundtripper)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/roundtripper/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/roundtripper.svg)](https://github.com/moul/roundtripper/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/roundtripper)](https://goreportcard.com/report/moul.io/roundtripper)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/roundtripper/badge)](https://www.codefactor.io/repository/github/moul/roundtripper)
[![codecov](https://codecov.io/gh/moul/roundtripper/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/roundtripper)
[![GolangCI](https://golangci.com/badges/github.com/moul/roundtripper.svg)](https://golangci.com/r/github.com/moul/roundtripper)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


## Usage

```golang
import "moul.io/roundtripper"

client := http.Client{
    Transport: &roundtripper.Transport{
        ExtraHeader: http.Header{
            "Authorization": []string{"Bearer LoremIpsumDolorSitAmet..."},
        },
    },
}
client.Get("...")
```

See https://pkg.go.dev/moul.io/roundtripper

## Install

### Using go

```console
$ go get -u moul.io/roundtripper
```

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
