curl -XPOST 'localhost:8080/_bulk' -d '
{ "index" : { "_index" : ".kibana", "_type" : "config", "_id" : "4.0.0" } }
{ "buildNum": 5888, "defaultIndex": "test*" }

{ "index" : { "_index" : ".kibana", "_type" : "index-pattern", "_id" : "test*" } }
{"title":"test*","timeFieldName":"@timestamp","customFormats":"{}","fields":"[{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_source\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"msg_type\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_index\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"pid\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"mbx\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"msg\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"agent\",\"count\":0,\"scripted\":false},{\"type\":\"date\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"@timestamp\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"name\":\"_type\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_id\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"path\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"tid\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"coll\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"fileId\",\"count\":0,\"scripted\":false}]"}

{ "index" : { "_index" : ".kibana", "_type" : "index-pattern", "_id" : "s3data" } }
{"title":"s3data","timeFieldName":"@timestamp","customFormats":"{}","fields":"[{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_source\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_index\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"customer\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"uri\",\"count\":0,\"scripted\":false},{\"type\":\"number\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"size\",\"count\":0,\"scripted\":false},{\"type\":\"date\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"@timestamp\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"name\":\"_type\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_id\",\"count\":0,\"scripted\":false},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"fileId\",\"count\":0,\"scripted\":false}]"}


{ "index" : { "_index" : ".kibana", "_type" : "visualization", "_id" : "TotalFiles", "_score" : 1.0 } }
{"title":"TotalFiles","visState":"{\"type\":\"metric\",\"params\":{\"fontSize\":60},\"aggs\":[{\"id\":\"2\",\"type\":\"cardinality\",\"schema\":\"metric\",\"params\":{\"field\":\"fileId\"}}],\"listeners\":{}}","description":"","version":1,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"index\":\"test*\",\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}},\"filter\":[]}"}}


{ "index" : { "_index" : ".kibana", "_type" : "visualization", "_id" : "TotalEvents", "_score" : 1.0 } }
{"title":"TotalEvents","visState":"{\"type\":\"metric\",\"params\":{\"fontSize\":\"23\"},\"aggs\":[{\"id\":\"1\",\"type\":\"count\",\"schema\":\"metric\",\"params\":{}}],\"listeners\":{}}","description":"","version":1,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"index\":\"test*\",\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}},\"filter\":[]}"}}
    


{ "index" : { "_index" : ".kibana", "_type" : "visualization", "_id" : "mailboxlist", "_score" : 1.0 } }
{"title":"mailboxlist","visState":"{\"type\":\"table\",\"params\":{\"perPage\":100,\"showPartialRows\":false,\"showMeticsAtAllLevels\":false},\"aggs\":[{\"id\":\"1\",\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"type\":\"terms\",\"schema\":\"bucket\",\"params\":{\"field\":\"mbx\",\"size\":1000,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}","description":"","version":1,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"index\":\"test*\",\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}},\"filter\":[]}"}}


{ "index" : { "_index" : ".kibana", "_type" : "visualization", "_id" : "NMailboxes", "_score" : 1.0 } }
{"title":"NMailboxes","visState":"{\"type\":\"metric\",\"params\":{\"fontSize\":60},\"aggs\":[{\"id\":\"1\",\"type\":\"cardinality\",\"schema\":\"metric\",\"params\":{\"field\":\"mbx\"}}],\"listeners\":{}}","description":"","version":1,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"index\":\"test*\",\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}},\"filter\":[]}"}}







{ "index" : { "_index" : ".kibana", "_type" : "dashboard", "_id" : "Main", "_score" : 1.0 } }
{"title":"Main","hits":0,"description":"","panelsJSON":"[{\"col\":4,\"id\":\"mailboxlist\",\"row\":1,\"size_x\":5,\"size_y\":7,\"type\":\"visualization\"},{\"col\":1,\"id\":\"NMailboxes\",\"row\":1,\"size_x\":3,\"size_y\":2,\"type\":\"visualization\"},{\"col\":1,\"id\":\"TotalFiles\",\"row\":3,\"size_x\":3,\"size_y\":2,\"type\":\"visualization\"}]","version":1,"kibanaSavedObjectMeta":{"searchSourceJSON":"{\"filter\":[{\"query\":{\"query_string\":{\"analyze_wildcard\":true,\"query\":\"*\"}}}]}"}}




'
