# aws-sns-sqs
Example of using Go with Amazon SNS and SQS

# Setup
1. Create new `.env` config file from dist
```shell
$ cp .env.dist .env
```

1.1 Update `.env` config file and run
```shell
$ . ./.env
```

2. Get dependencies by running
```shell
$ go mod vendor
```

3. Queue list
```shell
$ go run queuelist/main.go
```

4. Topic list
```shell
$ go run topiclist/main.go
```

5. Publish to SNS 
```shell
$ go run publisher/main.go topic <ARN>
```

6. Publish to SQS
```shell
$ go run publisher/main.go queue <ARN>
```

7. Consume from queue
```shell
$  go run consumer/main.go <ARN>
```