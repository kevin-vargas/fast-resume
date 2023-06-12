#!/bin/bash
# Need to login to registry.cloud.okteto.net
# Need cluster config for cloud okteto

DIRECTORY=$(pwd)
AUTHORIZATION_SERVER="${DIRECTORY}/authorization-server"
BFF_SERVER="${DIRECTORY}/bff-server"
SYNT_SERVER="${DIRECTORY}/synthesizer-server"
WEB_SERVER="${DIRECTORY}/web-server"
SECRET_MANIFEST="${DIRECTORY}/deploy/secrets.yaml"
DEPLOY_MANIFEST="${DIRECTORY}/deploy/deployment.yaml"
buildAndPushBackend (){
   local REGISTRY=registry.cloud.okteto.net/kevin-vargas
   local IMAGE_TAG="${REGISTRY}/$1"
   echo $IMAGE_TAG
   echo $2
   docker buildx build --platform linux/amd64 -t $IMAGE_TAG $2
   docker push $IMAGE_TAG
}

apply() {
    kubectl apply -f $1
}

log() {
    echo -e "${1}"
}

logt() {
    log "\t ♥♥♥ ${1} \t ♥♥♥"
}

{

    logt "building backend and pusing image"
    buildAndPushBackend authorization-server-fast $AUTHORIZATION_SERVER
    #buildAndPushBackend bff-server-fast $BFF_SERVER
    #buildAndPushBackend synthesizer-server-fast $SYNT_SERVER
    #buildAndPushBackend web-server-fast $WEB_SERVER

    logt "deploying kubernetes objects"
    apply $SECRET_MANIFEST
    apply $DEPLOY_MANIFEST
}