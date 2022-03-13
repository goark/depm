package subproc

import (
	"context"
	"io"

	"github.com/goark/errs"
)

//Cmd is context of subprocess
type Cmd struct {
	name        string          // command name
	args        []string        //command argments
	ctx         context.Context // context
	env         []string        //additional environment variables
	reader      io.Reader       // stdin
	errorWriter io.Writer       // stderr
	hideWindow  bool            // HideWindow attribution
}

//New creates new Cmd instance
func New(name string, arg ...string) *Cmd {
	return &Cmd{
		name: name,
		args: arg,
		env:  []string{},
	}
}

//WithContext method sets context element
func (c *Cmd) WithContext(ctx context.Context) *Cmd {
	c.ctx = ctx
	return c
}

//AddEnv method adds environment variables element
func (c *Cmd) AddEnv(s ...string) *Cmd {
	c.env = append(c.env, s...)
	return c
}

//WithStdin method sets stdin element
func (c *Cmd) WithStdin(r io.Reader) *Cmd {
	c.reader = r
	return c
}

//WithStderr method sets stderr element
func (c *Cmd) WithStderr(w io.Writer) *Cmd {
	c.errorWriter = w
	return c
}

//HideWindow method sets HideWindow attribution flag element (Windows only)
func (c *Cmd) HideWindow(enable bool) *Cmd {
	c.hideWindow = enable
	return c
}

//Output method runs subprocess and output result
func (c *Cmd) Output() ([]byte, error) {
	cmd, err := c.newExecCmd()
	if err != nil {
		return nil, errs.Wrap(
			err,
			errs.WithContext("name", c.name),
			errs.WithContext("args", c.args),
			errs.WithContext("env", c.env),
		)
	}
	b, err := cmd.Output()
	return b, errs.Wrap(
		err,
		errs.WithContext("cmd.Path", cmd.Path),
		errs.WithContext("cmd.Args", cmd.Args),
		errs.WithContext("cmd.Env", cmd.Env),
	)
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
