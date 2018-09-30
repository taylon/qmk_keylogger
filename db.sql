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

CREATE INDEX keyactions_time_index ON keyactions(timedate);
CREATE INDEX keyactions_keyboard_index ON keyactions(keyboard);
