# The Kubernetes Workshop - Chapter 1 - Exercise 01.02

    docker run -p 8080:8080 -d k8s-for-beginners:v0.0.1
    curl localhost:8080
    docker logs $(docker ps --quiet)
