package golist

import (
	"io"
	"io/ioutil"
	"runtime"
)

type config struct {
	optGOARCH      string    // GOARCH value
	optGOOS        string    // GOOS value
	optCgoEnabled  string    // CGO_ENABLED value
	errorWriter    io.Writer //Stderr
	flagHideWindow bool      // HideWindow flag (Windows only)
}

//OptFunc is self-referential function for functional options pattern
type OptEnv func(*config)

func newconfig(opts ...OptEnv) *config {
	cl := &config{
		optGOARCH:      runtime.GOARCH,
		optGOOS:        runtime.GOOS,
		optCgoEnabled:  "1",
		errorWriter:    ioutil.Discard,
		flagHideWindow: false,
	}
	for _, opt := range opts {
		opt(cl)
	}
	return cl
}

func (cl *config) GetEnv() []string {
	e := []string{}
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
	return func(cl *config) { cl.optGOARCH = s }
}

//WithGOOS returns setter function as type OptEnv
func WithGOOS(s string) OptEnv {
	return func(cl *config) { cl.optGOOS = s }
}

//WithCGOENABLED returns setter function as type OptEnv
func WithCGOENABLED(s string) OptEnv {
	return func(cl *config) { cl.optCgoEnabled = s }
}

//WithCGOENABLED returns setter function as type OptEnv
func WithErrorWriter(w io.Writer) OptEnv {
	return func(cl *config) { cl.errorWriter = w }
}

//WithCGOENABLED returns setter function as type OptEnv
func EnableHideWindow() OptEnv {
	return func(cl *config) { cl.flagHideWindow = true }
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
