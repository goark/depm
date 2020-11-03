package modjson

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/dotenc"
	"github.com/spiegel-im-spiegel/depm/modules"
	"github.com/spiegel-im-spiegel/errs"
)

type nodeJSON struct {
	Module *moduleJSON
	Deps   []*moduleJSON `json:",omitempty"`
}
type moduleJSON struct {
	Path     string
	Main     bool     `json:",omitempty"`
	Latest   string   `json:",omitempty"`
	Packages []string `json:",omitempty"`
}

func (nj *moduleJSON) label() string {
	name := nj.Path
	if len(nj.Latest) > 0 {
		name = fmt.Sprintf("%s (latest %s)", name, nj.Latest)
	}
	return name
}

func newModuleJSON(m *modules.Module) *moduleJSON {
	mod := &moduleJSON{
		Path: m.Name.String(),
		Main: m.Main,
	}
	if !m.Update.IsZero() {
		mod.Latest = m.Update.Version
	}
	if len(m.Packages) > 0 {
		mod.Packages = make([]string, len(m.Packages), cap(m.Packages))
		copy(mod.Packages, m.Packages)
		sort.SliceStable(mod.Packages, func(i, j int) bool {
			return mod.Packages[i] < mod.Packages[j]
		})

	}
	return mod
}

func newNodeJSON(deps []*dependency.NodeModule) []nodeJSON {
	nj := []nodeJSON{}
	for _, n := range deps {
		nd := nodeJSON{Module: newModuleJSON(n.Module), Deps: []*moduleJSON{}}
		for _, m := range n.Deps {
			nd.Deps = append(nd.Deps, newModuleJSON(m))
		}
		nj = append(nj, nd)
	}
	return nj
}

//EncodeJSON returns JSON formatted text from Node slice.
func Encode(deps []*dependency.NodeModule) ([]byte, error) {
	return json.Marshal(newNodeJSON(deps))
}

func EncodeDot(deps []*dependency.NodeModule, conf string) (string, error) {
	ejs := newNodeJSON(deps)
	ds := []*dotenc.Dep{}
	for _, ej := range ejs {
		if len(ej.Deps) > 0 {
			for _, d := range ej.Deps {
				ds = append(ds, dotenc.NewDep(ej.Module.label(), d.label()))
			}
		} else {
			ds = append(ds, dotenc.NewDep(ej.Module.label(), ""))
		}
	}
	dot, err := dotenc.New(conf)
	if err != nil {
		return "", errs.Wrap(err, errs.WithContext("conf", conf))
	}
	return dot.ImportDeps(ds...).String(), nil
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
