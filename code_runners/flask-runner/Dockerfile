FROM golang as build

WORKDIR /usr/app/code-runner
ENV GOPATH /usr/app

COPY go.mod go.sum ./
COPY main.go .
COPY code_runner_server/ code_runner_server/

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/code-runner

FROM python:3.9.13-slim
COPY --from=build /usr/local/bin/code-runner /usr/local/bin/code-runner

RUN pip install Flask --target=/libs

RUN mkdir /home/app
WORKDIR /home/app

COPY ./code_runners/flask-runner/start.sh /home/start.sh
COPY ./code_runners/flask-runner/stop.sh /home/stop.sh
COPY ./code_runners/flask-runner/install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV PROJECT_ID="project_id"
ENV FLASK_APP main.py
ENV FLASK_ENV development
ENV PYTHONPATH /libs
ENV PORT 5000
ENV HOST 0.0.0.0

ENV GIN_MODE=release
ENV SCRIPTS_PATH=/home
ENV RUNNER_PORT=8181

EXPOSE 8181

EXPOSE 5000
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /libs
ENTRYPOINT ["code-runner"]
