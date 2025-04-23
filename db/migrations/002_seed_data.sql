-- Insert sample users
INSERT INTO users (id, username, email, role, created_at, updated_at)
VALUES 
(1, 'admin_user', 'admin@example.com', 'admin', NOW(), NOW()),
(2, 'johndoe', 'john@example.com', 'user', NOW(), NOW()),
(3, 'janedoe', 'jane@example.com', 'user', NOW(), NOW());

-- Insert sample events
INSERT INTO events (id, organizer, title, long_description, short_description, date, location, is_published, created_at, updated_at)
VALUES
(1, 'admin_user', 'Tech Conference 2025', 'A big event about new tech trends.', 'Tech 2025', 1767264000, 'Buenos Aires', TRUE, NOW(), NOW()),
(2, 'admin_user', 'Go Workshop', 'Learn Go from scratch in one weekend.', 'GoLang 101', 1767350400, 'Córdoba', FALSE, NOW(), NOW());

-- Insert sample user_event relations
INSERT INTO user_event (user_id, event_id, joined_at)
VALUES
(2, 1, NOW()), -- John Doe se inscribe al evento 1
(3, 1, NOW()), -- Jane Doe también
(2, 2, NOW()); -- John Doe se inscribe al evento 2
