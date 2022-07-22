#!/bin/sh
kill -9 `ps faux | grep java | awk -F' ' '{print $2}'`
