#!/bin/sh
gcloud auth activate-service-account --key-file=/secrets/skaffold-service-account.json
gcloud config set project $PROJECT_DEV_NAME
gcloud container clusters get-credentials $CLUSTER_DEV_NAME --region $PROJECT_DEV_REGION
kubectl config get-contexts 
export STATE=$(git rev-list -1 HEAD --abbrev-commit)  
skaffold build -p prod --file-output build-$STATE.json
skaffold deploy -a build-$STATE.json