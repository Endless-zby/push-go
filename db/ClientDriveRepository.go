package db

import (
	"log"
	"push-go/entity"
)

const (
	findByClientIdQuery = "SELECT * FROM token WHERE client_id = ?"
)

func FindByClientId(clientId string) (entity.ClientDrive, error) {
	var clientDrive entity.ClientDrive
	if err := PushDb.Raw(findByClientIdQuery, 123456).Scan(&clientDrive).Error; err != nil {
		log.Printf("error querying database: %v", err.Error)
		return entity.ClientDrive{}, err
	}
	return clientDrive, nil
}
