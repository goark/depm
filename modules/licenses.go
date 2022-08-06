package modules

import (
	"os"
	"path/filepath"
	"regexp"

	licenseclassifier "github.com/google/licenseclassifier/v2"
	"github.com/google/licenseclassifier/v2/assets"
)

var (
	licenseRegexp = regexp.MustCompile(`^(?i)(LICEN(S|C)E|COPYING|README|NOTICE)(\..+)?$`)
	classifier    *licenseclassifier.Classifier
)

func init() {
	classifier, _ = assets.DefaultClassifier()
}

func FindLicense(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}
	for _, f := range files {
		if licenseRegexp.MatchString(f.Name()) {
			path := filepath.Join(dir, f.Name())
			if content, err := os.ReadFile(path); err == nil {
				if res := classifier.Match(content); len(res.Matches) > 0 {
					return res.Matches[0].Name
				}
			}
		}
	}
	return ""
}

/* Copyright 2021-2022 Spiegel
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
