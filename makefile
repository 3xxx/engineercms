BUILD_VERSION	:= 1.0.0
BUILD_TIME		:= $(shell date "+%F %T")
BUILD_NAME		:= go-version-sample
SOURCE			:= ./*.go
TARGET_DIR		:= /path-you-want/${BUILD_NAME}

all:
	go build -ldflags \
	"-X ${BUILD_NAME}/version.BuildVersion=${BUILD_VERSION} \
	-X '${BUILD_NAME}/version.BuildTime=${BUILD_TIME}' \
	-X ${BUILD_NAME}/version.BuildName=${BUILD_NAME}" \
	-o ${BUILD_NAME} ${SOURCE}

clean:
	rm ${BUILD_NAME} -f

install:
	# mkdir -p ${TARGET_DIR}
	# cp ${BUILD_NAME} ${TARGET_DIR} -f

.PHONY : all clean install ${BUILD_NAME}
