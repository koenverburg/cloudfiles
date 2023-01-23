`docker run -d -p 27017:27017 --name mongo mongo`

with persistent storage
`docker run -d -p 27017:27017 -v /path/to/local/folder:/data/db --name mongo mongo`


Running Ackee

`docker run -d -p 3000:3000 -e ACKEE_MONGODB='mongodb://mongo:27017/ackee' -e ACKEE_USERNAME=admin -e ACKEE_PASSWORD='admin' --link mongo --name ackee electerious/ackee`


