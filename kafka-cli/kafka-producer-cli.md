# Kafka Producer

```
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic third_topic
```
this will open up a terminal console and you can write on topic

```
>Hello disha...!
>Im writing on third topic
> ^C
```

Use control C to terminate.

### what if topic doesn't exist ??

> If one write to write on a topic which doesnot exist, in this case first topic will get created and data will be written there.\
> Above will genearate warning and this will take some time.

### How to write messages on kafka topic in same partition.

If you have multiple partition inside kafka topic, A message without message key can be wriiten to any of the partition. \
It the message is having key, then all message with same key sits in same partition.

Setting properties for keys
```
kafka-console-producer.sh --bootstrserver localhost:9092 --topic third_topic --property parse.key=true --property key.separator=:
>example key: example value                    
```
Here example key is key and example value is value, Key and value is separated by ':'

In this case if you don't provide key you will get exception in the terminal. saying key and value is expected.
```
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic third_topic --property parse.key=true --property key.separator=:
>asdfghj
org.apache.kafka.common.KafkaException: No key found on line 1: asdfghj
	at kafka.tools.ConsoleProducer$LineMessageReader.readMessage(ConsoleProducer.scala:292)
	at kafka.tools.ConsoleProducer$.main(ConsoleProducer.scala:51)
	at kafka.tools.ConsoleProducer.main(ConsoleProducer.scala)
```




