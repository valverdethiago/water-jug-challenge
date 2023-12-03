test: 
	{ \
	cd ./code/backend;\
	go test ./...;\
	}

build:
	{ \
	cd ./code/backend;\
	go build -o bin/waterJusgChallenge ;\
	} 

run-backend:
	{ \
	cd ./code/backend;\
	go run main.go --port=$(port);\
	} 

