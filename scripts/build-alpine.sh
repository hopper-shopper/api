WORKDIR=$(pwd)
cd ..
GOOS=linux GOARCH=amd64 go build -o bin/hopper-analytics-api.alpine main.go
cd $WORKDIR