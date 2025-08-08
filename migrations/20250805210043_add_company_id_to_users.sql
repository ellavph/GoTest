-- +goose Up
-- +goose StatementBegin
ALTER TABLE users 
ADD COLUMN company_id INTEGER,
ADD CONSTRAINT fk_users_company 
    FOREIGN KEY (company_id) 
    REFERENCES companies(id) 
    ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users 
DROP CONSTRAINT IF EXISTS fk_users_company,
DROP COLUMN IF EXISTS company_id;
-- +goose StatementEnd