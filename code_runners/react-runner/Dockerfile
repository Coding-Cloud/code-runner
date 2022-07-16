FROM golang as build

WORKDIR /usr/app/code-runner
ENV GOPATH /usr/app

COPY go.mod go.sum ./
COPY main.go .
COPY code_runner_server/ code_runner_server/

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/code-runner

FROM nospy/pnpm:16
RUN pnpm config set store-dir=/root/.local/share/pnpm

RUN mkdir /home/app

WORKDIR /home/app

COPY ./code_runners/react-runner/start.sh /home/start.sh
COPY ./code_runners/react-runner/install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV HOST=0.0.0.0
ENV PROJECT_ID=""
ENV PUBLIC_URL="localhost"
ENV PORT=3000

ENV GIN_MODE=release
ENV SCRIPTS_PATH=/home
ENV RUNNER_PORT=8181

EXPOSE 8181

EXPOSE 443
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
ENTRYPOINT ["code-runner"]
