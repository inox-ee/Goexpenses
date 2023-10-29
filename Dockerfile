FROM golang:1.21 as build

ARG _SLACK_SIGNING_SECRET="hoge"
ARG _SLACK_BOT_TOKEN="fuga"

RUN echo $_SLACK_SIGNING_SECRET

WORKDIR /app
# Copy dependencies list
COPY go.mod go.sum ./
# Build with optional lambda.norpc tag
COPY . .
RUN CGO_ENABLED=0 SLACK_BOT_TOKEN=_SLACK_BOT_TOKEN SLACK_SIGNING_SECRET=_SLACK_SIGNING_SECRET go build -tags lambda.norpc -o main main.go
# Copy artifacts to a clean image

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /app/main ./main
ENTRYPOINT [ "./main" ]
