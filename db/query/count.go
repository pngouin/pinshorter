package query

const CountCreate = `
INSERT INTO count (count_id, date, link_id) VALUES ($1, now(), $2)
`

const CountGet = `
SELECT count(*) FROM count where link_id=$1`
