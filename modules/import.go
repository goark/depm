package modules

import "github.com/spiegel-im-spiegel/depm/packages"

func ImportModules(ps *packages.Packages) *Modules {
	ms := &Modules{list: []*Module{}}
	for _, p := range ps.List() {
		m := newModule(p.Contained)
		if m != nil {
			m = ms.Set(m)
			m.SetPackage(p.Path)
			for _, path := range p.Imports {
				dp := ps.Get(path)
				if dp != nil {
					dm := newModule(dp.Contained)
					if dm != nil && !m.Equal(dm) {
						dm = ms.Set(dm)
						dm.SetPackage(dp.Path)
						m.SetDeps(dm)
					}
				}
			}
		}
	}
	return ms
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
