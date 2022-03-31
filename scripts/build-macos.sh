WORKDIR=$(pwd)
cd ..
go build -o bin/hopper-analytics-api.macos main.go
cd $WORKDIR