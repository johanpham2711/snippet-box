-- Create a `snippets` table with a user_id foreign key.
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL,
    user_id INTEGER NOT NULL,
    CONSTRAINT fk_snippets_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);