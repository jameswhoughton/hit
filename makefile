MAIN_PACKAGE_PATH=./cmd/crow
BINARY_NAME=crow

build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}
