all: build install service

install:
	cp -f supervisor.ini /etc/supervisor.d/wg-port.ini

build:
	go build -o wg-port

service:
	supervisorctl restart wg-port