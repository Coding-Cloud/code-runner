FROM maven:3.8.5-eclipse-temurin-17-alpine

RUN mkdir /home/app

WORKDIR /home/app

COPY ./quarkus-starter.sh /home/quarkus-starter.sh
RUN chmod +x /home/quarkus-starter.sh

ENV PROJECT_ID=""
ENV PUBLIC_URL=http://localhost:8080
ENV PORT=8080

EXPOSE 8080
VOLUME /home/app
VOLUME /logs/$PROJECT_ID.log
VOLUME /root/.m2
ENTRYPOINT ["sh", "/home/quarkus-starter.sh"]
