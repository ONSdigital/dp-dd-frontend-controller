EXECUTABLE=dp-dd-frontend-controller

build:
	go build -o build/${EXECUTABLE}

debug: build
	HUMAN_LOG=1 ./build/${EXECUTABLE}

.PHONY: build debug
