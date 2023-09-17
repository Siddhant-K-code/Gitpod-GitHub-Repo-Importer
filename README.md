# GitHub Repository Importer for Gitpod

## Introduction

Welcome to the GitHub Repo Importer for Gitpod, an experimental project pioneering the importing of GitHub repositories into Gitpod projects using the cutting-edge Gitpod API and the [`go-github`](https://github.com/google/go-github) SDK. By leveraging the power of these two remarkable tools, our project seeks to streamline and enhance your Gitpod project setup process, bringing your GitHub repositories into the Gitpod environment with unprecedented ease and efficiency.

> **Warning**: This project is still in its early stages of development, and as such, it is not yet ready for production use. However, I am working hard to bring it to a stable state as soon as possible. In the meantime, i welcome you to experiment with it and contribute to its development.

## Start working 🚀

> **Important**: Add [Gitpod PAT Token](https://www.gitpod.io/docs/references/gitpod-public-api) and [generate GitHub PAT](https://github.com/settings/tokens/new?description=Import+GitHub+Repos+to+Gitpod+Projects&scopes=repo) to [Gitpod Environment Variables](https://www.gitpod.io/docs/configure/projects/environment-variables#using-the-account-settings).

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/Siddhant-K-code/Gitpod-GitHub-Repo-Importer)

## Run the project

- Follow the instructions in [`main.go`](./main.go) to import your GitHub repositories into Gitpod projects.
- It requires Gitpod `Org Id`, you can find it via uncommenting [`main.go#L62`](./main.go#L62) and running the project. It will print the Org Id in the console.
- Add your `GITPOD_ORG_ID` in [`main.go#L45`](./main.go#L45).
- Add your GitHub Username in [`main.go#L68`](./main.go#L68)

  ```sh
  go run main.go
  ```

---

> **Note**: This Project is not maintained by Gitpod.
