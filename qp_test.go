package qp_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jsvensson/qp"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRequiredParams(t *testing.T) {
	Convey("When making a request", t, func() {
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
}

func TestRequiredParam(t *testing.T) {
	Convey("When making a request", t, func() {
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
}

func TestParams(t *testing.T) {
	Convey("When making a request", t, func() {

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
}
