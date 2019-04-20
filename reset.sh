#!/bin/sh
#move to directory for this file exists
cd `dirname $0`
#remove container
docker rm -f cown
#build image from Dockerfile
docker build ./ -t chrome_chat_server
#run server
docker run --env-file .env --name cown -v $(pwd)/src:/root/chrome_chat_server -ti -p 8080:8080 chrome_chat_server
#show names
docker ps -q

#
echo "look at localhost:8080 !!!"
