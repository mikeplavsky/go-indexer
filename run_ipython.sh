docker rm -f ipython

docker run -d  \
-v $(pwd)/notebook:/notebooks \
-e PASSWORD= \
-e USE_HTTP=1 \
--name=ipython \
-p 3001:8888 \
--link es0:es \
--restart=always ipython/scipyserver
