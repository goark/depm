package pkgjson

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/dotenc"
	"github.com/spiegel-im-spiegel/errs"
)

type edgeJSON struct {
	Package packageJSON
	Deps    []*nodeJSON `json:",omitempty"`
}
type packageJSON struct {
	ImportPath string
	Root       bool        `json:",omitempty"`
	Module     *moduleJSON `json:",omitempty"`
}
type moduleJSON struct {
	Path    string
	Version string
}
type nodeJSON struct {
	Package    packageJSON
	IsUnsafe   bool `json:",omitempty"`
	IsInternal bool `json:",omitempty"`
}

//EncodeJSON returns JSON formatted text from Node slice.
func Encode(deps []*dependency.EdgePackage) ([]byte, error) {
	return json.Marshal(newEdgeJSON(deps))
}

func EncodeDot(deps []*dependency.EdgePackage, conf string) (string, error) {
	ejs := newEdgeJSON(deps)
	ds := []*dotenc.Dep{}
	for _, ej := range ejs {
		for _, d := range ej.Deps {
			ds = append(ds, dotenc.NewDep(ej.Package.ImportPath, d.Package.ImportPath))
		}
	}
	dot, err := dotenc.New(conf)
	if err != nil {
		return "", errs.Wrap(err, errs.WithContext("conf", conf))
	}
	return dot.ImportDeps(ds...).String(), nil
}

func newEdgeJSON(deps []*dependency.EdgePackage) []edgeJSON {
	nj := []edgeJSON{}
	for _, n := range deps {
		nd := edgeJSON{Package: packageJSON{ImportPath: n.Package.Path, Root: n.Package.Root}, Deps: []*nodeJSON{}}
		if mod := n.Package.Contained; mod != nil {
			nd.Package.Module = &moduleJSON{Path: mod.Path, Version: mod.Version}
		}
		for _, p := range n.Deps {
			edge := &nodeJSON{Package: packageJSON{ImportPath: p.Path, Root: p.Root}, IsUnsafe: p.IsUnsafe(), IsInternal: p.IsInternal()}
			if mod := p.Contained; mod != nil {
				edge.Package.Module = &moduleJSON{Path: mod.Path, Version: mod.Version}
			}
			nd.Deps = append(nd.Deps, edge)
		}
		nj = append(nj, nd)
	}
	return nj
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
