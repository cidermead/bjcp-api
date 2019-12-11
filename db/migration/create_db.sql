DROP DATABASE bjcp_table_dev;
DROP USER dev_judge;

SET TIME ZONE 'America/Los_Angeles';

CREATE DATABASE bjcp_table_dev;
CREATE USER dev_judge WITH ENCRYPTED PASSWORD 'beerJudging';
GRANT ALL PRIVILEGES ON DATABASE bjcp_table_dev TO dev_judge;
ALTER DATABASE bjcp_table_dev OWNER TO dev_judge;
