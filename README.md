# Poke Battle

Simple Go project to try REST API, gRPC and RabbitMQ. Used SQLC with Postgres as database.

## Endpoints

POST /api/pokebattle -> Creates a pokemon battle and returns the id and the two pokemons that are fighting.
POST /api/pokebattle/long -> Sends 8 pokemons to fight to a RabbitMQ queue.
GET /api/pokebattle/{id} -> Returns the winner of a fight
GET /api/pokebattle -> Returns a list of winners
GET /api/pokebattle/stats -> Connects through gRPC and gets stats (wins, loses and winrate) for a given pokemon. Pokemon must be send through query parameter called 'pokemon'.
