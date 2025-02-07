package libpod

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/containers/common/libimage"
	"github.com/containers/common/libnetwork/types"
	"github.com/hanks177/podman/v4/libpod/define"
	"github.com/hanks177/podman/v4/pkg/errorhandling"
	"github.com/hanks177/podman/v4/pkg/rootless"
	"github.com/hanks177/podman/v4/pkg/util"
	"github.com/containers/storage"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// removeAllDirs removes all Podman storage directories. It is intended to be
// used as a backup for reset() when that function cannot be used due to
// failures in initializing libpod.
// It does not expect that all the directories match what is in use by Podman,
// as this is a common failure point for `system reset`. As such, our ability to
// interface with containers and pods is somewhat limited.
// This function assumes that we do not have a working c/storage store.
func (r *Runtime) removeAllDirs() error {
	var lastErr error

	// Grab the runtime alive lock.
	// This ensures that no other Podman process can run while we are doing
	// a reset, so no race conditions with containers/pods/etc being created
	// while we are resetting storage.
	// TODO: maybe want a helper for getting the path? This is duped from
	// runtime.go
	runtimeAliveLock := filepath.Join(r.config.Engine.TmpDir, "alive.lck")
	aliveLock, err := storage.GetLockfile(runtimeAliveLock)
	if err != nil {
		logrus.Errorf("Lock runtime alive lock %s: %v", runtimeAliveLock, err)
	} else {
		aliveLock.Lock()
		defer aliveLock.Unlock()
	}

	// We do not have a store - so we can't really try and remove containers
	// or pods or volumes...
	// Try and remove the directories, in hopes that they are unmounted.
	// This is likely to fail but it's the best we can do.

	// Volume path
	if err := os.RemoveAll(r.config.Engine.VolumePath); err != nil {
		lastErr = errors.Wrapf(err, "removing volume path")
	}

	// Tmpdir
	if err := os.RemoveAll(r.config.Engine.TmpDir); err != nil {
		if lastErr != nil {
			logrus.Errorf("Reset: %v", lastErr)
		}
		lastErr = errors.Wrapf(err, "removing tmp dir")
	}

	// Runroot
	if err := os.RemoveAll(r.storageConfig.RunRoot); err != nil {
		if lastErr != nil {
			logrus.Errorf("Reset: %v", lastErr)
		}
		lastErr = errors.Wrapf(err, "removing run root")
	}

	// Static dir
	if err := os.RemoveAll(r.config.Engine.StaticDir); err != nil {
		if lastErr != nil {
			logrus.Errorf("Reset: %v", lastErr)
		}
		lastErr = errors.Wrapf(err, "removing static dir")
	}

	// Graph root
	if err := os.RemoveAll(r.storageConfig.GraphRoot); err != nil {
		if lastErr != nil {
			logrus.Errorf("Reset: %v", lastErr)
		}
		lastErr = errors.Wrapf(err, "removing graph root")
	}

	return lastErr
}

// Reset removes all storage
func (r *Runtime) reset(ctx context.Context) error {
	var timeout *uint
	pods, err := r.GetAllPods()
	if err != nil {
		return err
	}
	for _, p := range pods {
		if err := r.RemovePod(ctx, p, true, true, timeout); err != nil {
			if errors.Cause(err) == define.ErrNoSuchPod {
				continue
			}
			logrus.Errorf("Removing Pod %s: %v", p.ID(), err)
		}
	}

	ctrs, err := r.GetAllContainers()
	if err != nil {
		return err
	}

	for _, c := range ctrs {
		if err := r.RemoveContainer(ctx, c, true, true, timeout); err != nil {
			if err := r.RemoveStorageContainer(c.ID(), true); err != nil {
				if errors.Cause(err) == define.ErrNoSuchCtr {
					continue
				}
				logrus.Errorf("Removing container %s: %v", c.ID(), err)
			}
		}
	}

	if err := r.stopPauseProcess(); err != nil {
		logrus.Errorf("Stopping pause process: %v", err)
	}

	rmiOptions := &libimage.RemoveImagesOptions{Filters: []string{"readonly=false"}}
	if _, rmiErrors := r.LibimageRuntime().RemoveImages(ctx, nil, rmiOptions); rmiErrors != nil {
		return errorhandling.JoinErrors(rmiErrors)
	}

	volumes, err := r.state.AllVolumes()
	if err != nil {
		return err
	}
	for _, v := range volumes {
		if err := r.RemoveVolume(ctx, v, true, timeout); err != nil {
			if errors.Cause(err) == define.ErrNoSuchVolume {
				continue
			}
			logrus.Errorf("Removing volume %s: %v", v.config.Name, err)
		}
	}

	// remove all networks
	nets, err := r.network.NetworkList()
	if err != nil {
		return err
	}
	for _, net := range nets {
		// do not delete the default network
		if net.Name == r.network.DefaultNetworkName() {
			continue
		}
		// ignore not exists errors because of the TOCTOU problem
		if err := r.network.NetworkRemove(net.Name); err != nil && !errors.Is(err, types.ErrNoSuchNetwork) {
			logrus.Errorf("Removing network %s: %v", net.Name, err)
		}
	}

	xdgRuntimeDir := filepath.Clean(os.Getenv("XDG_RUNTIME_DIR"))
	_, prevError := r.store.Shutdown(true)
	graphRoot := filepath.Clean(r.store.GraphRoot())
	if graphRoot == xdgRuntimeDir {
		if prevError != nil {
			logrus.Error(prevError)
		}
		prevError = errors.Errorf("failed to remove runtime graph root dir %s, since it is the same as XDG_RUNTIME_DIR", graphRoot)
	} else {
		if err := os.RemoveAll(graphRoot); err != nil {
			if prevError != nil {
				logrus.Error(prevError)
			}
			prevError = err
		}
	}
	runRoot := filepath.Clean(r.store.RunRoot())
	if runRoot == xdgRuntimeDir {
		if prevError != nil {
			logrus.Error(prevError)
		}
		prevError = errors.Errorf("failed to remove runtime root dir %s, since it is the same as XDG_RUNTIME_DIR", runRoot)
	} else {
		if err := os.RemoveAll(runRoot); err != nil {
			if prevError != nil {
				logrus.Error(prevError)
			}
			prevError = err
		}
	}
	runtimeDir, err := util.GetRuntimeDir()
	if err != nil {
		return err
	}
	tempDir := r.config.Engine.TmpDir
	if tempDir == runtimeDir {
		tempDir = filepath.Join(tempDir, "containers")
	}
	if filepath.Clean(tempDir) == xdgRuntimeDir {
		if prevError != nil {
			logrus.Error(prevError)
		}
		prevError = errors.Errorf("failed to remove runtime tmpdir %s, since it is the same as XDG_RUNTIME_DIR", tempDir)
	} else {
		if err := os.RemoveAll(tempDir); err != nil {
			if prevError != nil {
				logrus.Error(prevError)
			}
			prevError = err
		}
	}
	if storageConfPath, err := storage.DefaultConfigFile(rootless.IsRootless()); err == nil {
		if _, err = os.Stat(storageConfPath); err == nil {
			fmt.Printf("A storage.conf file exists at %s\n", storageConfPath)
			fmt.Println("You should remove this file if you did not modify the configuration.")
		}
	} else {
		if prevError != nil {
			logrus.Error(prevError)
		}
		prevError = err
	}

	return prevError
}
