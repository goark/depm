package packages

import (
	"strings"

	"github.com/spiegel-im-spiegel/depm/golist"
	"golang.org/x/tools/imports"
)

//Package information
type Package struct {
	Path       string         // import path of package in dir
	Root       bool           // root package
	Detail     bool           // if false Path only
	Goroot     bool           // is this package in the Go root?
	Standard   bool           // is this package part of the standard Go library?
	ForTest    string         // package is only for use in named test
	Incomplete bool           // this package or a dependency has an error
	DepOnly    bool           // package is only a dependency, not explicitly listed
	Contained  *golist.Module // info about package's containing module, if any (can be nil)
	Imports    []string       // import paths used by this package
}

func newPackageName(name string) *Package {
	return &Package{Path: imports.VendorlessPath(name), Root: true}
}

//Copy method copies elements of Package from golist.Package instance
func (p *Package) Copy(pp *golist.Package) *Package {
	if p == nil {
		p = newPackageName(pp.ImportPath)
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

//Equal returns true if left == right
func (p *Package) Equal(right *Package) bool {
	return p.Path == right.Path
}

//Equal returns true if standard Go library
func (p *Package) Valid() bool {
	return !p.Incomplete
}

//Equal returns true if standard Go library
func (p *Package) IsStandard() bool {
	return p.Goroot && p.Standard
}

//IsUnsafe returns true if unsafe package
func (p *Package) IsUnsafe() bool {
	return strings.EqualFold(p.Path, "unsafe") || strings.EqualFold(p.Path, "C")
}

//IsInternal returns true if internal package
func (p *Package) IsInternal() bool {
	return strings.Contains(p.Path, "internal")
}

//Packages is list of Packages.
type Packages struct {
	list []*Package
}

func importPackages(plist []golist.Package) *Packages {
	ps := &Packages{list: []*Package{}}
	for i := 0; i < len(plist); i++ {
		ps.Add(&plist[i])
		for _, s := range plist[i].Imports {
			p := ps.Get(s)
			if p == nil {
				p = ps.Set(newPackageName(s))
			}
			p.Root = false
		}
	}
	return ps
}

//Set method sets Package instance in Packages.
func (ps *Packages) Set(p *Package) *Package {
	for i := 0; i < len(ps.list); i++ {
		if ps.list[i].Equal(p) {
			if !ps.list[i].Detail {
				p.Root = ps.list[i].Root
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

//List method returns list of packages.
func (ps *Packages) List() []*Package {
	return ps.list
}

// //AddName method adds Package instance (Path element only) in Packages.
// func (ps *Packages) AddName(name string) {
// 	ps.Set(newPackageName(name))
// }

//Add method adds Package instance in Packages.
func (ps *Packages) Add(p *golist.Package) {
	if pp := ps.Get(p.ImportPath); pp != nil {
		pp.Copy(p)
		return
	}
	ps.Set(newPackageName(p.ImportPath).Copy(p))
}

//Merge method merges Package instance.
func (ps *Packages) Merge(pps *Packages) {
	for _, p := range pps.List() {
		ps.Set(p)
	}
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
