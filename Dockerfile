FROM public.ecr.aws/bitnami/golang:1.16 as build-image

WORKDIR /go/src
COPY go.mod go.sum ./
COPY backlog_to_slack/main.go ./

RUN go build -o ../bin

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["notification_slack"]