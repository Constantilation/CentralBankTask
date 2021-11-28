INSERT INTO bankfilleddates
SELECT date_trunc('day', dd):: date
FROM generate_series
         ( '1992-07-01'::timestamp
         , '2021-11-27'::timestamp
         , '1 day'::interval) dd
;