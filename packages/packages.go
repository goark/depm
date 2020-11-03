package packages

import (
	"strings"

	"github.com/spiegel-im-spiegel/depm/golist"
	"golang.org/x/tools/imports"
)

//Package information
type Package struct {
	Path       string         // import path of package in dir
	Node       bool           // node package
	Edge       bool           // edge package
	Detail     bool           // if false Path only
	Goroot     bool           // is this package in the Go root?
	Standard   bool           // is this package part of the standard Go library?
	ForTest    string         // package is only for use in named test
	Incomplete bool           // this package or a dependency has an error
	DepOnly    bool           // package is only a dependency, not explicitly listed
	Contained  *golist.Module // info about package's containing module, if any (can be nil)
	Imports    []string       // import paths used by this package
}

func newPackageName(name string, node, edge bool) *Package {
	return &Package{Path: imports.VendorlessPath(name), Node: node, Edge: edge}
}

//Copy method copies elements of Package from golist.Package instance
func (p *Package) Copy(pp *golist.Package) *Package {
	if p == nil {
		return nil
	}
	p.Detail = true
	p.Goroot = pp.Goroot
	p.Standard = pp.Standard
	p.ForTest = pp.ForTest
	p.Incomplete = pp.Incomplete
	p.DepOnly = pp.DepOnly
	p.Contained = pp.Module
	if len(pp.Imports) > 0 {
		p.Imports = make([]string, len(pp.Imports), cap(pp.Imports))
		copy(p.Imports, pp.Imports)
	}
	return p
}

//Valid returns true if is not Incomplete
func (p *Package) Valid() bool {
	return p != nil && !p.Incomplete
}

//Equal returns true if left == right
func (left *Package) Equal(right *Package) bool {
	return left.Valid() && right.Valid() && left.Path == right.Path
}

//EdgeOnly returns true if is not Node and is Edge
func (p *Package) EdgeOnly() bool {
	return !p.Valid() || (!p.Node && p.Edge)
}

//IsStandard returns true if standard Go library
func (p *Package) IsStandard() bool {
	return p.Valid() && p.Goroot && p.Standard
}

//IsUnsafe returns true if unsafe package
func (p *Package) IsUnsafe() bool {
	return p.Valid() && (strings.EqualFold(p.Path, "unsafe") || strings.EqualFold(p.Path, "C"))
}

//IsInternal returns true if internal package
func (p *Package) IsInternal() bool {
	return p.Valid() && strings.Contains(p.Path, "internal")
}

//Packages is list of Packages.
type Packages struct {
	list []*Package
}

//Set method sets Package instance in Packages.
func (ps *Packages) Set(p *Package) *Package {
	if ps == nil || p == nil {
		return nil
	}
	for i := 0; i < len(ps.list); i++ {
		if ps.list[i].Equal(p) {
			if !ps.list[i].Detail {
				p.Node = ps.list[i].Node
				p.Edge = ps.list[i].Edge
				ps.list[i] = p
			} else {
				p = ps.list[i]
			}
			return p
		}
	}
	ps.list = append(ps.list, p)
	return p
}

//Get method gets Package instance from Packages.
func (ps *Packages) Get(path string) *Package {
	if ps == nil {
		return nil
	}
	for i := 0; i < len(ps.list); i++ {
		if ps.list[i].Path == path {
			return ps.list[i]
		}
	}
	return nil
}

//Add method adds Package instance in Packages.
func (ps *Packages) Add(p *golist.Package) *Package {
	if ps == nil || p == nil {
		return nil
	}
	if pp := ps.Get(p.ImportPath); pp != nil {
		pp.Copy(p)
		return pp
	}
	return ps.Set(newPackageName(p.ImportPath, true, false).Copy(p))
}

//Merge method merges Package instance.
func (ps *Packages) Merge(pps *Packages) {
	for _, p := range pps.List() {
		ps.Set(p)
	}
}

//List method returns list of packages.
func (ps *Packages) List() []*Package {
	if ps == nil {
		return nil
	}
	return ps.list
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
