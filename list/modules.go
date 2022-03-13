package list

import "github.com/goark/depm/golist"

//Modules is list of Modules.
type Modules struct {
	list map[string]*golist.Module
}

func newModules() *Modules {
	return &Modules{list: map[string]*golist.Module{}}
}

//Set method sets Module instance to Modules
func (ms *Modules) Set(m *golist.Module) {
	if ms == nil || len(m.Name()) == 0 {
		return
	}
	ms.list[m.Name()] = m
}

//Get method gets Module instance form Modules
func (ms *Modules) Get(name string) *golist.Module {
	if m, ok := ms.list[name]; ok {
		return m
	}
	return nil
}

//Get method gets Module instance form Modules
func (ms *Modules) List() []*golist.Module {
	lst := []*golist.Module{}
	for _, v := range ms.list {
		lst = append(lst, v)
	}
	return lst
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
