# The Kubernetes Workshop - Chapter 2 - Exercise 02.05

linting:

    docker run -v .:/dir stackrox/kube-linter lint /dir

    kubectl apply -f k8s-for-beginners-deploy.yaml
    kubectl get service
    kubectl describe svc k8s-for-beginners
    kubectl get deploy
    kubectl delete pod k8s-for-beginners
    kubectl get pod

port forwarding to svc:

    kubectl port-forward service/k8s-for-beginners 9090:80

requests:

    for i in $(seq 1 30); do curl localhost:9090; done

install load balancer:

    kubectl label node kind-control-plane node.kubernetes.io/exclude-from-external-load-balancers-
    go install sigs.k8s.io/cloud-provider-kind@latest
    cloud-provider-kind
