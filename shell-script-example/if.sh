#!/usr/bin/env bash
if test "a"; then
    echo true
fi

a=1
b=2
if test $a -gt $b; then
    echo ">"
else
    echo "<="
fi