#!/bin/sh
mvn dependency:resolve 2>&1 | tee "/logs/${PROJECT_ID}.log"