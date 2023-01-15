# !/bin/ash

cd ../../
apk update && apk add git
go get github.com/gorilla/mux
go get github.com/gorilla/context
go get github.com/gorilla/handlers
go get gopkg.in/mgo.v2
go test -v 
status1=$?

if [ $status1 -ne 0 ]; then
  exit $status1
fi
exit 0

