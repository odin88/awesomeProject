package controllers

import (
	lib2 "awesomeProject/lib"
	"awesomeProject/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

func FacebookLogin(context *gin.Context) {
	var OAuth2Config = lib2.GetFacebookOAuthConfig()
	url := OAuth2Config.AuthCodeURL(lib2.GetRandomOAuthStateString())
	//http.Redirect(response, request, url, http.StatusTemporaryRedirect)

	fmt.Println(url)
	context.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleFacebookLogin function will handle the Facebook Login Callback
func HandleFacebookLogin(context *gin.Context) {
	//var state = context.PostForm("state")
	var state = context.Request.FormValue("state")
	var code = context.Request.FormValue("code")
	//var state = request.FormValue("state")
	//var code = request.FormValue("code")

	if state != lib2.GetRandomOAuthStateString() {
		//http.Redirect(response, request, "/?invalidlogin=true", http.StatusTemporaryRedirect)
		context.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}

	var OAuth2Config = lib2.GetFacebookOAuthConfig()

	token, err := OAuth2Config.Exchange(oauth2.NoContext, code)

	if err != nil || token == nil {
		//http.Redirect(response, request, "/?invalidlogin=true", http.StatusTemporaryRedirect)
		context.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}

	fbUserDetails, fbUserDetailsError := lib2.GetUserInfoFromFacebook(token.AccessToken)

	if fbUserDetailsError != nil {
		//http.Redirect(response, request, "/?invalidlogin=true", http.StatusTemporaryRedirect)
		context.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}

	authToken, authTokenError := SignInUser(fbUserDetails)

	if authTokenError != nil {
		//http.Redirect(response, request, "/?invalidlogin=true", http.StatusTemporaryRedirect)
		context.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}

	data := gin.H{"Name": "Authorization", "Value": "Bearer " + authToken, "Path": "/"}
	//http.SetCookie(response, cookie)

	//http.Redirect(response, request, "/profile", http.StatusTemporaryRedirect)

	//context.Redirect(http.StatusTemporaryRedirect,"/profile?cookie="+cookie.Value)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    data,
	})
}

func SignInUser(facebookUserDetails lib2.FacebookUserDetails) (string, error) {
	//var result lib.UserDetails

	if facebookUserDetails == (lib2.FacebookUserDetails{}) {
		return "", errors.New("User details Can't be empty")
	}

	if facebookUserDetails.Email == "" {
		return "", errors.New("Last Name can't be empty")
	}

	if facebookUserDetails.Name == "" {
		return "", errors.New("Password can't be empty")
	}

	userExists := models.User{}
	resultExists := db.Where("email = ?", facebookUserDetails.Email).First(&userExists)

	fmt.Println(resultExists)

	tokenString, _ := lib2.CreateJWT(facebookUserDetails.Email)

	if tokenString == "" {
		return "", errors.New("Unable to generate Auth token")
	}

	if resultExists.RowsAffected == 0 {
		user := models.User{}
		user.Email = facebookUserDetails.Email
		user.Name = facebookUserDetails.Name
		user.Token = tokenString
		result := db.Create(&user)
		if result.Error != nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
			return "", errors.New("Error occurred registration")
		}
	}

	if resultExists.RowsAffected == 1 {

		userExists.Token = tokenString
		result := db.Save(&userExists)
		if result.Error != nil {
			return "", result.Error
		}

		if result.Error != nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
			return "", errors.New("Error occurred registration")
		}
	}

	return tokenString, nil
}
