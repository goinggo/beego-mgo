export MGO_HOSTS=ds035428.mongolab.com:35428
export MGO_DATABASE=goinggo
export MGO_USERNAME=guest
export MGO_PASSWORD=welcome
export BUOY_DATABASE=goinggo

cd $GOPATH/src/github.com/goinggo/beego-mgo/test/serviceTests
go test -v