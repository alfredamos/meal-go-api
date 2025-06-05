package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alfredamos/go-meal-api/authenticate"
	"github.com/alfredamos/go-meal-api/models"
	"github.com/gin-gonic/gin"
)

func CreatePaymentController(context *gin.Context){
	//----> Get the origin.
  origin := authenticate.GetOrigin(context)

	//----> Get the cancel and success urls.
	cancelUrl, successUrl := models.MakeSuccessAndCancelUrls(origin)

	//----> Get the stripe secret.
	stripeSecretKey, exist := os.LookupEnv("STRIPE_SECRET_KEY")

	//----> Check if stripe secret is in the env.
	if !exist {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Stripe secret key is not available in the environment variable!"})
		return
	}

	//----> Instantiate Payment struct.
	payment := models.Payment{CancelUrl: cancelUrl, SuccessUrl: successUrl, StripeSecretKey: stripeSecretKey}

	//----> Initialize orderPayload
	orderPayload := models.PayloadOrder{}
	reqBody := context.Request.Body
	fmt.Printf("%+v ,request-body : ", reqBody)
	//----> Get the request payload
	err := context.ShouldBindJSON(&orderPayload)

	//----> Check for binding error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed!", "message": err.Error()})
		return
  }
	fmt.Printf("%+v", orderPayload)
	//----> Make the payment by stripe.
	sessionPayload, err := payment.CreatePayment(orderPayload)

	//----> Check for error in payment.
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	//-----> If there's sessionPayload, then store the order in the database.
	if sessionPayload.ID != string("") {
		orderPayload.PaymentId = sessionPayload.ID;
		orderPayload.CheckOutOrder()
	} 

	//----> Send back the response.
	//context.JSON(http.StatusCreated, sessionPayload.ExpiresAt)
	context.JSON(http.StatusCreated, sessionPayload.ID)
}

