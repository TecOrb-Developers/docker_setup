Assuming that you have copied my article-api folder inside your $GOPATH/src/.
To start unit tests , go inside directory "$GOPATH/src/article-api/tests/unit-tests" and then run "./start_test.sh"
This start_test.sh script is supposed to spin up a mongo db container and start unit testing. Make sure 27017 port is available.
Results can be seen in STDOUT

