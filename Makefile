all: build install service

install:
	cp -f supervisor.ini /etc/supervisord.d/wg-port.ini

build:
	go build -o wg-port

service:
	supervisorctl update
	supervisorctl restart wg-port