INSERT INTO departments (name, phone)
VALUES ($1, $2)
RETURNING id