package query

const CreatePostgresql = `
CREATE TABLE IF NOT EXISTS users (
                                     user_id uuid PRIMARY KEY,
                                     name varchar NOT NULL UNIQUE,
                                     password varchar NOT NULL,
                                     created_at date NOT NULL,
                                     deleted_at date
);

CREATE TABLE IF NOT EXISTS links (
                                     link_id uuid PRIMARY KEY ,
                                     title varchar NOT NULL,
                                     url varchar NOT NULL,
                                     api_point varchar NOT NULL,
                                     created_at date NOT NULL,
                                     deleted_at date,
                                     user_id uuid,
                                     FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS count (
                                     count_id uuid PRIMARY KEY,
                                     date date NOT NULL,
                                     link_id uuid NOT NULL,
                                     FOREIGN KEY (link_id) REFERENCES links(link_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS params (
                                    salt varchar
);
`
