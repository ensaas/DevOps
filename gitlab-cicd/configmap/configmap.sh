#!/usr/bin/env sh
kubectl delete configmap kubeconfig
kubectl create configmap kubeconfig --from-file=configmap/ews

