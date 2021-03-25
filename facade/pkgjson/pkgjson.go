package pkgjson

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/dotenc"
	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/modules"
	"github.com/spiegel-im-spiegel/depm/packages"
	"github.com/spiegel-im-spiegel/errs"
)

type nodeJSON struct {
	Package *packageJSON
	Deps    []*packageJSON `json:",omitempty"`
}
type packageJSON struct {
	ImportPath string
	Internal   bool        `json:",omitempty"`
	CGO        bool        `json:",omitempty"`
	Unsafe     bool        `json:",omitempty"`
	Module     *moduleJSON `json:",omitempty"`
}
type moduleJSON struct {
	Path    string `json:",omitempty"`
	Version string `json:",omitempty"`
	License string `json:",omitempty"`
}

//Encode returns JSON formatted text from Node slice.
func Encode(deps []*dependency.NodePackage) ([]byte, error) {
	return json.Marshal(newNodeJSON(deps))
}

//EncodeDot returns DOT lnguage formatted text from Node slice.
func EncodeDot(deps []*dependency.NodePackage, conf string) (string, error) {
	ejs := newNodeJSON(deps)
	ds := []*dotenc.Dep{}
	for _, ej := range ejs {
		if len(ej.Deps) > 0 {
			for _, d := range ej.Deps {
				ds = append(ds, dotenc.NewDep(ej.Package.ImportPath, d.ImportPath))
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

func newNodeJSON(deps []*dependency.NodePackage) []nodeJSON {
	nj := []nodeJSON{}
	for _, n := range deps {
		nd := nodeJSON{Package: newPackageJSON(n.Package), Deps: []*packageJSON{}}
		if mod := n.Package.Contained; mod != nil {
			nd.Package.Module = newModuleJSON(mod)
		}
		for _, p := range n.Deps {
			edge := newPackageJSON(p)
			if mod := p.Contained; mod != nil {
				edge.Module = newModuleJSON(mod)
			}
			nd.Deps = append(nd.Deps, edge)
		}
		nj = append(nj, nd)
	}
	return nj
}

func newPackageJSON(p *packages.Package) *packageJSON {
	return &packageJSON{ImportPath: p.Path, Internal: p.IsInternal(), CGO: p.UseCGO, Unsafe: p.UseUnsafe}
}

func newModuleJSON(m *golist.Module) *moduleJSON {
	var license string
	if len(m.Dir) > 0 {
		license = modules.FindLicense(m.Dir)
	}
	return &moduleJSON{Path: m.Path, Version: m.Version, License: license}
}

/* Copyright 2020-2021 Spiegel
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
