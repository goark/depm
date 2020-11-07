package lstjson

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/errs"
)

//Encode returns JSON formatted text from Module list.
func Encode(ms []*golist.Module) ([]byte, error) {
	buf := &bytes.Buffer{}
	je := json.NewEncoder(buf)
	je.SetIndent("", "\t")
	for _, m := range ms {
		if err := je.Encode(m); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	return buf.Bytes(), nil
}

//EncodeText returns plain text from Module list with compatible 'go list -m' format.
func EncodeText(ms []*golist.Module) ([]byte, error) {
	buf := &bytes.Buffer{}
	for _, m := range ms {
		fmt.Fprintln(buf, m)
	}
	return buf.Bytes(), nil
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
