-- name: GetMenu :many
SELECT * FROM menuItem AS m
WHERE  m.active = TRUE
and EXISTS (
    SELECT 1 FROM menuCollection AS mc
    WHERE mc.id = m.menuCollectionId 
    AND mc.active = TRUE
    AND mc.name = ?
)
ORDER BY sortOrder;
