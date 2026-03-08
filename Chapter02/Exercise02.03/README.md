# The Kubernetes Workshop - Chapter 2 - Exercise 02.03

linting:

    docker run -v .:/dir stackrox/kube-linter lint /dir

setup load balancer: https://github.com/kubernetes-sigs/cloud-provider-kind

    kubectl cluster-info --context kind-kind
    kubectl get pod --show-labels
    kubectl apply -f k8s-for-beginners-pod.yaml
    kubectl apply -f k8s-for-beginners-svc.yaml
    kubectl get service
    kubectl describe svc k8s-for-beginners

port forwarding to svc:

    kubectl port-forward service/k8s-for-beginners 9090:80

install load balancer:

    kubectl label node kind-control-plane node.kubernetes.io/exclude-from-external-load-balancers-
    go install sigs.k8s.io/cloud-provider-kind@latest
    cloud-provider-kind
