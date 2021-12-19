#!/bin/bash
appName="app-admin"
pid=`ps aux|grep $appName|grep -v grep|awk '{print $2}'`
kill -9 $pid