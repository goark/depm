package modules

import (
	"context"

	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/packages"
	"github.com/spiegel-im-spiegel/errs"
)

//ImportModules gets modules dependency information
func ImportModules(ctx context.Context, gctx golist.Context, name string, updFlag bool, withInternal bool) (*Modules, error) {
	ps, err := packages.ImportPackages(ctx, gctx, name)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("name", name), errs.WithContext("updFlag", updFlag))
	}
	ms := importModules(ps, withInternal)
	if updFlag {
		for _, m := range ms.List() {
			ml, err := gctx.GetModules(ctx, m.Name.Path, true)
			if err != nil {
				return nil, errs.Wrap(err, errs.WithContext("path", m.Name.Path), errs.WithContext("updFlag", updFlag))
			}
			if upd := searchModule(m.Name, ml); upd != nil && upd.Update != nil {
				if !m.Name.EqualAll(upd.Update.Path, upd.Update.Version) {
					m.Update = newName(upd.Update.Path, upd.Update.Version)
				}
			}
		}
	}
	return ms, nil
}

func importModules(ps *packages.Packages, withInternal bool) *Modules {
	ms := &Modules{list: []*Module{}}
	for _, p := range ps.List() {
		if m := ms.Add(p.Contained); m != nil {
			if withInternal || (!withInternal && !p.IsInternal()) {
				m.SetPackage(p)
			}
			for _, path := range p.Imports {
				if dp := ps.Get(path); dp != nil {
					if dm := m.SetDep(dp.Contained); dm != nil {
						dm = ms.Set(dm)
						if withInternal || (!withInternal && !dp.IsInternal()) {
							dm.SetPackage(dp)
						}
					}
				}
			}
		}
	}
	return ms
}

func searchModule(name Name, mlist []golist.Module) *golist.Module {
	for _, m := range mlist {
		if name.EqualAll(m.Path, m.Version) {
			return &m
		}
	}
	return nil
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
