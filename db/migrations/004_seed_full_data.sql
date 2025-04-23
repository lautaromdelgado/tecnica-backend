-- Nuevos usuarios
INSERT INTO users (username, email, role)
VALUES
('sofia_dev', 'sofia@dev.com', 'user'),
('lucas_admin', 'lucas@admin.com', 'admin'),
('ana_marketing', 'ana@marketing.com', 'user'),
('juan_support', 'juan@support.com', 'user'),
('marcos_pm', 'marcos@pm.com', 'admin');

-- Nuevos eventos (algunos públicos, otros no)
INSERT INTO events (organizer, title, long_description, short_description, date, location, is_published)
VALUES
('lucas_admin', 'Charlas de Go y Backend', 'Evento técnico sobre backend moderno', 'Backend Talks', 1768000000, 'UTN Córdoba', TRUE),
('lucas_admin', 'Curso intensivo de React', 'React desde cero a avanzado', 'React Avanzado', 1768800000, 'UBA Fadu', TRUE),
('lucas_admin', 'Taller de Inteligencia Artificial', 'Workshop con herramientas de IA', 'IA Workshop', 1769000000, 'UNR Rosario', FALSE),
('marcos_pm', 'Meetup de Product Managers', 'Encuentro de profesionales de gestión de producto', 'PM Meetup', 1768600000, 'Espacio Coworking BA', TRUE);

-- Relaciones user_event
INSERT INTO user_event (user_id, event_id)
VALUES
(1, 3),
(2, 4),
(3, 4),
(4, 2),
(5, 1),
(5, 4);

-- Logs simulados
INSERT INTO event_logs (title, organizer, action)
VALUES
('Charlas de Go y Backend', 'lucas_admin', 'create'),
('Curso intensivo de React', 'lucas_admin', 'create'),
('Taller de Inteligencia Artificial', 'lucas_admin', 'create'),
('Charlas de Go y Backend', 'lucas_admin', 'publish'),
('Taller de Inteligencia Artificial', 'lucas_admin', 'unpublish'),
('PM Meetup', 'marcos_pm', 'create'),
('PM Meetup', 'marcos_pm', 'publish');
