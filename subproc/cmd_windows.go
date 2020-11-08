package subproc

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/spiegel-im-spiegel/errs"
)

func (c *Cmd) newExecCmd() (*exec.Cmd, error) {
	path, err := LookPath(c.name)
	if err != nil {
		return nil, errs.Wrap(
			err,
			errs.WithContext("os", runtime.GOOS),
			errs.WithContext("name", c.name),
			errs.WithContext("path", path),
			errs.WithContext("args", c.args),
			errs.WithContext("env", c.env),
			errs.WithContext("hideWindow", c.hideWindow),
		)
	}
	var cmd *exec.Cmd
	if c.ctx != nil {
		cmd = exec.CommandContext(c.ctx, path, c.args...)
	} else {
		cmd = exec.Command(path, c.args...)
	}
	if c.reader != nil {
		cmd.Stdin = c.reader
	}
	if c.errorWriter != nil {
		cmd.Stderr = c.errorWriter
	}
	if len(c.env) > 0 {
		cmd.Env = append(os.Environ(), c.env...)
	}
	if c.hideWindow == true { //Windows only
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	return cmd, nil
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
