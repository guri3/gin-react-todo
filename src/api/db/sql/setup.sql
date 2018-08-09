DROP TABLE IF EXISTS todo;
CREATE TABLE IF NOT EXISTS todo (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  done BOOLEAN
);

INSERT INTO todo VALUES
  (1, 'todo1', false),
  (2, 'todo2', false),
  (3, 'todo3', false);
