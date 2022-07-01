#!/bin/bash

# generate k8s deployment yaml
goctl kube deploy -name deployment-my-zero -namespace default -image deployment-my-zero -o deployment-my-zero.yaml -port 80 --serviceAccount find-endpoints
#goctl kube deploy -secret docker-login -replicas 1 -nodePort 30080 -requestCpu 50 -requestMem 100 -limitCpu 500 -limitMem 1000 -name deployment-my-zero -namespace default -image my-zero -o deployment-my-zero.yaml -port 80 --serviceAccount find-endpoints

# generate this service configmap
kubectl create configmap conf-my-zero --from-file=./../my_zero/etc/my-zero.k8s.yaml
