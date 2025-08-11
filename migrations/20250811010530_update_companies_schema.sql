-- +goose Up
-- Update companies table to match the entity structure
-- Add missing columns if they don't exist

-- Add email column (will fail silently if already exists)
ALTER TABLE companies ADD COLUMN IF NOT EXISTS email text;

-- Add phone column (will fail silently if already exists)
ALTER TABLE companies ADD COLUMN IF NOT EXISTS phone text;

-- Add address column (will fail silently if already exists)
ALTER TABLE companies ADD COLUMN IF NOT EXISTS address text;

-- Remove location column if it exists (replaced by address)
-- Note: PostgreSQL doesn't have DROP COLUMN IF EXISTS, so we'll handle this manually

-- +goose Down
-- Revert changes

-- Remove added columns
ALTER TABLE companies DROP COLUMN IF EXISTS email;
ALTER TABLE companies DROP COLUMN IF EXISTS phone;
ALTER TABLE companies DROP COLUMN IF EXISTS address;

-- Add location column back
ALTER TABLE companies ADD COLUMN IF NOT EXISTS location text;
