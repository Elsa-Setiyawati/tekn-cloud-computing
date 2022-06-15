## _215611104 - Elsa Setiyawati_

# LATIHAN

# Orchestration Docker Orchestration Hands-on Lab

# Section 1: What is Orchestration

# Section 2: Configure Swarm Mode

## 1. docker run -dt ubuntu sleep infinity

## 2. docker ps

![20](images/1.png)

![20](images/2.png)

## Step 2.1 - Create a Manager node

### 1. docker swarm init --advertise-addr $(hostname -i)

### 2. docker info

![20](images/3.png)

![20](images/4.png)

## Step 2.2 - Join Worker nodes to the Swarm

### 1. docker node ls

![20](images/5.png)

# Section 3: Deploy applications across multiple hosts

## Step 3.1 - Deploy the application components as Docker services

### 1. docker service create --name sleep-app ubuntu sleep infinity

### 2. docker service ls

![20](images/6.png)

![20](images/7.png)

# Section 4: Scale the application

## 1. docker service update --replicas 7 sleep-app

## 2. docker service ps sleep-app

## 3. docker service update --replicas 4 sleep-app

## 4. docker service ps sleep-app

![20](images/8.png)

![20](images/9.png)

![20](images/10.png)

![20](images/11.png)

# Section 5: Drain a node and reschedule the containers

## 1. docker node ls

## 2. docker ps

## 3. docker node ls

## 4. docker node ls

## 5. docker ps

## 6. docker service ps sleep-app

![20](images/12.png)

![20](images/13.png)

![20](images/14.png)

![20](images/15.png)

![20](images/16.png)

![20](images/17.png)

# Cleaning Up

## 1. docker service rm sleep-app

## 2. docker ps

## 3. docker swarm leave –force

## 4. docker swarm leave --force

## 5. docker swarm leave –force

![20](images/18.png)

![20](images/19.png)

![20](images/20.png)

![20](images/21.png)

![20](images/22.png)
