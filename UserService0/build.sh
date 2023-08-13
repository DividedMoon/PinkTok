#!/bin/bash

cd ./UserService || exit
hz update --idl ./idl/user_service.proto --handler_by_method
hz client --idl ./idl/user_service.proto
echo "UserService build success"
cd ../RelationService || exit
hz update --idl ./idl/relation_service.proto --handler_by_method
hz client --idl ./idl/relation_service.proto
echo "RelationService build success"
