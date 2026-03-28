# The Kubernetes Workshop - Chapter 5 - Exercise 05.01

linting:

    docker run --rm -v .:/dir stackrox/kube-linter lint /dir

running:

    kubectl create namespace "the-kubernetes-workshop"
    kubectl config set-context --current --namespace="the-kubernetes-workshop"
    kubectl create -f single-container-pod.yaml
    kubectl get pods
    kubectl describe pod
    kubectl delete pod first-pod
