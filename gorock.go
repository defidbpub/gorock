package gorock

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
	"regexp"
)

func CreateDockerFile(app string, port string, param string) bool {

	err := ioutil.WriteFile("Dockerfile", []byte(""+
		"FROM golang:latest\n"+
		"RUN mkdir -p /app\n"+
		"ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin\n"+
		"RUN apt-get install git coreutils\n"+
		"COPY . /app\n"+
		"WORKDIR /app\n"+
		"RUN go test\n"+
		"RUN go build -a -o "+app+" .\n"+
		"RUN find . -name '*.go' -delete\n"+
		"RUN chmod +x /app/"+app+"\n"+
		"EXPOSE "+port+"\n"+
		"ENTRYPOINT [\"/app/"+app+"\" "+param+"]"), 0755)
	if err != nil {
		log.Errorf("Unable to write file: %v", err)
	}
	log.Info("Dockerfile created.")
	return true

}

func DockerCheck() bool {
	out, err := exec.Command("docker", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	matched, err := regexp.MatchString(`Server: Docker Engine`, string(out))
	if err != nil {
		log.Fatal(err)
	}
	if matched {
		log.Info("Docker installed.")
	}
	return matched
}

func DockerImageBuild(tag string) bool {
	out, err := exec.Command("docker", "build", "-t", tag, ".").Output()
	if err != nil {
		log.Fatal(err)
	}
	matched, err := regexp.MatchString(tag, string(out))
	if err != nil {
		log.Fatal(err)
	}
	if matched {
		log.Info("Docker image created.\nDetail:\n", string(out))
	}
	return matched
}

func DockerImagePush(tag string) bool {
	out, err := exec.Command("docker", "push", tag).Output()
	if err != nil {
		log.Fatal(err)
	}
	matched, err := regexp.MatchString("Pushed", string(out))
	if matched {
		log.Info("Docker image uploaded.\nDetail:\n", string(out))
	}
	return matched
}

func GkeDeploymentFile(app string, port string, tag string, namespace string) bool {
	err := ioutil.WriteFile("deployment.yaml", []byte(""+
		"apiVersion: apps/v1\n"+
		"kind: Deployment\n"+
		"metadata:\n"+
		"  name: "+app+"\n"+
		"spec:\n"+
		"  selector:\n"+
		"    matchLabels:\n"+
		"      app: "+app+"\n"+
		"  template:\n"+
		"    metadata:\n"+
		"      labels:\n"+
		"        app: "+app+"\n"+
		"    spec:\n"+
		"      containers:\n"+
		"        - name: "+app+"\n"+
		"          image: "+tag+"\n"+
		"          imagePullPolicy: Always\n"+
		"          ports:\n"+
		"            - containerPort: "+port+"\n"+
		"---\n"+
		"apiVersion: v1\n"+
		"kind: Service\n"+
		"metadata:\n"+
		"  name: "+app+"-service\n"+
		"  namespace: "+namespace+"\n"+
		"spec:\n"+
		"  ports:\n"+
		"    - protocol: TCP\n"+
		"      port: "+port+"\n"+
		"      targetPort: "+port+"\n"+
		"  selector:\n"+
		"    app: "+app+"\n"+
		"  type: LoadBalancer\n"+
		"  loadBalancerIP: \"\"\n"), 0755)
	if err != nil {
		log.Errorf("Unable to write file: %v", err)
	}
	log.Info("Kubernetes deployment and service file created.")
	return true
}
