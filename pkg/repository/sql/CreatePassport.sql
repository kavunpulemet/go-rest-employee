INSERT INTO passports (type, number)
VALUES ($1, $2)
RETURNING id