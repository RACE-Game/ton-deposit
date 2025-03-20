package rest

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	webapps "github.com/Fuchsoria/telegram-webapps"
	"github.com/go-chi/cors"

	"github.com/RACE-Game/ton-deposit/cmd"
)

func corsMiddleware(next http.Handler) http.Handler {
	h := cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		//ExposedHeaders:   []string{""},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		//Debug:            true,
	})

	return h(next)
}

func versionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-App-Version", cmd.Version)

		next.ServeHTTP(w, r)
	})
}

func showHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Printf("Header field %q, Value %q\n", k, v)
		}

		next.ServeHTTP(w, r)
	})
}

func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "No basic auth present"}`))
			return
		}

		if username != "admin" || password != "likeABoss" {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid username or password"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

const TelegramUserKey = "telegram_user"

func checkTelegramAuthMiddleware(next http.Handler, tgAPIKey, appSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		secret := r.Header.Get("App-Secret")
		if secret == appSecret && secret != "" && appSecret != "" {
			next.ServeHTTP(w, r)
			return
		}

		value := r.Header.Get("Init-Data")

		err, user := webapps.VerifyWebAppData(value, tgAPIKey)
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte(`{"message": "init invalid user"}`))
		// 	return
		// }
		if err != nil {
			fmt.Printf("wrong init data, error: %s, %s %s \n data: %+v \n raw id: %s\n",
				err.Error(), r.Method, r.URL.Path, user, value)

			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), TelegramUserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetHash(value string) (string, error) {
	data, err := url.ParseQuery(value)
	if err != nil {
		return "", err
	}

	if len(data["hash"]) == 0 {
		return "", fmt.Errorf("no hash in %s", value)
	}

	return data["hash"][0], nil
}
