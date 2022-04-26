package pkg

import (
	"github.com/anchore/stereoscope/pkg/file"
	"github.com/anchore/syft/syft/linux"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/source"
)

type Context struct {
	Source         *source.Metadata
	Distro         *linux.Release
	Files          []file.Reference
	PackageCatalog []pkg.Package
	Secrets        interface{}
}
