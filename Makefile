build:
	go build -buildmode=c-shared -o lib.so main.go

run:
	python3 test.py

clean:
	rm -rf lib.h lib.so