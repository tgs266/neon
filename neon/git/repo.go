package git

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/tgs266/neon/neon/api"
	"github.com/tgs266/neon/neon/errors"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/neon/store/entities"
	"gopkg.in/yaml.v2"
)

func cloneRepo(c *gin.Context, repo string, credentials entities.Credentials) (*git.Repository, error) {

	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))

	options := &git.CloneOptions{
		URL:  repo,
		Auth: credentials.GetGitCreds(c),
	}

	r, err := git.PlainClone(repoPath, false, options)
	return r, err
}

func WriteAppFile(app *api.CreateAppRequest) error {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(app.Repository))
	yamlData, err := yaml.Marshal(&app)
	if err != nil {
		return err
	}
	fileName := path.Join(repoPath, "app.yaml")
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadAppFile(repo string) (*api.CreateAppRequest, error) {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))
	yfile, err := ioutil.ReadFile(path.Join(repoPath, "app.yaml"))

	if err != nil {
		return nil, err
	}
	var data *api.CreateAppRequest
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		return nil, err2
	}
	return data, nil
}

func CreateOverride(repo string, productName string) error {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))

	os.MkdirAll(path.Join(repoPath, "neon", productName), os.ModePerm)

	d1 := []byte("# add config overrides here")
	err := os.WriteFile(path.Join(repoPath, "neon", productName, "overrides.yaml"), d1, 0644)
	return err
}

func CommitAll(repo string) error {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	w.AddGlob("*")
	commit, err := w.Commit("system update", &git.CommitOptions{
		Author: &object.Signature{
			Name: "Neon",
			When: time.Now(),
		},
	})
	if err != nil {
		return err
	}
	_, err = r.CommitObject(commit)
	return err
}

func Push(c *gin.Context, repo string, creds entities.Credentials) error {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	err = r.Push(&git.PushOptions{
		Auth: creds.GetGitCreds(c),
	})
	return err
}

func wipeAndAddFiles(c *gin.Context, req api.CreateAppRequest, creds entities.Credentials, repo *git.Repository) error {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(req.Repository))

	yamlData, err := yaml.Marshal(&req)

	if err != nil {
		return err
	}

	fileName := path.Join(repoPath, "app.yaml")
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Add("app.yaml")
	if err != nil {
		return err
	}

	commit, err := w.Commit("Add app.yaml file", &git.CommitOptions{
		Author: &object.Signature{
			Name: "Neon",
			When: time.Now(),
		},
	})
	if err != nil {
		return err
	}

	_, err = repo.CommitObject(commit)
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		Auth: creds.GetGitCreds(c),
	})

	return err

}

func FillRepository(c *gin.Context, req api.CreateAppRequest) error {
	credentials, err := store.CredentialsRepository().GetByName(req.CredentialName)
	if err != nil {
		errors.NewNotFound("credentials not found", err).Abort(c)
		return err
	}

	r, err := cloneRepo(c, req.Repository, credentials)
	if err != nil {
		return err
	}

	return wipeAndAddFiles(c, req, credentials, r)
}

func AddProduct(c *gin.Context, productName string, app entities.App) error {
	credentials, err := store.CredentialsRepository().GetByName(app.Credentials)
	if err != nil {
		return err
	}
	appData, err := ReadAppFile(app.Repository)
	if err != nil {
		return err
	}
	appData.Products = append(appData.Products, productName)
	err = WriteAppFile(appData)
	if err != nil {
		return err
	}
	err = CommitAll(appData.Repository)
	if err != nil {
		return err
	}
	return Push(c, appData.Repository, credentials)
}
