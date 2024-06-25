# gateway-graph
Graph Formatting

Description:
This repository contains a solution in Go for formatting transaction data into specified time intervals (day, week, month, hour).

Requirements: 

Go programming language
Understanding of handling timestamps and data structures in Go

Installation:

Clone the repository:
git clone https://github.com/mashinolol/gateway-test-final.git

Navigate into the project directory:
cd gateway-test-final

Usage:
To use the application, follow these steps:

Set up your environment:
Ensure you have Go installed on your machine.

Run the application:
Execute the main Go file:
go run cmd/main.go

Testing:
Navigate to the transactions folder:
cd pkg/transactions/

gateway-graph/
├── go.mod
├── go.sum
├── main.go
└── pkg/
    └── transactions/
        ├── transactions.go
        └── transactions_test.go

Run tests to verify functionality:
go test -v

Test coverage:
go test -cover 

Functionality
The main functionality of this project is to format transaction data into specified time intervals:

Day: Transactions grouped by day.
Week: Transactions grouped by week.
Month: Transactions grouped by month.
Hour: Transactions grouped by hour.
Example
Here's an example of how the data is formatted when interval == day:

go
Copy code
[
	{
		4456,
		1616025600,
	},
	{
		4231,
		1615939200,
	},
	{
		4321,
		1615852800,
	},
]