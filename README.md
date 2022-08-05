This project is designed to showcase an event-driven/clean architecture implementation of a simple browser-based chat application using Go.

This application should allow several users to talk in a chatroom and also to get stock quotes from an API using `/stock=stock_code`.

Stock API: https://stooq.com/q/l/?s=aapl.us&f=sd2t2ohlcv&h&e=csv

## Run instructions

1. Install docker compose: https://docs.docker.com/compose/install/
2. From the root directory, run: `docker-compose up`
3. Open your browser and navigate to http://localhost:8080/

> Note: environment variables are set in `/docker/app/.env` and `docker-compose.yml`
