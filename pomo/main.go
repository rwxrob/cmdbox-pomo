package main

import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/rwxrob/cmdbox-pomo"
)

func main() {
	cmdtab.Execute("pomo")
}
