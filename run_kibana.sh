docker create \
-p 5601:5601 \
--name=kibana --link=es0:es --restart=always mikeplavsky/kibana

docker start kibana

./init-kibana-settings.sh
