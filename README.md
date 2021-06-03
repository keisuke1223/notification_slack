# Notification_slack
* backlog_to_slack
    * We will notify you at 9 am from Monday to Friday by slack of tickets created on the previous business day .
* 
* 

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

You may need the following for local testing.
* [Golang](https://golang.org)

**set Environment Variable in template.yaml**
```yaml
Environment:
  Variables:
    ...
    SlackChannelName:
    ...
```

## Local development

**Invoking function locally**

```bash
sam build
sam local invoke BacklogToSlackFunction
```

## Packaging and deployment

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

## Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests locally:

```shell
go test -v ./backlog_to_slack/
```