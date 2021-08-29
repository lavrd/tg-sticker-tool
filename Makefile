modules:
	go mod tidy
	go mod vendor

run:
	go run . -src ${SRC} -dst ${DST}
