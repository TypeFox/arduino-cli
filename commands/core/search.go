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

package core

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/arduino/arduino-cli/arduino/cores"
	"github.com/arduino/arduino-cli/commands"
	"github.com/arduino/arduino-cli/rpc"
)

type platformSearchResult struct {
	Release *cores.PlatformRelease
	Package *cores.Package
}

func PlatformSearch(ctx context.Context, req *rpc.PlatformSearchReq) (*rpc.PlatformSearchResp, error) {
	pm := commands.GetPackageManager(req)
	if pm == nil {
		return nil, errors.New("invalid instance")
	}

	search := req.SearchArgs

	res := make([]platformSearchResult, 0)
	if isUsb, _ := regexp.MatchString("[0-9a-f]{4}:[0-9a-f]{4}", search); isUsb {
		vid, pid := search[:4], search[5:]
		boards := pm.FindPlatformReleaseProvidingBoardsWithVidPid(vid, pid)
		for _, b := range boards {
			res = append(res, platformSearchResult{Release: b})
		}
	} else {
		match := func(line string) bool {
			return strings.Contains(strings.ToLower(line), search)
		}
		for _, targetPackage := range pm.GetPackages().Packages {
			for _, platform := range targetPackage.Platforms {
				platformRelease := platform.GetLatestRelease()
				if platformRelease == nil {
					continue
				}
				if match(platform.Name) || match(platform.Architecture) {
					res = append(res, platformSearchResult{
						Release: platformRelease,
						Package: targetPackage,
					})
					continue
				}
				for _, board := range platformRelease.BoardsManifest {
					if match(board.Name) {
						res = append(res, platformSearchResult{
							Release: platformRelease,
							Package: targetPackage,
						})
						break
					}
				}
			}
		}
	}

	out := []*rpc.SearchOutput{}
	for _, r := range res {
		plt := &rpc.SearchOutput{
			ID:       r.Release.Platform.String(),
			Name:     r.Release.Platform.Name,
			Version:  r.Release.Version.String(),
			Sentence: "Boards included in this package: ",
		}

		i := 0
		boardNames := make([]string, len(r.Release.Boards))
		for _, b := range r.Release.Boards {
			boardNames[i] = b.Name()
			i++
		}

		plt.Paragragh = strings.Join(boardNames, ", ")
		if r.Package != nil {
			plt.Author = r.Package.Maintainer
		}
		out = append(out, plt)
	}
	return &rpc.PlatformSearchResp{SearchOutput: out}, nil
}
