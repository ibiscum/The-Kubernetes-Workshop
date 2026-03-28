# The Kubernetes Workshop - Chapter 5 - Exercise 05.05

linting:

    docker run --rm -v .:/dir stackrox/kube-linter lint /dir

running:

    kubectl create -f pod-with-container-command.yaml
    kubectl describe pods
    kubectl logs command-pod -f
    kubectl delete pod command-pod
