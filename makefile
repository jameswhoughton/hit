MAIN_PACKAGE_PATH=./cmd/hit
BINARY_NAME=hit

build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}
