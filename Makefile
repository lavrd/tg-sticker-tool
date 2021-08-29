BIN=./tg-sticker-tool

modules:
	go mod tidy
	go mod vendor

run:
	go build -mod vendor -o ${BIN}
	${BIN} -src ${SRC} -dst ${DST}
	rm -rf ${BIN}
