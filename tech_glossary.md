# Tech Glossary

## A

## B

## C

## D

### DynamoDB (AWS)

A Key-Value and Document type database. See [Databases](Databases.md)

## E

### ElasticSearch

A Search type database. See [Databases](Databases.md)

## F

## G

## H

## I

## J

## K

### Kafka

(Apache)

Event based streaming system. [Message Broker](#message-broker)

Uses [PubSub](#pubsub)

[Intro](https://kafka.apache.org/intro)

Data stored in logs sorted into streams called Topics.

Topics can be large or small, can remember things for a short time or forever.

Tends to be used for microservice architectures. Different services process and/or produce events from one or more Topics.

Kafka Connect used to collect and export data from external data sources and to external outputs. Plug and play connectors that convert source to a topic and convert a topic to the output.

Kafka Streams - Java API for doing various things such as:

- Grouping
- Aggregating
- Filtering
- Enriching (combining Topics, like Joins in SQL)

### Kinesis

(AWS)

[Message Broker](#message-broker)

## L

## M

### Message Queue

See [Queue](#queue)

### Message Broker

Manages one or more [Queues](#queue) to for example:

- allow [PubSub](#pubsub) pattern
- smooth data rate during spikes
- guard against misbehaving hosts
- translate between protocols

See [article](message_[brokers.png) from [here](https://dzone.com/articles/evaluating-message-brokers-kafka-vs-kinesis-vs-sqs)

Such as: [Kafka](#kafka), [Kinesis](#kinesis), [SQS](#sqs)

## N

## O

## P

### PubSub

Publisher - Subscriber pattern.

Compared to Observer pattern, this involves a intermediate, the [Message Broker](#message-broker). The Publishers and Subscribers do not communicate directly and don't know about each other.

Asynchronous, whereas Observer is Synchronous.

### Postgres

A Relational type database. See [Databases](Databases.md)

## Q

### Queue

First In First Out data structure. Data is is only consumed once.

## R

### Redis

A Key-Value type database. See [Databases](Databases.md)

Also a [Message Broker](#message-broker).

## S

### SQL

Structured Query Language

Used for Relational Databases. See [Databases](Databases.md)

Such as: [PostgresSQL](#postgres)

### SQLAlchemy

Python wrapper for SQL queries. "SQL toolkit and Object Relational Mapper". See [SQL](#sql)

### SQS

(AWS)

Simple Queuing Service. [Message Broker](#message-broker)

Can smooth data rate

Destroys message once processed from a queue

Single producer, single consumer

For multi consumer need to duplicate queues.

## T

## U

## V

## W

## X

## Y

## Z
