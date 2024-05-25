CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE pokemons(
    Id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    Name VARCHAR(255),
    Url VARCHAR(255)
)