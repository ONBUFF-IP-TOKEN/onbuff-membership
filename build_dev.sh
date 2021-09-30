set -x

sh ./prebuild_dev.sh

go build -o bin/onbuff-membership rest_server/main.go

cd bin
./onbuff-membership -c=config.yml