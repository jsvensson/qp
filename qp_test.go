package qp_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jsvensson/qp"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPackage(t *testing.T) {
	Convey("When testing the qp package", t, func() {

		Convey("When testing RequiredParams()", func() {
			Convey("With one required parameter", func() {
				Convey("And one value exists", func() {
					url := "http://example.com/?required=1"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params, err := qp.RequiredParams(req, "required")

					Convey("The value is extracted", func() {
						So(err, ShouldBeNil)
						So(params, ShouldHaveLength, 1)
						So(params["required"], ShouldHaveLength, 1)
						So(params["required"][0], ShouldEqual, "1")
					})
				})

				Convey("And two values exist", func() {
					url := "http://example.com/?required=1&required=2"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params, err := qp.RequiredParams(req, "required")

					Convey("The values are extracted", func() {
						So(err, ShouldBeNil)
						So(params, ShouldHaveLength, 1)
						So(params["required"], ShouldHaveLength, 2)
						So(params["required"][0], ShouldEqual, "1")
						So(params["required"][1], ShouldEqual, "2")
					})
				})

				Convey("And the parameter is missing", func() {
					url := "http://example.com/"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params, err := qp.RequiredParams(req, "required")

					Convey("We get an error", func() {
						So(err, ShouldBeError)
						So(params, ShouldBeNil)
					})
				})
			})

		})

		Convey("When testing RequiredParam()", func() {
			Convey("With one required parameter", func() {
				Convey("And one value exists", func() {
					url := "http://example.com/?required=1"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params, err := qp.RequiredParam(req, "required")

					Convey("The value is extracted", func() {
						So(err, ShouldBeNil)
						So(params, ShouldHaveLength, 1)
						So(params["required"], ShouldEqual, "1")
					})
				})

				Convey("And two values exist", func() {
					url := "http://example.com/?required=1&required=2"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params, err := qp.RequiredParam(req, "required")

					Convey("The first value is extracted", func() {
						So(err, ShouldBeNil)
						So(params, ShouldHaveLength, 1)
						So(params["required"], ShouldEqual, "1")
					})
				})

				Convey("And the parameter is missing", func() {
					url := "http://example.com/"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params, err := qp.RequiredParam(req, "required")

					Convey("We get an error", func() {
						So(err, ShouldBeError)
						So(params, ShouldBeNil)
					})
				})
			})

		})

		Convey("When testing Params()", func() {

			Convey("With one required parameter", func() {
				Convey("And one value exists", func() {
					url := "http://example.com/?required=1"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params := qp.Params(req, "required")

					Convey("The value is extracted", func() {
						So(params, ShouldNotBeNil)
						So(params, ShouldHaveLength, 1)
						So(params["required"][0], ShouldEqual, "1")
					})
				})

				Convey("And two values exist", func() {
					url := "http://example.com/?required=1&required=2"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params := qp.Params(req, "required")

					Convey("The values are extracted", func() {
						So(params, ShouldNotBeNil)
						So(params, ShouldHaveLength, 1)
						So(params["required"][0], ShouldEqual, "1")
						So(params["required"][1], ShouldEqual, "2")
					})
				})

				Convey("And the parameter is missing", func() {
					url := "http://example.com/"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					params := qp.Params(req, "required")

					Convey("We get a nil slice", func() {
						So(params, ShouldBeNil)
					})
				})
			})

		})

		// Test Param()
		Convey("When testing Param()", func() {

			Convey("With one parameter", func() {
				Convey("And one value exists", func() {
					url := "http://example.com/?required=1"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					param, ok := qp.Param(req, "required")

					Convey("The value is extracted", func() {
						So(ok, ShouldBeTrue)
						So(param, ShouldEqual, "1")
					})
				})

				Convey("And two values exist", func() {
					url := "http://example.com/?required=1&required=2"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					param, ok := qp.Param(req, "required")

					Convey("The values are extracted", func() {
						So(ok, ShouldBeTrue)
						So(param, ShouldHaveLength, 1)
						So(param, ShouldEqual, "1")
					})
				})

				Convey("And the parameter is missing", func() {
					url := "http://example.com/"
					req := httptest.NewRequest(http.MethodGet, url, nil)

					param, ok := qp.Param(req, "required")

					Convey("We get no value", func() {
						So(param, ShouldHaveLength, 0)
						So(ok, ShouldBeFalse)
					})
				})
			})

		})
	})
}
