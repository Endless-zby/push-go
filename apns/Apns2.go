package apns

import (
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"log"
	"push-go/config"
)

var Client *apns2.Client

func InitApns(cfg *config.ApnsConfig) {
	authKey, err := token.AuthKeyFromFile(cfg.AuthKeyFile)
	if err != nil {
		log.Fatal("token error:", err)
	}

	apnsToken := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: cfg.KeyID,
		// TeamID from developer account (View Account -> Membership)
		TeamID: cfg.TeamID,
	}
	Client = apns2.NewTokenClient(apnsToken).Production()
}
