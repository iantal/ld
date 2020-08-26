package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/internal/data"
	"github.com/iantal/ld/internal/files"
	"github.com/iantal/ld/protos/ld"
)

type Breakdown struct {
	log      hclog.Logger
	basePath string
	rkHost   string
	store    files.Storage
}

func NewBreakdown(l hclog.Logger, basePath, rkHost string, store files.Storage) *Breakdown {
	return &Breakdown{l, basePath, rkHost, store}
}

func (b *Breakdown) GetLanguages(projectID string) ([]*ld.Language, error) {
	project, err := b.getProjectName(projectID)
	if err != nil {
		b.log.Error("Could not get project name", "projectID", projectID, "err", err)
		return []*ld.Language{}, nil
	}

	err = b.downloadRepository(project)
	if err != nil {
		b.log.Error("Could not download zip from rk for project", "projectID", projectID, "err", err)
		return []*ld.Language{}, nil
	}

	return b.executeLinguist(filepath.Join(projectID, "unzip", project.Name)), nil
}

func (b *Breakdown) getProjectName(projectID string) (*data.Project, error) {
	resp, err := http.DefaultClient.Get("http://" + b.rkHost + "/api/v1/projects/" + projectID)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	project := &data.Project{}
	err = json.Unmarshal(body, project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (b *Breakdown) downloadRepository(project *data.Project) error {
	b.log.Info("Project", "p", project)
	projectID := project.ProjectID.String()
	resp, err := http.DefaultClient.Get("http://" + b.rkHost + "/api/v1/projects/" + projectID + "/download")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}

	b.save(projectID, project.Name, resp.Body)
	resp.Body.Close()

	return nil
}

func (b *Breakdown) save(id, projectName string, r io.ReadCloser) {
	b.log.Info("Save project - storage", "id", id)

	unzippedPath := filepath.Join(id, "unzip")

	zp := id + ".zip"
	fp := filepath.Join(id, "zip", zp)
	err := b.store.Save(fp, r)

	if err != nil {
		b.log.Error("Unable to save file", "error", err)
	} else {
		b.log.Info("Unzipping", "id", id, "path", unzippedPath)
		err := b.store.Unzip(b.store.FullPath(fp), b.store.FullPath(unzippedPath), projectName)
		if err != nil {
			b.log.Error("Unable to unzip file", "error", err)
		}
	}
}

func (b *Breakdown) executeLinguist(repo string) []*ld.Language {
	repoPath := filepath.Join(b.basePath, repo)
	b.log.Info("Base path", "repo", repoPath)
	cmd := exec.Command("github-linguist", repoPath, "--json")
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
