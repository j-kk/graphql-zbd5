DROP TABLE IF EXISTS GeoPositions, Interests, Users, UsersInterests, ads, adwords, views;
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
    id          SERIAL PRIMARY KEY UNIQUE                           NOT NULL,
    interest_id INTEGER REFERENCES Interests (id) ON UPDATE CASCADE NOT NULL,
    user_id     INTEGER REFERENCES Users (id) ON UPDATE CASCADE     NOT NULL
);

ALTER TABLE GeoPositions
    ADD user_id INTEGER REFERENCES Users (id) ON UPDATE CASCADE ON DELETE CASCADE;


CREATE TABLE ads
(
    id         SERIAL PRIMARY KEY UNIQUE NOT NULL,
    width      INTEGER CHECK ( width > 0 ),
    height     INTEGER CHECK ( height > 0 ),
    main_color VARCHAR
);

CREATE TABLE adwords
(
    id    SERIAL PRIMARY KEY UNIQUE NOT NULL,
    word  VARCHAR                   NOT NULL,
    ad_id INTEGER REFERENCES ads (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE views
(
    id      SERIAL PRIMARY KEY UNIQUE                       NOT NULL,
    ads_id  INTEGER REFERENCES ads (id) ON UPDATE CASCADE   NOT NULL,
    user_id INTEGER REFERENCES Users (id) ON UPDATE CASCADE NOT NULL,
    t       TIMESTAMP DEFAULT now()                         NOT NULL
)