package generator

import "embed"

type Skeleton struct {
	root     string
	template embed.FS
}

//go:embed templates/root
var serviceTpl embed.FS

func NewSkeleton(root string) *Skeleton {
	return &Skeleton{root: root, template: serviceTpl}
}
