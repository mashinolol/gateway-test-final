package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mashinolol/gateway-graph/pkg/transactions"
)

func main() {
	transactionsList := []*transactions.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}

	intervals := []string{"day", "week", "month", "hour"} // изменены порядок для соответствия примерам выше
	for _, interval := range intervals {
		fmt.Printf("\nTransactions in a/an %s:\n", interval)
		result, err := transactions.Group(transactionsList, interval) // используем правильный вызов Group
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, tr := range result {
			fmt.Printf("Value: %d, Timestamp: %v\n", tr.Value, tr.Timestamp)
		}
	}
}
