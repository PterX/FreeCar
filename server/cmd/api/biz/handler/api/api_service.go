// Code generated by hertz generator.

package api

import (
	"context"
	"time"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/server/cmd/api"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/config"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/pkg"
	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/car"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/profile"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/trip"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/user"
	"github.com/CyanAsterisk/FreeCar/server/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt"
)

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	// rpc to get accountID
	resp, err := config.GlobalUserClient.Login(ctx, &user.LoginRequest{Code: req.Code})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}
	// create a JWT
	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		ID: resp.AccountId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		errno.SendResponse(c, errno.ServiceErr, nil)
		return
	}
	// return token
	errno.SendResponse(c, errno.Success, api.LoginResponse{
		Token:     token,
		ExpiredAt: time.Now().Unix() + consts.ThirtyDays,
	})
}

// GetUserInfo .
// @router /user/info [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalUserClient.GetUser(ctx, &user.GetUserRequest{AccontId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp.UserInfo)
}

// UpdateUserInfo .
// @router /user/info [POST]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	_, err = config.GlobalUserClient.UpdateUser(ctx, &user.UpdateUserRequest{
		AccountId:   aid.(int64),
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, api.UpdateUserResponse{})
}

// UploadAvatar .
// @router /user/avatar [POST]
func UploadAvatar(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalUserClient.UploadAvatar(ctx, &user.UploadAvatarRequset{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, api.UploadAvatarResponse{UploadUrl: resp.UploadUrl})
}

// CreateCar .
// @router /car [POST]
func CreateCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalCarClient.CreateCar(ctx, &car.CreateCarRequest{
		AccountId: aid.(int64),
		PlateNum:  req.PlateNum,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCCarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCar .
// @router /car [GET]
func GetCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalCarClient.GetCar(ctx, &car.GetCarRequest{
		AccountId: aid.(int64),
		Id:        req.Id,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCCarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetCars .
// @router /cars [GET]
func GetCars(ctx context.Context, c *app.RequestContext) {
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.AuthorizeFail, nil)
		return
	}
	resp, err := config.GlobalCarClient.GetCars(ctx, &car.GetCarsRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCCarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp.Cars)
}

// GetProfile .
// @router /profile [GET]
func GetProfile(ctx context.Context, c *app.RequestContext) {
	var err error

	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.GetProfile(ctx, &profile.GetProfileRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp.Profile)
}

// SubmitProfile .
// @router /profile [POST]
func SubmitProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SubmitProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.SubmitProfile(ctx, &profile.SubmitProfileRequest{
		AccountId: aid.(int64),
		Identity: &profile.Identity{
			LicNumber:       req.Identity.LicNumber,
			Name:            req.Identity.Name,
			Gender:          profile.Gender(req.Identity.Gender),
			BirthDateMillis: req.Identity.BirthDateMillis,
		},
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// ClearProfile .
// @router /profile [DELETE]
func ClearProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ClearProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.ClearProfile(ctx, &profile.ClearProfileRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp.Profile)
}

// GetProfilePhoto .
// @router /profile/photo [GET]
func GetProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.GetProfilePhoto(ctx, &profile.GetProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp.Url)
}

// CreateProfilePhoto .
// @router /profile/photo [POST]
func CreateProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.CreateProfilePhoto(ctx, &profile.CreateProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CompleteProfilePhoto .
// @router /profile/photo/complete [POST]
func CompleteProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.CompleteProfilePhoto(ctx, &profile.CompleteProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp.Identity)
}

// ClearProfilePhoto .
// @router /profile/photo [DELETE]
func ClearProfilePhoto(ctx context.Context, c *app.RequestContext) {
	var err error
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.ClearProfilePhoto(ctx, &profile.ClearProfilePhotoRequest{AccountId: aid.(int64)})
	if err != nil {
		errno.SendResponse(c, errno.RPCProfileSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// CreateTrip .
// @router /trip [POST]
func CreateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CreateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalTripClient.CreateTrip(ctx, &trip.CreateTripRequest{
		Start: &trip.Location{
			Latitude:  req.Start.Latitude,
			Longitude: req.Start.Longitude,
		},
		CarId:     req.CarId,
		AvatarUrl: req.AvatarUrl,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCTripSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrip .
// @router /trip/:id [GET]
func GetTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetTripRequest
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	req.Id = c.Param("id")

	resp, err := config.GlobalTripClient.GetTrip(ctx, &trip.GetTripRequest{
		Id:        req.Id,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCTripSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetTrips .
// @router /trips [GET]
func GetTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetTripsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalTripClient.GetTrips(ctx, &trip.GetTripsRequest{
		Status:    trip.TripStatus(req.Status),
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCTripSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// UpdateTrip .
// @router /trip/:id [PUT]
func UpdateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	req.Id = c.Param(consts.ID)

	resp, err := config.GlobalTripClient.UpdateTrip(ctx, &trip.UpdateTripRequest{
		Id: req.Id,
		Current: &trip.Location{
			Latitude:  req.Current.Latitude,
			Longitude: req.Current.Longitude,
		},
		EndTrip:   req.EndTrip,
		AccountId: aid.(int64),
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCTripSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// AdminLogin .
// @router /admin/login [POST]
func AdminLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AdminLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	// rpc to get accountID
	resp, err := config.GlobalUserClient.AdminLogin(ctx, &user.AdminLoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}
	// create a JWT
	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		ID: resp.AccountId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		errno.SendResponse(c, errno.ServiceErr, nil)
		return
	}
	// return token
	errno.SendResponse(c, errno.Success, api.AdminLoginResponse{
		Token:     token,
		ExpiredAt: time.Now().Unix() + consts.ThirtyDays,
	})
}

// ChangeAdminPassword .
// @router /admin/password [POST]
func ChangeAdminPassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ChangeAdminPasswordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	aid, flag := c.Get(consts.AccountID)
	if !flag {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalUserClient.ChangeAdminPassword(ctx, &user.ChangeAdminPasswordRequest{
		AccountId:    aid.(int64),
		OldPassword:  req.OldPassword,
		NewPassword_: req.NewPassword,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// AddUser .
// @router /user [POST]
func AddUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AddUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalUserClient.AddUser(ctx, &user.AddUserRequest{
		AccountId:    req.AccountId,
		Username:     req.Username,
		PhoneNumber:  req.PhoneNumber,
		AvatarBlobId: req.AvatarBlobId,
		OpenId:       req.OpenId,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// DeleteUser .
// @router /user [DELETE]
func DeleteUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalUserClient.DeleteUser(ctx, &user.DeleteUserRequest{AccountId: req.AccountId})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetSomeUsers .
// @router /user/some [GET]
func GetSomeUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetSomeUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalUserClient.GetSomeUsers(ctx, &user.GetSomeUsersRequest{})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// GetAllUsers .
// @router /user/all [GET]
func GetAllUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetAllUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalUserClient.GetAllUsers(ctx, &user.GetAllUsersRequest{})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}

	errno.SendResponse(c, errno.Success, resp)
}

// DeleteCar .
// @router /car [DELETE]
func DeleteCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalCarClient.DeleteCar(ctx, &car.DeleteCarRequest{
		Id: req.Id,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCCarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)

}

// UpdateCar .
// @router /car/update [POST]
func UpdateCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateCarRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalCarClient.AdminUpdateCar(ctx, &car.AdminUpdateCarRequest{
		Id:  req.Id,
		Car: pkg.ConvertCar(req.Car),
	})
	if err != nil {
		errno.SendResponse(c, errno.CarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetSomeCars .
// @router /cars/some [GET]
func GetSomeCars(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetSomeCarsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalCarClient.GetSomeCars(ctx, &car.GetSomeCarsRequest{})
	if err != nil {
		errno.SendResponse(c, errno.CarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetAllCars .
// @router /cars/all [GET]
func GetAllCars(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetAllCarsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalCarClient.GetAllCars(ctx, &car.GetAllCarsRequest{})
	if err != nil {
		errno.SendResponse(c, errno.CarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// DeleteProfile .
// @router /profile [DELETE]
func DeleteProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalProfileClient.DeleteProfile(ctx, &profile.DeleteProfileRequest{AccountId: req.AccountId})
	if err != nil {
		errno.SendResponse(c, errno.CarSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
	return
}

// GetAllProfile .
// @router /profiles/all [GET]
func GetAllProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetAllProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, err)
		return
	}

	resp, err := config.GlobalProfileClient.GetAllProfile(ctx, &profile.GetAllProfileRequest{})
	if err != nil {
		errno.SendResponse(c, errno.ProfileSrvErr, nil)
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetSomeProfile .
// @router /profiles/some [GET]
func GetSomeProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetSomeProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, err)
		return
	}

	resp, err := config.GlobalProfileClient.GetSomeProfile(ctx, &profile.GetSomeProfileRequest{})
	if err != nil {
		errno.SendResponse(c, errno.ProfileSrvErr, nil)
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetPendingProfile .
// @router /profile/pending [GET]
func GetPendingProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetPendingProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalProfileClient.GetPendingProfile(ctx, &profile.GetPendingProfileRequest{})
	if err != nil {
		errno.SendResponse(c, errno.ServiceErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetAllTrips .
// @router /trips/all [GET]
func GetAllTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetAllTripsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalTripClient.GetAllTrips(ctx, &trip.GetAllTripsRequest{})
	if err != nil {
		errno.SendResponse(c, errno.TripSrvErr, nil)
	}
	errno.SendResponse(c, errno.Success, resp)
}

// GetSomeTrips .
// @router /trips/some [GET]
func GetSomeTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetSomeTripsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	resp, err := config.GlobalTripClient.GetSomeTrips(ctx, &trip.GetSomeTripsRequest{})
	if err != nil {
		errno.SendResponse(c, errno.TripSrvErr, nil)
	}
	errno.SendResponse(c, errno.Success, resp)
}

// DeleteTrip .
// @router /trip [DELETE]
func DeleteTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DeleteTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalTripClient.DeleteTrip(ctx, &trip.DeleteTripRequest{Id: req.Id})
	if err != nil {
		errno.SendResponse(c, errno.TripSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
}

// CheckProfile .
// @router /profile/check [POST]
func CheckProfile(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CheckProfileRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}
	resp, err := config.GlobalProfileClient.CheckProfile(ctx, &profile.CheckProfileRequest{
		AccountId: req.AccountId,
		Accept:    req.Accept,
	})
	if err != nil {
		errno.SendResponse(c, errno.ProfileSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, resp)
	return
}

// UpdateUser .
// @router /user [PUT]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UpdateUserByIDRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		errno.SendResponse(c, errno.ParamsErr, nil)
		return
	}

	_, err = config.GlobalUserClient.UpdateUser(ctx, &user.UpdateUserRequest{
		AccountId:   req.AccountId,
		Username:    req.Username,
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		errno.SendResponse(c, errno.RPCUserSrvErr, nil)
		return
	}
	errno.SendResponse(c, errno.Success, api.UpdateUserByIDResponse{})
}
