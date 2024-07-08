package apis

import "net/http"

func Login() http.HandlerFunc {
	route := "/api/users/login"

	return func(w http.ResponseWriter, r *http.Request) {

		resp := Response{
			Success: true,
		}

		RespondWith(w, r, route, resp)
	}
}
