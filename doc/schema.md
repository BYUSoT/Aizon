# Schema
The Aizon application is structured around the concept of an order.
Those orders are associated with customers, and hours and material usages are logged against them.
The state of an order is tracked, until the invoice is generated and the client is charged.

## Orders
An order has the following characteristics:
- id
- customer w/ department
- ordered date
- completed date
- customer PO
- team
- description
- status
- auxilliary contact
- billing contact
- operating unit
- payment type
