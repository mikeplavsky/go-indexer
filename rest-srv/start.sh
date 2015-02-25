#/bin/sh                                                                                                                                               
                                                                                                                                                       
echo "starting indexing..."                                                                                                                            

ORIGIN_DIR=$(pwd)
   
SCRIPT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $SCRIPT_DIR/..

docker build -t go_indexer .

if [ -z "$(docker images | grep go_indexer)" ]  ; then
	#todo:retry
	echo "failed to build"
	exit 1
fi

# todo:remove hardcode
if [ -z "$(docker ps -a | grep es[0-9])" ]  ; then
	echo "creating stask"
	./create_stack.sh;
fi


if [ -z "$(docker ps -a | grep loader[0-9])" ]  ; then
	echo "creating loaders"
        ./create_loaders.sh;
fi

./start_stack.sh
./start_loaders.sh
while ! curl localhost:8080/_cat/nodes?v
do
    echo -n .
    sleep 1
done

if [ -z "$(docker ps | grep kibana)" ]  ; then
	./run_kibana.sh
fi

cd $ORIGIN_DIR                                                                                                                                         
