#!/bin/sh
mvn dependency:resolve 2>&1 | tee -a "/logs/${PROJECT_ID}.log"