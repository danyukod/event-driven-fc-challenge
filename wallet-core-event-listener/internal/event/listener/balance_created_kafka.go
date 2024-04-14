package listener

import (
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/danyukod/wallet-core-event-listener/internal/usecase/create_balance"
	"github.com/danyukod/wallet-core-event-listener/pkg/kafka"
)

type CreateBalanceKafkaListener struct {
	Kafka   *kafka.Consumer
	Usecase *create_balance.CreateBalanceUseCase
}

type BalanceUpdated struct {
	Name    string
	Payload interface{}
}

func NewCreateBalanceKafkaListener(kafka *kafka.Consumer, usecase *create_balance.CreateBalanceUseCase) *CreateBalanceKafkaListener {
	return &CreateBalanceKafkaListener{
		Kafka:   kafka,
		Usecase: usecase,
	}
}

func (l *CreateBalanceKafkaListener) Listen() {
	kafkaMessageChan := make(chan *ckafka.Message)
	go l.Kafka.Consume(kafkaMessageChan)
	for msg := range kafkaMessageChan {
		var balanceUpdated BalanceUpdated
		err := json.Unmarshal(msg.Value, &balanceUpdated)
		if err != nil {
			println(err.Error())
			continue
		}

		accountIdFrom := balanceUpdated.Payload.(map[string]interface{})["account_id_from"]
		balanceAccountIDFrom := balanceUpdated.Payload.(map[string]interface{})["balance_account_id_from"]

		balanceFrom := create_balance.CreateBalanceInputDTO{
			AccountID: accountIdFrom.(string),
			Amount:    balanceAccountIDFrom.(float64),
		}

		accountIdTo := balanceUpdated.Payload.(map[string]interface{})["account_id_to"]
		balanceAccountIDTo := balanceUpdated.Payload.(map[string]interface{})["balance_account_id_to"]

		balanceTo := create_balance.CreateBalanceInputDTO{
			AccountID: accountIdTo.(string),
			Amount:    balanceAccountIDTo.(float64),
		}
		err = l.Usecase.Execute(balanceFrom)
		if err != nil {
			println(err.Error())
			return
		}
		err = l.Usecase.Execute(balanceTo)
		if err != nil {
			println(err.Error())
			return
		}
	}

}
