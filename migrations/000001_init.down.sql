DROP INDEX IF EXISTS subscriptions.idx_subscriptions_user_id;
DROP INDEX IF EXISTS subscriptions.idx_subscriptions_service_name;
DROP INDEX IF EXISTS subscriptions.idx_subscriptions_start_date;

DROP TABLE IF EXISTS subscriptions.subscriptions;

DROP SCHEMA IF EXISTS subscriptions CASCADE;