package modules

import (
	"github.com/spiegel-im-spiegel/depm/golist"
)

//Module information
type Module struct {
	Name     Name     // module name (path and version)
	Main     bool     // is this the main module?
	Indirect bool     // is this module only an indirect dependency of main module?
	Update   Name     // available update, if any (with -u)
	Packages []string // pakcages in this module
	Deps     []Name   // dependency module names
}

func newModule(m *golist.Module) *Module {
	if m == nil {
		return nil
	}
	mm := &Module{
		Name:     newName(m.Path, m.Version),
		Main:     m.Main,
		Indirect: m.Indirect,
		Update:   Name{},
		Packages: []string{},
		Deps:     []Name{},
	}
	if m.Update != nil {
		mm.Update = newName(m.Update.Path, m.Update.Version)
	}
	return mm
}

//Equal returns true if left == right
func (left *Module) Equal(right *Module) bool {
	return left.Name.Equal(right.Name)
}

//SetPackage sets package name to Module
func (m *Module) SetPackage(pkg string) {
	for _, s := range m.Packages {
		if s == pkg {
			return
		}
	}
	m.Packages = append(m.Packages, pkg)
}

//SetDeps sets dependency module name to Module
func (m *Module) SetDeps(dm *Module) {
	for _, nm := range m.Deps {
		if nm.Equal(dm.Name) {
			return
		}
	}
	m.Deps = append(m.Deps, dm.Name)
}

//Modules is list of Modules.
type Modules struct {
	list []*Module
}

//Set method sets Module instance to Modules
func (ms *Modules) Set(m *Module) *Module {
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
	for _, mm := range ms.list {
		if mm.Name.Equal(name) {
			return mm
		}
	}
	return nil
}

//List method returns list of modules.
func (ms *Modules) List() []*Module {
	return ms.list
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
