package board

import (
	"context"
	"errors"
	"time"

	"github.com/arduino/arduino-cli/arduino/cores"
	"github.com/arduino/arduino-cli/arduino/cores/packagemanager"
	"github.com/arduino/arduino-cli/commands"
	"github.com/arduino/arduino-cli/rpc"
	discovery "github.com/arduino/board-discovery"
	log "github.com/sirupsen/logrus"
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
		b := &rpc.AttachedSerialBoard{
			Port:         s.Port,
			SerialNumber: s.SerialNumber,
			ProductID:    s.ProductID,
			VendorID:     s.VendorID,
		}
		completeInfoForSerial(pm, b)

		serial[i] = b
		i++
	}

	nd := monitor.Network()
	network := make([]*rpc.AttachedNetworkBoard, len(nd))
	i = 0
	for _, s := range nd {
		network[i] = &rpc.AttachedNetworkBoard{
			Name:    s.Name,
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

func completeInfoForSerial(pm *packagemanager.PackageManager, b *rpc.AttachedSerialBoard) {
	log.SetLevel(log.DebugLevel)
	var matchingBoard *cores.Board
	for _, pkg := range pm.GetPackages().Packages {
		for _, platform := range pkg.Platforms {
			platformRelease := pm.GetInstalledPlatformRelease(platform)
			if platformRelease == nil {
				continue
			}

			for _, brd := range platformRelease.Boards {
				if !brd.HasUsbID(b.VendorID, b.ProductID) {
					continue
				}

				matchingBoard = brd
				break
			}
		}
	}

	if matchingBoard == nil {
		log.WithField("port", b.Port).Debug("did not find installed package")
		return
	}

	b.Fqbn = matchingBoard.FQBN()
	b.Name = matchingBoard.Name()
}
