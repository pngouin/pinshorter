package query

const UserCreate = `
INSERT INTO users (user_id, name, password, created_at) VALUES ($1, lower($2), $3, now())
`

const UserExist = `
SELECT count(*) FROM users WHERE user_id=$1 and name=lower($2) AND deleted_at is null
`

const UserById = `
SELECT user_id, name FROM user WHERE user_id=$1 AND deleted_at IS NULL
`

const UserConnection = `
SELECT user_id, count(*) from users where name=lower($1) and password=$2 and deleted_at is null group by user_id
`
