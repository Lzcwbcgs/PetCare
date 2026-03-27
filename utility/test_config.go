package utility

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

func ConfigureTestConfig(tb testing.TB) {
	tb.Helper()

	adapter, ok := g.Cfg().GetAdapter().(*gcfg.AdapterFile)
	if !ok || adapter == nil {
		newAdapter, err := gcfg.NewAdapterFile()
		if err != nil {
			tb.Fatalf("create config adapter: %v", err)
		}
		g.Cfg().SetAdapter(newAdapter)
		adapter = newAdapter
	}
	adapter.SetFileName(filepath.Join(ProjectRoot(tb), "manifest", "config", "config.yaml"))
}

func ProjectRoot(tb testing.TB) string {
	tb.Helper()

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		tb.Fatal("resolve project root: runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(currentFile), ".."))
}
