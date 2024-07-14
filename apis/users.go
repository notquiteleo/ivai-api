package apis

import (
	"ivai-api/models"
	"net/http"
	"time"
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
			Mobile: mobile,
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

func MobileLogin() http.HandlerFunc {
	route := "/api/users/mobile_login"

	return func(w http.ResponseWriter, r *http.Request) {
		mobile := r.URL.Query().Get("mobile")
		if mobile == "" {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "mobile is required",
			})
			return
		}

		verityCode := r.URL.Query().Get("verity_code")
		if verityCode == "" {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "verity_code is required",
			})
			return
		}

		user, err := models.GetUserByMobile(mobile)
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		log, err := models.GetLatestLogByContent(user.ID, verityCode)
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
		if log.Content != verityCode {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: "verity code is incorrect",
			})
			return
		}


		_, err = GenerateToken(int64(user.ID))
		if err != nil {
			RespondWith(w, r, route, Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		resp := Response{
			Success: true,
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
