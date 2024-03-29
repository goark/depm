package modules

import (
	"github.com/goark/depm/golist"
	"github.com/goark/depm/packages"
)

//Module information
type Module struct {
	Name      Name     // module name (path and version)
	Node      bool     // node module
	Edge      bool     // edge module
	License   string   // license of module
	Replace   Name     // replaced by this module
	Main      bool     // is this the main module?
	Indirect  bool     // is this module only an indirect dependency of main module?
	Update    Name     // available update, if any (with -u)
	Packages  []string // pakcages in this module
	UseCGO    bool     // using cgo in this module
	UseUnsafe bool     // include unsafe package in this module
	Deps      []Name   // dependency module names
	Error     error    // error loading module
}

func newModule(m *golist.Module, node, edge bool) *Module {
	if m == nil {
		return nil
	}
	mm := &Module{
		Name:      newName(m.Path, m.Version),
		Node:      node,
		Edge:      edge,
		Replace:   Name{},
		Main:      m.Main,
		Indirect:  m.Indirect,
		Update:    Name{},
		Packages:  []string{},
		UseCGO:    false,
		UseUnsafe: false,
		Deps:      []Name{},
		Error:     m.GetError(),
	}
	if m.Replace != nil {
		mm.Replace = newName(m.Replace.Path, m.Replace.Version)
	}
	if m.Update != nil {
		mm.Update = newName(m.Update.Path, m.Update.Version)
	}
	if len(m.Dir) > 0 {
		// fmt.Println(m.Dir)
		mm.License = FindLicense(m.Dir)
	}

	return mm
}

//Valid returns true if is not Incomplete
func (m *Module) Valid() bool {
	return m != nil && m.Error == nil
}

//Equal returns true if left == right
func (left *Module) Equal(right *Module) bool {
	return left.Valid() && right.Valid() && left.Name.Equal(right.Name)
}

//EdgeOnly returns true if is not Node and is Edge
func (m *Module) EdgeOnly() bool {
	return !m.Valid() || (!m.Node && m.Edge)
}

//SetPackage sets package name to Module
func (m *Module) SetPackage(p *packages.Package) {
	if m == nil {
		return
	}
	for _, s := range m.Packages {
		if s == p.Path {
			return
		}
	}
	m.Packages = append(m.Packages, p.Path)
	m.UseCGO = p.UseCGO
	m.UseUnsafe = p.UseUnsafe
}

//SetDeps sets dependency module name to Module
func (m *Module) SetDep(mm *golist.Module) *Module {
	if m == nil || mm == nil {
		return nil
	}
	if m.Name.EqualAll(mm.Path, mm.Version) {
		return nil
	}
	dm := newModule(mm, false, true)
	for _, nm := range m.Deps {
		if nm.EqualAll(mm.Path, mm.Version) {
			return dm
		}
	}
	m.Deps = append(m.Deps, dm.Name)
	return dm
}

//Modules is list of Modules.
type Modules struct {
	list []*Module
}

//Set method sets Module instance to Modules
func (ms *Modules) Set(m *Module) *Module {
	if ms == nil {
		return nil
	}
	for _, mm := range ms.list {
		if mm.Equal(m) {
			return mm
		}
	}
	ms.list = append(ms.list, m)
	return m
}

//Get method gets Module instance from Modules
func (ms *Modules) Get(name Name) *Module {
	if ms == nil {
		return nil
	}
	for _, mm := range ms.list {
		if mm.Name.Equal(name) {
			return mm
		}
	}
	return nil
}

//Set method sets Module instance to Modules
func (ms *Modules) Add(m *golist.Module) *Module {
	if ms == nil || m == nil {
		return nil
	}
	mm := ms.Get(newName(m.Path, m.Version))
	if mm == nil {
		mm = newModule(m, true, false)
		ms.list = append(ms.list, mm)
	}
	return mm
}

//List method returns list of modules.
func (ms *Modules) List() []*Module {
	if ms == nil {
		return nil
	}
	return ms.list
}

/* Copyright 2020-2021 Spiegel
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
