package dependency

import (
	"sort"

	"github.com/spiegel-im-spiegel/depm/packages"
)

//EdgePackage is Graph of dependent packages
type EdgePackage struct {
	Package *packages.Package
	Deps    []*packages.Package
}

func newEdgePackage(p *packages.Package) *EdgePackage {
	return &EdgePackage{Package: p, Deps: []*packages.Package{}}
}

//New creates slice if EdgePackage instances.
func NewPackages(ps *packages.Packages, withStd bool, withInternal bool) []*EdgePackage {
	nd := []*EdgePackage{}
	for _, p := range ps.List() {
		if p.Valid() && !p.IsStandard() && (!withInternal && !p.IsInternal() || withInternal) && len(p.Imports) > 0 {
			n := newEdgePackage(p)
			for _, path := range p.Imports {
				dp := ps.Get(path)
				if p.Valid() && (((!withStd && !dp.IsStandard()) || withStd) && (!withInternal && !dp.IsInternal() || withInternal)) {
					n.Deps = append(n.Deps, dp)
				}
			}
			if len(n.Deps) > 0 {
				nd = append(nd, n)
			}
		}
	}
	sort.SliceStable(nd, func(i, j int) bool {
		return nd[i].Package.Path < nd[j].Package.Path
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
