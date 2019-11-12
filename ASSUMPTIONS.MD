# Database
- In locations table, the province column should be foreign key from another table that contains provinces data. Because lack of time, I save just the province name.


# Ticket
- When customer purchase ticket, I insert the transaction to database, and I didn't handle if the customer already paid the ticket or not. If the transaction inserted, the remaining quota of ticket that available is the substraction of ticket quantity (quantity from the first time ticket created) and sum of ticket's quantity that already purchased in transactions table. 