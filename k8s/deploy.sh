#!/bin/bash -e

kubectl apply -f ld-service.yml
kubectl apply -f ld-deployment.yml

