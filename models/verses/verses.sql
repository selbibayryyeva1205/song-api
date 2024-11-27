CREATE TABLE verses (
    id SERIAL PRIMARY KEY,
    song_id int,
    verse_number int,
    text TEXT NOT NULL
);
