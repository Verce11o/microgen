package generator

import "embed"

type Skeleton struct {
	root     string
	template embed.FS
}

//go:embed templates/root
var serviceTpl embed.FS

func NewSkeleton() *Skeleton {
	return &Skeleton{root: "templates/root", template: serviceTpl}
}
