# Gobblr

Make your development databases gobble up known data

<!-- toc -->

- [Quickstart](#quickstart)
  * [CLI](#cli)
  * [Web server](#web-server)
- [Why?](#why)
  * [Why might you need test data?](#why-might-you-need-test-data)

<!-- tocstop -->

## Quickstart

### CLI

```bash
gobblr db sql mysql -u <user> -p <password> -d <database> --path /path/to/data
```

### Web server

```bash
gobblr db sql mysql -u <user> -p <password> -d <database> --path /path/to/data --run
```

## Why?

I've worked on so many projects where there's a need to have a known data set for development
and testing purposes. Most database ORMs provide a way of migrating the database structure and
seed data, but they can be a bit difficult to target data specifically for testing purposes.

I've used this basic approach for years where the data is driven from JSON files or JS scripts,
but this is a way of making this cross-platform.

**NB**. This should not be used for populating production databases. If you need to seed your
production databases with data, use an ORM for the language you're using.

### Why might you need test data?

Having a known data set allows developers to get started fast. If you want to see how a `GET`
endpoint works, you need data in there. If you want to update or delete a record, you need
pre-existing data.

If you're writing integration/end-to-end tests (and you probably should be for most web apps)
then you need a known data set so your tests pass. And of course, when writing these types of
tests, you want to ensure that your data set is reset before each run - this ensures truly
isolated tests and avoids the problem of having a suite of 10,000 tests failing, but passing
when run individually.
