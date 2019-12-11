CREATE TYPE STYLE_TYPE AS ENUM ('beer', 'mead', 'cider');

CREATE TABLE categories
(
    id                 SERIAL,
    category_id        VARCHAR(2) NOT NULL,
    type               STYLE_TYPE NOT NULL,
    name               VARCHAR(256) UNIQUE NOT NULL,
    notes              TEXT NOT NULL,

    deleted_at         TIMESTAMP,
    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id),
    CONSTRAINT unique_category UNIQUE (category_id)
);
ALTER TABLE categories OWNER TO dev_judge;
GRANT ALL PRIVILEGES ON TABLE categories TO dev_judge;

CREATE TABLE styles
(
    id                 SERIAL,
    category_id        INTEGER NOT NULL,
    style_id           VARCHAR(3) NOT NULL,
    name               VARCHAR(256) UNIQUE NOT NULL,
    aroma              TEXT,
    appearance         TEXT,
    flavor             TEXT,
    mouthfeel          TEXT,
    impression         TEXT,
    comments           TEXT,
    history            TEXT,
    ingredients        TEXT,
    comparison         TEXT,
    entry_instructions TEXT,
    examples           VARCHAR(64)[] NOT NULL DEFAULT '{}',
    varieties          VARCHAR(24)[] NOT NULL DEFAULT '{}',
    tags               VARCHAR(24)[] NOT NULL DEFAULT '{}',
    stats              JSON NOT NULL DEFAULT '{}',
    beer_exam          BOOLEAN NOT NULL DEFAULT FALSE,

    deleted_at         TIMESTAMP,
    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id),
    CONSTRAINT unique_style UNIQUE (style_id),
    CONSTRAINT unique_style_and_category UNIQUE (style_id, category_id),
    CONSTRAINT styles_to_category FOREIGN KEY (category_id) REFERENCES categories (id)
);
ALTER TABLE styles OWNER TO dev_judge;
GRANT ALL PRIVILEGES ON TABLE styles TO dev_judge;

CREATE TABLE similar_styles
(
    style_id           INTEGER NOT NULL,
    similar_id         INTEGER NOT NULL,

    deleted_at         TIMESTAMP,
    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW(),

    CONSTRAINT similar_pair PRIMARY KEY (style_id, similar_id),
    CONSTRAINT base_style_constraint FOREIGN KEY (style_id) REFERENCES styles (id),
    CONSTRAINT similar_style_constraint FOREIGN KEY (similar_id) REFERENCES styles (id)
);
ALTER TABLE similar_styles OWNER TO dev_judge;
GRANT ALL PRIVILEGES ON TABLE similar_styles TO dev_judge;

CREATE TABLE questions
(
    id                 SERIAL,
    question           VARCHAR(355) NOT NULL DEFAULT 'Question',
    options            TEXT [] NOT NULL DEFAULT '{"Answer"}',
    answer             INTEGER NOT NULL DEFAULT 0,
    topic              VARCHAR(20) NOT NULL DEFAULT 'general',
    exam               STYLE_TYPE NOT NULL DEFAULT 'beer',

    stats              JSON NOT NULL DEFAULT '{ "correct": 0, "wrong": 0 }',
    active             BOOLEAN NOT NULL DEFAULT FALSE,
    deleted            BOOLEAN NOT NULL DEFAULT FALSE,

    deleted_at         TIMESTAMP,
    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id)
);
ALTER TABLE questions OWNER TO dev_judge;
GRANT ALL PRIVILEGES ON TABLE questions TO dev_judge;





