## go 实战教程

### sql

```sql
-- user表
CREATE TABLE users
(
    id INT
    UNSIGNED AUTO_INCREMENT,
   login_name VARCHAR
    (64) UNIQUE KEY,
   pwd TEXT,
   PRIMARY KEY
    ( id )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


    -- video表
    CREATE TABLE video_info
    (
        id VARCHAR(64) NOT NULL,
        author_id INT
        UNSIGNED,
   name TEXT,
   display_ctime  TEXT,
   create_ctime  DATETIME,
   PRIMARY KEY
        (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

        --评论
        CREATE TABLE comments
        (
            id VARCHAR(64) NOT NULL,
            video_id VARCHAR(64),
            author_id INT
            UNSIGNED,
   content  TEXT,
   ctime  DATETIME,
   PRIMARY KEY
            (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

            --sessions
            CREATE TABLE sessions
            (
                session_id VARCHAR(255) NOT NULL,
                TTL TINYTEXT,
                login_name VARCHAR(64),
                PRIMARY KEY(session_id)
            )
            ENGINE=InnoDB DEFAULT CHARSET=utf8;


```