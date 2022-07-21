#!/bin/sh
kill -9 `ps faux | grep ng | awk -F' ' '{print $2}'`
