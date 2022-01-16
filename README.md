

To build docker file locally, for example.
```bash
docker build -t akedev7/go-ms-bff:latest -f "bff/Dockerfile" .
```

Then you can run using

```bash
docker network create akedev7
docker-compose up -d
```