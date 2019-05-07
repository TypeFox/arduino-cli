/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO SA (http://www.arduino.cc/)
 *
 * This software is released under the GNU General Public License version 3,
 * which covers the main part of arduino-cli.
 * The terms of this license can be found at:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * You can be released from the requirements of the above licenses by purchasing
 * a commercial license. Buying such a license is mandatory if you want to modify or
 * otherwise use the software for commercial activities involving the Arduino
 * software without disclosing the source code of your own applications. To purchase
 * a commercial license, send an email to license@arduino.cc.
 */

package lib

import (
	"context"

	"github.com/arduino/arduino-cli/arduino/libraries/librariesindex"
	"github.com/arduino/arduino-cli/arduino/libraries/librariesmanager"
	"github.com/arduino/arduino-cli/commands"
	"github.com/arduino/arduino-cli/common/formatter/output"
	"github.com/arduino/arduino-cli/rpc"
	log "github.com/sirupsen/logrus"
)

// ListLibraries returns the list of installed libraries. If updatable is true it
// returns only the libraries that may be updated.
func ListLibraries(lm *librariesmanager.LibrariesManager, updatable bool) *output.InstalledLibraries {
	res := &output.InstalledLibraries{}
	for _, libAlternatives := range lm.Libraries {
		for _, lib := range libAlternatives.Alternatives {
			var available *librariesindex.Release
			if updatable {
				available = lm.Index.FindLibraryUpdate(lib)
				if available == nil {
					continue
				}
			}
			res.Libraries = append(res.Libraries, &output.InstalledLibary{
				Library:   lib,
				Available: available,
			})
		}
	}
	return res
}

func List(ctx context.Context, req *rpc.LibraryListReq) (*rpc.LibraryListResp, error) {
	log.SetLevel(log.DebugLevel)

	lm := commands.GetLibraryManager(req)

	res := make([]*rpc.InstalledLibrary, 0)
	for _, libAlternatives := range lm.Libraries {
		for _, lib := range libAlternatives.Alternatives {
			librelease := &rpc.LibraryRelease{
				Author:        lib.Author,
				Version:       lib.Version.String(),
				Maintainer:    lib.Maintainer,
				Sentence:      lib.Sentence,
				Paragraph:     lib.Paragraph,
				Website:       lib.Website,
				Category:      lib.Category,
				Architectures: lib.Architectures,
				Types:         lib.Types,
			}

			res = append(res, &rpc.InstalledLibrary{
				Name:      lib.Name,
				Installed: librelease,
			})
		}
	}
	return &rpc.LibraryListResp{
		Libraries: res,
	}, nil
}
