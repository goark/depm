package packages

import (
	"context"

	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/errs"
)

//ImportPackages gets packages dependency information
func ImportPackages(ctx context.Context, gctx golist.Context, name string) (*Packages, error) {
	plist, err := gctx.GetPackages(ctx, name)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("name", name))
	}
	ps := importPackages(plist)
	for {
		done := true
		for _, p := range ps.List() {
			if !p.Detail && p.Path != "C" {
				plist, err := gctx.GetPackages(ctx, p.Path)
				if err != nil {
					return ps, errs.Wrap(err, errs.WithContext("Package.Path", p.Path))
				}
				ps.Merge(importPackages(plist))
				done = false
				break
			}
		}
		if done {
			break
		}
	}
	return ps, nil
}

func importPackages(plist []golist.Package) *Packages {
	ps := &Packages{list: []*Package{}}
	for i := 0; i < len(plist); i++ {
		ps.Add(&plist[i])
		for _, s := range plist[i].Imports {
			p := ps.Get(s)
			if p == nil {
				ps.Set(newPackageName(s, false, true))
			} else {
				p.Edge = true
			}
		}
	}
	return ps
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
