# edison_test_task
Test task for [Edison](https://edsd.ru)

# How to run
1. docker network create app_main
2. docker compose up --build(Windows: docker-compose up --build)

# Details
Frontend is made on ReactJS, backend - Golang(net/http + gorilla/(mux + sessions)).
Code isn't optimized well, but still runs fast.
Session size isn't limited.
