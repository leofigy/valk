#!/usr/bin/env sh

# Creates a cluster of K3S

MASTER=${1-MASTER}
WORKER1=${2-WORKER1}
WORKER2=${3-WORKER2}
# mem defaults 
MEM=${4-1GB}

multipass launch --name $MASTER  --mem=${MEM};
multipass launch --name $WORKER1 --mem=${MEM};
multipass launch --name $WORKER2 --mem=${MEM};

multipass exec $MASTER -- /bin/bash -c "curl -sfL https://get.k3s.io | sh -";
K3S_NODEIP_MASTER="https://$(multipass info $MASTER | grep "IPv4" | awk -F' ' '{print $2}'):6443";
K3S_TOKEN="$(multipass exec $MASTER -- /bin/bash -c "sudo cat /var/lib/rancher/k3s/server/node-token")";
multipass exec $WORKER1 -- /bin/bash -c "curl -sfL https://get.k3s.io | K3S_TOKEN=${K3S_TOKEN} K3S_URL=${K3S_NODEIP_MASTER} sh -";
multipass exec $WORKER2 -- /bin/bash -c "curl -sfL https://get.k3s.io | K3S_TOKEN=${K3S_TOKEN} K3S_URL=${K3S_NODEIP_MASTER} sh -";

# copying K3s info
CONFIG="$(multipass exec $MASTER -- /bin/bash -c "sudo cat /etc/rancher/k3s/k3s.yaml")";
mkdir -p ~/.kube
echo "$CONFIG" > ${HOME}/.kube/k3s.yaml;
sed -ie s,https://127.0.0.1:6443,${K3S_NODEIP_MASTER},g ${HOME}/.kube/k3s.yaml;
kubectl --kubeconfig=${HOME}/.kube/k3s.yaml get nodes;
