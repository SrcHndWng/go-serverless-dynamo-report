# go-serverless-dynamo-report

## About This

Sample program for below functions.

- Get data from DynamoDB.
- Put pdf report file on S3.
- Using Serverless Framework for deploy.

## Requirement

* Serverless Framework
* aws cli, configure credential(~/.aws/credentials)
* serverless-dynamodb-local
* go-assets-builder
* gopdf

## Usage

### Before Usage

Download "times.ttf" from below url to ./report/assets/times.ttf, or create ttf file.
("ttf file" is a font format file.)

https://github.com/oneplus1000/gopdfsample/tree/master/ttf

### Build

```
$ make clean
$ make build
```

### Test

Before test, require S3 Bucket(Bucket name is configured in serverless.yml)

```
$ sls dynamodb start
$ make test
```

or

```
$ sls dynamodb start
$ go test -v ./report/...
```

### Deploy, Remove

```
$ make deploy
$ sls remove --verbose
```

### Lambda execute

```
$ sls invoke --function report
```
