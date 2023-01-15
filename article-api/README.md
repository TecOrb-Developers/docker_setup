# Article API test

#Go version 1.10

#Docker version 18.06.0-ce, build 0ffa825

#docker-compose version 1.21.2, build a133471

## Installing / Getting started

1. For setting up go environment , I pulled this script from github: 
   wget https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh
  * Edit VERSION to 1.10 and execute the script "./goinstall.sh --darwin" for mac. 
  * If you already have a go setup or $GOPATH , you can copy article-api/ folder into your $GOPATH/src folder which i will share.
  Now install following packages if not present already:
    
    - go get "github.com/gorilla/mux"
    - go get "github.com/gorilla/context"
    - go get "github.com/gorilla/handlers"
    - go get "gopkg.in/mgo.v2"

2. You need to have docker and docker-compose installed, versions I used are mentioned at top.

3. Use "docker-compose up -d" to bring up two containers for mongo db and article-api app respectively. I have uploaded the article-api image to docker hub so that you can easily download and port my codebase easily anywhere.
4. Alternatively, you can execute "docker image build -t article-api ." to build image yourself. Please make sure office firewall is not blocing dns queries of docker build process. I had to edit daemon.json to add dns and restart docker engine. 

#cat /etc/docker/daemon.json

{
  "dns": ["192.168.1.6"]
}

If you face issues while building docker image , please try dns setting for docker engine. After building this image , you will need to edit image name in docker-compose.yml and start docker-compose. Make sure port 27017 and 8000 are free. Your setup should be ready.

5. To start unit testing of this service , please check /article-api/tests/unit-tests/README.md

## Description of this solution

1. I have tried to adhere to the instructions given in the test link. Three endpoints have been developed.
   * POST /articles 
   * GET /articles/{id}
   * GET /tags/{tagName}/{date}

   Articles will only accept the fields given in test link i.e id, title, date, body, tags. Any other tag will be greeted with 400 bad request message. For getting the articles with a tag and date , I noticed that data format did not have hyphens, so this api will accept 20160922 and not 2016-09-22. As i said i have tried to follow the instructions as much as possible. For third endpoint , I am crafting a custom response struct where articles field is a string array of last 10 ids of articles with tag and date. 

2. Even though I had zilch experience in Golang but the line-"submission written in Go will be looked on favourably" caught my eye. I took it as a challenge to develop API. I took help from various blogs, stackoverflow, github to learn and improve about myriads of roadblock I faced along the way. I had an high level idea of NoSQL and that probably helped in setting the DB. 
  * Error handling:
     *  GET request for non existent article ID will respond with 404 NOT FOUND.
     * POST request for articles with unknown and random fields will respond with 400 BAD REQUEST.
     * POST request for articles with existing article ID will respond with 409 CONFLICT.
     * GET request for tag&date  will respond with 404 NOT FOUND, if there is no such article with given tag and date.
  * Testing strategy:
     * Tried to follow test driven approach of writing unit test cases in parallel to developing api endpoints.
     * Logged few details to help in debugging.
     * Ran "go test -v" from /article-api folder (make sure you have installed go packages i mentioned above). I have used testing package to simulate unit tests. Kept improving after finding issues.
     * Installed mongoshell to manually login to mongodb and verify the documents in collections. 

3. Listing out the api calls for reference purpose, which you can make from local machine once you docker containers are up and running.
   * curl -X POST http://localhost:8000/articles --data '{"id": "someuniqueid", "title": "latest science shows that potato chips are better for you than sugar", "date" : "2016-09-22", "tags" : ["health", "fitness", "media"]}' -v

   * curl -X GET http://localhost:8000/articles/someuniqueid -v

   * curl -X GET http://localhost:8000/tags/health/20160922 -v

 
## Assumptions

1. I have not done thorough negative testing and have tried to stick with specific desired functionality of these API endpoints.
2. I have done development on ubuntu machine and giving you docker image hoping the code will work fine for you on macOs.
3. Assuming the you have installed docker, docker compose and go setup as well.
4. You have internet connectivity :) 





