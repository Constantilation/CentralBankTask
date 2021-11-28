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