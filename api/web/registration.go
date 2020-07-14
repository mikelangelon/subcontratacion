package web

import (
	"bytes"
	"io"
	"mikelangelon/m/v2/internal/app/company"
	"mikelangelon/m/v2/internal/app/user"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type RegistrationResource struct {
	CompanyService CompanyService
	UserService    user.UserService
	OkCall         func(w http.ResponseWriter, r *http.Request)
}

func (r RegistrationResource) Register(resp http.ResponseWriter, req *http.Request) {
	u := GetUser(req)
	logrus.WithField("username", u).Error("current user")
	err := r.UserService.Register(u)
	if err != nil {
		logrus.WithError(err).WithField("username", u.User).Error("Problem creating user")
	}
	c := GetCompany(req)
	err = r.CompanyService.CreateCompany(c)
	if err != nil {
		// TODO Handle
	}
	r.OkCall(resp, req)
}

func GetUser(req *http.Request) user.User {
	username := req.FormValue("username")
	password := req.FormValue("password")

	return user.User{
		User:     username,
		Password: password,
	}
}

func GetCompany(req *http.Request) company.Company {
	n, err := strconv.ParseInt(req.FormValue("fundationYear"), 10, 64)
	if err == nil {
		logrus.WithError(err).Info("problem converting fundation year")
	}

	fn, _, err := req.FormFile("logo")
	defer fn.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, fn)
	logrus.Println(buf.Bytes())

	return company.Company{
		Name:          req.FormValue("name"),
		CIF:           req.FormValue("cif"),
		City:          req.FormValue("city"),
		CP:            req.FormValue("cp"),
		Description:   req.FormValue("description"),
		Employees:     req.FormValue("employees"),
		FundationYear: n,
		Quantity:      req.FormValue("quantity"),
		Street:        req.FormValue("street"),
		Web:           req.FormValue("web"),
	}
}
