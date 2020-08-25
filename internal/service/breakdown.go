package service

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/protos/ld"
)

type Breakdown struct {
	log hclog.Logger
}

func NewBreakdown(l hclog.Logger) *Breakdown {
	return &Breakdown{l}
}

func (b *Breakdown) GetLanguages(projectID string) ([]*ld.Language, error) {

	return executeLinguist("portapps"), nil
}

func executeLinguist(repo string) []*ld.Language {
	cmd := exec.Command("github-linguist", "/opt/data/"+repo, "--json")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	output := string(cmdOutput.Bytes())

	var result map[string][]interface{}
	json.Unmarshal([]byte(output), &result)

	languages := []*ld.Language{}
	for key, value := range result {
		language := &ld.Language{
			Name: key,
		}

		files := []string{}
		for _, fp := range value {
			files = append(files, fp.(string))
		}
		language.Files = files
		languages = append(languages, language)
	}

	return languages
}
