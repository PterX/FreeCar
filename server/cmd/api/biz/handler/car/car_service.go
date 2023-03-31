// Code generated by hertz generator.

package car

import (
	"context"
	"net/http"

	hcar "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/car"
	"github.com/CyanAsterisk/FreeCar/server/cmd/api/config"
	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/errno"
	kcar "github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/car"
	"github.com/CyanAsterisk/FreeCar/server/shared/tools"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// AdminCreateCar .
// @router /admin/car/car [PUT]
func AdminCreateCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hcar.AdminCreateCarRequest
	resp := new(kcar.CreateCarResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalCarClient.CreateCar(ctx, &kcar.CreateCarRequest{PlateNum: req.PlateNum})
	if err != nil {
		hlog.Error("rpc car service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// AdminDeleteCar .
// @router /admin/car/car [DELETE]
func AdminDeleteCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hcar.AdminDeleteCarRequest
	resp := new(kcar.DeleteCarResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalCarClient.DeleteCar(ctx, &kcar.DeleteCarRequest{
		AccountId: c.MustGet(consts.AccountID).(int64),
		Id:        req.ID,
	})
	if err != nil {
		hlog.Error("rpc car service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// AdminGetSomeCars .
// @router /admin/car/some [GET]
func AdminGetSomeCars(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hcar.AdminGetSomeCarsRequest
	resp := new(kcar.GetSomeCarsResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalCarClient.GetSomeCars(ctx, &kcar.GetSomeCarsRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc car service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// AdminGetAllCars .
// @router /admin/car/all [GET]
func AdminGetAllCars(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hcar.AdminGetAllCarsRequest
	resp := new(kcar.GetAllCarsResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalCarClient.GetAllCars(ctx, &kcar.GetAllCarsRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc car service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetCars .
// @router /mini/car/cars [GET]
func GetCars(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hcar.GetCarsRequest
	resp := new(kcar.GetCarsResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalCarClient.GetCars(ctx, &kcar.GetCarsRequest{AccountId: c.MustGet(consts.AccountID).(int64)})
	if err != nil {
		hlog.Error("rpc car service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetCar .
// @router /mini/car/car [GET]
func GetCar(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hcar.GetCarRequest
	resp := new(kcar.GetCarResponse)

	if err = c.BindAndValidate(&req); err != nil {
		resp.BaseResp = tools.BuildBaseResp(errno.ParamsErr)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp, err = config.GlobalCarClient.GetCar(ctx, &kcar.GetCarRequest{
		AccountId: c.MustGet(consts.AccountID).(int64),
		Id:        req.ID,
	})
	if err != nil {
		hlog.Error("rpc car service err", err)
		resp.BaseResp = tools.BuildBaseResp(errno.ServiceErr)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
