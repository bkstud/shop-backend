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
	"github.com/stripe/stripe-go/v72/product"
)

var cred auth.Credentials

func init() {
	var err error
	cred, err = auth.ReadOauthSecrets("./secrets/stripe-creds.json", "STRIPE")
	if err != nil {
		log.Panic("Failed to initialize stripe credentials")
	}
	stripe.Key = cred.Csecret
}

func CreateCheckoutSession(c *gin.Context) {
	c.Request.ParseForm()
	// TODO: prices and names can't be retrieved from form as this can be exploited
	// They have to be taken from Items database by ID.
	prices := c.Request.PostForm["price"]
	names := c.Request.PostForm["name"]
	ids := c.Request.PostForm["id"]

	items := []*stripe.CheckoutSessionLineItemParams{}
	for index, prodID := range ids {
		floatPrice, _ := strconv.ParseFloat(prices[index], 64)
		item := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:     stripe.String(names[index]),
					Metadata: map[string]string{"InternalID": prodID},
				},
				UnitAmount: stripe.Int64(int64(floatPrice * 100)),
			},
			Quantity: stripe.Int64(1),
		}
		items = append(items, item)
	}

	session := sessions.Default(c)
	email := fmt.Sprintf("%v", session.Get("user-id"))
	params := &stripe.CheckoutSessionParams{
		Mode:          stripe.String(string(stripe.CheckoutSessionModePayment)),
		CustomerEmail: stripe.String(email),
		LineItems:     items,
		SuccessURL:    stripe.String("https://localhost:5000/api/v1/payment/success?session_id={CHECKOUT_SESSION_ID}"), //stripe.String(config.FRONTEND_ADDRESS + "/checkout?success=true"),
		CancelURL:     stripe.String(config.FRONTEND_ADDRESS + "/checkout?canceled=true"),
	}

	s, err := stripeSession.New(params)

	if err != nil {
		log.Panic("Failed to create session: %v", err)
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}

func HandleSuccess(c *gin.Context) {
	query := c.Request.URL.Query()
	s, _ := stripeSession.Get(query.Get("session_id"), nil)
	// log.Println("Customer:=", s.Customer)
	// log.Println("CustomerEmail:=", s.CustomerEmail)
	i := stripeSession.ListLineItems(s.ID, &stripe.CheckoutSessionListLineItemsParams{})
	for i.Next() {
		line := i.LineItem()
		p, _ := product.Get(line.Price.Product.ID, nil)
		log.Println("Got product:=", p)
		log.Println("Need to create Transaction with ID:=", p.Metadata["InternalID"])

	}

	c.Redirect(http.StatusSeeOther, config.FRONTEND_ADDRESS+"/checkout?success=true")
}
