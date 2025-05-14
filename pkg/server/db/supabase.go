package db

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/supabase-go"
)

var (
	url string
	key string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding without it")
	}

	key = os.Getenv("SUPABASE_KEY")
	url = os.Getenv("SUPABASE_URL")
	if key == "" || url == "" {
		log.Fatal("SUPABASE_KEY and SUPABASE_URL must be set")
	}
}

func NewClient() *supabase.Client {
	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatalf("failed to create Supabase client: %v", err)
	}

	return client
}

func NewUserClient(token string) (*gotrue.Client, error) {
	client, err := supabase.NewClient(url, key, nil)
	userClient := client.Auth.WithToken(token)
	if err != nil {
		return nil, errors.New("failed to create Supabase client")
	}
	return &userClient, nil
}
