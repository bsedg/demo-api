FROM golang:1.13 as build-step

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . .

RUN git log -n1 --format=format:'%H'
RUN git rev-parse --abbrev-ref HEAD

RUN GIT_COMMIT=$(git log -n1 --format=format:'%H') \
    GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.gitCommit=${GIT_COMMIT},main.gitBranch=${GIT_BRANCH}" -o demo-api

FROM scratch

COPY --from=build-step /app/demo-api /demo-api

EXPOSE 8080
ENTRYPOINT ["/demo-api"]