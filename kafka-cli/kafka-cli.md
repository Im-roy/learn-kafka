# Kafka topics

### To list down name of all the topics 
```kafka-topics.sh --bootstrap-server localhost:9092 --list```

### create a topic with 
> topic name : first_topic \
> No. of partitions : 1 \
> Replication-factor : 1 

```
$ kafka-topics.sh --bootstrap-server localhost:9092 --create --topic first_topic --partitions 1 --replication-factor 1

WARNING: Due to limitations in metric names, topics with a period ('.') or underscore ('_') could collide. To avoid issues it is best to use either, but not both.
Created topic first_topic.
```

Warning says that if you go on creating a topic with name first.name it will not be created bcz of collision issue.

By default number of partitions and replication factor is 1.

### To describe a topic

```
kafka-topics.sh --bootstrap-server localhost:9092 --describe
Topic: first_topic	TopicId: -c0dnmIqQ7KYE3UMecdWfA	PartitionCount: 1	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: first_topic	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
Topic: second_topic	TopicId: 88hO5LRBTmCBmV-GmI_QHA	PartitionCount: 1	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: second_topic	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
Topic: third_topic	TopicId: 2PyadBtqQCONjrLylMjO3g	PartitionCount: 3	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: third_topic	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	Topic: third_topic	Partition: 1	Leader: 0	Replicas: 0	Isr: 0
	Topic: third_topic	Partition: 2	Leader: 0	Replicas: 0	Isr: 0
```

### How to delete a topic 

This command will delete a topic named first_topic
```
kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic first_topic
```