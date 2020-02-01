package query

const LinkAllByUser = `
SELECT l.link_id, title, url, api_point, l.created_at, u.user_id, u.name, count(c.count_id)
    FROM links l
    INNER JOIN users u
        ON l.user_id = u.user_id
    INNER JOIN count c
        ON l.user_id = u.user_id
    WHERE
        u.user_id=$1 AND u.deleted_at is null
    GROUP BY l.link_id, u.user_id;
`

const LinkById = `
SELECT link_id, title, url, api_point, created_at, links.user_id
    FROM links
    WHERE link_id=$1 AND deleted_at is null
`

const LinkDelete = `
UPDATE links
    SET deleted_at=now()
    WHERE link_id=$1 AND user_id=$2 AND deleted_at is NULL
`

const LinkCreate = `
INSERT INTO links (link_id, title, url, api_point, created_at, user_id) VALUES ($1, $2, $3, $4, now(), $5)
`

const LinkIsApiPointExist = `
SELECT count(*)
    FROM links WHERE api_point=$1
`

const LinkByAPIPoint = `
SELECT link_id, title, url, api_point, links.created_at, links.user_id
    FROM links
    WHERE api_point=$1 AND deleted_at is null
`
