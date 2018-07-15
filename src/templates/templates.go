package templates

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"github.com/alejandroq/go-explicit-semver.io/src/log"
	"github.com/olekukonko/tablewriter"
)

var mu sync.Mutex
var errs = make(chan error)
var done = make(chan struct{}, 1)

const directoryName = ".semver"

// Error arguments
var (
	errIncompleteTemplate   = errors.New("template lacks either an input, output or both")
	errTemplateDoesNotExist = errors.New("requested template does not exist")
	errWriteError           = errors.New("500 write error")
)

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

// Templates handles the invocation of the `templates` subcommand
func Templates(cmd, templateInput, templateOutput string) error {
	template := Template{
		Input:  templateInput,
		Output: templateOutput,
	}

	switch cmd {
	case "add":
		add(template)
	case "rm":
		rm(template)
	default:
		list()
	}

	select {
	case <-done:
		return nil
	case err := <-errs: // do we have to clean up goroutines to avoid leak? use `-runtime` flag - verify main goroutine is the last one to exit 1
		return err
	}
}

func validateTemplate(template Template) {
	if template.Input == "" || template.Output == "" {
		errs <- errIncompleteTemplate
	}
}

func readConfig() Config {
	f, err := ioutil.ReadFile(directoryName + "/config.json")
	if err != nil {
		errs <- err
	}

	var config Config
	err = json.Unmarshal(f, config)
	if err != nil {
		errs <- err
	}

	return config
}

func writeConfig(config Config) int {
	f, err := os.OpenFile(directoryName+"/config.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		errs <- err
	}
	defer f.Close()

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		errs <- err
	}

	n, err := f.Write(jsonConfig)
	if err != nil {
		errs <- err
	}

	log.Log("updated "+f.Name()+" file", map[string]interface{}{
		"bytes":  n,
		"update": config,
	}) // log file update

	return n
}

// mutex locks in the following functions are meant to
// account for mutual exclusion - make global file handling
// more deterministic in the concurrent setting
func add(template Template) {
	mu.Lock()
	defer mu.Unlock()

	validateTemplate(template)

	config := readConfig()
	config.Templates = append(config.Templates, template)
	// n := writeConfig(config)
	writeConfig(config)

	done <- struct{}{}
	// if n > 0 {
	// 	done <- struct{}{}
	// } else {
	// 	errs <- errWriteError
	// }
}

func rm(template Template) {
	mu.Lock()
	defer mu.Unlock()

	validateTemplate(template)

	config := readConfig()
	for i, t := range config.Templates {
		if t.Input == template.Input && t.Output == template.Output {
			a := config.Templates
			a = append(a[:i], a[i+1:]...) // delete entry into slice
			config.Templates = a
		} else {
			errs <- errTemplateDoesNotExist
		}
	}
	n := writeConfig(config)

	if n > 0 {
		done <- struct{}{}
	} else {
		errs <- errWriteError
	}

	done <- struct{}{}
}

func list() {
	config := readConfig()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"template", "output"})
	for _, t := range config.Templates {
		table.Append([]string{
			t.Input,
			t.Output,
		})
	}
	table.Render()

	done <- struct{}{}
}
