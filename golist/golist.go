package golist

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/spiegel-im-spiegel/errs"
)

type cmdLine struct {
	argument      string
	optGOARCH     string
	optGOOS       string
	optCgoEnabled string
	errorWriter   io.Writer
}

//OptFunc is self-referential function for functional options pattern
type OptEnv func(*cmdLine)

func newCmdLine(arg string, opts ...OptEnv) *cmdLine {
	cl := &cmdLine{argument: arg, optGOARCH: runtime.GOARCH, optGOOS: runtime.GOOS, optCgoEnabled: "1", errorWriter: ioutil.Discard}
	for _, opt := range opts {
		opt(cl)
	}
	return cl
}

func (cl *cmdLine) GetEnv() []string {
	e := os.Environ()
	if len(cl.optGOARCH) > 0 {
		e = append(e, "GOARCH="+cl.optGOARCH)
	}
	if len(cl.optGOOS) > 0 {
		e = append(e, "GOOS="+cl.optGOOS)
	}
	if len(cl.optCgoEnabled) > 0 {
		e = append(e, "CGO_ENABLED="+cl.optCgoEnabled)
	}
	return e
}

//WithGOARCH returns setter function as type OptEnv
func WithGOARCH(s string) OptEnv {
	return func(cl *cmdLine) { cl.optGOARCH = s }
}

//WithGOOS returns setter function as type OptEnv
func WithGOOS(s string) OptEnv {
	return func(cl *cmdLine) { cl.optGOOS = s }
}

//WithCGOENABLED returns setter function as type OptEnv
func WithCGOENABLED(s string) OptEnv {
	return func(cl *cmdLine) { cl.optCgoEnabled = s }
}

//WithCGOENABLED returns setter function as type OptEnv
func WithErrorWriter(w io.Writer) OptEnv {
	return func(cl *cmdLine) { cl.errorWriter = w }
}

//GetPackagesRaw returns package information by JSON string
func GetPackagesRaw(ctx context.Context, name string, opts ...OptEnv) ([]byte, error) {
	cl := newCmdLine(name, opts...)
	cmd := exec.CommandContext(ctx, "go", "list", "-json", cl.argument)
	cmd.Env = cl.GetEnv()
	cmd.Stderr = cl.errorWriter
	b, err := cmd.Output()
	return b, errs.Wrap(err, errs.WithContext("name", name))
}

func GetPackages(ctx context.Context, name string, opts ...OptEnv) ([]Package, error) {
	b, err := GetPackagesRaw(ctx, name, opts...)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("name", name))
	}
	decoder := json.NewDecoder(bytes.NewReader(b))
	ps := []Package{}
	for {
		var p Package
		if err := decoder.Decode(&p); err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return ps, errs.Wrap(err, errs.WithContext("name", name))
		}
		ps = append(ps, p)
	}
	return ps, nil
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
