FROM golang as build

WORKDIR /usr/app/code-runner
ENV GOPATH /usr/app

COPY go.mod go.sum ./
COPY main.go .
COPY code_runner_server/ code_runner_server/

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/code-runner

FROM maven:3.8.5-openjdk-17
COPY --from=build /usr/local/bin/code-runner /usr/local/bin/code-runner
RUN yum install -y procps

RUN mkdir /home/app

WORKDIR /home/app

COPY ./code_runners/quarkus-runner/start.sh /home/start.sh
COPY ./code_runners/quarkus-runner/stop.sh /home/stop.sh
COPY ./code_runners/quarkus-runner/install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV PORT=8080

ENV GIN_MODE=release
ENV SCRIPTS_PATH=/home
ENV RUNNER_PORT=8181

EXPOSE 8181

EXPOSE 8080
VOLUME /home/app
VOLUME /logs
VOLUME /root/.m2
ENTRYPOINT ["code-runner"]
