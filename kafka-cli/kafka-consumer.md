# kafka consumer

```
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic third_topic
```
This will start reading the kafka topic from tail of the topic.

```
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic third_topic --from-beginning
Hello disha...!
Im writing on third topic
 example value
Hello consumer
testig
Hello topic
```

### To Read kafka topic with timestamp key and value.

```
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic third_topic --from-beginning --formatter kafka.tools.DefaultMessageFormatter --property print.timestamp=true --property print.key=true --property print.value=true
CreateTime:1665806866772	null	Hello disha...!
CreateTime:1665806910535	null	Im writing on third topic
CreateTime:1665807709884	example key	 example value
CreateTime:1665808285751	null	Hello consumer
CreateTime:1665808312706	null	testig
CreateTime:1665808369370	null	Hello topic
```

