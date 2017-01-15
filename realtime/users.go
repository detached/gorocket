package realtime

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/detached/gorocket/api"
)

type ddpLoginRequest struct {
	User     ddpUser `json:"user"`
	Password ddpPassword `json:"password"`
}

type ddpTokenLoginRequest struct {
	Token string `json:"resume"`
}

type ddpUser struct {
	Email string `json:"email"`
}

type ddpPassword struct {
	Digest    string `json:"digest"`
	Algorithm string `json:"algorithm"`
}


// Register a new user on the server. This function does not need a logged in user. The registered user gets logged in
// to set its username.
//
// The ddp methods 'registerUser' and 'setUsername' are not documented.
func (c *Client) RegisterUser(credentials *api.UserCredentials) error {

	if _, err := c.ddp.Call("registerUser", credentials); err != nil {
		return err
	}

	if err := c.Login(credentials); err != nil {
		return err
	}

	if _, err := c.ddp.Call("setUsername", credentials.Name); err != nil {
		return err
	}

	return nil
}

// Login a user. The password and the email are not allowed to be nil.
//
// https://rocket.chat/docs/developer-guides/realtime-api/method-calls/login/
func (c *Client) Login(credentials *api.UserCredentials) (error) {

	digest := sha256.Sum256([]byte(credentials.Password))

	_, err := c.ddp.Call("login", ddpLoginRequest{
		User:     ddpUser{Email: credentials.Email},
		Password: ddpPassword{Digest: hex.EncodeToString(digest[:]), Algorithm:"sha-256"}})

	return err
}