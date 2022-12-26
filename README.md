# Gobblr

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Maintainer](https://img.shields.io/badge/maintainer-mrsimonemms-blue)](https://github.com/MrSimonEmms)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrsimonemms/gobblr)](https://goreportcard.com/report/github.com/mrsimonemms/gobblr)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/mrsimonemms/gobblr.svg)](https://github.com/mrsimonemms/gobblr)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mrsimonemms/gobblr)

Make your development databases gobble up known data

<!-- toc -->

* [Compatability](#compatability)
* [Data](#data)
  * [Tips](#tips)
  * [Supported formats](#supported-formats)
    * [JSON](#json)
    * [JavaScript](#javascript)
  * [Meta data](#meta-data)
* [Running Gobblr](#running-gobblr)
  * [CLI](#cli)
  * [Docker Compose](#docker-compose)
  * [Web server](#web-server)
  * [Docker](#docker)
* [Resetting data](#resetting-data)
  * [CLI](#cli-1)
  * [Web server](#web-server-1)
  * [Use in integration/end-to-end tests](#use-in-integrationend-to-end-tests)
* [Why?](#why)
  * [Why might you need test data?](#why-might-you-need-test-data)
  * [Why shouldn't I use my ORM's migration in development?](#why-shouldnt-i-use-my-orms-migration-in-development)

<!-- Regenerate with "pre-commit run -a markdown-toc" -->

<!-- tocstop -->

## Compatability

This supports the following databases:

 - [MongoDB](https://www.mongodb.com)
 - [MySQL](https://www.mysql.com)
 - [PostgreSQL](https://www.postgresql.org/)
 - [SQLite](https://www.sqlite.org)
 - [SQL Server](https://www.microsoft.com/sql-server)

To request additional drivers, please raise a PR.

## Data

Your test data will live in your repo, usually in a folder called `/data`.

### Tips

 - if you need to link data between datasets, use a hard-coded identifying key
 - any data that matches an [RFC3339](https://pkg.go.dev/time) or YYYY-MM-DD (eg, `2022-10-23`) format is sent as a date object

### Supported formats

#### JSON

This would be used to add static data.

```json
[
  {
    "registered": true,
    "confirmed": true,
    "username": "test1",
    "email_address": "test@test.com",
    "name": "Test Testington",
    "password": "13905934cb5503269ca5613b72a8aa8a:57a1846d12716308be489ec25f2a017e22aba88d591f5c5184e046cb42ec921f1ac5f5249b212efe5c8a644d45d7c63d3ee075cce79ba3a663c731971a56f4c3"
  }
]
```

#### JavaScript

This would be used to add dynamic data and must be returned from a `data` function.

```js
function data() {
  return [
    {
      item: 2,
      some_date: new Date(),
    },
    {
      item: 3,
      some_date: new Date(),
    },
  ]
}

```

### Meta data

By default, each data item will add a `createdAt` and `updatedAt` entry with the current datetime.

If you need to amend this, either to remove either or both keys or to rename them, you can set the data file to include a `meta` object and `data` array. The `data` array will remain the same as per the above examples and the `meta` key controls how the data is added.

```json
{
  "meta": {
    "created": true,
    "createdKey": "createdAt",
    "updated": true,
    "updatedKey": "updatedAt"
  },
  "data": []
}
```

| Key | Type | Default Value | |
| --- | --- | --- | --- |
| `created` | `boolean` | `true` | Controls whether the `createdKey` appears in the dataset |
| `createdKey` | `strng` | `createdAt` | The name of the key |
| `updated` | `boolean` | `true` | Controls whether the `updatedKey` appears in the dataset |
| `updatedKey` | `strng` | `updatedAt` | The name of the key |

## Running Gobblr

For ease, all these example use MongoDB. The principles are the same for all database types - please run `gobblr db --help` to see the documentation for the individual databases.

You will need one instance of Gobblr running for each database you have in your development stack.

For the connection URI, see the [MongoDB docs](https://www.mongodb.com/docs/manual/reference/connection-string).

### CLI

```bash
gobblr db mongodb -u <connection-uri> -d <database> --path /path/to/data
```

### Docker Compose

```yaml
services:
  gobblr-mongodb:
    image: ghcr.io/mrsimonemms/gobblr
    ports:
      - 5670:5670
    environment:
      GOBBLR_CONNECTION_URI: mongodb://mongodb:27017/gobblr
      GOBBLR_DATABASE: gobblr
      GOBBLR_PATH: /data
    links:
      - mongodb
    volumes:
      - ./data:/data # Path to your data files
    restart: on-failure
    command: db mongodb --run

  mongodb:
    image: library/mongo
    ports:
      - 27017:27017
```

### Web server

```bash
gobblr db mongodb -u <connection-uri> -d <database> --path /path/to/data --run
```

### Docker

```bash
docker run -it --rm ghcr.io/mrsimonemms/gobblr \
  db mongodb -u <connection-uri> -d <database> --path /path/to/data --run
```

## Resetting data

The key advantage of Gobblr is that the data can be reset to the known state very quickly. This can be done in two ways:

### CLI

Every time a `gobblr db` command is run, the data is cleared down before the data is ingested again. Just run the command again

### Web server

Gobblr can be run as a web server to make resetting the data possible via an HTTP call. The call is `POST:/data/reset`. This call accepts no arguments and returns the tables affected and gives a count of the data inserted. For example:

```json
[
  {
    "table": "users",
    "count": 1
  },
  {
    "table": "items",
    "count": 2
  }
]
```

### Use in integration/end-to-end tests

This can be used in integration/end-to-end tests to check that your application correctly interacts with the database(s). Most test frameworks have a `beforeEach` type method that runs an arbitrary function prior to each test being executed - you will need to make a call to Gobblr in one of these test hooks to reset the data to ensure the data is consistent.

Here is an example using [Jest](https://jestjs.io). In the your tests, create a file similar to this:

```js
import axios from 'axios';

// Reset the dataset before every test is run
beforeEach(() =>
  axios({
    // GOBBLR_URL envvar will typically be http://localhost:5670 (local running) or http://gobblr:5670 (Docker Compose)
    url: `${process.env.GOBBLR_URL}/data/reset`,
    method: 'post',
    headers: {
      'content-type': 'application/json',
    },
    data: {},
  }),
);
```

## Why?

I've worked on so many projects where there's a need to have a known data set for development and testing purposes. Most database ORMs provide a way of migrating the database structure and seed data, but they can be a bit difficult to target data specifically for testing purposes.

I've used this basic approach for years where the data is driven from JSON files or JS scripts, but this is a way of making this cross-platform.

**NB**. This should not be used for populating production databases. If you need to seed your production databases with data, use an ORM for the language you're using.

### Why might you need test data?

Having a known data set allows developers to get started fast. If you want to see how a `GET` endpoint works, you need data in there. If you want to update or delete a record, you need pre-existing data.

If you're writing integration/end-to-end tests (and you probably should be for most web apps) then you need a known data set so your tests pass. And of course, when writing these types of tests, you want to ensure that your data set is reset before each run - this ensures truly isolated tests and avoids the problem of having a suite of 10,000 tests failing, but passing when run individually.

### Why shouldn't I use my ORM's migration in development?

If your current setup works for you, keep using it.

Typically, ORM migrations require your dependency tree installing before they can run. This will take time that you can avoid with a zero-dependency binary.

In addition to that, you typically don't want to seed your dev/test data into your production seed data and not all ORM migration tools allow you to segregate data by deployment type.

Lastly, ORM migrations are there to migrate data, not set it to a known state. When running an integration test against known data in a database, you should set your databases to a known state before each test it run. With a migration tool, there are usually no guarantees that any changed data has been put back to the known state - with Gobblr, the tables are wiped clean before the data is ingested making certain that your data is in a known state.
