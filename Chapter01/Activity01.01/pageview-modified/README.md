# The Kubernetes Workshop - Chapter 1 - Activity 01.01

    docker run --rm -i hadolint/hadolint < Dockerfile

    go mod init pageview-modified
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pageview-modified

    docker build -t activity01.01:0.0.2 .
    docker run --rm -p 8080:8080 activity01.01:0.0.2

    curl localhost:8080
