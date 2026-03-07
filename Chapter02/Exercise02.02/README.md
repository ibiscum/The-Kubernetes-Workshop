# The Kubernetes Workshop - Chapter 02 - Exercise 02.02

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
    docker ps -a
