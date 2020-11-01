package dotenc

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spiegel-im-spiegel/errs"
)

//Config is configuration class
type Config struct {
	Node map[string]interface{} `toml:"node"`
	Edge map[string]interface{} `toml:"edge"`
}

//Decode returns Config instance from stream
func DecodeConfig(path string) (*Config, error) {
	if len(path) == 0 {
		return &Config{Node: map[string]interface{}{}, Edge: map[string]interface{}{}}, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("path", path))
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("path", path))
	}
	c := &Config{}
	if err := toml.Unmarshal(data, c); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("path", path))
	}
	return c, nil
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
