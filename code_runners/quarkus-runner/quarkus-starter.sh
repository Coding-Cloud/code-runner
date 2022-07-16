#!/bin/sh
mvn clean quarkus:dev -Dquarkus.http.host=0.0.0.0 -Dquarkus.http.port=8080 2>&1 | tee "/logs/${PROJECT_ID}.log"
