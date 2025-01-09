CREATE OR REPLACE FUNCTION log_account_creation_date()
RETURNS TRIGGER AS $$
BEGIN
    NEW.open_date := CURRENT_DATE;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_log_account_creation_date
BEFORE INSERT ON account_details
FOR EACH ROW
EXECUTE FUNCTION log_account_creation_date();

CREATE OR REPLACE FUNCTION prevent_customer_deletion_with_loans()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM loan
        WHERE customer_id = OLD.customer_id
          AND loan_status = 'Active'
    ) THEN
        RAISE EXCEPTION 'cannot delete customer with active loans.';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_prevent_customer_deletion
BEFORE DELETE ON customer_details
FOR EACH ROW
EXECUTE FUNCTION prevent_customer_deletion_with_loans();

CREATE OR REPLACE FUNCTION update_account_balance_after_transaction()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        IF NEW.transaction_type IN ('Withdrawal', 'Transfer') THEN
            UPDATE account_details
            SET balance = balance - NEW.amount
            WHERE account_id = NEW.source_account_id;
        END IF;

        IF NEW.transaction_type IN ('Deposit', 'Transfer') AND NEW.destination_account_id IS NOT NULL THEN
            UPDATE account_details
            SET balance = balance + NEW.amount
            WHERE account_id = NEW.destination_account_id;
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_account_balance
AFTER INSERT ON transaction
FOR EACH ROW
EXECUTE FUNCTION update_account_balance_after_transaction();

CREATE OR REPLACE FUNCTION check_sufficient_balance()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.transaction_type IN ('Withdrawal', 'Transfer') THEN
        PERFORM balance
        FROM account_details
        WHERE account_id = NEW.source_account_id
          AND balance >= NEW.amount;

        IF NOT FOUND THEN
            RAISE EXCEPTION 'insufficient balance for this transaction.';
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_sufficient_balance
BEFORE INSERT ON transaction
FOR EACH ROW
EXECUTE FUNCTION check_sufficient_balance();

CREATE OR REPLACE FUNCTION calculate_total_balance(customer_id INTEGER)
RETURNS DECIMAL(15, 2) AS $$
DECLARE
    total_balance DECIMAL(15, 2);
BEGIN
    SELECT COALESCE(SUM(balance), 0)
    INTO total_balance
    FROM account_details
    WHERE customer_id = customer_id;

    RETURN total_balance;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION check_loan_status(loan_id INTEGER)
RETURNS TEXT AS $$
DECLARE
    loan_status TEXT;
BEGIN
    SELECT CASE
        WHEN end_date <= CURRENT_DATE THEN 'Settled'
        ELSE loan_status
    END
    INTO loan_status
    FROM loan
    WHERE loan_id = loan_id;

    RETURN loan_status;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION count_active_loans(customer_id INTEGER)
RETURNS INTEGER AS $$
DECLARE
    active_loan_count INTEGER;
BEGIN
    SELECT COUNT(*)
    INTO active_loan_count
    FROM loan
    WHERE customer_id = customer_id
      AND loan_status = 'Active';

    RETURN active_loan_count;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION calculate_total_loan_payments(loan_id INTEGER)
RETURNS DECIMAL(15, 2) AS $$
DECLARE
    total_payments DECIMAL(15, 2);
BEGIN
    SELECT COALESCE(SUM(payment_amount), 0)
    INTO total_payments
    FROM loanPayment
    WHERE loan_id = loan_id;

    RETURN total_payments;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_customer_name(customer_id INTEGER)
RETURNS TEXT AS $$
DECLARE
    customer_name TEXT;
BEGIN
    SELECT CONCAT(first_name, ' ', last_name)
    INTO customer_name
    FROM customer_details
    WHERE customer_id = customer_id;

    RETURN customer_name;
END;
$$ LANGUAGE plpgsql;
