
CREATE TABLE songs(
    id SERIAL PRIMARY KEY,
    duration INT,
    songname VARCHAR(30),
    author VARCHAR(50)
);

INSERT INTO songs(duration, songname, author) VALUES 
(35, 'Lazy song', 'Bruno Mars'),
(45, 'Complicated', 'Avril Lavigne'),
(55, 'Kigeki', 'Gen Hoshino'),
(65, 'Drunk', 'Keshi'),
(75, 'Stay', 'The Kid LAROI and Justin Bieber');

CREATE TABLE playlists(
    id SERIAL PRIMARY KEY,
    playlistname VARCHAR(30),
);

INSERT INTO playlists(playlistname) VALUES 
('Pop'),
('Jazz'),
('Meloman');

CREATE TABLE songs_playlists(
    id_song INT
    id_playlist INT
    INDEX `FK_films_actors_actors` (`id_song`),
    INDEX `FK_films_actors_films` (`id_playlist`),
    CONSTRAINT `FK_films_actors_actors` FOREIGN KEY (`id_song`) REFERENCES `songs` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `FK_films_actors_films` FOREIGN KEY (`id_playlist`) REFERENCES `playlists` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
)

INSERT INTO songs_playlists(id_song, id_playlist) VALUES 
(1, 1),
(3, 1),
(2, 2),
(4, 2),
(1, 3),
(2, 3);