# Spaceports Leaderboard API

As part of my assignment in CS205 Operating Systems with Android, I have created an API server for the game (developed in Java) to call to get the updates for the leaderboard.

### Set up

1. Copy `.env.example` and rename to be `.env`

   ```bash
   cp .env.example .env
   ```

1. Run Docker compose

   ```bash
   docker compose up -d --build
   ```

### Other Commands

docker exec -it your_postgres_container_name psql -U your_username -d your_database_name
