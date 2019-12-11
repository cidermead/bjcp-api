'use strict';

const fs = require('fs');
const beerQuestions = require('./beer-questions-db.json');
const ciderQuestions = require('./cider-questions-db.json');
const meadQuestions = require('./mead-questions-db.json');

const sntz = str => str.replace('\'', '\'\'');

const FILE_NAME = './questions/questions-db.sql';
/*
const createTableSQL = `
DROP TABLE IF EXISTS questions;
DROP TYPE IF EXISTS STYLE_TYPE;

CREATE TYPE STYLE_TYPE AS ENUM ('beer', 'mead', 'cider');
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

`; */


let questionsSQL = [...beerQuestions, ...ciderQuestions, ...meadQuestions].map( ({question, options, answer, topic, exam = 'beer', active = true, deleted = false}) =>
  `   ('${sntz(question)}', '{${options.map(val => `"${sntz(val)}"`).join(', ')}}', ${answer}, '${topic}', '${exam}', ${active}, ${deleted})`).join(',\n');

questionsSQL = `\nINSERT INTO questions (question, options, answer, topic, exam, active, deleted) VALUES \n${questionsSQL};`;

fs.writeFileSync(FILE_NAME, questionsSQL);


console.log(`${FILE_NAME} created successfully`);
