# Create customer

As a client

I want to create a customer record

### API

PUT /api/customer

    {
        customer_id
        name
        email
        phone_no
    }

### Notes

- API must do upsert. If customer_id is absent, create new customer, else update existing customer
- for new customer, generate sequential customer_id as CUST-xxxx

# Show customer

As a client

I want to get all details of a customer

So that I can seek any information regarding my customer

### API

Get /api/customers/<customer_id>

    {
        name
        email
        phone_no
    }

# List customers

As a client

I want to list all my customers

### API

GET /api/customers

    [{
        name
        email
        phone_no
        gstin
    }]

# Paginate customers

As a client

I want to list all my customers by pages

### API

GET /api/customers?page=3

    [{
        name
        email
        phone_no
    }]
