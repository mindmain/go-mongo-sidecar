#!/bin/bash

export mongo_port=27017
export mongo_url="mongodb://localhost:${mongo_port}"


rs_is_init() {
    v=mongosh $mongo_url --eval "rs.status()" --quiet
    if [ $v == *"no replset config has been received"* ]; then
        return 0
    else
        return 1
    fi
}


echo "Checking if replica set is initialized"
if rs_is_init; then
    echo "Replica set is not initialized"
    echo "Initializing replica set"
    mongosh $mongo_url --eval "rs.initiate()"
    echo "Replica set initialized"
else

    echo "Replica set is already initialized"
fi

mongosh $mongo_url --eval "rs.status()" --quiet

