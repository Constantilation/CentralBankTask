BINARY=engine
engine:
	CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ${BINARY} app/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t centralbanktask .

run:
	docker-compose up --build -d

stop:
	docker-compose down

runall:
	sudo make stop; sudo docker volume prune; sudo make docker;  sudo docker-compose --verbose up