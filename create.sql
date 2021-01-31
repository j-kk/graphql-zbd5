DROP TABLE IF EXISTS GeoPositions, Interests, Users, UsersInterests;
CREATE TABLE GeoPositions
(
    id     SERIAL PRIMARY KEY UNIQUE NOT NULL,
    width  float                     NOT NULL,
    height float                     NOT NULL

);

CREATE TABLE Interests
(
    id   SERIAL PRIMARY KEY UNIQUE NOT NULL,
    name VARCHAR                   NOT NULL UNIQUE
);

INSERT INTO Interests(name)
VALUES ('cyberpunk'),
       ('C++');

CREATE TABLE Users
(
    id         SERIAL PRIMARY KEY UNIQUE NOT NULL,
    gender     VARCHAR(1),
    birth_year INTEGER CHECK ( birth_year > 1900 ),
    income     INTEGER CHECK ( income >= 0 ),
    geopos_id  INTEGER                   REFERENCES GeoPositions (id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE UsersInterests
(
    id          SERIAL PRIMARY KEY UNIQUE         NOT NULL,
    interest_id INTEGER REFERENCES Interests (id) NOT NULL,
    user_id     INTEGER REFERENCES Users (id)     NOT NULL
);

ALTER TABLE GeoPositions
    ADD user_id INTEGER REFERENCES Users (id) ON UPDATE CASCADE ON DELETE CASCADE;



