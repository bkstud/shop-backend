package controller

import (
	"fmt"
	"log"
	"net/http"
	"shop/config"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	stripeSession "github.com/stripe/stripe-go/v72/checkout/session"
)

func CreateCheckoutSession(c *gin.Context) {
	stripe.Key = "sk_test_51LMZKLIYEkANHE3L5vMzodC8FMtjo99Nm0RecxG1VOntoohvnp7Tgn0TGbaIVKOQQoSFT3OUF72roSWEmjrOJcjc00DmrqsJ4Y"
	session := sessions.Default(c)
	email := fmt.Sprintf("%v", session.Get("user-id"))
	params := &stripe.CheckoutSessionParams{
		Mode:          stripe.String(string(stripe.CheckoutSessionModePayment)),
		CustomerEmail: stripe.String(email),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Item1"),
					},
					UnitAmount: stripe.Int64(200),
				},
				Quantity: stripe.Int64(1),
			},
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Item2"),
					},
					UnitAmount: stripe.Int64(200),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(config.FRONTEND_ADDRESS + "/checkout?success=true"),
		CancelURL:  stripe.String(config.FRONTEND_ADDRESS + "/checkout?cancel=true"),
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

func handleFailure(c *gin.Context) {

}
