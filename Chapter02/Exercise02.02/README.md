# The Kubernetes Workshop - Chapter 2 - Exercise 02.02

## linter

    docker run -v .:/dir stackrox/kube-linter lint /dir

## minikube

    eval $(minikube docker-env)
    minikube image load k8s-for-beginners:v0.0.1

## kind

    kind load docker-image k8s-for-beginners:v0.0.1
    kubectl apply -f k8s-for-beginners-pod.yaml

    kubectl get pod
    kubectl get pod -o wide
    kubectl describe pod k8s-for-beginners
    kubectl get node
    kubectl port-forward pod/k8s-for-beginners 9090:8080

    curl localhost:9090
