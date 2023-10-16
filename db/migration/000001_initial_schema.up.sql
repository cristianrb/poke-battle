CREATE TABLE pokemon_battle (
    id uuid DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    pokemon1 VARCHAR NOT NULL,
    pokemon2 VARCHAR NOT NULL,
    winner VARCHAR DEFAULT NULL
);
