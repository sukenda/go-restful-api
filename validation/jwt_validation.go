package validation

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/sukenda/go-restful-api/config"
	"github.com/sukenda/go-restful-api/entity"
	"time"
)

func CreateToken(user entity.User) (string, error) {
	configuration := config.New()
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte(configuration.Get("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, err

}
