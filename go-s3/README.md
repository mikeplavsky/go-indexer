# go-s3

You need to set environment variables:

```
export AWS_ACCESS_KEY_ID=<YOUR ID>
export AWS_SECRET_ACCESS_KEY=<YOUR SECRET>
```

Then just start it:

```
./go-s3 -bucket_name=<s3 bucket> [-folder_name=<folder>]
```

If you want to build it and run tests:

```
docker run -ti --rm -v $(pwd):/go/src/go-s3 mikeplavsky/docker-golang
go get -d
go test
```
