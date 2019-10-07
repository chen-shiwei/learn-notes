#!/usr/bin/env bash
echo "脚本名称"$0
echo $#
echo $*
echo $@
echo $-
echo $?
echo "脚本运行的当前进程ID号" $$
echo "后台运行的最后一个进程的ID号"$!
