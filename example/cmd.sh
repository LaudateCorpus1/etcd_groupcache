#!/bin/bash
hostip=`ip route show 0.0.0.0/0 | grep -Eo 'via \S+' | awk '{ print $2 }'`
PUBLIC_HOSTIP="$hostip" /go/bin/main
