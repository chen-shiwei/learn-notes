#!/usr/bin/env bash
read input
case ${input} in
        1|2|3|4) echo ${input}
        ;;
        6)echo "breack"
        ;;
    esac

#breack 跳出所有循环
while true; do
    read input
    case ${input} in
        1|2|3|4) echo ${input}
        ;;
        *)echo "breack"
        break
        ;;
    esac
done
#continue 跳出当前循环
while true; do
    read input
    case ${input} in
        1|2|3|4) echo ${input}
        ;;
        *)echo "breack"
        continue
        ;;
    esac
done
