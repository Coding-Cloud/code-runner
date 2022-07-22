#!/bin/sh
pip install -r requirements.txt --target=/libs 2>&1 | tee -a "/logs/${PROJECT_ID}.log"