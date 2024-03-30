-- для создания подходящей бд следует использовать этот sql скрипт
CREATE TABLE user_groups (
                             user_id INTEGER NOT NULL,
                             username TEXT(32) NOT NULL,
                             group_id INTEGER NOT NULL,
                             mailing INTEGER,
                             CONSTRAINT user_groups_pk PRIMARY KEY (user_id)
);