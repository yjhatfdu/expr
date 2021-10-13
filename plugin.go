package expr

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"plugin"
)

func LoadPlugin(fp string) error {
	_, err := plugin.Open(fp)
	return err
}

func LoadPluginBinary(data []byte) error {
	fp := path.Join(os.TempDir(), fmt.Sprintf("%X_%X.so", rand.Uint64(), rand.Uint64()))
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	f.Close()
	return LoadPlugin(fp)
}
