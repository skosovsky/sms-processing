package main

import ( //nolint:gci // false positive
	"sms-processing/config"
	log "sms-processing/pkg/logger" //nolint:goimports // false positive
	//nolint:gci // false positive
	_ "github.com/joho/godotenv/autoload" // load .env - for test
)

func main() {
	log.NewLogger(
		log.WithLevel("DEBUG"),
		log.WithAddSource(false),
		log.WithIsJSON(true),
		log.WithMiddleware(false),
		log.WithSetDefault(true))

	appInfo, err := config.NewAppInfo()
	if err != nil {
		log.Fatal("app info", log.ErrAttr(err))
	}
	log.Info("get app name and revision: OK", log.AnyAttr("appInfo", appInfo))

	cfg, err := config.New()
	if err != nil {
		log.Fatal("config", log.ErrAttr(err))
	}
	log.Info("get config: OK", log.AnyAttr("cfg", cfg))
}
