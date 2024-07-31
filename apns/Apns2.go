package apns

import (
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"log"
)

var Client *apns2.Client

func init() {
	authKey, err := token.AuthKeyFromFile("AuthKey_LH4T9V5U4R_5U8LBRXG3A.p8")
	if err != nil {
		log.Fatal("token error:", err)
	}

	apnsToken := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: "LH4T9V5U4R",
		// TeamID from developer account (View Account -> Membership)
		TeamID: "5U8LBRXG3A",
	}
	Client = apns2.NewTokenClient(apnsToken).Production()
}
