# The Kubernetes Workshop - Chapter 5 - Exercise 05.06

linting:

    docker run --rm -v .:/dir stackrox/kube-linter lint /dir

running:

    kubectl create -f pod-with-exposed-port.yaml
    kubectl get pods
    kubectl describe pods
    kubectl port-forward pod/port-exposed-pod 8080

    kubectl delete pod command-pod
