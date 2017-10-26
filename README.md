# XDCRWithKafka

XDCR(Cross datacenter replication) is unstable for wan network environment, and network delay will be more than seconds.
In this case, a message queue can be transform the data in a reliable way.

## XDCR for redis using kafka

- Producer:
feed stream data use appendonly file and produce a message to kafka topic

- Consumer:
consume message from kafka topic and push data to sink cluster
