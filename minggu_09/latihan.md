## _215611104 - Elsa Setiyawati_

# LATIHAN

## 1. Langkah pertama yaitu Clone the Lab’s GitHub Repo.

![20](images/1.png)

## 2. Run a single task in an Alpine Linux container

### a. “hostname” untuk mengecek hostname

![20](images/2.png)

### b. docker container run alpine hostname

![20](images/3.png)

### c. docker container ls

![20](images/4.png)

## 3. Run an interactive Ubuntu container

### Berikan perintah :

### a. winpty docker container run --interactive --tty --rm ubuntu bash

### b. ls /

### c. ps aux

### d. cat /etc/issue

![20](images/5.png)

## 4. Run a background MySQL container

### a. Berikan perintah ini pada CLI

docker container run \
 --detach \
 --name mydb \
 -e MYSQL_ROOT_PASSWORD=my-secret-pw \
 mysql:latest
![20](images/6.png)
![20](images/7.png)

### b. docker container ls

### c. docker container logs mydb

### d. docker container top mydb

### e. docker exec -it mydb

mysql --user=root --password=$MYSQL_ROOT_PASSWORD --version
![20](images/8.png)

## 5. Build a simple website image

### a. cd ~/linux_tweet_app

### b. cat Dockerfile

### c. export DOCKERID=your docker id

### d. echo $DOCKERID

### e. docker image build --tag $DOCKERID/linux_tweet_app:1.0 .

### f. masukkan perintah ini pada CLI

docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0

### g. docker container rm --force linux_tweet_app

![20](images/9.png)

![20](images/10.png)
