# Databases

[7 Database Paradigms - Fireship](https://www.youtube.com/watch?v=W2Z7fbCLSTw)

There are at least 2 others, such asn Data Warehouse and Time-Series

## Key-Value

eg. Redis, memcached, etcd

SET and GET only

Very fast because usually all data stored in memory, but as a result not a lot can be stored (compared to disk)

No queries, such as joining and searching

Often used as a cache. not as the main database

Best For

- caching
- pub/sub
- leaderboards

## Wide Column

eg. Cassandra, Hbase

each row key points to a varying number of columns
CQL language

No schema -> can handle unstructured data
No joins

Scales Horizontally

Best For

- time series
- historical records
- high write, low read

## Document

eg. MongoDB, firestore, dynamoDB, couchDB

collection of documents

documents are sets of key values

sub collections can be groups of related documents and indexed

schema less and no joins again

unstructured but decent level of relationism.

Best For

- most apps
- games
- IoT

Not good for graphs which need joins

## Relational

eg. postgres, MySQL, SQLServer, CockroachDB

SQL = Structured Query Language

Separate tables for different things but relations between them

primary key for things from this table, foreign key identifying things from other tables

Schema required up front, ie data shape

ACID compliant, mostly difficult to scale (CockroachDB is a modern one that scales well)

Best for Most Apps

Not Ideal for unstructured data

## Graph

eg. Neo4j, dgraph

Where the relations *are* the data

Great for when you have a lot of joins

Best for

- graphs
- knowlegde graphs
- recommendation engines

## Search

eg. Elastic search, Solr, algolia, MeiliSearch

Like Document database, but under the hood the system analyses the document text for key words and indexes them

Best for
typeahead
search engine

## Multi-model

eg. FaunaDB

describe how data will be provided and consumed using GraphQL. system decides how best to implement

ACID compliant

Best for Everything?
