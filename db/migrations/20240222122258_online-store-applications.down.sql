-- Drop tables
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS users;

-- Drop ENUM types
DROP TYPE IF EXISTS users_status_enum;
DROP TYPE IF EXISTS users_role_enum;