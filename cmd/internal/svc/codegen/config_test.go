package codegen

import (
	"github.com/unionj-cloud/go-doudou/v2/toolkit/astutils"
	"os"
	"path/filepath"
	"testing"
)

const testDir = "testdata"

func TestGenConfig(t *testing.T) {
	dir := testDir + "config"
	InitSvc(dir)
	defer os.RemoveAll(dir)
	var ic astutils.InterfaceCollector
	GenConfig(dir, ic)
}

func TestGenConfig1(t *testing.T) {
	var ic astutils.InterfaceCollector
	GenConfig(filepath.Join(testDir), ic)
}
