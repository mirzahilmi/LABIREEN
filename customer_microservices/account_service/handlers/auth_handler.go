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

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	svc services.AuthService
	ml  mail.EmailSender
}

func NewAuthHandler(svc services.AuthService, ml mail.EmailSender) *authHandler {
	return &authHandler{svc, ml}
}

func (aH *authHandler) RegisterCustomer(ctx *gin.Context) {
	var request entities.CustomerRegister
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.FailOrError(ctx, http.StatusBadRequest, "Bad request", err)
		return
	}

	if request.Password != request.PasswordConfirm {
		response.FailOrError(ctx, http.StatusBadRequest, "Password do not match", nil)
		return
	}

	//Generate Verification Code
	tempCode := uniuri.NewLen(20)
	code := crypto.Encode(tempCode)
	request.VerificationCode = code

	if err := aH.svc.RegisterCustomer(request); err != nil {
		response.FailOrError(ctx, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	request.VerificationCode = ""

	emailData := mail.EmailData{
		Email:   []string{request.Email},
		URL:     os.Getenv("CLIENT_ORIGIN") + "/verifyemail/" + code,
		Name:    request.Name,
		Subject: "Your account verification code",
	}

	aH.ml.SendEmail(&emailData)

	response.Success(ctx, http.StatusOK, "User successfuly created, please check your email for email verification", request)
}

func (aH *authHandler) LoginCustomer(ctx *gin.Context) {
	var request entities.CustomerLogin
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.FailOrError(ctx, http.StatusBadRequest, "Bad request", err)
		return
	}

	err := aH.svc.LoginCustomer(request)
	if err != nil {
		response.FailOrError(ctx, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}

	token, err := jwtx.GenerateToken(request)
	if err != nil {
		response.FailOrError(ctx, http.StatusInternalServerError, "Server error, failed to generate token", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Login Successful", token)
}

func (aH *authHandler) VerifyEmail(ctx *gin.Context) {
	code := ctx.Params.ByName("verification-code")

	if err := aH.svc.VerifyCustomer(code); err != nil {
		response.FailOrError(ctx, http.StatusBadRequest, "User already verified", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Email verified successfully", nil)
}
