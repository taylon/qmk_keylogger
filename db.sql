CREATE TABLE key_actions (
  created_at INTEGER,
  keyboard_name TEXT,
  col INTEGER,
  row INTEGER,
  press INTEGER,
  tap_count INTEGER,
  tap_interrupted INTEGER,
  key_code INTEGER,
  layer INTEGER
);

CREATE INDEX key_actions_time_index ON key_actions(created_at);
CREATE INDEX key_actions_keyboard_index ON key_actions(keyboard_name);
