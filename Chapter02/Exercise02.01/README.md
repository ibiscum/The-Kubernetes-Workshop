# The Kubernetes Workshop - Chapter 02 - Exercise 2.01

    which minikube
    minikube version
    which kubectl

    useradd k8suser
    su - root
    adduser k8suser sudo

    su - k8suser
    sudo usermod -aG docker $USER

    docker context use rootless
    docker info | grep "Docker Root Dir"
    curl -fsSL https://get.docker.com/rootless | sh

    minikube start --driver=docker
    minikube status

    kubectl version
    kubectl get node

    minikube ssh
