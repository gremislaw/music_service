
CREATE TABLE playlist(
    id SERIAL PRIMARY KEY,
    duration INT,
    songname VARCHAR(30),
    author VARCHAR(50)
);

INSERT INTO playlist(duration, songname, author) VALUES 
(35, 'Lazy song', 'Bruno Mars'),
(45, 'Complicated', 'Avril Lavigne'),
(55, 'Kigeki', 'Gen Hoshino'),
(65, 'Drunk', 'Keshi'),
(75, 'Stay', 'The Kid LAROI and Justin Bieber');