
- docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine

- docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:12.5-alpine

- migrate -path db/ -database postgres://user:password@localhost:5432/dbname?sslmode=disable  force 15

[Populate DB]
- go install github.com/MarioCarrion/complex-pipelines/part5
- DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable" part5
