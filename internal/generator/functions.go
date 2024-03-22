package generator

import (
	"html/template"
	"strings"
)

var FuncMap = template.FuncMap{
	"ToPublicName":  ToPublicName,
	"ToPrivateName": ToPrivateName,
}

func ToPrivateName(name string) string {
	return strings.ToLower(string(name[0])) + name[1:]
}

func ToPublicName(name string) string {
	return strings.ToUpper(string(name[0])) + name[1:]
}
