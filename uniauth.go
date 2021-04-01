package fiber_middleware

import (
	"fmt"

	"github.com/UniAuth/fiber-middleware/minion"
	"github.com/UniAuth/fiber-middleware/models"
	"github.com/gofiber/fiber/v2"
)

var uniAuthInstance models.UniAuthInstance
var appConfigs []models.ApplicationConfig

// AuthenticateMiddleware redirects user to authorization page
func AuthenticateMiddleware(configName string) fiber.Handler {
	// return instance of fiber middleware
	return func(c *fiber.Ctx) error {
		config, err := minion.GetConfigByName(appConfigs, configName)
		if err != nil {
			fmt.Println(fmt.Sprintf("Config named %s was not found, using %s", configName, config.Name))
		}
		// generate url of UniAuth Server
		authLink := fmt.Sprintf("%s/%s?client_id=%s&redirect_uri=%s", config.Url, config.AuthEndpoint, config.ClientId, config.RedirectUri)
		// take a permanent redirect to auth server
		return c.Redirect(authLink, 302)
	}
}

func Init(configs []models.ApplicationConfig) models.UniAuthInstance {
	uniAuthInstance = models.UniAuthInstance{
		Authenticate: AuthenticateMiddleware,
	}
	fmt.Println("UniAuth Module Injected")
	appConfigs = configs

	return uniAuthInstance
}
