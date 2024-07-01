SELECT
    *
FROM biodata.students s
WHERE s.created_at < NOW() OR s.updated_at < NOW()
ORDER BY s.id