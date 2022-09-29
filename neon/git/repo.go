package git

import (
	"fmt"
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

func cloneRepo(c *gin.Context, repo string, credentials entities.Credentials) *git.Repository {

	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))

	options := &git.CloneOptions{
		URL:  repo,
		Auth: credentials.GetGitCreds(c),
	}

	r, err := git.PlainClone(repoPath, false, options)
	errors.Check(err).NewInternal("could not clone repository").Panic()
	return r
}

func WriteAppFile(app *api.CreateAppRequest) {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(app.Repository))
	yamlData, err := yaml.Marshal(&app)
	errors.Check(err).NewInternal("could not marshal yaml").Panic()

	fileName := path.Join(repoPath, "app.yaml")
	err = ioutil.WriteFile(fileName, yamlData, 0644)
	errors.Check(err).NewInternal("could not write app.yaml file for git repository").Panic()
}

func ReadAppFile(repo string) *api.CreateAppRequest {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))
	fmt.Println(path.Join(repoPath, "app.yaml"))
	yfile, err := ioutil.ReadFile(path.Join(repoPath, "app.yaml"))
	fmt.Println(err)
	errors.Check(err).NewInternal("failed to read app.yaml for repository").Panic()

	var data *api.CreateAppRequest
	err2 := yaml.Unmarshal(yfile, &data)
	errors.Check(err2).NewInternal("failed to unmarshal app.yaml").Panic()

	return data
}

func CreateOverride(repo string, productName string) {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))

	os.MkdirAll(path.Join(repoPath, "neon", productName), os.ModePerm)

	d1 := []byte("# add config overrides here")
	err := os.WriteFile(path.Join(repoPath, "neon", productName, "overrides.yaml"), d1, 0644)
	errors.Check(err).NewInternal("could not create override file").Panic()

}

func CommitOverrides(repo string) error {
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
	w.AddGlob("neon/*")
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

func Push(c *gin.Context, repo string, creds entities.Credentials) {
	dir := os.Getenv("NEON_HOME")
	repoPath := path.Join(dir, path.Base(repo))
	r, err := git.PlainOpen(repoPath)
	errors.Check(err).NewInternal("couldnt open local repository").Panic()
	err = r.Push(&git.PushOptions{
		Auth: creds.GetGitCreds(c),
	})
	errors.Check(err).NewInternal("couldnt push to remote repository").Panic()
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

func FillRepository(c *gin.Context, req api.CreateAppRequest) {
	credentials, err := store.CredentialsRepository().GetByName(req.CredentialName)
	errors.Check(err).NewNotFound("credentials not found").Panic()

	r := cloneRepo(c, req.Repository, credentials)

	wipeAndAddFiles(c, req, credentials, r)
}

func AddProduct(c *gin.Context, productName string, app entities.App) error {
	credentials, err := store.CredentialsRepository().GetByName(app.Credentials)
	if err != nil {
		return err
	}
	appData := ReadAppFile(app.Repository)

	appData.Products = append(appData.Products, productName)
	WriteAppFile(appData)
	CreateOverride(app.Repository, productName)

	err = CommitOverrides(appData.Repository)
	errors.Check(err).NewInternal("failed to commit to repository").Panic()

	Push(c, appData.Repository, credentials)
	return nil
}
