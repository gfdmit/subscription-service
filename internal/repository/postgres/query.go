package postgres

const createSubscriptionQuery = `
	INSERT INTO subscriptions.subscriptions (service_name, price, user_id, start_date, end_date)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id, service_name, price, user_id, start_date, end_date
`

const getSubscriptionQuery = `
	SELECT * FROM subscriptions.subscriptions 
	WHERE id = $1
`

const getSubscriptionsQuery = `
	SELECT * FROM subscriptions.subscriptions
`

const updateSubscriptionQuery = `
	UPDATE subscriptions.subscriptions 
	SET service_name = $1, price = $2, user_id = $3, start_date = $4, end_date = $5
	WHERE id = $6
	RETURNING id, service_name, price, user_id, start_date, end_date
`

const deleteSubscriptionQuery = `
	DELETE FROM subscriptions.subscriptions 
	WHERE id = $1
	RETURNING id, service_name, price, user_id, start_date, end_date
`
