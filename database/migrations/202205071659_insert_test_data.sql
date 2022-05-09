TRUNCATE party;
TRUNCATE "user";

INSERT INTO "user" (username, password) VALUES ('user', '$2a$10$pg0jdTWDU1mXrEBKDvIgLeVTf/u4BL9PdVc6kP8IP4QDU.goTUoZ2');

INSERT INTO party (name, city, address, date_time)  VALUES ('Easter celebrations', 'All', 'World', '2022-04-17 24:00:00.000');
INSERT INTO party (name, city, address, date_time)  VALUES ('Charity concert party', 'Lviv', 'Rynok square', '2022-05-31 19:00:00.000');
INSERT INTO party (name, city, address, date_time)  VALUES ('Independence day party', 'Kyiv', 'Central park memorial', '2022-08-24 09:00:00.000');
INSERT INTO party (name, city, address, date_time)  VALUES ('New year eve party', 'Kharkiv', 'Svobody square', '2022-12-31 22:00:00.000');