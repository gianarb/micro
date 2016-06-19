.PHONY: compile

compile:
	rm -rf micro_*
	env GOOS=linux GOARCH=386 go build
	mv micro micro_linux_386
	docker build -t gianarb/micro .
