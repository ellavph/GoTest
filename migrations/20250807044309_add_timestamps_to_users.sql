-- +goose Up
-- +goose StatementBegin
ALTER TABLE users 
ADD COLUMN created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users 
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;
-- +goose StatementEnd
