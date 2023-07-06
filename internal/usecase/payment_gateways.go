package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/the-go-dragons/final-project/internal/domain"
	"github.com/the-go-dragons/final-project/pkg/config"
)

var getRedirectURL = func() string {
	return config.Config.Payment.Redirect.Url
}

type Bank string

const (
	Saderat Bank = "saderat"
)

func (b Bank) String() string {
	return string(b)
}

var BANKS = map[Bank]func() Gateway{
	Saderat: NewSaderatGateway,
}

func validateAndGetBank(bank string) (Bank, error) {
	if _, ok := BANKS[Bank(bank)]; ok {
		return Bank(bank), nil
	}
	return "", errors.New("invalid bank name")
}

type PaymentPage struct {
	URL      string
	Method   string
	DataType string
	Data     map[string]interface{}
}
type Gateway interface {
	GetName() string
	GetToken(order domain.Order) string
	GetPaymentPage(token string) PaymentPage
	VerifyPayment(data map[string][]string) (domain.Payment, error)
}

type SaderatGateway struct {
	name       Bank
	tokenURL   string
	paymentURL string
	verifyURL  string
	terminalId string
}

func NewSaderatGateway() Gateway {

	return SaderatGateway{
		name:       Saderat,
		tokenURL:   config.Config.Payment.Gateways.Saderat.Urls.Token,
		paymentURL: config.Config.Payment.Gateways.Saderat.Urls.Payment,
		verifyURL:  config.Config.Payment.Gateways.Saderat.Urls.Verify,
		terminalId: config.Config.Payment.Gateways.Saderat.Terminal.Id,
	}
}
func (s SaderatGateway) GetName() string {
	return s.name.String()
}

func (s SaderatGateway) GetPaymentPage(token string) PaymentPage {
	data := map[string]interface{}{
		"token":      token,
		"terminalID": s.terminalId,
	}
	return PaymentPage{
		URL:      s.paymentURL,
		Method:   "POST",
		DataType: "FORM",
		Data:     data,
	}
}

func (s SaderatGateway) VerifyPayment(data map[string][]string) (domain.Payment, error) {
	payment := domain.Payment{}
	for key, val := range data {
		fmt.Println(key, " : ", val)
	}
	vdata := url.Values{
		"Tid":            {s.terminalId},
		"digitalreceipt": {data["digitalreceipt"][0]},
	}
	res, err := http.PostForm(s.verifyURL, vdata)
	if err != nil {
		return payment, err
	}
	var result map[string]string
	json.NewDecoder(res.Body).Decode(&result)
	for key, val := range result {
		fmt.Println(key, " : ", val)
	}
	if result["Status"] != "Ok" {
		return payment, errors.New("payment faild")
	}
	amount, _ := strconv.ParseInt(result["ReturnId"], 10, 64)
	// Todo: should remove the following if statment, its here because of a bug in testbank
	if amount <= 0 {
		amount, _ = strconv.ParseInt(data["amount"][0], 10, 64)
	}

	invoiceID, _ := strconv.Atoi(data["invoiceid"][0])
	payment.OrderID = uint(invoiceID)

	payment.PayAmount = amount
	return payment, nil
}

func (s SaderatGateway) GetToken(order domain.Order) string {

	data := url.Values{
		"InvoiceID":   {fmt.Sprint(order.ID)},
		"TerminalID":  {s.terminalId},
		"Amount":      {fmt.Sprint(order.Amount)},
		"callbackURL": {getRedirectURL() + "?bank=" + s.GetName()},
	}
	res, err := http.PostForm(s.tokenURL, data)

	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	return fmt.Sprint(result["Accesstoken"])
}
