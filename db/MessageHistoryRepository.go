package db

import (
	"log"
	"push-go/entity"
)

const (
	findAllByClientIdLimit = "SELECT * FROM history WHERE client_id = ? order by `timestamp` desc LIMIT ? OFFSET ?"
)

func FindAllByClientIdLimit(clientId string, page int, size int) ([]entity.MessageHistory, error) {
	var results []entity.MessageHistory
	offset := (page - 1) * size
	if err := PushDb.Raw(findAllByClientIdLimit, clientId, size, offset).Scan(&results).Error; err != nil {
		log.Printf("error querying database: %v", err.Error)
		return nil, err
	}
	return results, nil
}
