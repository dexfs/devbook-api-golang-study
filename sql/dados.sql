-- Inserting data into the usuarios table
INSERT INTO usuarios (nome, nick, email, senha) VALUES
                                                    ('John Doe', 'john_doe', 'john@example.com', '$2a$10$8V67qwvgRkwfaXSbIqdwCuQ4Eekg6Ch3ZG9sS/KfLVkcGyl38kg0K'),
                                                    ('Jane Smith', 'jane_smith', 'jane@example.com', '$2a$10$8V67qwvgRkwfaXSbIqdwCuQ4Eekg6Ch3ZG9sS/KfLVkcGyl38kg0K'),
                                                    ('Alice Johnson', 'alice_johnson', 'alice@example.com', '$2a$10$8V67qwvgRkwfaXSbIqdwCuQ4Eekg6Ch3ZG9sS/KfLVkcGyl38kg0K');

-- Inserting data into the seguidores table to represent followers
INSERT INTO seguidores (usuario_id, seguidor_id) VALUES
                                                     (1, 2), -- John Doe follows Jane Smith
                                                     (1, 3), -- John Doe follows Alice Johnson
                                                     (2, 3); -- Jane Smith follows Alice Johnson