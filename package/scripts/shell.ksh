if [ -z "$DOCKER_HOST" ]; then
   DOCKER_TOOL=docker
else
   DOCKER_TOOL=docker-legacy
fi

# set the definitions
INSTANCE=sqs-db-inserter
NAMESPACE=uvadave

$DOCKER_TOOL run -it $NAMESPACE/$INSTANCE /bin/bash -l

# return status
exit $?
