package modules

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/google/licenseclassifier"
	"github.com/google/licenseclassifier/stringclassifier"
)

var (
	licenseRegexp = regexp.MustCompile(`^(?i)(LICEN(S|C)E|COPYING|README|NOTICE)(\..+)?$`)
	classifier    *licenseclassifier.License
)

func init() {
	classifier, _ = licenseclassifier.New(licenseclassifier.DefaultConfidenceThreshold)
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
				matches := func() stringclassifier.Matches {
					defer func() {
						if r := recover(); r != nil {
							return
						}
					}()
					return classifier.MultipleMatch(string(content), true)
				}()
				if len(matches) > 0 {
					return matches[0].Name
				}
			}
		}
	}
	return ""
}

/* Copyright 2021 Spiegel
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
