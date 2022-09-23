package testdata

import "go-proxy/internal/module"

var Modules = []*module.Module{
	{
		Name: "hello-world",
		Git:  "https://gitlab.com/anilmisirlioglu/hello-world",
		Source: []string{
			"https://gitlab.com/anilmisirlioglu/hello-world",
			"https://gitlab.com/anilmisirlioglu/hello-world/tree/main{/dir}",
			"https://gitlab.com/anilmisirlioglu/hello-world/tree/main{/dir}/{file}#L{line}",
		},
	},
}
