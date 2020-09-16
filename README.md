# hollein

**AWS Serverless stack(with Golang) to handle my Project.**

features:

- Extract GitHub contribution count every day
- ...

```bash
.
|── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── deploy.sh
├── docker-compose.yml
├── dynamodb                    <-- DynamoDB config
│   └── config.go
├── github                      <-- GitHub Scraper
│   └── client.go
├── go.mod
├── go.sum
├── handler                     <-- Main handler
│   └── handler.go
├── main.go                     <-- EntryPoint
├── model                       <-- Data Definition
│   └── data.go
├── repository                  <-- Hanldle Database Manipulation
│   └── data.go
├── samconfig.toml
└── template.yam
```

## Requirements

- AWS CLI already configured with Administrator permission
- [Docker installed](https://www.docker.com/community-edition)
- [Golang](https://golang.org)
- SAM CLI -
  [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Packaging and deployment

To deploy your application for the first time, run the following in your shell:

```bash
$ make build
$ make deploy
```
