package main

import (
	. "polydawn.net/gosh/psh"
)

func main() {
	echo := Sh("echo")("-n", "-e")

	echo("wat\n", "\t\033[0;31mred and indented\033[0m\n")()

	Sh("bash")(17)

	cat := Sh("cat")
	catIn := cat.BakeArgs("-")
	catIn()
}