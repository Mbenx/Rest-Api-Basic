package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mbenx/Rest-Api-Basic/config"
	"github.com/Mbenx/Rest-Api-Basic/models"
	"github.com/danilopolani/gocialite/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

var JWT_SECRET = "secret"

// var Gocial = gocialite.NewDispatcher()

// func main() {
// router := gin.Default()

// router.GET("/", indexHandler)
// router.GET("/auth/:provider", RedirectHandler)
// router.GET("/auth/:provider/callback", CallbackHandler)

// router.Run("127.0.0.1:9090")
// }

// Show homepage with login URL
// func indexHandler(c *gin.Context) {
// 	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
// 		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
// 		"<a href='/auth/linkedin'><button>Login with LinkedIn</button></a><br>" +
// 		"<a href='/auth/facebook'><button>Login with Facebook</button></a><br>" +
// 		"<a href='/auth/google'><button>Login with Google</button></a><br>" +
// 		"<a href='/auth/bitbucket'><button>Login with Bitbucket</button></a><br>" +
// 		"<a href='/auth/amazon'><button>Login with Amazon</button></a><br>" +
// 		"<a href='/auth/amazon'><button>Login with Slack</button></a><br>" +
// 		"</body></html>"))
// }

func init() {
	gotenv.Load()
}

// RedirectHandler .. Redirect to correct oAuth URL
func RedirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("GITHUB_CLIENT_ID"),
			"clientSecret": os.Getenv("GITHUB_CLIENT_SECRET"),
			"redirectURL":  "http://localhost:8080/api/v1/auth/github/callback",
		},
		"facebook": {
			"clientID":     os.Getenv("FACEBOOK_CLIENT_ID"),
			"clientSecret": os.Getenv("FACEBOOK_CLIENT_SECRET"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/facebook/callback",
		},
		"google": {
			"clientID":     os.Getenv("GOOGLE_CLIENT_ID"),
			"clientSecret": os.Getenv("GOOGLE_CLIENT_SECRET"),
			"redirectURL":  "http://localhost:8080/api/v1/auth/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"github":   []string{"public_repo"},
		"facebook": []string{},
		"google":   []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		log.Printf("Log provider", err.Error())
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// CallbackHandler ... Handle Callback of provider
func CallbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	var newUser = getOrRegisterUser(provider, user)
	var jwtToken = createToken(&newUser)

	c.JSON(200, gin.H{
		"message":  "Login Berhasil",
		"data":     newUser,
		"token":    token,
		"jwtToken": jwtToken,
	})

	// Print in terminal user information
	fmt.Printf("%#v", token)
	// fmt.Printf("%#v", user)
	// fmt.Printf("%#v", provider)

	// If no errors, show provider name
	// c.Writer.Write([]byte("Hi, " + user.FullName))

}

func getOrRegisterUser(provider string, user *structs.User) models.User {
	var userData models.User

	config.DB.Where("provider = ? AND social_id = ?", provider, user.ID).First(&userData)

	if userData.ID == 0 {
		newUser := models.User{
			Username: user.Username,
			FullName: user.FullName,
			Email:    user.Email,
			SocialID: user.ID,
			Provider: provider,
			Avatar:   user.Avatar,
		}

		config.DB.Create(&newUser)
		return newUser
	} else {
		return userData
	}
}

func createToken(user *models.User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().AddDate(0, 0, 3).Unix(),
		"iat":       time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString([]byte(JWT_SECRET))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenString)
	return tokenString

}
