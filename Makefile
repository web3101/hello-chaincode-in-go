export GOOS=linux
export GOARCH=amd64
export GOBIN=./bin


regres_v2:
	go build -o bin/ chaincodes/regress_v2.go