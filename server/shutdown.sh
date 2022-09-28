#bin/bash

pid=`ps ax | grep -i 'magic' | grep -v grep | awk '{print $1}'`
if [ -z "${pid}" ] ; then
        echo "No magic running."
        exit -1;
fi

echo "The magic(${pid}) is running..."

kill ${pid}

echo "Send shutdown request to magic(${pid}) OK"