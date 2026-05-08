package dotenc

import (
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/goark/errs"
)

// Config is configuration class
type Config struct {
	Node map[string]interface{} `toml:"node"`
	Edge map[string]interface{} `toml:"edge"`
}

// Decode returns Config instance from stream
func DecodeConfig(path string) (cfg *Config, err error) {
	if len(path) == 0 {
		cfg = &Config{Node: map[string]interface{}{}, Edge: map[string]interface{}{}}
		return
	}
	file, ferr := os.Open(filepath.Clean(path))
	if ferr != nil {
		err = errs.Wrap(ferr, errs.WithContext("path", path))
		return
	}
	defer func() {
		err = errs.Join(err, file.Close())
	}()

	data, ferr := io.ReadAll(file)
	if ferr != nil {
		err = errs.Wrap(ferr, errs.WithContext("path", path))
		return
	}
	cfg = &Config{}
	if terr := toml.Unmarshal(data, cfg); terr != nil {
		err = errs.Wrap(terr, errs.WithContext("path", path))
		return
	}
	return
}

/* Copyright 2020-2026 Spiegel
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
