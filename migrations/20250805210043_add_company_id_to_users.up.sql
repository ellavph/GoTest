ALTER TABLE users 
ADD COLUMN company_id INTEGER,
ADD CONSTRAINT fk_users_company 
    FOREIGN KEY (company_id) 
    REFERENCES companies(id) 
    ON DELETE SET NULL;