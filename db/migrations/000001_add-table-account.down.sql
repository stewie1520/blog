START TRANSACTION;

DROP INDEX ix_accounts_email;
DROP TABLE accounts;

COMMIT;
