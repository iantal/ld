package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/internal/files"
	"github.com/iantal/ld/protos/ld"
)

type Breakdown struct {
	log      hclog.Logger
	basePath string
	rmHost   string
	store    files.Storage
}

func NewBreakdown(l hclog.Logger, basePath, rmHost string, store files.Storage) *Breakdown {
	return &Breakdown{l, basePath, rmHost, store}
}

func (b *Breakdown) GetLanguages(projectID, commit string) ([]*ld.Language, error) {

	projectPath := filepath.Join(b.store.FullPath(projectID), "bundle")

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := b.downloadRepository(projectID, commit)
		if err != nil {
			b.log.Error("Could not download bundled repository", "projectID", projectID, "commit", commit, "err", err)
			return []*ld.Language{}, fmt.Errorf("Could not download bundled repository for %s", projectID)
		}
	}

	bp := projectID + ".bundle"
	srcPath := b.store.FullPath(filepath.Join(projectID, "bundle", bp))
	destPath := b.store.FullPath(filepath.Join(projectID, "unbundle"))

	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		err := b.store.Unbundle(srcPath, destPath)
		if err != nil {
			b.log.Error("Could not unbundle repository", "projectID", projectID, "commit", commit, "err", err)
			return []*ld.Language{}, fmt.Errorf("Could not unbundle repository for %s", projectID)
		}
	}

	return b.executeLinguist(filepath.Join(destPath, projectID)), nil
}

func (b *Breakdown) downloadRepository(projectID, commit string) error {
	resp, err := http.DefaultClient.Get("http://" + b.rmHost + "/api/v1/projects/" + projectID + "/" + commit + "/download")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}

	b.log.Info("Content-Dispozition", "file", resp.Header.Get("Content-Disposition"))

	b.save(projectID, resp.Body)
	resp.Body.Close()

	return nil
}

func (b *Breakdown) save(projectID string, r io.ReadCloser) error {
	b.log.Info("Save project - storage", "projectID", projectID)

	bp := projectID + ".bundle"
	fp := filepath.Join(projectID, "bundle", bp)
	err := b.store.Save(fp, r)

	if err != nil {
		b.log.Error("Unable to save file", "error", err)
		return fmt.Errorf("Unable to save file %s", err)
	}

	return nil
}

func (b *Breakdown) executeLinguist(repo string) []*ld.Language {
	b.log.Info("Executing linguist for project", "project", repo)
	cmd := exec.Command("github-linguist", repo, "--json")
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
