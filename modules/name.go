package modules

import "fmt"

//Name is module name class
type Name struct {
	Path    string
	Version string
}

func newName(path, ver string) Name {
	return Name{Path: path, Version: ver}
}

func (nm Name) String() string {
	if len(nm.Version) > 0 {
		return fmt.Sprintf("%s@%s", nm.Path, nm.Version)
	}
	return nm.Path
}

//EqualAll method returns true if equals path and version
func (nm Name) EqualAll(path, ver string) bool {
	return nm.Path == path && nm.Version == ver
}

//Equal method returns true if left == right
func (left Name) Equal(right Name) bool {
	return left.EqualAll(right.Path, right.Version)
}

//Equal method returns true if left == right
func (nm Name) IsZero() bool {
	return len(nm.Path) == 0
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
