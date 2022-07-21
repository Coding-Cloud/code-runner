#!/bin/sh
kill -9 `ps faux | grep mvn | awk -F' ' '{print $2}'`
