package pkgjson

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/dotenc"
	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/packages"
	"github.com/spiegel-im-spiegel/errs"
)

type nodeJSON struct {
	Package *packageJSON
	Deps    []*edgeJSON `json:",omitempty"`
}
type edgeJSON struct {
	Package    *packageJSON
	IsUnsafe   bool `json:",omitempty"`
	IsInternal bool `json:",omitempty"`
}
type packageJSON struct {
	ImportPath string
	Module     *moduleJSON `json:",omitempty"`
}
type moduleJSON struct {
	Path    string `json:",omitempty"`
	Version string `json:",omitempty"`
}

//EncodeJSON returns JSON formatted text from Node slice.
func Encode(deps []*dependency.NodePackage) ([]byte, error) {
	return json.Marshal(newNodeJSON(deps))
}

func EncodeDot(deps []*dependency.NodePackage, conf string) (string, error) {
	ejs := newNodeJSON(deps)
	ds := []*dotenc.Dep{}
	for _, ej := range ejs {
		if len(ej.Deps) > 0 {
			for _, d := range ej.Deps {
				ds = append(ds, dotenc.NewDep(ej.Package.ImportPath, d.Package.ImportPath))
			}
		} else {
			ds = append(ds, dotenc.NewDep(ej.Package.ImportPath, ""))
		}
	}
	dot, err := dotenc.New(conf)
	if err != nil {
		return "", errs.Wrap(err, errs.WithContext("conf", conf))
	}
	return dot.ImportDeps(ds...).String(), nil
}

func newPackageJSON(p *packages.Package) *packageJSON {
	return &packageJSON{ImportPath: p.Path}
}

func newModuleJSON(m *golist.Module) *moduleJSON {
	return &moduleJSON{Path: m.Path, Version: m.Version}
}

func newNodeJSON(deps []*dependency.NodePackage) []nodeJSON {
	nj := []nodeJSON{}
	for _, n := range deps {
		nd := nodeJSON{Package: newPackageJSON(n.Package), Deps: []*edgeJSON{}}
		if mod := n.Package.Contained; mod != nil {
			nd.Package.Module = newModuleJSON(mod)
		}
		for _, p := range n.Deps {
			edge := &edgeJSON{Package: newPackageJSON(p), IsUnsafe: p.IsUnsafe(), IsInternal: p.IsInternal()}
			if mod := p.Contained; mod != nil {
				edge.Package.Module = newModuleJSON(mod)
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
