#!/bin/bash
sh configmap/configmap.sh

for i in `ls`;
do
kubectl apply -f $i -n tekton-pipelines
done