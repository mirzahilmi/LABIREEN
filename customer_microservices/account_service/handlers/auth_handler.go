package handlers

import (
	"labireen/customer_microservices/account_service/entities"
	"labireen/customer_microservices/account_service/services"
	"os"

	"labireen/customer_microservices/account_service/utilities/crypto"
	"labireen/customer_microservices/account_service/utilities/jwtx"
	"labireen/customer_microservices/account_service/utilities/mail"
	"labireen/customer_microservices/account_service/utilities/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	RegisterCustomer(ctx *gin.Context)
	LoginCustomer(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
}

type authHandlerImpl struct {
	svc services.AuthService
	ml  mail.EmailSender
}

func NewAuthHandler(svc services.AuthService, ml mail.EmailSender) *authHandlerImpl {
	return &authHandlerImpl{svc, ml}
}

func (aH *authHandlerImpl) RegisterCustomer(ctx *gin.Context) {
	var request entities.CustomerRegister
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log := response.ErrorLog{
			Field:   "request",
			Message: "Request does not fullfil requirements",
		}
		response.Error(ctx, http.StatusBadRequest, "Bad request", log)
		return
	}

	if request.Password != request.PasswordConfirm {
		log := response.ErrorLog{
			Field:   "Password",
			Message: "Password should be the same as Confirm Password",
		}
		response.Error(ctx, http.StatusForbidden, "Password mismatch", log)
		return
	}

	//Generate Verification Code
	code := crypto.Encode(request.Email)
	request.VerificationCode = code

	if err := aH.svc.RegisterCustomer(request); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	request.VerificationCode = ""

	emailData := mail.EmailData{
		Email:   []string{request.Email},
		URL:     os.Getenv("CLIENT_ORIGIN") + "/auth/verify/" + code,
		Name:    request.Name,
		Subject: "Your account verification code",
	}

	aH.ml.SendEmail(&emailData)

	response.Success(ctx, http.StatusOK, "User successfuly created, please check your email for email verification", request)
}

func (aH *authHandlerImpl) LoginCustomer(ctx *gin.Context) {
	var request entities.CustomerLogin
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log := response.ErrorLog{
			Field:   "request",
			Message: "Request does not fullfil requirements",
		}
		response.Error(ctx, http.StatusBadRequest, "Bad request", log)
		return
	}

	id, err := aH.svc.LoginCustomer(request)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "Failed to logged in", err.Error())
		return
	}

	token, err := jwtx.GenerateToken(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Server error, failed to generate token", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Login Successful", token)
}

func (aH *authHandlerImpl) VerifyEmail(ctx *gin.Context) {
	code := ctx.Params.ByName("verification-code")

	if err := aH.svc.VerifyCustomer(code); err != nil {
		response.Error(ctx, http.StatusBadRequest, "User verification failed", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Email verified successfully", nil)
}
