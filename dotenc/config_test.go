package dotenc

import (
	"testing"
)

func TestGgm(t *testing.T) {
	testCases := []struct {
		inp      string
		fontname string
		color    string
	}{
		{inp: "./sample.toml", fontname: "Inconsolata", color: "red"},
		{inp: "", fontname: "", color: ""},
	}
	for _, tc := range testCases {
		res, err := DecodeConfig(tc.inp)
		if err != nil {
			t.Errorf("Decode() err = \"%v\", want nil.", err)
		}
		fn := getValueFrom(res.Node, "fontname")
		if fn != tc.fontname {
			t.Errorf("res.Node[\"fontname\"] = \"%v\", want \"%v\".", fn, tc.fontname)
		}
		clr := getValueFrom(res.Edge, "color")
		if clr != tc.color {
			t.Errorf("res.Node[\"color\"] = \"%v\", want \"%v\".", clr, tc.color)
		}
	}
}

func getValueFrom(m map[string]interface{}, k string) string {
	if v, ok := m[k]; ok {
		if s, ok := v.(string); ok {
			return string(s)
		}
		return ""
	}
	return ""
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
