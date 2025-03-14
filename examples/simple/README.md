# Simple Example

The simple example exists to show how Gobblr works.

* `data`: the data that will be ingested to the database.
* `migrations`: a [GORM migration](https://gorm.io/docs/migration.html) script to
  build the SQL databases - this will be handled by your database library (which
  may or may not be GORM).

You can use [Docker Compose](https://docs.docker.com/compose) to run the example
stack:

```shell
docker-compose run --rm --service-ports gobblr-<db-type>
```
