CREATE TABLE sdk_names (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE sdk_repos (
    id TEXT PRIMARY KEY,
    github TEXT NOT NULL
);

CREATE TABLE sdk_languages (
    id TEXT,
    language TEXT NOT NULL,
    PRIMARY KEY (id, language)
);

CREATE TABLE sdk_types (
    id TEXT PRIMARY KEY,
    type TEXT NOT NULL,
    FOREIGN KEY (type) REFERENCES sdk_type_info(type)
);


CREATE TABLE sdk_type_info (
    type TEXT PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE sdk_features (
    id TEXT,
    feature TEXT NOT NULL,
    introduced DATE NOT NULL,
    deprecated DATE,
    removed DATE,
    PRIMARY KEY (id, feature),
    FOREIGN KEY (feature) REFERENCES sdk_feature_info(feature)
);

CREATE TABLE sdk_feature_info (
    feature TEXT PRIMARY KEY,
    description TEXT NOT NULL
);


INSERT INTO sdk_type_info (type, description) VALUES
                                              ('client-side', 'Primarily used for user-facing application.'),
                                              ('server-side', 'Primarily used for server-side applications.');

INSERT INTO sdk_feature_info (feature, description) VALUES
    ('u2c', 'The concept of Users is replaced with Contexts, which can be used to represent users, devices, or other entities.');
