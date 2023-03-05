CURRENT_DIR=$(shell pwd)

# APP=$(shell basename ${CURRENT_DIR})

# APP_CMD_DIR=${CURRENT_DIR}/cmd

# REGISTRY=${REGISTRY}
# TAG=latest
# ENV_TAG=latest
# PROJECT_NAME=${PROJECT_NAME}


go :
	go run ${CURRENT_DIR}/cmd/main.go