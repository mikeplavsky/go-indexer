docker rm -f ipython

docker run -d  \
-v $(pwd)/notebook:/notebooks \
-e PASSWORD= \
-e USE_HTTP=1 \
--name=ipython \
--net=host \
--restart=always ipython/scipyserver
