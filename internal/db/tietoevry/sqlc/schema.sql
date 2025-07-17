-- Main entities:
-- USERS
-- EXERCISES
-- SYMPTOMS
-- MEASUREMENTS
-- TEST_RESULTS
-- QUESTIONNAIRES
-- ACTIVITY_ZONES

-- <USERS>
-- in 360° Training, we have user_profiles as a separate table,
-- for KUHA3, it is simpler to have its data in the same table.
-- The need for the unique constraint in KUHA3 perspective is debatable.

CREATE EXTENSION IF NOT EXISTS citext;


create table users
(
    id                          uuid    not null primary key,
    sportti_id                  integer not null unique,
    profile_gender              citext,
    profile_birthdate           date,
    profile_weight              real,
    profile_height              real,
    profile_resting_heart_rate  integer,
    profile_maximum_heart_rate  integer,
    profile_aerobic_threshold   integer,
    profile_anaerobic_threshold integer,
    profile_vo2max              integer
);
-- </USERS>


-- <EXERCISES>
-- in 360° Training, the sport_type column references the sport_types table
-- from KUHA3 perspective, having it just as text without reference is more flexible
-- The need for the unique constraint in KUHA3 perspective is debatable.
create table exercises
(
    id                  uuid                     not null
        primary key,
    created_at          timestamp with time zone not null,
    updated_at          timestamp with time zone not null,
    user_id             uuid                     not null
        references users
            on delete cascade,
    start_time          timestamp with time zone not null,
    duration            interval                 not null,
    comment             text,
    sport_type          text,
    detailed_sport_type text,
    distance            real,
    avg_heart_rate      real,
    max_heart_rate      real,
    trimp               real,
    sprint_count        integer,
    avg_speed           real,
    max_speed           real,
    source              text                     not null,
    status              text,
    calories            integer,
    training_load       integer,
    raw_id              text,
    raw_data            jsonb,
    feeling             integer,
    recovery            integer,
    rpe                 integer,
    constraint exercises_source_raw_id_unique
        unique (source, raw_id)
);

create table exercise_hr_zones
(
    exercise_id     uuid                     not null
        references exercises
            on delete cascade,
    zone_index      integer                  not null,
    seconds_in_zone integer                  not null,
    lower_limit     integer                  not null,
    upper_limit     integer                  not null,
    created_at      timestamp with time zone not null,
    updated_at      timestamp with time zone not null,
    primary key (exercise_id, zone_index)
);

-- In 360° Training, the sample_type column references the exercise_sample_types table,
-- from KUHA3 perspective, having it just as text without reference is more flexible
create table exercise_samples
(
    id             uuid    not null,
    user_id        uuid    not null
        references users
            on delete cascade,
    exercise_id    uuid    not null
        references exercises
            on delete cascade,
    sample_type    text    not null,
    recording_rate integer not null,
    samples        double precision[],
    source         citext  not null,
    primary key (exercise_id, sample_type)
);

-- In 360° Training, the section_type column references the exercise_section_types table and source references the data_sources table,
-- from KUHA3 perspective, having them just as citext without reference is more flexible
create table exercise_sections
(
    id           uuid                     not null
        primary key,
    user_id      uuid                     not null
        references users
            on delete cascade,
    exercise_id  uuid                     not null
        references exercises
            on delete cascade,
    created_at   timestamp with time zone not null,
    updated_at   timestamp with time zone not null,
    start_time   timestamp with time zone not null,
    end_time     timestamp with time zone not null,
    section_type citext,
    name         text,
    comment      text,
    source       citext                   not null,
    raw_id       text,
    raw_data     jsonb
);
-- </EXERCISES>

-- <SYMPTOMS>
-- In 360° Training, the source column references the data_sources table,
-- from KUHA3 perspective, having it just as citext without reference is more flexible.
-- The need for the unique constraint in KUHA3 perspective is debatable.
create table symptoms
(
    id              uuid                     not null
        primary key,
    user_id         uuid                     not null
        references users
            on delete cascade,
    date            date                     not null,
    symptom         text                     not null,
    severity        integer                  not null,
    comment         text,
    source          citext                   not null,
    created_at      timestamp with time zone not null,
    updated_at      timestamp with time zone not null,
    raw_id          text,
    original_id     uuid
        references symptoms,
    recovered       boolean,
    pain_index      integer,
    side            citext,
    category        text,
    additional_data jsonb,
    constraint symptoms_source_user_id_date_raw_id_unique
        unique (source, user_id, date, raw_id)
);
-- </SYMPTOMS>

