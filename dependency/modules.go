package dependency

import (
	"sort"

	"github.com/goark/depm/modules"
)

//EdgePackage is Graph of dependent packages
type NodeModule struct {
	Module *modules.Module
	Deps   []*modules.Module
}

func newNodeModule(m *modules.Module) *NodeModule {
	return &NodeModule{Module: m, Deps: []*modules.Module{}}
}

//NewModules creates slice if NodeModule instances.
func NewModules(ms *modules.Modules) []*NodeModule {
	nd := []*NodeModule{}
	for _, m := range ms.List() {
		if m.Valid() {
			n := newNodeModule(m)
			for _, nm := range m.Deps {
				if dm := ms.Get(nm); dm != nil {
					n.Deps = append(n.Deps, dm)
				}
			}
			if len(n.Deps) > 0 || !m.EdgeOnly() {
				nd = append(nd, n)
			}
		}
	}
	sort.SliceStable(nd, func(i, j int) bool {
		return nd[i].Module.Name.String() < nd[j].Module.Name.String()
	})
	return nd
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
