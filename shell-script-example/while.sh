#!/usr/bin/env bash
a=5
while (($a > 1)); do
    echo ${a}
    a=`expr ${a} - 1`
done