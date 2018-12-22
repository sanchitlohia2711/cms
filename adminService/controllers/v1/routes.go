package v1

import "github.com/gin-gonic/gin"

//SetUpRoutes set up all the routes
func SetUpRoutes(r *gin.Engine) {
	//Oauth routes
	r.POST("/v1/product", ProductCreate)
	r.POST("/v1/oauth/token", TokenHandler)
	r.POST("/v1/oauth/revoke", RevokeHandler)
	r.POST("/v1/oauth/introspect", IntrospectHandler)
	r.POST("/v1/oauth/validate", ValidateHandler)

	//OTP Routes
	r.POST("/v1/oauth/otp/send", GenerateOTP)
	r.POST("/v1/oauth/otp/verify", VerifyOTP)

	//SecurityPinRoutes
	r.POST("/v1/oauth/securitycode/set", SetSecurityCode)
	r.POST("/v1/oauth/securitycode/verify", VerifySecurityCode)

}
