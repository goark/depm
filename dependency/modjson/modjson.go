package modjson

import (
	"encoding/json"
	"sort"

	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/dotenc"
	"github.com/spiegel-im-spiegel/depm/modules"
	"github.com/spiegel-im-spiegel/errs"
)

type edgeJSON struct {
	Module *nodeJSON
	Deps   []*nodeJSON `json:",omitempty"`
}
type nodeJSON struct {
	Path     string
	Packages []string `json:",omitempty"`
	Main     bool     `json:",omitempty"`
}

func newNodeJSON(mod *modules.Module) *nodeJSON {
	node := &nodeJSON{
		Path: mod.Name.String(),
		Main: mod.Main,
	}
	if len(mod.Packages) > 0 {
		node.Packages = make([]string, len(mod.Packages), cap(mod.Packages))
		copy(node.Packages, mod.Packages)
		sort.SliceStable(node.Packages, func(i, j int) bool {
			return node.Packages[i] < node.Packages[j]
		})

	}
	return node
}

func newEdgeJSON(deps []*dependency.EdgeModule) []edgeJSON {
	nj := []edgeJSON{}
	for _, n := range deps {
		nd := edgeJSON{Module: newNodeJSON(n.Module), Deps: []*nodeJSON{}}
		for _, m := range n.Deps {
			nd.Deps = append(nd.Deps, newNodeJSON(m))
		}
		nj = append(nj, nd)
	}
	return nj
}

//EncodeJSON returns JSON formatted text from Node slice.
func Encode(deps []*dependency.EdgeModule) ([]byte, error) {
	return json.Marshal(newEdgeJSON(deps))
}

func EncodeDot(deps []*dependency.EdgeModule, conf string) (string, error) {
	ejs := newEdgeJSON(deps)
	ds := []*dotenc.Dep{}
	for _, ej := range ejs {
		for _, d := range ej.Deps {
			ds = append(ds, dotenc.NewDep(ej.Module.Path, d.Path))
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
