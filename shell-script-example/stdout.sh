#!/usr/bin/env bash

#输出重定向覆盖
who > users
#top > top
#输出重定向追加
#a=1
#while true; do
#    a=`expr $a + 2`
#    echo ${a} >> a
#done


# dev/null 不输出
echo "abc" > /dev/null
#屏蔽 stdout 和 stderr
eda  > /dev/null 2>&1
