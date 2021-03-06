<p align="center"><img width 100% src="https://github.com/godefi/medadata/blob/main/gorocklogo.png"></p>

## GoRock Kubernetes Module
Support module for application deployment to Google Kubernetes Engine using Google Container Registry.

Release:
1.0.1

Requirements:
Golang 1.16.5+ | Docker 19.03.13+ (Desktop 2.5+) | Kubernetes 1.19+

```bash
go get "github.com/defidbpub/gorock"
```

Purpose:

Documentation:
- import:
```go
(in your Go project)
mkdir deploy
(in deploy directory)
vi deploy.go

package main

import (
	"github.com/defidbpub/gorock"
)

func main() {

	var repo = "<docker-image-tag>:1.0"
	var app = "<app-name>"
	var port = "<app-port>"
	var namespace = "<kubernetes-namespace>"
	var param = "<params>" // in double quotes comma separated

	//playbook

	gorock.CreateDockerFile(app, port, param)
	if gorock.DockerCheck() {
		if gorock.DockerImageBuild(repo) {
			gorock.DockerImagePush(repo)
			gorock.GkeDeploymentFile(app, port, repo, namespace)
		}
	}

}
```

- run:
```bash
go run deploy/deploy.go
```

TODO:

Contribution:

License: 
Jan Rock (c) 2021
