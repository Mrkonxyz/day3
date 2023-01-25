#!/bin/sh

MONGO_USER="root"
MONGO_PASSWORD="password"
MONGO_HOST="127.0.0.1"
MONGO_DB="sxexpo"
MONGO_COLLECTION="privilege_type"

mongosh "mongodb://$MONGO_USER:$MONGO_PASSWORD@$MONGO_HOST:27017/$MONGO_DB?authSource=admin" <<EOF
db.$MONGO_COLLECTION.insertMany([
    {
        "type": "Non Coupon"
    },
    {
        "type": "Coupon"
    },
    {
        "type": "Information"
    }
])
EOF
exit $?