FROM maven:3.8.5-eclipse-temurin-17-alpine

RUN mkdir /home/app

WORKDIR /home/app

COPY ./start.sh /home/start.sh
COPY ./install-dependencies.sh /home/install-dependencies.sh
RUN chmod +x /home/start.sh
RUN chmod +x /home/install-dependencies.sh

ENV PROJECT_ID=""
ENV PUBLIC_URL=http://localhost:8080
ENV PORT=8080
ENV SCRIPTS_PATH=/home

EXPOSE 8080
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /root/.m2
ENTRYPOINT ["sh", "/home/quarkus-starter.sh"]
