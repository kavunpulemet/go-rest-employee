SELECT
    e.id AS id,
    e.name AS name,
    e.surname AS surname,
    e.phone AS phone,
    e.company_id AS company_id,
    p.id AS passport_id,
    p.type AS passport_type,
    p.number AS passport_number,
    d.id AS department_id,
    d.name AS department_name,
    d.phone AS department_phone
FROM
    employees e
    JOIN passports p ON e.passport_id = p.id
    JOIN departments d ON e.department_id = d.id
WHERE
    e.company_id = $1;
