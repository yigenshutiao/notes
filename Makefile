
clean:
	rm -rf note

build:
	go build -tags="Pprof" -gcflags="-N -l" -o note ./
