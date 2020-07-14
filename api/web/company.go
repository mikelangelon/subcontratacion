package web

import (
	"bytes"
	"io"
	"mikelangelon/m/v2/internal/app/company"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type CompanyResource struct {
	CompanyService CompanyService
	OkCall         func(w http.ResponseWriter, r *http.Request)
}

type CompanyService interface {
	CreateCompany(company company.Company) error
	UpdateCompany(company company.Company) error
	GetCompany(companyId string) (*company.Company, error)
	GetLastCompanies() ([]*company.Company, error)
}

func (r CompanyResource) Register(resp http.ResponseWriter, req *http.Request) {
	n, err := strconv.ParseInt(req.FormValue("fundationYear"), 10, 64)
	if err == nil {
		logrus.WithError(err).Info("problem converting fundation year")
	}

	logrus.Info("register!")
	c := company.Company{
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
	fn, _, err := req.FormFile("logo")
	defer fn.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, fn)
	logrus.Println(buf.Bytes())

	err = r.CompanyService.CreateCompany(c)
	if err != nil {
		// TODO Handle
	}
	r.OkCall(resp, req)
}
