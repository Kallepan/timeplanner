#!/bin/bash

export $(grep -v '^#' .staging.env | xargs)

kubectl create namespace $NAMESPACE --kubeconfig=$KUBECONFIG

kubectl delete secret secrets -n $NAMESPACE --kubeconfig=$KUBECONFIG
kubectl create secret generic secrets -n $NAMESPACE \
    --from-env-file=.staging.env \
    --kubeconfig=$KUBECONFIG
kubectl delete secret regcred -n $NAMESPACE --kubeconfig=$KUBECONFIG
kubectl create secret docker-registry -n $NAMESPACE \
    regcred \
    --docker-server=$DOCKER_REGISTRY_SERVER \
    --docker-username=$DOCKER_REGISTRY_USERNAME \
    --docker-password=$DOCKER_REGISTRY_PASSWORD \
    --kubeconfig=$KUBECONFIG

# Deploy to staging
echo $DOCKER_REGISTRY_PASSWORD | docker login -u $DOCKER_REGISTRY_USERNAME --password-stdin

cd api-gateway
docker build --platform linux/amd64 -t $DOCKER_REGISTRY_SERVER/$DOCKER_REGISTRY_USERNAME/timeplanner-api-gateway:$VERSION .
docker push $DOCKER_REGISTRY_USERNAME/timeplanner-api-gateway:$VERSION

cd ../planner-backend
docker build --platform linux/amd64 -t $DOCKER_REGISTRY_SERVER/$DOCKER_REGISTRY_USERNAME/timeplanner-planner-backend:$VERSION .
docker push $DOCKER_REGISTRY_USERNAME/timeplanner-planner-backend:$VERSION

cd ../frontend
docker build --platform linux/amd64 -t $DOCKER_REGISTRY_SERVER/$DOCKER_REGISTRY_USERNAME/timeplanner-planner-frontend:$VERSION .
docker push $DOCKER_REGISTRY_USERNAME/timeplanner-planner-frontend:$VERSION


kubectl rollout restart deployment api-gateway -n $NAMESPACE --kubeconfig=$KUBECONFIG
kubectl rollout restart deployment planner-backend -n $NAMESPACE --kubeconfig=$KUBECONFIG
kubectl rollout restart deployment planner-frontend -n $NAMESPACE --kubeconfig=$KUBECONFIG

cd infrastructure/staging
kubectl kustomize . --kubeconfig=$KUBECONFIG | kubectl apply -f - --kubeconfig=$KUBECONFIG