## _215611104 - Elsa Setiyawati_

# LATIHAN

# Application Containerization And Microservice Orchestration

## Stage Setup

### 1. git clone https://github.com/ibnesayeed/linkextractor.git

### 2. cd linkextractor

### 3. git checkout demo

![20](images/1.png)

## Step 0: Basic Link Extractor Script

### 1. git checkout step0

### 2. tree

### 3. cat linkextractor.py

### 4. ./linkextractor.py http://example.com/

### 5. ls -l linkextractor.py

### 6. python3 linkextractor.py

![20](images/2.png)

![20](images/3.png)

![20](images/4.png)

## Step 1: Containerized Link Extractor Script

### 1. git checkout step1

### 2. tree

### 3. cat Dockerfile

### 4. docker image build -t linkextractor:step1 .

### 5. docker image ls

### 6. docker container run -it --rm linkextractor:step1 http://example.com/

### 7. docker container run -it --rm linkextractor:step1 https://training.play-with-docker.com/

![20](images/5.png)

![20](images/6.png)

![20](images/7.png)

![20](images/8.png)

![20](images/9.png)

![20](images/10.png)

## Step 2: Link Extractor Module with Full URI and Anchor Text

### 1. git checkout step2

### 2. tree

### 3. cat linkextractor.py

### 4. docker image build -t linkextractor:step2 .

### 5. docker image ls

### 6. docker container run -it --rm linkextractor:step2 https://training.play-with-docker.com/

![20](images/11.png)

![20](images/12.png)

![20](images/13.png)

![20](images/14.png)

![20](images/15.png)

![20](images/16.png)

## Step 3: Link Extractor API Service

### 1. git checkout step3

### 2. tree

### 3. cat Dockerfile

### 4. cat main.py

### 5. docker image build -t linkextractor:step3 .

### 6. docker container run -d -p 5000:5000 --name=linkextractor linkextractor:step3

### 7. docker container ls

### 8. curl -i http://localhost:5000/api/http://example.com/

### 9. docker container logs linkextractor

### 10. docker container rm -f linkextractor

![20](images/17.png)

![20](images/18.png)

![20](images/19.png)

![20](images/20.png)

![20](images/21.png)

![20](images/22.png)

![20](images/23.png)

![20](images/24.png)

![20](images/25.png)

## Step 4: Link Extractor API and Web Front End Services

### 1. git checkout step4

### 2. tree

### 3. cat docker-compose.yml

### 4. cat www/index.php

### 5. docker-compose up -d --build

### 6. docker container ls

### 7. curl -i http://localhost:5000/api/http://example.com/

### 8. sed -i 's/Link Extractor/Super Link Extractor/g' www/index.php

### 9. git reset --hard

### 10. docker-compose down

![20](images/26.png)

![20](images/27.png)

![20](images/28.png)

![20](images/29.png)

![20](images/30.png)

![20](images/31.png)

![20](images/32.png)

![20](images/33.png)

## Step 5: Redis Service for Caching

### 1. git checkout step5

### 2. tree

### 3. cat www/Dockerfile

### 4. cat api/main.py

### 5. cat docker-compose.yml

### 6. docker-compose up -d --build

### 7. docker-compose exec redis redis-cli monitor

### 8. sed -i 's/Link Extractor/Super Link Extractor/g' www/index.php

### 9. git reset --hard

### 10. docker-compose down

![20](images/34.png)

![20](images/35.png)

![20](images/36.png)

![20](images/37.png)

![20](images/38.png)

![20](images/39.png)

![20](images/40.png)

## Step 6: Swap Python API Service with Ruby

### 1. git checkout step6

### 2. tree

### 3. cat api/linkextractor.rb

### 4. cat api/Dockerfile

### 5. cat docker-compose.yml

### 6. docker-compose up -d --build

### 7. curl -i http://localhost:4567/api/http://example.com/

### 8. docker-compose down

### 9. cat logs/extraction.log

![20](images/41.png)

![20](images/42.png)

![20](images/43.png)

![20](images/44.png)

![20](images/45.png)

![20](images/46.png)

![20](images/47.png)
