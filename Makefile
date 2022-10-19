.SILENT:

build:
	go get gopkg.in/gomail.v2 && go build -o mailer-k .

run: build
	./mailer-k