OBJ_NAME = 
LDFLAGS = 
# install:
# 	$(eval OBJ_NAME += raycasting)
# 	$(eval LDFLAGS += "-w -s")
# 	cd ./cmd/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
run:
	cd ./cmd/api/; go run *.go
migrate-up:
	migrate -path=./migrations -database=postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable up
