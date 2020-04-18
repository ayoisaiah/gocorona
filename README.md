# Gocorona - Coronavirus Statistics Dashboard for your Terminal

[![Go Report Card](https://goreportcard.com/badge/github.com/ayoisaiah/gocorona)](https://goreportcard.com/report/github.com/ayoisaiah/gocorona)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/7136493cf477467387381890cb25dc9e)](https://www.codacy.com/manual/ayoisaiah/gocorona?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ayoisaiah/gocorona&amp;utm_campaign=Badge_Grade)
[![HitCount](http://hits.dwyl.com/ayoisaiah/gocorona.svg)](http://hits.dwyl.com/ayoisaiah/gocorona)
[![PR's Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)

View the latest Coronavirus (COVID-19) statistics in your terminal.

## Features

- View worldwide stats for cases, deaths, recoveries, active cases and
mortality rate
- View stats for Coronavirus disease reports across the US states
- Sort the data by cases, cases today, deaths, deaths today, recoveries, active,
critical and mortality rate
- See prevention tips and other info about the Coronavirus pandemic

## Demo

[![asciicast](https://asciinema.org/a/6FfMPrBKsEqpiO7YyaEdrZoV0.svg)](https://asciinema.org/a/6FfMPrBKsEqpiO7YyaEdrZoV0)

## Installation and usage

You can download the precompiled binaries for Linux, Windows, and macOS [here](https://github.com/ayoisaiah/gocorona/releases) (only for amd64)

Ensure the binary is executable by the current user, then execute it:

```bash
$ chmod +x gocorona-linux
$ ./gocorona-linux
```

Or if you have `go` installed, clone the repo and install from source:

```bash
$ git clone https://github.com/ayoisaiah/gocorona
$ cd gocorona
$ go install ./cmd/gocorona/...
```

Then run the `gocorona` command:

```bash
$ gocorona
```

## Credit and sources

Gocorona relies heavily on other open source software listed below:

- [Termui](https://github.com/gizak/termui)
- [NovelCovid API](https://github.com/NovelCovid/API)

## Contribute

Bug reports, or pull requests are much welcome!

## Licence

Created by Ayooluwa Isaiah and released under the terms of the [MIT Licence](http://opensource.org/licenses/MIT).
