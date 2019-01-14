#!/bin/bash

version=v1.12.2
image=k8s.gcr.io/kube-apiserver:$version
container=kube-apiserver
MASTER_IP=192.168.70.140

docker rm -f $container
docker run -d \
  --name $container \
  -p 6443:6443 \
  -p 8080:8080 \
  -e KUBECONFIG=/etc/kubernetes/admin.conf \
  --link etcd:etcd \
  -v /etc/kubernetes:/etc/kubernetes \
  -v /etc/ssl/certs:/etc/ssl/certs \
  -v /etc/pki:/etc/pki \
  --restart always \
  $image \
  kube-apiserver --logtostderr=true --v=256 --etcd-servers=http://etcd:2379 --insecure-bind-address=0.0.0.0 --secure-port=6443 --port=8080 --bind-address=0.0.0.0 --service-cluster-ip-range=10.254.0.0/16 --enable-admission-plugins=NodeRestriction,NamespaceLifecycle,NamespaceExists,LimitRanger,SecurityContextDeny,ServiceAccount,ResourceQuota --advertise-address=$MASTER_IP --allow-privileged=true --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key --client-ca-file=/etc/kubernetes/pki/ca.crt --service-account-key-file=/etc/kubernetes/pki/sa.pub --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-ca.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-ca.key --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key --enable-bootstrap-token-auth=true --authorization-mode=Node,RBAC