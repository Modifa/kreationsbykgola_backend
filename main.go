package main

import (
	router "Modifa/kreationsbykgola_backend/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	gin.DisableConsoleColor()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())
	// r.Use(CORSMiddleware())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// OPTIONS method for ReactJS
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "x-access-token", "content-type", "Content-Length", "Authorization", "Cache-Control"}
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	/************/
	/* Users */
	/************/
	router.Init(r)

	return r
}

func setupConfigs() {
	os.Setenv("CURRENTDOMAIN", "https://906b-102-32-62-145.in.ngrok.io")

	os.Setenv("EmailFrom", "projectcommunication123@gmail.com")
	//sk_test_e4186c254bd9ccdcf8b9c012db2624fb8ad1b239

	os.Setenv("KreationsByKgolaEmail", "kreationsbykgola@gmail.com")

	os.Setenv("PaymentAuthSecret", "sk_test_e4186c254bd9ccdcf8b9c012db2624fb8ad1b239")
	// os.Setenv("PaymentAuthSecretPublic", "pk_test_3bedcaf8d5d75c8c16923944d9066aeeb29735b2")

	os.Setenv("PaymentChargeUrl", "https://api.paystack.co/transaction/initialize")
	os.Setenv("PaymentChargeAuthorizationUrl", "https://api.paystack.co/transaction/charge_authorization")
	// https://api.paystack.co/transaction/verify/:reference
	os.Setenv("VerifyUrl", "https://api.paystack.co/transaction/verify/")

	os.Setenv("VerifyCardUrl", "https://api.paystack.co/transferrecipient")

	os.Setenv("ProjectMain", "postgres://cogjgedlgavael:cf43a86f559ebdd296331ca10991a0bfc87dfcf1fb7c83d3407698719348a669@ec2-18-204-74-74.compute-1.amazonaws.com:5432/d7jnruc4m8g23q")
	os.Setenv("WEBSERVER_PORT", "8080")

}

func main() {
	//Uncommented When Not Debugging
	gin.SetMode(gin.ReleaseMode)
	// export GIN_MODE=release

	// gocron.Start()
	// s := gocron.NewScheduler()
	// gocron.Every(2).Seconds().Do(c.CheckNewUser)
	//  <-s.Start()

	r := setupRouter()

	setupConfigs()

	r.Run(":" + os.Getenv("WEBSERVER_PORT"))
}
