package software_test

import (
	"fmt"
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	"github.com/brainupdaters/drlm-common/pkg/os"
	"github.com/brainupdaters/drlmctl/software"
)

func TestCompile(t *testing.T) {
	fs.Init()
	fmt.Println(software.SoftwareCore.Compile(os.Linux, os.ArchAmd64, "develop"))
	t.Fail()
}
