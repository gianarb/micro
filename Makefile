.PHONY: compile

compile:
	rm -rf micro_*
	env GOOS=linux GOARCH=386 go build -o micro
	mv micro micro_linux_386
	docker build -t gianarb/micro .
