package facade

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/dependency/modjson"
	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/modules"
	"github.com/spiegel-im-spiegel/depm/packages"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newModuleCmd returns cobra.Command instance for show sub-command
func newModuleCmd(ui *rwi.RWI) *cobra.Command {
	moduleCmd := &cobra.Command{
		Use:     "module [flags] [package import path]",
		Aliases: []string{"mod", "m"},
		Short:   "Visualize depndency modules",
		Long:    "Visualize depndency modules.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//package path
			path := "all" //local all packages
			if len(args) > 0 {
				path = args[0]
			}

			//Run command
			ps, err := packages.ImportPackages(
				context.Background(),
				path,
				golist.WithGOARCH(goarchString),
				golist.WithGOOS(goosString),
				golist.WithCGOENABLED(cgoenabledString),
				golist.WithErrorWriter(ui.ErrorWriter()),
			)
			if err != nil {
				return debugPrint(ui, errs.Wrap(
					err,
					errs.WithContext("path", path),
				))
			}
			ds := dependency.NewModules(modules.ImportModules(ps))
			if dotFlag {
				s, err := modjson.EncodeDot(ds, dotConfFile)
				if err != nil {
					return debugPrint(ui, errs.Wrap(
						err,
						errs.WithContext("path", path),
					))
				}
				return ui.Outputln(s)
			} else {
				b, err := modjson.Encode(ds)
				if err != nil {
					return debugPrint(ui, errs.Wrap(
						err,
						errs.WithContext("path", path),
					))
				}
				return ui.OutputBytes(b)
			}
		},
	}

	return moduleCmd
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
