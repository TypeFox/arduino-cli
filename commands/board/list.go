package board

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arduino/arduino-cli/arduino/cores/packagemanager"
	"github.com/arduino/arduino-cli/commands"
	"github.com/arduino/arduino-cli/rpc"
	discovery "github.com/arduino/board-discovery"
)

// List lists all attached boards and matches them with installed platforms
func List(ctx context.Context, req *rpc.BoardListReq) (*rpc.BoardListResp, error) {
	pm := commands.GetPackageManager(req)
	if pm == nil {
		return nil, errors.New("invalid instance")
	}

	monitor := discovery.New(time.Millisecond)

	// This is a bit of a hack, but akin to how the list command does it.
	// TODO: check if this function gets called in a Go routine of the main handler
	monitor.Start()
	time.Sleep(5 * time.Second)
	monitor.Stop()

	sd := monitor.Serial()
	serial := make([]*rpc.AttachedSerialBoard, len(sd))
	i := 0
	for _, s := range sd {
		fqbn, err := findFqbnForSerial(pm, s.VendorID, s.ProductID)
		if err != nil && !isNotFoundError(err) {
			return nil, err
		}

		serial[i] = &rpc.AttachedSerialBoard{
			Port:         s.Port,
			Fqbn:         fqbn,
			SerialNumber: s.SerialNumber,
			ProductID:    s.ProductID,
			VendorID:     s.VendorID,
		}
		i++
	}

	nd := monitor.Network()
	network := make([]*rpc.AttachedNetworkBoard, len(nd))
	i = 0
	for _, s := range nd {
		network[i] = &rpc.AttachedNetworkBoard{
			Name:    s.Name,
			Fqbn:    "unknown",
			Info:    s.Info,
			Address: s.Address,
			Port:    uint64(s.Port),
		}
		i++
	}

	return &rpc.BoardListResp{
		Serial:  serial,
		Network: network,
	}, nil
}

func findFqbnForSerial(pm *packagemanager.PackageManager, vid, pid string) (string, error) {
	for _, pkg := range pm.GetPackages().Packages {
		for _, platform := range pkg.Platforms {
			platformRelease := pm.GetInstalledPlatformRelease(platform)
			if platformRelease == nil {
				continue
			}

			for _, brd := range platformRelease.Boards {
				if !brd.HasUsbID(vid, pid) {
					continue
				}

				return brd.FQBN(), nil
			}
		}
	}

	return "", &fqbnNotFoundError{VID: vid, PID: pid}
}

type fqbnNotFoundError struct {
	VID string
	PID string
}

func (e *fqbnNotFoundError) Error() string {
	return fmt.Sprintf("no installed board found for %s:%s", e.VID, e.PID)
}

func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(*fqbnNotFoundError)
	return ok
}
