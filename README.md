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

### API Endpoints

Base URL: /api/v1

1.  Insert score

    Allows players to submit their score along with a name.

    - Method: POST /scores
    - Payload:
      ```json
      {
        "name": "PlayerName",
        "score": 12345
      }
      ```

1.  View leaderboard

    Get the list of top scores. You can limit the number of results by specifying the limit parameter.

    - Method: GET /leaderboard
    - Query Parameters:
      - limit (optional): Specifies the number of top scores to return.

### Other Commands

docker exec -it your_postgres_container_name psql -U your_username -d your_database_name


### gcloud

```bash
gcloud auth configure-docker

docker buildx build -t lead
erboard-api --platform linux/amd64 . 

docker build -t gcr.io/PROJECT_ID/REPOSITORY_NAME/MY_GO_APP:v1 .

docker push gcr.io/PROJECT_ID/REPOSITORY_NAME/MY_GO_APP:v1
```