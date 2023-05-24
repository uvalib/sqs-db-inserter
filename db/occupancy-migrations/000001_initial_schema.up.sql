-- camera data
CREATE TABLE cameras(
   id             serial PRIMARY KEY,

   serial_no      VARCHAR( 32 ) UNIQUE NOT NULL,
   location_short VARCHAR( 32 ) UNIQUE NOT NULL,
   location_long  VARCHAR( 255 ) UNIQUE NOT NULL,

   -- create information
   created_at     timestamp DEFAULT NOW()
);

-- enum type definition
CREATE TYPE data_source AS ENUM ('occupancy', 'sum');

-- raw metrics data
CREATE TABLE rawmetrics(
   id              serial PRIMARY KEY,

   -- camera telemetry
   source          data_source NOT NULL,
   serial_no       VARCHAR( 32 ) NOT NULL,
   occupancy       integer NOT NULL,
   count_in        integer NOT NULL,
   count_out       integer NOT NULL,
   source_unixtime VARCHAR( 32 ) NOT NULL,

   -- create information
   created_at      timestamp DEFAULT NOW()
);

-- raw metrics indexes
CREATE INDEX serial_idx ON rawmetrics(serial_no);
CREATE INDEX source_idx ON rawmetrics(source);

-- occupancy view
CREATE VIEW rawmetrics_occupancy AS SELECT
   id,
   serial_no,
   occupancy,
   count_in as total_in,
   count_out as total_out,
   to_date(to_char(to_timestamp(CAST(source_unixtime as bigint)), 'YYYYMMDD'), 'YYYYMMDD') as date,
   to_char(to_timestamp(CAST(source_unixtime as bigint)), 'HH24:MI:SS') as time,
   to_char(to_timestamp(CAST(source_unixtime as bigint)), 'Day') as day,
   to_char(to_timestamp(CAST(source_unixtime as bigint)), 'D') as dow
FROM rawmetrics WHERE source = 'occupancy';

-- sum view
CREATE VIEW rawmetrics_sum AS SELECT
   id,
   serial_no,
   count_in,
   count_out,
   to_date(to_char(to_timestamp(CAST(source_unixtime as bigint)), 'YYYYMMDD'), 'YYYYMMDD') as date,
   to_char(to_timestamp(CAST(source_unixtime as bigint)), 'HH24:MI:SS') as time,
   to_char(to_timestamp(CAST(source_unixtime as bigint)), 'Day') as day,
   to_char(to_timestamp(CAST(source_unixtime as bigint)), 'D') as dow
FROM rawmetrics WHERE source = 'sum';
