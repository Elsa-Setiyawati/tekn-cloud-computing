## _215611104 - Elsa Setiyawati_

# LATIHAN

# Section #1 - Networking Basics

## 1. Step 1: The Docker Network Command

docker network

![20](images/1.png)

## 2. Step 2: List networks

docker network ls

![20](images/2.png)

## 3. Step 3: Inspect a network

docker network inspect bridge

![20](images/3.png)

## 4. Step 4: List network driver plugins

docker info

![20](images/4.png)

# Section #2 - Bridge Networking

## 5. Step 1: The Basics

1. docker network ls
2. apk update
3. apk add bridge
4. brctl show
5. ip a

![20](images/5.png)

![20](images/6.png)

![20](images/7.png)

![20](images/8.png)

## 6. Step 2: Connect a container

1. docker run -dt ubuntu sleep infinity
2. docker ps
3. brctl show
4. docker network inspect bridge

![20](images/9.png)

![20](images/10.png)

![20](images/11.png)

![20](images/12.png)

## 7. Step 3: Test network connectivity

1. ping -c5 172.17.0.2 .
2. docker ps
3. ping -c5 www.github.com

![20](images/13.png)

![20](images/14.png)

## 8 Step 4: Configure NAT for external connectivity

1. docker run --name web1 -d -p 8080:80 nginx
2. docker ps
3. curl 127.0.0.1:8080

![20](images/15.png)

# Section #3 - Overlay Networking

## 9. Step 1: The Basics

1. docker swarm init --advertise-addr $(hostname -i)
2. docker node ls

![20](images/16.png)
![20](images/17.png)

## 10. Step 2: Create an overlay network

1. docker network create -d overlay overnet
2. docker network ls
3. docker network inspect overnet

![20](images/18.png)

![20](images/19.png)

![20](images/20.png)

## 11. Step 3: Create a service

1. docker service create --name myservice \ --network overnet \ --replicas 2 \ ubuntu sleep infinity
2. docker network ls
3. docker service ps myservice
4. docker network inspect overnet

![20](images/21.png)

![20](images/22.png)

![20](images/23.png)

## 12. Step 4: Test the network

1. docker network inspect overnet
2. docker ps
3. docker exec -it yourcontainerid /bin/bash
4. apt-get update && apt-get install -y iputils-ping
5. ping -c5 172.17.0.2

![20](images/24.png)

## 13. Cleaning Up

1. docker service rm myservice
2. docker ps
3. docker kill c7bf9743d493 b46bb4cec484

![20](images/25.png)

![20](images/26.png)
