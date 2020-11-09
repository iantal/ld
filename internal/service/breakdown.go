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

	"github.com/iantal/ld/internal/files"
	"github.com/iantal/ld/internal/util"
	"github.com/iantal/ld/protos/ld"
	"github.com/sirupsen/logrus"
)

type Breakdown struct {
	log      *util.StandardLogger
	basePath string
	rmHost   string
	store    files.Storage
}

func NewBreakdown(l *util.StandardLogger, basePath, rmHost string, store files.Storage) *Breakdown {
	return &Breakdown{l, basePath, rmHost, store}
}

func (b *Breakdown) GetLanguages(projectID, commit string) ([]*ld.Language, error) {
	projectPath := filepath.Join(b.store.FullPath(projectID), commit, "bundle")

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := b.downloadRepository(projectID, commit)
		if err != nil {
			b.log.WithFields(logrus.Fields{
				"projectID": projectID,
				"commit":    commit,
				"error":     err,
			}).Error("Could not download bundled repository")
			return []*ld.Language{}, fmt.Errorf("Could not download bundled repository for %s", projectID)
		}
	}

	bp := commit + ".bundle"
	srcPath := b.store.FullPath(filepath.Join(projectID, commit, "bundle", bp))
	destPath := b.store.FullPath(filepath.Join(projectID, commit, "unbundle"))

	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		err := b.store.Unbundle(srcPath, destPath)
		if err != nil {
			b.log.WithFields(logrus.Fields{
				"projectID": projectID,
				"commit":    commit,
				"error":     err,
			}).Error("Could not unbundle repository")
			return []*ld.Language{}, fmt.Errorf("Could not unbundle repository for %s", projectID)
		}
	}

	return b.executeLinguist(projectID, commit, filepath.Join(destPath, commit)), nil
}

func (b *Breakdown) downloadRepository(projectID, commit string) error {
	resp, err := http.DefaultClient.Get("http://" + b.rmHost + "/api/v1/projects/" + projectID + "/" + commit + "/download")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected error code 200 got %d", resp.StatusCode)
	}

	b.log.WithField("file", resp.Header.Get("Content-Disposition")).Info("Content-Dispozition")

	b.save(projectID, commit, resp.Body)
	resp.Body.Close()

	return nil
}

func (b *Breakdown) save(projectID, commit string, r io.ReadCloser) {
	b.log.WithFields(logrus.Fields{
		"projectID": projectID,
		"commit":    commit,
	}).Info("Save project to storage")

	bp := commit + ".bundle"
	fp := filepath.Join(projectID, commit, "bundle", bp)
	err := b.store.Save(fp, r)

	if err != nil {
		b.log.WithFields(logrus.Fields{
			"projectID": projectID,
			"commit":    commit,
			"error":     err,
		}).Error("Unable to save file")
		return
	}
}

func (b *Breakdown) executeLinguist(projectID, commit, repo string) []*ld.Language {
	b.log.WithFields(logrus.Fields{
		"projectID": projectID,
		"commit":    commit,
		"path":      repo,
	}).Info("Executing linguist")
	cmd := exec.Command("github-linguist", repo, "--json")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()

	if err != nil {
		b.log.WithFields(logrus.Fields{
			"projectID": projectID,
			"commit":    commit,
			"error":     err,
		}).Error("Error executing linguist")
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

	for _, l := range languages {
		b.log.WithFields(logrus.Fields{
			"projectID":   projectID,
			"commit":      commit,
			"language":    l.Name,
			"total_files": len(l.Files),
		}).Info("Linguist result")
	}
	return languages
}