-- <MEASUREMENTS>
-- In 360° Training, the name column references the measurement_types table and source references the data_sources table,
-- from KUHA3 perspective, having them just as citext without reference is more flexible.
-- The measurement_types table has some metadata relating to the type, but only thing that is relevant here is "type", so I have denormalized it to here.
-- The need for the unique constraint in KUHA3 perspective is debatable.
create table measurements
(
    id              uuid                     not null
        primary key,
    created_at      timestamp with time zone not null,
    updated_at      timestamp with time zone not null,
    user_id         uuid                     not null
        references users
            on delete cascade,
    date            date                     not null,
    name            citext                   not null,
    name_type       citext                   not null, -- denormalized
    source          citext                   not null,
    value           text                     not null,
    value_numeric   real,
    comment         text,
    raw_id          text,
    raw_data        jsonb,
    additional_info jsonb,
    constraint measurements_source_user_id_date_raw_id_unique
        unique (source, user_id, date, name, raw_id)
);
-- </MEASUREMENTS>

-- <TEST_RESULTS>
-- In 360° Training, the type_id column references the test_result_types table and test_event_id references the test_events table,
-- from KUHA3 perspective, I've denormalized the relevant bits as type_* columns here.
-- Same thing with test_event_* and test_event_template_* columns with their respective tables.
create table test_results
(
    id                              uuid                     not null
        primary key,
    user_id                         uuid                     not null
        references users
            on delete cascade,
    type_id                         uuid                     not null,
    type_type                       citext,                            -- denormalized
    type_result_type                citext                   not null, -- denormalized
    type_name                       text,                              -- denormalized
    timestamp                       timestamp with time zone not null,
    name                            text,
    comment                         text,
    data                            jsonb                    not null,
    created_at                      timestamp with time zone not null,
    updated_at                      timestamp with time zone not null,
    test_event_id                   uuid,
    test_event_name                 text,                              -- denormalized
    test_event_date                 date,                              -- denormalized
    test_event_template_test_id     uuid,
    test_event_template_test_name   text,                              -- denormalized
    test_event_template_test_limits jsonb                              -- denormalized
);
-- </TEST_RESULTS>

-- <QUESTIONNAIRES>
-- Again, some denormalization here (questionnaire_instances, questions, options), for the sake of simpler dumps.
create table question_answers
(
    user_id                   uuid                     not null
        references users
            on delete cascade,
    questionnaire_instance_id uuid                     not null,
    questionnaire_name_fi     text,                              -- denormalized
    questionnaire_name_en     text,                              -- denormalized
    questionnaire_key         text                     not null, -- denormalized
    question_id               uuid                     not null,
    question_label_fi         text,                              -- denormalized
    question_label_en         text,                              -- denormalized
    question_type             text                     not null, -- denormalized
    option_id                 uuid,
    option_value              integer,                           -- denormalized
    option_label_fi           text,                              -- denormalized
    option_label_en           text,                              -- denormalized
    free_text                 text,
    created_at                timestamp with time zone not null,
    updated_at                timestamp with time zone not null,
    value                     jsonb,
    primary key (questionnaire_instance_id, question_id, user_id)
);
-- </QUESTIONNAIRES>

-- <ACTIVITY_ZONES>
create table activity_zones
(
    user_id           uuid                     not null
        references users
            on delete cascade,
    date              date                     not null,
    created_at        timestamp with time zone not null,
    updated_at        timestamp with time zone not null,
    seconds_in_zone_0 real,
    seconds_in_zone_1 real,
    seconds_in_zone_2 real,
    seconds_in_zone_3 real,
    seconds_in_zone_4 real,
    seconds_in_zone_5 real,
    source            text                     not null,
    raw_data          jsonb,
    primary key (user_id, date, source)
);
-- </ACTIVITY_ZONES>
