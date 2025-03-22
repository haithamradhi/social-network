package auth

// endpoint definition

// function to handle the request

// register request type
// {
//     "username": "test",
//     "password": "test",
//     "email": "test",
//     "name": "test"
// }

// register response type
// {
//     "status": "success",
//     "message": "user created",
//     "data": {
//         "id": 1,
//         "username": "test",
//         "email": "test",
//         "name": "test"
//     }
// }

//login
//logout
//get details

// // RegisterHandlers registers handlers for different HTTP requests.
// func RegisterHandlers(rg *routing.RouteGroup, service Service, logger log.Logger) {
// 	rg.Post("/login", login(service, logger))
// }

func RegisterHandlers() {

}

// // login returns a handler that handles user login request.
// func login(service Service, logger log.Logger) routing.Handler {
// 	return func(c *routing.Context) error {
// 		var req struct {
// 			Username string `json:"username"`
// 			Password string `json:"password"`
// 		}

// 		if err := c.Read(&req); err != nil {
// 			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
// 			return errors.BadRequest("")
// 		}

// 		token, err := service.Login(c.Request.Context(), req.Username, req.Password)
// 		if err != nil {
// 			return err
// 		}
// 		return c.Write(struct {
// 			Token    string `json:"token"`
// 			Username string `json:"username"`
// 		}{token, req.Username})
// 	}
// }

// // authenticate returns a handler that validates a JWT token.
// func authenticate(service Service, logger log.Logger) routing.Handler {
// 	return func(c *routing.Context) error {
// 		return c.Write(struct {
// 			Valid bool `json:"valid"`
// 		}{true})
// 	}
// }
