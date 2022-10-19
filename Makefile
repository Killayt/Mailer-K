.SILENT:

build:
	go get gopkg.in/gomail.v2

run: build
	go build -o mailer-k .
	./mailer-k