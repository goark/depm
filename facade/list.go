package facade

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/depm/facade/lstjson"
	"github.com/spiegel-im-spiegel/depm/golist"
	"github.com/spiegel-im-spiegel/depm/list"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gocli/signal"
)

//newModuleCmd returns cobra.Command instance for show sub-command
func newListCmd(ui *rwi.RWI) *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list [flags] [package import path]",
		Aliases: []string{"lst", "l"},
		Short:   "list modules",
		Long:    "list modules, compatible 'go list -m' command",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			jsonFlag, err := cmd.Flags().GetBool("json")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --json option", errs.WithCause(err)))
			}
			updFlag, err := cmd.Flags().GetBool("check-update")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --check-update option", errs.WithCause(err)))
			}

			//package path
			path := "all" //local all packages
			if len(args) > 0 {
				path = args[0]
			}

			//Run command
			ms, err := list.ImportModules(
				signal.Context(context.Background(), os.Interrupt),
				path,
				updFlag,
				golist.WithGOARCH(goarchString),
				golist.WithGOOS(goosString),
				golist.WithCGOENABLED(cgoenabledString),
				golist.WithErrorWriter(ui.ErrorWriter()),
			)
			if err != nil {
				return debugPrint(ui, errs.Wrap(
					err,
					errs.WithContext("jsonFlag", jsonFlag),
					errs.WithContext("updFlag", updFlag),
					errs.WithContext("path", path),
				))
			}

			//Output
			if jsonFlag {
				b, err := lstjson.Encode(ms)
				if err != nil {
					return debugPrint(ui, errs.Wrap(
						err,
						errs.WithContext("jsonFlag", jsonFlag),
						errs.WithContext("updFlag", updFlag),
						errs.WithContext("path", path),
					))
				}
				return ui.OutputBytes(b)
			} else {
				b, err := lstjson.EncodeText(ms)
				if err != nil {
					return debugPrint(ui, errs.Wrap(
						err,
						errs.WithContext("jsonFlag", jsonFlag),
						errs.WithContext("updFlag", updFlag),
						errs.WithContext("path", path),
					))
				}
				return ui.OutputBytes(b)
			}
		},
	}
	listCmd.Flags().BoolP("json", "j", false, "output by JSON format")
	listCmd.Flags().BoolP("check-update", "u", false, "check updating module")

	return listCmd
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
