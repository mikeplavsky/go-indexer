ORIGIN_DIR=$(pwd)

cd ..
./stop_loaders.sh
./stop_stack.sh

docker stop $(docker ps -aq)

echo "Removing containers"
docker rm -f $(docker ps -aq)

if [ -z "$(docker ps -q)" ]  ; then
        echo "containers has been successfully stopped"
else
	echo "some containers are not stopped"
	docker ps
	exit 1
fi

cd $ORIGIN_DIR
