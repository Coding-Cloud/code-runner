FROM golang as build

WORKDIR /usr/app/code-runner
ENV GOPATH /usr/app

COPY go.mod go.sum ./
COPY main.go .
COPY code_runner_server/ code_runner_server/

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/code-runner

FROM nospy/pnpm:16
COPY --from=build /usr/local/bin/code-runner /usr/local/bin/code-runner

RUN pnpm config set auto-install-peers=true
RUN pnpm config set store-dir=/root/.local/share/pnpm
RUN pnpm install -g @nestjs/cli

RUN mkdir /home/app
WORKDIR /home/app

COPY ./code_runners/nest-runner/start.sh /home/start.sh
COPY ./code_runners/nest-runner/stop.sh /home/stop.sh
COPY ./code_runners/nest-runner/install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV GIN_MODE=release
ENV SCRIPTS_PATH=/home
ENV RUNNER_PORT=8181

EXPOSE 8181

EXPOSE 3000
VOLUME /home/app
VOLUME /logs
VOLUME /root/.local/share/pnpm
ENTRYPOINT ["code-runner"]
