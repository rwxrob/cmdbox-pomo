# Pomodoro Timer Command in Go

[![GoDoc](https://godoc.org/cmdbox-pomo?status.svg)](https://godoc.org/cmdbox-pomo)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/cmdbox-pomo)](https://goreportcard.com/report/cmdbox-pomo)

## Install `pomo` as Standalone

The `pomo` command can be used as a standalone program

```
go get github.com/rwxrob/cmdbox-pomo/pomo
```

That's it. It will download, compile and install `pomo` (provided you
have Go 1.16 or later installed).

## Usage

```
pomo help
```

## Add Pomodoro to TMUX

Here's an example of how to add `pomo` to your TMUX configuration. Your
mileage may vary.

```tmux
set -g status-interval 1
set -g status-left "%A, %B %-e, %Y, %-l:%M%P %Z%0z #(pomo)" 
```
