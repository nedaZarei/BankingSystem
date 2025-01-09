CREATE VIEW customer_accounts AS
SELECT 
    cd.first_name || ' ' || cd.last_name AS customer_name,
    cd.phone_number,
    an.account_number,
    ad.account_type,
    ad.balance
FROM customer_details cd
JOIN customer_login cl ON cd.customer_id = cl.customer_id
JOIN account_details ad ON cl.customer_id = ad.customer_id
JOIN account_numbers an ON ad.account_id = an.account_id;


CREATE VIEW bank_transactions AS
SELECT 
    b.name AS bank_name,
    t.transaction_id,
    src_an.account_number AS source_account_number,
    dst_an.account_number AS destination_account_number,
    t.amount AS transaction_amount,
    t.transaction_date
FROM transaction t
JOIN account_numbers src_an ON t.source_account_id = src_an.account_id
JOIN account_details src_ad ON src_an.account_id = src_ad.account_id
JOIN customer_details src_cd ON src_ad.customer_id = src_cd.customer_id
JOIN bank b ON src_cd.bank_id = b.bank_id
LEFT JOIN account_numbers dst_an ON t.destination_account_id = dst_an.account_id;

CREATE VIEW bank_member AS
SELECT 
    b.name AS bank_name,
    COALESCE(ed.first_name || ' ' || ed.last_name, cd.first_name || ' ' || cd.last_name) AS full_name,
    COALESCE(el.employee_id::TEXT, cl.customer_id::TEXT) AS member_id,
    CASE 
        WHEN ed.employee_id IS NOT NULL THEN 'Employee'
        ELSE 'Customer'
    END AS role,
    COALESCE(ce.email, '') AS email,
    COALESCE(ed.position, '') AS position,
    COALESCE(cd.phone_number, '') AS phone_number
FROM bank b
LEFT JOIN employee_details ed ON ed.branch_id IN (SELECT branch_id FROM branch WHERE bank_id = b.bank_id)
LEFT JOIN employee_login el ON ed.employee_id = el.employee_id
LEFT JOIN customer_details cd ON cd.bank_id = b.bank_id
LEFT JOIN customer_login cl ON cd.customer_id = cl.customer_id
LEFT JOIN customer_email ce ON cl.customer_id = ce.customer_id;