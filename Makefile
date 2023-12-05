test: 
	{ \
	cd ./code/backend;\
	go test ./...;\
	}

build-backend:
	{ \
	cd ./code/backend;\
	go build -o bin/waterJusgChallenge ;\
	} 

run-backend:
	{ \
	cd ./code/backend;\
	go run main.go --port=$(port);\
	} 

run:
	docker compose -f ./docker-compose.yml up -d --build

stop:
	docker compose -f ./docker-compose.yml down

