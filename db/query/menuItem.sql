-- name: GetMenu :many
SELECT * FROM menu_item AS m
WHERE  m.active = TRUE
and EXISTS (
    SELECT 1 FROM menu_collection AS mc
    WHERE mc.id = m.menuCollectionId 
    AND mc.active = TRUE
    AND mc.name = ?
)
ORDER BY sortOrder;
