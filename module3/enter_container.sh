# Check input
docker ps |grep -w $1
if [ $? -ne 0 ];then
echo "Container $1 not found."
exit 1
fi
# Enter container
PID=`docker inspect --format "{{ .State.Pid }}" $1`
nsenter --target ${PID} -n ifconfig
