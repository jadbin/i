#!/usr/bin/env bash

bin=`dirname "${BASH_SOURCE-$0}"`
bin=`cd "$bin"; pwd`

kill -9 `ps -A|grep webapp|awk '{print $1}'`
"$bin"/webapp &

