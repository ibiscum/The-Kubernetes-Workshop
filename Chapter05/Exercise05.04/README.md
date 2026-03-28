# The Kubernetes Workshop - Chapter 5 - Exercise 05.04

linting:

    docker run --rm -v .:/dir stackrox/kube-linter lint /dir

running:

    kubectl get namespaces
    kubectl config set-context $(kubectl config current-context) --namespace the-kubernetes-workshop
    kubectl get pods
    kubectl get nodes
