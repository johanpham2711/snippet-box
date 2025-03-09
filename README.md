# <p style="text-align: center">Snippet Box</p>

<div style="display: flex; justify-content: center; flex-wrap: wrap; gap: 10px; text-align: center;">

  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go" />

</div>

## Description

[Snippet Box](https://github.com/johanpham2711/snippet-box) with [Golang](https://go.dev/).

by [Johan Pham](https://github.com/johanpham2711)

## Installation

```bash
go get
```

## Run

1. Basic

```bash

$ go run main.go

```

2. Live reload

```bash
# Home
$ go install github.com/air-verse/air@latest
$ export PATH=$(go env GOPATH)/bin:$PATH
$ source ~/.zshrc  # or source ~/.bash_profile

$ cd <project_folder>
$ export PATH="$HOME/go/bin:$PATH"
$ source ~/.zshrc  # or source ~/.bash_profile

$ air init (Optional)
$ air
```

