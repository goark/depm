package dotenc

import (
	"fmt"
	"io"
	"strings"

	"github.com/emicklei/dot"
	"github.com/spiegel-im-spiegel/errs"
)

//Ctx is context class for parsing
type Ctx struct {
	graph   *dot.Graph
	config  *Config
	count   int
	mapNode map[string]string
}

//New returns new Cxt instance
func New(conf string) (*Ctx, error) {
	//parse config file
	cf, err := DecodeConfig(conf)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("conf", conf))
	}
	return &Ctx{
		graph:   dot.NewGraph(dot.Directed).ID("G"),
		config:  cf,
		count:   0,
		mapNode: map[string]string{},
	}, nil
}

//ImportDeps method makes graph data
func (ctx *Ctx) ImportDeps(deps ...*Dep) *Ctx {
	for _, dep := range deps {
		ctx.addEdgeAttr(ctx.getNode(dep.From).Edge(ctx.getNode(dep.To)))
	}
	return ctx
}

//Write method writes to io.Writer
func (ctx *Ctx) Write(w io.Writer) {
	ctx.graph.Write(w)
}

func (ctx *Ctx) String() string {
	return ctx.graph.String()
}

func (ctx *Ctx) getNode(label string) dot.Node {
	if ctx == nil {
		return dot.Node{}
	}
	if n, ok := ctx.mapNode[label]; ok {
		return ctx.graph.Node(n)
	}
	ctx.count++
	n := fmt.Sprintf("N%d", ctx.count)
	ctx.mapNode[label] = n
	return ctx.addNodeAttr(ctx.graph.Node(n)).Attr("label", strings.ReplaceAll(label, "@", "\n"))
}

func (ctx *Ctx) addNodeAttr(n dot.Node) dot.Node {
	if ctx == nil || ctx.config == nil {
		return n
	}
	for k, v := range ctx.config.Node {
		n.Attr(k, v)
	}
	return n
}

func (ctx *Ctx) addEdgeAttr(e dot.Edge) dot.Edge {
	if ctx == nil || ctx.config == nil {
		return e
	}
	for k, v := range ctx.config.Edge {
		e.Attr(k, v)
	}
	return e
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
