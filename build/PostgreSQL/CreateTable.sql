CREATE TABLE IF NOT EXISTS BankFilledDates
(
    fulldate DATE,
    isFilled BOOLEAN default false
);

CREATE TABLE IF NOT EXISTS BankInfo
(
    fulldate    DATE,
    valutename  TEXT,
    valuteValue FLOAT
);

CREATE INDEX bankinfo_3c_idx ON bankinfo (fulldate, valutename, valuteValue DESC);