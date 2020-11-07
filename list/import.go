package list

import (
	"context"

	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/packages"
	"github.com/spiegel-im-spiegel/errs"
)

//ImportModules gets modules dependency information
func ImportModules(ctx context.Context, name string, updFlag bool, opts ...golist.OptEnv) ([]*golist.Module, error) {
	ps, err := packages.ImportPackages(ctx, name, opts...)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("name", name), errs.WithContext("updFlag", updFlag))
	}
	ms := newModules()
	for _, p := range ps.List() {
		ms.Set(p.Contained)
	}
	if updFlag {
		for _, m := range ms.List() {
			ml, err := golist.GetModules(ctx, m.Path, opts...)
			if err != nil {
				return nil, errs.Wrap(err, errs.WithContext("path", m.Path), errs.WithContext("updFlag", updFlag))
			}
			if upd := searchModule(m, ml); upd != nil && upd.Update != nil {
				ms.Set(upd)
			}
		}
	}
	return ms.List(), nil
}

func searchModule(m *golist.Module, mlist []golist.Module) *golist.Module {
	name := m.Name()
	for _, mm := range mlist {
		if mm.Name() == name {
			return &mm
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
