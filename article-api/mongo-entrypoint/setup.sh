#!/bin/bash 
#Please make sure you have bash shell or change path to bash in mac
echo "Creating separate database and users for Articles"
mongo admin --host localhost -u 'test' -p 'test' --eval "db.getSiblingDB('Articles').createUser({ user : 'test', pwd :  'test', roles : [{ role: 'readWrite', db: 'Articles'}] ,passwordDigestor:'server'});"
mongo admin --host localhost -u 'test' -p 'test' --eval "db.getSiblingDB('Articles').createCollection('store');"
echo "Mongo db:Articles and its users created."
