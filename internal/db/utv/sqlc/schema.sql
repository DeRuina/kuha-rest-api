CREATE TABLE user_data(
  user_id UUID PRIMARY KEY,
  data jsonb NOT NULL
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE notifications(
  id UUID PRIMARY KEY,
  to_id UUID NOT NULL,
  from_id UUID NOT NULL,
  status INT NOT NULL,
  expires INT NOT NULL,
  notification JSONB NOT NULL
);


CREATE TABLE utv_groups(
  id UUID PRIMARY KEY,
  group_name TEXT NOT NULL,
  created INT NOT NULL,
  active INT NOT NULL,
  deleted INT NOT NULL
);


CREATE TABLE utv_group_members(
  group_id UUID REFERENCES utv_groups(id) ON DELETE CASCADE,
  user_id UUID NOT NULL,
  added INT NOT NULL,
  UNIQUE(group_id, user_id)
);

CREATE TABLE oura_data(
  user_id UUID NOT NULL,
  summary_date DATE NOT NULL,
  PRIMARY KEY (user_id, summary_date),
  data JSONB NOT NULL
);

CREATE TABLE oura_tokens(
   user_id UUID PRIMARY KEY,
   data jsonb NOT NULL
);

CREATE TABLE coachtech_data(
  coachtech_id INT NOT NULL,
  summary_date DATE NOT NULL,
  test_id text NOT NULL,
  primary key(coachtech_id, summary_date, test_id),
  data JSONB NOT NULL
);

CREATE TABLE coachtech_ids(
  user_id UUID PRIMARY KEY,
  coachtech_id INT NOT NULL
);

CREATE TABLE polar_data(
  user_id UUID NOT NULL,
  summary_date DATE NOT NULL,
  PRIMARY KEY (user_id, summary_date),
  data JSONB NOT NULL
);

CREATE TABLE polar_tokens(
  user_id UUID PRIMARY KEY,
  data JSONB NOT NULL
);

CREATE TABLE suunto_tokens(
  user_id UUID PRIMARY KEY,
  data JSONB NOT NULL
);

CREATE TABLE suunto_data(
  user_id UUID NOT NULL,
  summary_date DATE NOT NULL,
  PRIMARY KEY (user_id, summary_date),
  data JSONB NOT NULL
);

CREATE TABLE source_cache(
  source TEXT PRIMARY KEY,
  data JSONB NOT NULL
);

CREATE TABLE resource_data(
  resource_id TEXT PRIMARY KEY,
  data JSONB NOT NULL
);

CREATE TABLE app_data(
  app_id TEXT NOT NULL,
  field_name TEXT NOT NULL,
  PRIMARY KEY (app_id, field_name),
  data JSONB NOT NULL
);