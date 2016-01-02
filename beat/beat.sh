#!/usr/bin/env bash
### BEGIN INIT INFO
# Provides:           beat
# Required-Start:     $remote_fs $syslog
# Required-Stop:      $remote_fs $syslog
# Default-Start:      2 3 4 5
# Default-Stop:       0 1 6
# Short-Description:  heartbeat.centaurwarchief.com
# Description:        heartbeat.centaurwarchief.com
### END INIT INFO
case "$1" in
  start)
    /home/pi/beat https://heartbeat.centaurwarchief.com &
    ;;
  stop)
    killall beat
    ;;
  *)
    echo "/etc/init.d/beat {start|stop}"
    exit 1
    ;;
esac

exit 0
