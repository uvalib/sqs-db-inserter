-- drop the views
DROP VIEW IF EXISTS rawmetrics_sum;
DROP VIEW IF EXISTS rawmetrics_occupancy;

-- occupancy view
CREATE VIEW rawmetrics_occupancy AS SELECT
   id,
   serial_no,
   occupancy,
   count_in as total_in,
   count_out as total_out,
   to_date(to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'YYYYMMDD'), 'YYYYMMDD') as date,
   to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'HH24:MI:SS') as time,
   to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'Day') as day,
   to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'D') as dow
FROM rawmetrics WHERE source = 'occupancy';

-- sum view
CREATE VIEW rawmetrics_sum AS SELECT
   id,
   serial_no,
   count_in,
   count_out,
   to_date(to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'YYYYMMDD'), 'YYYYMMDD') as date,
   to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'HH24:MI:SS') as time,
   to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'Day') as day,
   to_char(to_timestamp(CAST(source_unixtime as bigint)) AT TIME ZONE 'EST5EDT', 'D') as dow
FROM rawmetrics WHERE source = 'sum';

