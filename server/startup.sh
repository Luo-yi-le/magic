#bin/bash

execfile=./magic

pid=`ps ax | grep -i 'magic' | grep -v grep | awk '{print $1}'`
if [ ! -z "${pid}" ] ; then
        echo "重启"
        kill ${pid}
fi

if [ ! -x "${execfile}" ]; then
  sudo chmod +x "${execfile}"
fi

nohup "${execfile}" &

echo "启动"