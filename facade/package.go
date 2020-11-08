package facade

import (
	"context"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/depm/dependency"
	"github.com/spiegel-im-spiegel/depm/facade/pkgjson"
	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/packages"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gocli/signal"
)

//newPackageCmd returns cobra.Command instance for show sub-command
func newPackageCmd(ui *rwi.RWI) *cobra.Command {
	packageCmd := &cobra.Command{
		Use:     "package [flags] [package import path]",
		Aliases: []string{"pkg", "p"},
		Short:   "analyze depndency packages",
		Long:    "analyze depndency packages.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			dotFlag, err := cmd.Flags().GetBool("dot")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --dot option", errs.WithCause(err)))
			}
			dotConfFile, err := cmd.Flags().GetString("dot-config")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --dot-config option", errs.WithCause(err)))
			}
			includeInternal, err := cmd.Flags().GetBool("include-internal")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --include-internal option", errs.WithCause(err)))
			}
			includeStd, err := cmd.Flags().GetBool("include-standard")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --include-standard option", errs.WithCause(err)))
			}

			//package path
			path := "all" //local all packages
			if len(args) > 0 {
				path = args[0]
			}

			//Run command
			ps, err := packages.ImportPackages(
				signal.Context(context.Background(), os.Interrupt, syscall.SIGTERM),
				golist.New(
					golist.WithGOARCH(goarchString),
					golist.WithGOOS(goosString),
					golist.WithCGOENABLED(cgoenabledString),
					golist.WithErrorWriter(ui.ErrorWriter()),
					golist.EnableHideWindow(), //Windows only
				),
				path,
			)
			if err != nil {
				return debugPrint(ui, errs.Wrap(
					err,
					errs.WithContext("path", path),
					errs.WithContext("dotFlag", dotFlag),
					errs.WithContext("dotConfFile", dotConfFile),
					errs.WithContext("includeInternal", includeInternal),
					errs.WithContext("includeStd", includeStd),
				))
			}
			ds := dependency.NewPackages(ps, includeStd, includeInternal)

			//Output
			if dotFlag {
				s, err := pkgjson.EncodeDot(ds, dotConfFile)
				if err != nil {
					return debugPrint(ui, errs.Wrap(
						err,
						errs.WithContext("path", path),
						errs.WithContext("dotFlag", dotFlag),
						errs.WithContext("dotConfFile", dotConfFile),
						errs.WithContext("includeInternal", includeInternal),
						errs.WithContext("includeStd", includeStd),
					))
				}
				return ui.Outputln(s)
			} else {
				b, err := pkgjson.Encode(ds)
				if err != nil {
					return debugPrint(ui, errs.Wrap(
						err,
						errs.WithContext("path", path),
						errs.WithContext("dotFlag", dotFlag),
						errs.WithContext("dotConfFile", dotConfFile),
						errs.WithContext("includeInternal", includeInternal),
						errs.WithContext("includeStd", includeStd),
					))
				}
				return ui.OutputBytes(b)
			}
		},
	}
	packageCmd.Flags().BoolP("dot", "", false, "output by DOT language")
	packageCmd.Flags().StringP("dot-config", "", "", "config file for DOT language")
	packageCmd.Flags().BoolP("include-internal", "i", false, "include internal packages")
	packageCmd.Flags().BoolP("include-standard", "s", false, "include standard Go library")

	return packageCmd
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
