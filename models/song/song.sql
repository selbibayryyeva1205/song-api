CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_name TEXT NOT NULL,
    song_name TEXT NOT NULL,
    release_date time ,
    song_text TEXT NOT NULL,
    link TEXT NOT NULL
);
