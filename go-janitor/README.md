**Garbage Collector**

As soon as Indexer Instances are terminated there are bunch of SQS queues left. 
go-janitor should be (which is not the case now) run periodically to clean them up.
