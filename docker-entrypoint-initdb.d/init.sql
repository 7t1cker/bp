CREATE DATABASE vk_v2;
\c vk_v2;

CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR
);

CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    skill VARCHAR
);

CREATE TABLE IF NOT EXISTS divisions (
    id SERIAL PRIMARY KEY,
    division_name VARCHAR
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    division_id INTEGER REFERENCES divisions(id),
    group_id INTEGER REFERENCES groups(id),
    skill_tasks JSONB,
    balance NUMERIC DEFAULT 0.0,
    status VARCHAR CHECK (status IN ('работает', 'перерыв', 'не работает')),
    access_token VARCHAR,
    login VARCHAR UNIQUE,
    password VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    middle_name VARCHAR,
    role VARCHAR
);

CREATE TABLE IF NOT EXISTS quests_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS learn (
    id SERIAL PRIMARY KEY,
    skill_id INTEGER REFERENCES skills(id),
    learn_title VARCHAR
);

CREATE TABLE IF NOT EXISTS quests (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    description TEXT,
    deadline DATE,
    creator_id INTEGER REFERENCES users(id),
    cost NUMERIC,
    priority INTEGER,
    skills_required JSONB,
    recurrence_limit INTEGER
);
CREATE TABLE IF NOT EXISTS assignee_quests (
    id SERIAL PRIMARY KEY,
    assignee_id INTEGER REFERENCES users(id),
    quest_id INTEGER REFERENCES quests(id),
    recurrence INTEGER,
    creation_timestamp TIMESTAMP,
    closing_timestamp TIMESTAMP,
    done BOOLEAN DEFAULT FALSE,
    quest_type_id INTEGER
);

CREATE TABLE IF NOT EXISTS hot_tasks (
    id SERIAL PRIMARY KEY,
    assignee_quest_id INTEGER REFERENCES assignee_quests(id),
    creation_time TIMESTAMP,
    hot NUMERIC,
    end_time TIMESTAMP,
    fire BOOLEAN
);
INSERT INTO quests_types (name) VALUES ('common');
INSERT INTO quests_types (name) VALUES ('hot');
INSERT INTO quests_types (name) VALUES ('fire');
INSERT INTO quests_types (name) VALUES ('learn');
INSERT INTO quests_types (name) VALUES ('multi-plex');



