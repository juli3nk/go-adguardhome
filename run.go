package adguardhome

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func Run(listenIPAddr string) (*os.Process, error) {
	dataPath := "/opt/adguardhome"
	confDir := path.Join(dataPath, "conf")
	workDir := path.Join(dataPath, "work")

	configFile := filepath.Join(confDir, "AdGuardHome.yaml")

	if !dirExists(confDir) {
		return nil, fmt.Errorf("directory %s does not exist", confDir)
	}
	if !dirExists(workDir) {
		return nil, fmt.Errorf("directory %s does not exist", workDir)
	}

	args := []string{
		"-c", configFile,
		"-w", workDir,
		"-h", listenIPAddr,
	}

	cmd := exec.Command("/usr/local/bin/AdGuardHome", args...)

	go func() {
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}()

	for {
		if cmd.Process != nil {
			break
		}
	}

	return cmd.Process, nil
}
