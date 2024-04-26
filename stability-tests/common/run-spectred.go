package common

import (
	"fmt"
	"os"
	"sync/atomic"
	"syscall"
	"testing"

	"github.com/spectre-project/spectred/domain/dagconfig"
)

// RunSpectredForTesting runs spectred for testing purposes
func RunSpectredForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	spectredRunCommand, err := StartCmd("SPECTRED",
		"spectred",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Spectred started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := spectredRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Spectred closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := spectredRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Spectred stopped")
	}
}
