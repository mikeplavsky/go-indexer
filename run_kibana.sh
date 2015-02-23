docker run -d \
-p 5601:5601 \
--name=kibana --link=es0:es mikeplavsky/kibana
