# Pomodoro Timer Command in Go

*See [cmdtab](https://github.com/rwxrob/cmdtab) for how to compose
`pomo` command into your own monolith utilities.*

![WIP](https://img.shields.io/badge/status-wip-red.svg)
[![GoDoc](https://godoc.org/cmdtab-pomo?status.svg)](https://godoc.org/cmdtab-pomo)
[![License](https://img.shields.io/badge/license-MPLv2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/cmdtab-pomo)](https://goreportcard.com/report/cmdtab-pomo)
[![Coverage](https://gocover.io/_badge/cmdtab-pomo)](https://gocover.io/cmdtab-pomo)

## Install `pomo` as Standalone

The `pomo` command can be used as a standalone program (even though it
is also designed to be modularly used in a
[`cmdtab`](https://github.com/rwxrob/cmdtab) monolith command).

```
go get github.com/rwxrob/cmdtab-pomo/pomo
```

That's it. It will download, compile and install `pomo` (provided you
have Go 1.16 or later installed).

*Other packaging and distribution methods are being considered.*

### Add Pomodoro to TMUX

```tmux
set -g status-left "%A, %B %-e, %Y, %-l:%M%P %Z%0z #(pomo)" 
```
