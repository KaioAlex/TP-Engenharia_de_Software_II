package main

import (
	"fmt"
	"ll/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (c PsqlConfig) ConeccaoInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

type PsqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type MC struct {
	APIKey       string `json:"api_key"`
	PublicAPIKey string `json:"public_api_key"`
	Domain       string `json:"domain"`
}

type Configuracao struct {
	Port     int        `json:"port"`
	Env      string     `json:"env"`
	Pepper   string     `json:"pepper"`
	HMACKey  string     `json:"hmac_key"`
	Database PsqlConfig `json:"database"`
	Mailgun  MC         `json:"mailgun"`
}

func RetornaPsqlConfig() PsqlConfig {
	return PsqlConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "kaio",
		Password: "123",
		Name:     "tp_test",
	}
}

func RetornaConfig() Configuracao {
	return Configuracao{
		Port:     3000,
		Env:      "dev",
		Pepper:   "secret-random-string",
		HMACKey:  "secret-hmac-key",
		Database: RetornaPsqlConfig(),
	}
}

func testingUserService() (*models.Services, error) {
	cfg := RetornaConfig()
	dbCfg := cfg.Database
	services, err := models.NewServices(
		models.WithGorm("postgres", dbCfg.ConeccaoInfo()),
		models.WithLogMode(!(cfg.Env == "prod")),
		models.WithUser(cfg.Pepper, cfg.HMACKey),
		models.WithGallery(),
		models.WithImage(),
	)
	if err != nil {
		panic(err)
	}
	services.AutoMigrate()

	return services, err
}

func TestCreateUser(t *testing.T) {
	serv, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}

	user := models.User{
		Name:     "Caio",
		Email:    "caio@gmail.com",
		Password: "12345678",
	}

	serv.DestructiveReset()

	if err := serv.User.Create(&user); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, user.Name, "Caio", "OK response is expected")

	defer serv.Close()
}
