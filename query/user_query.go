package query

const (
	FindUserByEmail = "SELECT id, email, password, name, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL"
	FindUserByID    = "SELECT name, email,image FROM users WHERE id = $1"
	AddNewUser      = "INSERT INTO users (email, password, verified_at) VALUES ($1, $2, $3) RETURNING id"
	UpdateUser      = "UPDATE users SET name = $1, updated_at = NOW() WHERE id= $2 AND deleted_at IS NULL RETURNING id"

	AddNewCart = "INSERT INTO carts (user_id) VALUES ($1)"

	FindUnverifiedUser        = "SELECT email FROM unverified_users WHERE email = $1 AND code = $2 AND deleted_at IS NULL"
	FindUnverifiedUserByEmail = "SELECT email FROM unverified_users WHERE email = $1 AND deleted_at IS NULL"
	DeleteUnverifiedUser      = "UPDATE unverified_users SET deleted_at = NOW() WHERE email = $1 AND deleted_at IS NULL RETURNING id"
	AddUnverifiedUser         = "INSERT INTO unverified_users (email, code, expired_at) VALUES ($1, $2, $3) RETURNING code"

	ListUser       = "SELECT id, name, email FROM users WHERE deleted_at IS NULL ORDER BY {{ .SortBy }} {{ .Sort }} LIMIT $1 OFFSET $2"
	UpdateUserName = "UPDATE users SET name = $1, updated_at = NOW() WHERE id = $2 AND deleted_at IS NULL RETURNING id"
)
