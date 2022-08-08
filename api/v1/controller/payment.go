package controller

import (
	"fmt"
	"log"
	"net/http"
	"shop/api/v1/auth"
	"shop/api/v1/model"
	"shop/config"
	"strconv"

	"github.com/gin-contrib/sessions"
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
	ids := c.Request.PostForm["id"]

	items := []*stripe.CheckoutSessionLineItemParams{}
	for _, prodID := range ids {
		internalItem := FindItemById(c, prodID)
		if internalItem == nil {
			return
		}
		itemPrice := internalItem.Price
		itemName := internalItem.Name

		item := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("usd"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:     stripe.String(itemName),
					Metadata: map[string]string{"InternalID": prodID},
				},
				UnitAmount: stripe.Int64(int64(itemPrice * 100)),
			},
			Quantity: stripe.Int64(1),
		}
		items = append(items, item)
	}

	session := sessions.Default(c)
	email := fmt.Sprintf("%v", session.Get("user-id"))
	successUrl := fmt.Sprintf("%s:%d/api/v1/payment/success?session_id={CHECKOUT_SESSION_ID}", config.SERVER_ADDRESS, config.SERVER_PORT)
	params := &stripe.CheckoutSessionParams{
		Mode:          stripe.String(string(stripe.CheckoutSessionModePayment)),
		CustomerEmail: stripe.String(email),
		LineItems:     items,
		SuccessURL:    stripe.String(successUrl),
		CancelURL:     stripe.String(config.FRONTEND_ADDRESS + "/checkout?canceled=true"),
	}

	s, err := stripeSession.New(params)

	if err != nil {
		log.Panic("Failed to create session:", err)
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}

func HandleSuccess(c *gin.Context) {
	query := c.Request.URL.Query()
	s, _ := stripeSession.Get(query.Get("session_id"), nil)
	user := model.User{}
	Db.First(&user, "email = ?", s.CustomerEmail)
	i := stripeSession.ListLineItems(s.ID, &stripe.CheckoutSessionListLineItemsParams{})
	for i.Next() {
		line := i.LineItem()
		p, _ := product.Get(line.Price.Product.ID, nil)
		id64, err := strconv.ParseUint(p.Metadata["InternalID"], 10, 32)
		if err != nil {
			log.Panic("Failed to convert id:", err)
		}
		itemID := uint(id64)

		t := model.Transaction{
			ItemID:  itemID,
			UserID:  user.ID,
			Payment: string(s.PaymentStatus),
			Type:    "purchase",
			Status:  "pending",
		}
		if err := Db.Create(&t).Error; err != nil {
			log.Panic("Failed to create new Transaction:", err)
		}

	}

	c.Redirect(http.StatusSeeOther, config.FRONTEND_ADDRESS+"/checkout?success=true")
}
