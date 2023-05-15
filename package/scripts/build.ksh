if [ -z "$DOCKER_HOST" ]; then
   DOCKER_TOOL=docker
else
   DOCKER_TOOL=docker-17.04.0
fi

# set the definitions
INSTANCE=sqs-db-inserter
NAMESPACE=uvadave

# build the image
$DOCKER_TOOL build -f package/Dockerfile -t $NAMESPACE/$INSTANCE .

# return status
exit $?
