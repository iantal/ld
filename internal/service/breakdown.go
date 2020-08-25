package service

import (
	"bytes"
	"fmt"
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
	executeLinguist("portapps")
	return []*ld.Language{}, nil
}

func executeLinguist(repo string) {
	cmd := exec.Command("github-linguist", "/opt/data/"+repo, "--json")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))
}
