# The Kubernetes Workshop - Chapter 5 - Exercise 05.03

linting:

    docker run --rm -v .:/dir stackrox/kube-linter lint /dir

running:

    kubectl create -f single-container-pod-with-namespace.yaml
    kubectl get pods
    kubectl describe pod -n kube-public
    kubectl delete pod first-pod-with-namespace -n kube-public
