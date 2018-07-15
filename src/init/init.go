package init

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/alejandroq/go-explicit-semver.io/src/log"
)

const directoryName = ".semver"

// Config for the application. Typically housed in the `.semver/config.json` file.
type Config struct {
	Versioning []Artifact `json:"versioning"`
	Templates  []Template `json:"templates"`
}

// Artifact that is being versioned
// If directory, children need be `diff`d as per Patch releases
type Artifact string

// Template input (template file) to output file
type Template struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

var (
	errNotEnoughArguments = errors.New("please provide a list of files/directories to maintain a semantic version for")
)

// Init creates the `.semver` directory as well as it's accompanying files
// supplemental arguments may be provided to initalize configuration
func Init(args []string) error {
	// Initializing goroutine channels
	errs := make(chan error)
	done := make(chan struct{}, 1)

	if len(args) == 0 {
		return errNotEnoughArguments
	}

	if err := createDirIfNotExist(directoryName); err != nil {
		return err
	}

	go func() {
		_, err := os.OpenFile(directoryName+"/ledger.json", os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		f, err := os.OpenFile(directoryName+"/config.json", os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			errs <- err
		}
		defer f.Close()

		config := createConfig(args)
		jsonConfig, err := json.Marshal(config)
		if err != nil {
			errs <- err
		}

		n, err := f.Write(jsonConfig)
		if err != nil {
			errs <- err
		}

		log.Log("wrote "+f.Name()+" file", map[string]interface{}{
			"bytes": n,
			"file":  config,
		}) // log file creation
		log.Log("initialized vanilla go-explicit-semver in "+f.Name(), make(map[string]interface{})) // log command completion

		// indicate operation is complete
		done <- struct{}{}
	}()

	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}
}

func createConfig(args []string) Config {
	var c Config
	for _, v := range args {
		c.Versioning = append(c.Versioning, Artifact(v))
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
