#!/bin/bash

appName="app-admin"
pid=`ps aux|grep $appName|grep -v grep|awk '{print $2}'`
kill -9 $pid
chmod 777 $appName
time=$(date "+%Y%m%d")
#创建文件夹
if [ -d "./logs/$time" ];then
      nohup ./$appName >>logs/$time/$time.log 2>&1 &
    else
        mkdir -p ./logs/$time/
        nohup ./$appName >>logs/$time/$time.log 2>&1 &
    fi