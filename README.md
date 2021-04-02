<p align="center"><img width 100% src="https://github.com/godefi/medadata/blob/main/gorocklogo.png"></p>

## GoRock Kubernetes Module
Support module for application deployment to Google Kubernetes Engine using Google Container Registry.

Release:
1.0.0

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=defidbpub_gorock&metric=alert_status)](https://sonarcloud.io/dashboard?id=defidbpub_gorock)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=defidbpub_gorock&metric=bugs)](https://sonarcloud.io/dashboard?id=defidbpub_gorock)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=defidbpub_gorock&metric=security_rating)](https://sonarcloud.io/dashboard?id=defidbpub_gorock)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=defidbpub_gorock&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=defidbpub_gorock)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=defidbpub_gorock&metric=code_smells)](https://sonarcloud.io/dashboard?id=defidbpub_gorock)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=defidbpub_gorock&metric=ncloc)](https://sonarcloud.io/dashboard?id=defidbpub_gorock)

Requirements:
Golang 1.15+ | Docker 19.03.13+ (Desktop 2.5+) | Kubernetes 1.19+

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

	// PLAYBOOK

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
