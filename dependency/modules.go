package dependency

import (
	"sort"

	"github.com/spiegel-im-spiegel/depm/modules"
)

//EdgePackage is Graph of dependent packages
type EdgeModule struct {
	Module *modules.Module
	Deps   []*modules.Module
}

func newEdgeModule(m *modules.Module) *EdgeModule {
	return &EdgeModule{Module: m, Deps: []*modules.Module{}}
}

//NewModules creates slice if EdgeModule instances.
func NewModules(ms *modules.Modules) []*EdgeModule {
	nd := []*EdgeModule{}
	for _, m := range ms.List() {
		if len(m.Deps) > 0 {
			n := newEdgeModule(m)
			for _, nm := range m.Deps {
				dm := ms.Get(nm)
				if dm != nil {
					n.Deps = append(n.Deps, dm)
				}
			}
			if len(n.Deps) > 0 {
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
