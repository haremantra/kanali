#!/bin/bash

# silently check if helm is installed
which helm > /dev/null
# install helm if not present
if [ $? != 0 ]; then
   echo "installing helm"
   curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh
   chmod 700 get_helm.sh
   ./get_helm.sh
fi

# deploy tiller
helm init

# add necessary rbac permissions for helm
./scripts/helm-rbac.sh

# add helm repositories for dependencies
helm repo add incubator https://kubernetes-charts-incubator.storage.googleapis.com/

# wait for the tiller deployment to be ready
kubectl rollout status -w deployment/tiller-deploy --namespace=kube-system

# while sleep 1
# do
#     helm install ./helm --name kanali &>/dev/null && break || continue
# done

# install dependencies
helm dep up ./helm

# start kanali and dependencies
helm install ./helm --name kanali

kubectl get pods --all-namespaces

# wait for deployments to be ready
kubectl rollout status -w deployment/kube-dns --namespace=kube-system
kubectl rollout status -w deployment/etcd --namespace=default
kubectl rollout status -w deployment/kanali --namespace=default

kubectl get pods --all-namespaces