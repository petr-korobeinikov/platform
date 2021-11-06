# Building with stdin

```shell
cat ../../../hack/dockerfile/go/Dockerfile | docker build -t service -f - .
```

```shell
docker run --rm -p 9000:9000 --name service service
```
