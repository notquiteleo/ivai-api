package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ivai-api/models"
	"net/http"
	"time"

	"ivai-api/middlewares/jwt"
)

func VerityCode() http.HandlerFunc {
	route := "/api/users/verity_code"

	return func(w http.ResponseWriter, r *http.Request) {
		mobile := r.URL.Query().Get("mobile")
		if mobile == "" {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "mobile is required",
			})
			return
		}

		user, err := models.FindOrCreateUser(&models.Users{
			UID:    GenerateUUID(),
			Mobile: mobile,
			Name:   mobile,
		})
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		// TODO 短信服务api
		verityCode := "123456"
		err = models.CreateLog(&models.UserLogs{
			UserID:    user.ID,
			Content:   verityCode,
			CreatedAt: time.Now(),
			Action:    models.UserLogMobileLogin,
		})
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		resp := Response{
			Success: true,
			Data:    verityCode,
		}

		RespondWith(w, r, route, resp)
	}
}

type MobileLoginRequest struct {
	Mobile     string `json:"mobile"`
	VerityCode string `json:"verity_code"`
}

func MobileLogin() http.HandlerFunc {
	route := "/api/users/mobile_login"

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO get body from post api
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var req MobileLoginRequest
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		// mobile := r.URL.Query().Get("mobile")
		if req.Mobile == "" {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "mobile is required",
			})
			return
		}

		// verityCode := r.URL.Query().Get("verity_code")
		if req.VerityCode == "" {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "verity_code is required",
			})
			return
		}

		user, err := models.GetUserByMobile(req.Mobile)
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		log, err := models.GetLatestLogByContent(user.ID, req.VerityCode)
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if time.Now().Sub(log.CreatedAt) > 5*time.Minute {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "verity code is expired",
			})
			return
		}

		if log.Content != req.VerityCode {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "verity code is incorrect",
			})
			return
		}

		token, err := jwt.GenerateToken(fmt.Sprintf("%d", user.ID))
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		resp := Response{
			Success: true,
			Data: map[string]string{
				"token": token,
			},
		}

		RespondWith(w, r, route, resp)
	}
}

func WxLogin() http.HandlerFunc {
	route := "/api/users/wx_login"

	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "code is required",
			})
			return
		}
		resp := Response{
			Success: true,
		}

		RespondWith(w, r, route, resp)
	}
}
