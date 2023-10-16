--
--
--

-- camera data
CREATE TABLE cameras(
   id             serial PRIMARY KEY,

   serial_no      VARCHAR( 32 ) UNIQUE NOT NULL,
   location_short VARCHAR( 32 ) NOT NULL,
   location_long  VARCHAR( 255 ) NOT NULL,

   -- create information
   created_at     timestamp DEFAULT NOW()
);

-- raw metrics data
CREATE TABLE rawmetrics(
   id              serial PRIMARY KEY,

   -- camera telemetry
   source          ENUM('occupancy', 'sum'),
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
    date_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%Y-%m-%d') as date,
    time_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%T') as time,
    date_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%W') as day,
    date_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%w') as dow
FROM rawmetrics WHERE source = 'occupancy';

-- sum view
CREATE VIEW rawmetrics_sum AS SELECT
    id,
    serial_no,
    count_in,
    count_out,
    date_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%Y-%m-%d') as date,
    time_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%T') as time,
    date_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%W') as day,
    date_format(convert_tz(from_unixtime(source_unixtime), 'UTC', 'EST5EDT'), '%w') as dow
FROM rawmetrics WHERE source = 'sum';

--
-- end of file
--
