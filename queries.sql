INSERT INTO customer_login (username, password, customer_id)
VALUES ('markrobertsss', 'password', DEFAULT);

INSERT INTO customer_details (customer_id, first_name, last_name, birth_date, phone_number, address, customer_type, bank_id)
VALUES (currval('customer_login_customer_id_seq'), 'mark', 'roberts', '1956-01-01', '1234567890', '56 Main St', 'Natural', 1
);

INSERT INTO account_numbers (account_number, account_id)
VALUES ('6534123456', DEFAULT);

INSERT INTO account_details (account_id, customer_id, account_type, account_password, balance, open_date)
VALUES (currval('account_numbers_account_id_seq'), 3, 'Savings', '2456', 4500, CURRENT_DATE);

INSERT INTO transaction (source_account_id, destination_account_id, amount, transaction_type, transaction_date)
SELECT ad.account_id, NULL, 1000.00, 'Deposit', CURRENT_TIMESTAMP
FROM account_details ad
WHERE ad.customer_id = 3
LIMIT 1;

UPDATE account_details 
SET balance = balance + 1000.00
WHERE customer_id = 3;

INSERT INTO transaction (source_account_id, destination_account_id, amount, transaction_type, transaction_date)
SELECT ad.account_id, NULL, 500.00, 'Withdrawal', CURRENT_TIMESTAMP
FROM account_details ad
WHERE ad.customer_id = 3
LIMIT 1;

UPDATE account_details 
SET balance = balance - 500.00
WHERE customer_id = 3;

INSERT INTO transaction (source_account_id, destination_account_id, amount, transaction_type, transaction_date)
SELECT ad.account_id, 1, 300.00, 'Transfer', CURRENT_TIMESTAMP
FROM account_details ad
WHERE ad.customer_id = 3
LIMIT 1;

UPDATE account_details 
SET balance = balance - 300.00
WHERE customer_id = 3;

banking_system=# UPDATE account_details 
SET balance = balance + 300.00
WHERE account_id = 1;

SELECT t.*
FROM transaction t
WHERE t.source_account_id = 3 
   OR t.destination_account_id = 3
ORDER BY t.transaction_date DESC;

SELECT t.*
FROM transaction t
WHERE t.source_account_id = 1 
   OR t.destination_account_id = 1
ORDER BY t.transaction_date DESC;

INSERT INTO loan (customer_id, loan_type, amount, interest_rate, duration, start_date, end_date, loan_status)
VALUES 
    (1, 'Home', 250000.00, 4.5, 360, '2024-01-01', '2054-01-01', 'Active'),
    (1, 'Car', 35000.00, 6.0, 60, '2024-01-01', '2029-01-01', 'Active'),
    (2, 'Personal', 10000.00, 8.5, 24, '2024-01-01', '2026-01-01', 'Active'),
    (3, 'Business', 100000.00, 7.0, 120, '2024-01-01', '2034-01-01', 'Active'),
    (3, 'Education', 50000.00, 5.5, 84, '2024-01-01', '2031-01-01', 'Paidoff');

SELECT *
FROM loan
WHERE loan_status = 'Active';

INSERT INTO loanPayment (loan_id, payment_amount, due_date, payment_date, payment_status)
SELECT 
    loan_id,
    amount / duration as payment_amount,
    start_date + interval '1 month' as due_date,
    start_date + interval '1 month' as payment_date,
    'Paid' as payment_status
FROM loan;

SELECT ad.*, an.account_number
FROM account_details ad
JOIN account_numbers an ON ad.account_id = an.account_id
WHERE ad.balance > 5000;

SELECT ad.*, an.account_number
FROM account_details ad
JOIN account_numbers an ON ad.account_id = an.account_id
WHERE ad.balance > 5000;
banking_system=# SELECT 
    cd.first_name,
    cd.last_name,
    ad.account_type,
    SUM(ad.balance) as total_balance
FROM customer_details cd
JOIN account_details ad ON cd.customer_id = ad.customer_id
JOIN account_numbers an ON ad.account_id = an.account_id
GROUP BY cd.customer_id, cd.first_name, cd.last_name, ad.account_type;

SELECT 
    ed.first_name,
    ed.last_name,
    SUM(l.amount) as total_loan_amount
FROM employee_details ed
JOIN customer_details cd ON CONCAT(ed.first_name, ed.last_name) = CONCAT(cd.first_name, cd.last_name)
JOIN loan l ON cd.customer_id = l.customer_id
WHERE l.loan_status = 'Active'
GROUP BY ed.employee_id, ed.first_name, ed.last_name;

 SELECT 
    cd.first_name,
    cd.last_name,
    COUNT(ad.account_id) as account_count
FROM customer_details cd
JOIN account_details ad ON cd.customer_id = ad.customer_id
GROUP BY cd.customer_id, cd.first_name, cd.last_name
HAVING COUNT(ad.account_id) > 1;

SELECT 
    cd.first_name,
    cd.last_name,
    COUNT(l.loan_id) as active_loan_count
FROM customer_details cd
JOIN loan l ON cd.customer_id = l.customer_id
WHERE l.loan_status = 'Active'
GROUP BY cd.customer_id, cd.first_name, cd.last_name
ORDER BY active_loan_count DESC
LIMIT 5;

SELECT 
    l.*,
    COUNT(lp.payment_id) as paid_installments
FROM loan l
LEFT JOIN loanPayment lp ON l.loan_id = lp.loan_id
WHERE lp.payment_status = 'Paid'
GROUP BY l.loan_id
ORDER BY paid_installments ASC
LIMIT 5;

SELECT 
    cd.first_name,
    cd.last_name,
    l.loan_id,
    l.amount as loan_amount
FROM customer_details cd
JOIN loan l ON cd.customer_id = l.customer_id
JOIN loanPayment lp ON l.loan_id = lp.loan_id
WHERE lp.payment_status = 'Unpaid'
AND lp.due_date < CURRENT_DATE;

SELECT 
    cd.first_name,
    cd.last_name,
    SUM(ad.balance) as total_balance
FROM customer_details cd
JOIN account_details ad ON cd.customer_id = ad.customer_id
GROUP BY cd.customer_id, cd.first_name, cd.last_name
ORDER BY total_balance DESC
LIMIT 5;