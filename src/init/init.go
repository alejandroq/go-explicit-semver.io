package init

import (
	"encoding/json"
	"os"

	"github.com/alejandroq/go-explicit-semver.io/src/log"
	"github.com/sirupsen/logrus"
)

const directoryName = ".gosemver"

// Config for the application. Typically housed in the `.gosemver/config.json` file.
type Config struct {
	Versioning []Artifact `json:"versioning"`
	Templates  []Template `json:"templates"`
}

// Artifact that is being versioned
// If directory, children need be `diff`d as per Patch releases
type Artifact struct {
	Source       string `json:"source"` // can derive a unique ID from source, as is expected to be unique.
	GitTagMirror bool   `json:"primary"`
}

// Template input (template file) to output file
type Template struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

// Init creates the `.gosemver` directory as well as it's accompanying files
// supplemental arguments may be provided to initalize configuration
func Init(args []string) error {
	// TODO execute with go routines. Create "error" and "done" channel.
	// Any meets of "error" in a switch case, will cause the application
	// to logout and error.

	if err := createDirIfNotExist(directoryName); err != nil {
		return err
	}

	_, err := os.OpenFile(directoryName+"/ledger.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(directoryName+"/config.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	config := createConfig(args)
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	n, err := f.Write(jsonConfig)
	if err != nil {
		return err
	}
	log.Elog.WithFields(logrus.Fields{
		"bytes":  n,
		"config": config,
	}).Info("wrote .gosemver/config.json file")

	return nil
}

func createConfig(args []string) Config {
	var c Config
	for i, v := range args {
		if i == 0 {
			c.Versioning = append(c.Versioning, Artifact{
				Source:       v,
				GitTagMirror: true,
			})
		} else {
			c.Versioning = append(c.Versioning, Artifact{
				Source:       v,
				GitTagMirror: false,
			})
		}
	}
	c.Templates = []Template{}
	return c
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
