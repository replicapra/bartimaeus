package util

import (
	"github.com/charmbracelet/log"
)

func CheckErr(msg interface{}) {
	if msg != nil {
		log.Fatal(msg)
	}
}
