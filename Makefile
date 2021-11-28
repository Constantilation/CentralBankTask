BINARY=engine
engine:
	go build -o ${BINARY} app/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t interviewtask .

run:
	docker-compose up --build -d

stop:
	docker-compose down

.PHONY: clean install unittest build docke