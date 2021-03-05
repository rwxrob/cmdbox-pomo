package main

import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/rwxrob/cmdtab-pomo"
)

func main() {
	cmdtab.Execute("pomo")
}
