/* 
  As queries devem conter o nome da query, o tipo de retorno e a query em si.
  -- name: <QueryName> :returnType<one | many | exec | batchs | copyfrom>
  Os tipos sendo, respectivamente, retorno de uma linha, retorno de várias linhas e execução de uma query.
  E tipos especiais, como batch e copyfrom

  Definido a `query.sql` com as queries, é necessário rodar o comando `sqlc generate` para gerar o código em Go.
 */

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;