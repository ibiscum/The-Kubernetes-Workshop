# The Kubernetes Workshop - Chapter 01 - Exercise 1.03

    docker run --rm -i hadolint/hadolint < Dockerfile

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o k8s-for-beginners main.go
    docker rm -f $(docker ps -aq)
    docker build -t k8s-for-beginners:v0.0.1 .
    docker run -p 8080:8080 -d k8s-for-beginners:v0.0.1
    docker images
