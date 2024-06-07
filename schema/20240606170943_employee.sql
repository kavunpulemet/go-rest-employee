-- +goose Up
-- +goose StatementBegin
CREATE TABLE passports (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    number VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    phone VARCHAR(255) UNIQUE NOT NULL,
    company_id INT NOT NULL,
    passport_id INT UNIQUE REFERENCES passports(id),
    department_id INT REFERENCES departments(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employees;

DROP TABLE departments;

DROP TABLE passports;
-- +goose StatementEnd
