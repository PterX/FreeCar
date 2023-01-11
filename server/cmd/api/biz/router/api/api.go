// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	api "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/car", append(_createcarMw(), api.CreateCar)...)
	root.GET("/car", append(_getcarMw(), api.GetCar)...)
	root.POST("/profile", append(_submitprofileMw(), api.SubmitProfile)...)
	root.DELETE("/profile", append(_clearprofileMw(), api.ClearProfile)...)
	root.GET("/trips", append(_gettripsMw(), api.GetTrips)...)
	{
		_auth := root.Group("/auth", _authMw()...)
		_auth.POST("/login", append(_loginMw(), api.Login)...)
	}
	root.GET("/profile", append(_profileMw(), api.GetProfile)...)
	_profile := root.Group("/profile", _profileMw()...)
	_profile.POST("/photo", append(_createprofilephotoMw(), api.CreateProfilePhoto)...)
	_profile.DELETE("/photo", append(_clearprofilephotoMw(), api.ClearProfilePhoto)...)
	_profile.GET("/photo", append(_photoMw(), api.GetProfilePhoto)...)
	_photo := _profile.Group("/photo", _photoMw()...)
	_photo.POST("/complete", append(_completeprofilephotoMw(), api.CompleteProfilePhoto)...)
	root.POST("/trip", append(_tripMw(), api.CreateTrip)...)
	_trip := root.Group("/trip", _tripMw()...)
	_trip.GET("/:id", append(_gettripMw(), api.GetTrip)...)
	_trip.PUT("/:id", append(_updatetripMw(), api.UpdateTrip)...)
}