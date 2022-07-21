#!/bin/sh
kill -9 `ps faux | grep nest | awk -F' ' '{print $2}'`
