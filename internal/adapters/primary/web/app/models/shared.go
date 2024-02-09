package models

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	cents = 100
)

var (
	printer = message.NewPrinter(language.English)
)
