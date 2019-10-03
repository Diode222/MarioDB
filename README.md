# MarioDB
MarioDB is a levelDB-based DBMS (Database Management System). 

MarioDB has implemented basic stand-alone function based on C/S, and is considering joining the distributed service. 
As a k/v database, marioDB supports unstructured data.

## Client
Currently MarioDB has only one test client implemented with go (https://github.com/Diode222/MarioDB_Client), 
it's not a real dababase client, just for testing purposes. But when I finish all the design of MarioDB, 
I will complete the implementation of the client.

## Usage
MarioDB supports both local startup and docker startup with some custom parameters.

### 1. Local startup

####Requires:
##### (1) golang 1.13 (Should work if golang>=1.11 or higher.)
##### (2) GO111MODULE=on (go module environment)
####Steps:
##### (1) go get github.com/Diode222/MarioDB
##### (2) go build main.go (In the path of MarioDB)
##### (3) ./main -h (This command can see the specific meanings of all parameters, printed as follows)

Usage:

    -ip           string    IP                              (Default: 127.0.0.1)
    -port         string    Port                            (Default: 45555)
    -dbPath       string    Database's data path            (Default: /home/diode)
    -maxClient    string    Max count of connected clients  (Default: 100)
    -dbLRUMax     string    Max count of db caches in LRU   (Default: 100)
    
Example: ./main.go -ip 127.0.0.1 -port 50000 -dbPath /home/diode/levelDB_database_root/ -maxClient 20 -dbLRUMax 20

### 2. Docker startup
####Requires:
##### (1) Docker
####Steps:
###### (1) sudo docker build -t mariodb:latest .
###### (2) docker run -dit --network host -v /home/diode/dbdatapath:/root/data mariodb -ip 127.0.0.1 -port 45555 -dbPath /root/data -maxClient 20 -dbLRUMax 20

## TODO
#### 1. Implement a logging system to increase availability.
#### 2. Load balancing, read/write separation.
#### 3. Basic distributed services based on log system and raft.
#### 4. Cluster deployment of distributed services.
