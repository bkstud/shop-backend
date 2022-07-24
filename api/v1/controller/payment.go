package controller

import (
	"fmt"
	"log"
	"net/http"
	"shop/api/v1/auth"
	"shop/config"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	stripeSession "github.com/stripe/stripe-go/v72/checkout/session"
)

func CreateCheckoutSession(c *gin.Context) {
	c.Request.ParseForm()
	prices := c.Request.PostForm["price"]
	names := c.Request.PostForm["name"]

	items := []*stripe.CheckoutSessionLineItemParams{}
	for index := range names {
		floatPrice, _ := strconv.ParseFloat(prices[index], 64)
		items = append(items,
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(names[index]),
					},
					UnitAmount: stripe.Int64(int64(floatPrice * 100)),
				},
				Quantity: stripe.Int64(1),
			})
	}

	cred, err := auth.ReadOauthSecrets("./secrets/stripe-creds.json", "STRIPE")
	if err != nil {
		log.Panic("Failed to initialize facebook credentials")
	}
	stripe.Key = cred.Csecret
	session := sessions.Default(c)
	email := fmt.Sprintf("%v", session.Get("user-id"))
	params := &stripe.CheckoutSessionParams{
		Mode:          stripe.String(string(stripe.CheckoutSessionModePayment)),
		CustomerEmail: stripe.String(email),
		LineItems:     items,
		SuccessURL:    stripe.String(config.FRONTEND_ADDRESS + "/checkout?success=true"),
		CancelURL:     stripe.String(config.FRONTEND_ADDRESS + "/checkout?canceled=true"),
	}

	s, err := stripeSession.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	// c.Redirect(http.StatusSeeOther, s.URL)
	fmt.Println("Redirecting to:=", s.URL)
	c.Redirect(http.StatusSeeOther, s.URL)
}

func handleSuccess(c *gin.Context) {

}
