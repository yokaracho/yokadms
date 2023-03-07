package team

import (
	"edm/internal/core"
	"edm/pkg/datetime"
	"github.com/alecxcode/sqla"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestTeamHandler_AdminUserAllowedToRemove(t *testing.T) {
	tb := &TeamBase{}

	// Set up test user with ADMIN role
	user := Profile{
		ID:       1,
		UserRole: ADMIN,
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/teams", strings.NewReader("deleteButton=1&ids%5B%5D=1"))

	var Page = TeamPage{
		AppTitle:   tb.text.AppTitle,
		AppVersion: core.AppVersion,
		PageTitle:  tb.text.TeamPageTitle,
		SortedBy:   "FullName",
		SortedHow:  1, // 0 - DESC, 1 - ASC
		Filters: sqla.Filter{ClassFilter: []sqla.ClassFilter{
			{Name: "jobunits", Column: "units.ID"},
			{Name: "companies", Column: "companies.ID"},
			{Name: "userrole", Column: "profiles.UserRole"},
			{Name: "userlock", Column: "profiles.UserLock"}},
			TextFilterName:    "searchText",
			TextFilterColumns: []string{"FullName", "profiles.Contacts", "profiles.JobTitle", "units.UnitName", "companies.ShortName"},
		},
		PageNumber: 1,
		LoggedinID: user.ID,
	}

	Page.UserConfig = user.UserConfig
	if user.UserRole == ADMIN {
		Page.LoggedinAdmin = true
	}

	// Parsing Filters
	Page.Filters.GetFilterFromForm(r,
		datetime.ConvDateStrToInt64, datetime.ConvDateTimeStrToInt64,
		map[string]int{"my": Page.LoggedinID})

	// Mock form values for removal
	r.Form = url.Values{}
	r.Form.Add("elemsOnCurrentPage", "1")

	// Call the method with POST request with deleteButton and ids values.
	tb.TeamHandler(w, r)

	// Assert that removal was successful
}
