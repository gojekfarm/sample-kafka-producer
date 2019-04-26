# Sample Kafka Producer

Build and push messages to kafka
```
go build && ./sample-kafka-producer --total 20
```

The code pushes TestMessage with fields (order_number, order_url, success) populated.

This is meant for producing data to kafka, as a means for testing with [beast](https://github.com/gojekfarm/beast)
