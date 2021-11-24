package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"bytes"

	rice "github.com/GeertJohan/go.rice"
	"github.com/campoy/svg-badge/badge"
	"github.com/joho/godotenv"
	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"golang.org/x/image/font/gofont/goregular"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func stripeConfigHandler(w http.ResponseWriter, r *http.Request) {
	type Config struct {
		StripePK            string
		StripeACC           string
		StripeApiVersion    string
		StripeLocale        string
	}

	switch r.Method {
	case http.MethodGet:
		config := Config{
			os.Getenv("STRIPE_PK"),
			os.Getenv("STRIPE_ACC"),
			os.Getenv("STRIPE_API_VERSION"),
			os.Getenv("STRIPE_LOCALE"),
		}

		// log.Fatal(configJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(config)
	default:
		http.Error(w, "This endpoint only supports GET requrests.", http.StatusMethodNotAllowed)
		return
	}
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	type Config struct {
		MetaTitle         string
		Title             string
		Description       string
		DefaultCurrency   string
		BackLink          string
		EmailLink         string
		GitLink           string
		GithubLink        string
		GitlabLink        string
		FacebookLink      string
		TwitterLink       string
		InstagramLink     string
		LinkedinLink      string
		BGColor           string
		TextColor         string
		ButtonColor       string
		TopNavColor       string
		BGColorDark       string
		TextColorDark     string
		ButtonColorDark   string
		TopNavColorDark   string
	}

	switch r.Method {
	case http.MethodGet:
		config := Config{
			os.Getenv("META_TITLE"),
			os.Getenv("TITLE"),
			os.Getenv("DESCRIPTION"),
			os.Getenv("DEFAULT_CURRENCY"),
			os.Getenv("BACK_LINK"),
			os.Getenv("EMAIL_LINK"),
			os.Getenv("GIT_LINK"),
			os.Getenv("GITHUB_LINK"),
			os.Getenv("GITLAB_LINK"),
			os.Getenv("FACEBOOK_LINK"),
			os.Getenv("TWITTER_LINK"),
			os.Getenv("INSTAGRAM_LINK"),
			os.Getenv("LINKEDIN_LINK"),
			os.Getenv("BG_COLOR"),
			os.Getenv("TEXT_COLOR"),
			os.Getenv("BUTTON_COLOR"),
			os.Getenv("TOP_NAV_COLOR"),
			os.Getenv("BG_COLOR_DARK"),
			os.Getenv("TEXT_COLOR_DARK"),
			os.Getenv("BUTTON_COLOR_DARK"),
			os.Getenv("TOP_NAV_COLOR_DARK"),
		}

		// log.Fatal(configJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(config)
	default:
		http.Error(w, "This endpoint only supports GET requrests.", http.StatusMethodNotAllowed)
		return
	}
}

func sendOwncastNotification(message string) {
	type ReqBody struct {
		Body string `json:"body"`
	}

	if os.Getenv("OWNCAST_HOSTNAME") != "" {
		reqUrl :=  "https://" + os.Getenv("OWNCAST_HOSTNAME") + "/api/integrations/chat/system"

		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + os.Getenv("OWNCAST_TOKEN")
						
		ownCastBody := &ReqBody{
			Body: message,
		}
		
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(ownCastBody)

		// Create a new request using http
		req, err := http.NewRequest(http.MethodPost, reqUrl, payloadBuf)

		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json")

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		log.Println(string([]byte(body)))
		log.Println(resp)
	}
}

func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		type LineItems struct {
			Amount      int64
			Currency    string
			Description string
		}

		type CheckoutData struct {
			ClientSecret string `json:"client_secret"`
		}

		type Response struct {
			SessionID string
			StripePK  string
		}

		var items LineItems
		_ = json.NewDecoder(r.Body).Decode(&items)

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(items.Amount),
			Currency: stripe.String(items.Currency),
			Description: stripe.String(items.Description),
		}

		intent, _ := paymentintent.New(params)
		
		data := CheckoutData{
			ClientSecret: intent.ClientSecret,
		}

		sendOwncastNotification("Donation of " + strconv.FormatInt(items.Amount / 100, 10) + items.Currency + ". Message: " + items.Description)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	default:
		http.Error(w, "This endpoint only supports POST requrests.", http.StatusMethodNotAllowed)
		return
	}
}

func badgeHandler(w http.ResponseWriter, r *http.Request) {
	m, err := badge.NewMaker("Go Regular", goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	badgeFontSize, err := strconv.ParseFloat(os.Getenv("BADGE_FONT_SIZE"), 32)
	badgeHeight, err := strconv.Atoi(os.Getenv("BADGE_HEIGHT"))

	b, err := m.New("EL", os.Getenv("BADGE_TEXT"), os.Getenv("BADGE_COLOR"), badgeFontSize, badgeHeight)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health Check"))
}

func main() {
	loadEnv()
	stripe.Key = os.Getenv("STRIPE_SK")
	assets := http.FileServer(http.Dir("./"))

	http.Handle("/", http.FileServer(rice.MustFindBox("frontend/dist").HTTPBox()))
	http.Handle("/assets/", assets)
	http.HandleFunc("/api/site/config", configHandler)
	http.HandleFunc("/api/pay/config", stripeConfigHandler)
	http.HandleFunc("/api/pay/secret", getSecretHandler)
	http.HandleFunc("/api/badge", badgeHandler)
	http.HandleFunc("/api/healthcheck", healthCheckHandler)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
