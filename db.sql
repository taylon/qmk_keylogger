CREATE TABLE keyactions (
  timedate INTEGER,
  keyboard TEXT,
  column INTEGER,
  row INTEGER,
  press INTEGER,
  tapCount INTEGER,
  tapInterrupted INTEGER,
  keycode INTEGER,
  layer INTEGER
);

CREATE INDEX keyactions_index ON keyactions(timedate);

CREATE TABLE keyaction_errors (
  timedate INTEGER,
  input TEXT
);
