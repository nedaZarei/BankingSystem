CREATE TABLE bank (
    bank_id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    headquarter_address VARCHAR(255) NOT NULL,
    PRIMARY KEY (bank_id)
);

CREATE TABLE branch (
    branch_id SERIAL NOT NULL,
    bank_id INTEGER NOT NULL,
    address VARCHAR(255) NOT NULL,
    PRIMARY KEY (branch_id),
    CONSTRAINT branch_bank_id_foreign FOREIGN KEY (bank_id) REFERENCES bank (bank_id)
);

CREATE TABLE employee_login (
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    employee_id SERIAL NOT NULL,
    PRIMARY KEY (username),
    UNIQUE (employee_id)
);

CREATE TABLE employee_details (
    employee_id SERIAL NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    department VARCHAR(255) NOT NULL,
    salary DECIMAL(8, 2) NOT NULL,
    branch_id INTEGER NOT NULL,
    PRIMARY KEY (employee_id),
    CONSTRAINT employee_branch_id_foreign FOREIGN KEY (branch_id) REFERENCES branch (branch_id),
    CONSTRAINT employee_details_employee_id_foreign FOREIGN KEY (employee_id) REFERENCES employee_login (employee_id)
);

CREATE TABLE customer_login (
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    customer_id SERIAL NOT NULL,
    PRIMARY KEY (username),
    UNIQUE (customer_id)
);

CREATE TABLE customer_email (
    email VARCHAR(255) NOT NULL,
    customer_id INTEGER NOT NULL,
    PRIMARY KEY (email),
    CONSTRAINT customer_email_customer_id_foreign FOREIGN KEY (customer_id) REFERENCES customer_login (customer_id)
);

CREATE TABLE customer_details (
    customer_id SERIAL NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    customer_type VARCHAR(255) NOT NULL CHECK (customer_type IN ('Individual', 'Corporate')),
    bank_id INTEGER NOT NULL,
    PRIMARY KEY (customer_id),
    CONSTRAINT customer_bank_id_foreign FOREIGN KEY (bank_id) REFERENCES bank (bank_id),
    CONSTRAINT customer_details_customer_id_foreign FOREIGN KEY (customer_id) REFERENCES customer_login (customer_id)
);

CREATE TABLE account_numbers (
    account_number VARCHAR(255) NOT NULL,
    account_id SERIAL NOT NULL,
    PRIMARY KEY (account_number),
    UNIQUE (account_id)
);

CREATE TABLE account_details (
    account_id SERIAL NOT NULL,
    customer_id INTEGER NOT NULL,
    account_type VARCHAR(255) NOT NULL CHECK (account_type IN ('Savings', 'Checking')),
    account_password VARCHAR(255) NOT NULL,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
    account_status VARCHAR(255) NOT NULL CHECK (account_status IN ('Active', 'Closed', 'Suspended')) DEFAULT 'Active',
    open_date DATE NOT NULL,
    close_date DATE,
    PRIMARY KEY (account_id),
    CONSTRAINT account_details_account_id_foreign FOREIGN KEY (account_id) REFERENCES account_numbers (account_id),
    CONSTRAINT account_customer_id_foreign FOREIGN KEY (customer_id) REFERENCES customer_login (customer_id)
);
COMMENT ON COLUMN account_details.close_date IS 'can be null';

CREATE TABLE manages (
    employee_id INTEGER NOT NULL,
    account_id INTEGER NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    PRIMARY KEY (employee_id, account_id),
    CONSTRAINT manages_employee_id_foreign FOREIGN KEY (employee_id) REFERENCES employee_login (employee_id),
    CONSTRAINT manages_account_id_foreign FOREIGN KEY (account_id) REFERENCES account_numbers (account_id)
);
COMMENT ON COLUMN manages.end_date IS 'can be null';

CREATE TABLE transaction (
    transaction_id SERIAL NOT NULL,
    source_account_id INTEGER NOT NULL,
    destination_account_id INTEGER,
    amount DECIMAL(10,2) NOT NULL,
    transaction_type VARCHAR(255) NOT NULL CHECK (transaction_type IN ('Deposit', 'Withdrawal', 'Transfer', 'Interest')),
    transaction_date TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY (transaction_id),
    CONSTRAINT transaction_source_account_id_foreign FOREIGN KEY (source_account_id) REFERENCES account_numbers (account_id),
    CONSTRAINT transaction_destination_account_id_foreign FOREIGN KEY (destination_account_id) REFERENCES account_numbers (account_id)
);
COMMENT ON COLUMN transaction.destination_account_id IS 'is null for deposit and withdrawal';

CREATE TABLE loan (
    loan_id SERIAL NOT NULL,
    customer_id INTEGER NOT NULL,
    loan_type VARCHAR(255) NOT NULL,
    amount DECIMAL(8, 2) NOT NULL,
    interest_rate DECIMAL(8, 2) NOT NULL,
    duration INTEGER NOT NULL CHECK (duration > 0),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    loan_status VARCHAR(255) NOT NULL,
    PRIMARY KEY (loan_id),
    CONSTRAINT loan_customer_id_foreign FOREIGN KEY (customer_id) REFERENCES customer_login (customer_id)
);
COMMENT ON COLUMN loan.duration IS 'in months';

CREATE TABLE loanPayment (
    payment_id SERIAL NOT NULL,
    loan_id INTEGER NOT NULL,
    payment_amount DECIMAL(8, 2) NOT NULL,
    due_date DATE NOT NULL,
    payment_date DATE NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    PRIMARY KEY (payment_id),
    CONSTRAINT loanpayment_loan_id_foreign FOREIGN KEY (loan_id) REFERENCES loan (loan_id)
);