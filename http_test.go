package jw_scraper_test

import (
	"github.com/QSCTech/jw-scraper"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func init() {
	viper.AutomaticEnv()
}

func TestHttpService_GetLoginPage(t *testing.T) {
	server := jw_scraper.NewHttpService("https://jw.zjuqsc.com")
	page, err := server.GetLoginPage()
	assert.Nil(t, err)
	assert.NotZero(t, page)
}

func TestHttpService_Login(t *testing.T) {
	server := jw_scraper.NewHttpService("https://jw.zjuqsc.com")
	jwbCookie, err := server.Login(
		viper.GetString("TEST_STU_ID"),
		viper.GetString("TEST_PASSWORD"),
		"dDwxNTc0MzA5MTU4Ozs+b5wKASjiu+fSjITNzcKuKXEUyXg=")
	assert.Nil(t, err)
	assert.NotZero(t, jwbCookie)
}

func TestHttpService_Login_Fail(t *testing.T) {
	server := jw_scraper.NewHttpService("https://jw.zjuqsc.com")
	_, err := server.Login(
		viper.GetString("TEST_STU_ID"),
		viper.GetString("TEST_WRONG_PASSWORD"),
		"dDwxNTc0MzA5MTU4Ozs+b5wKASjiu+fSjITNzcKuKXEUyXg=")
	assert.NotNil(t, err)
	stat, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, stat.Code())
}

func TestHttpService_GetDefaultCourses(t *testing.T) {
	server := jw_scraper.NewHttpService("https://jw.zjuqsc.com")
	jwbCookie, err := server.Login(
		viper.GetString("TEST_STU_ID"),
		viper.GetString("TEST_PASSWORD"),
		"dDwxNTc0MzA5MTU4Ozs+b5wKASjiu+fSjITNzcKuKXEUyXg=")
	assert.Nil(t, err)
	page, err := server.GetDefaultCourses(viper.GetString("TEST_STU_ID"), jwbCookie)
	assert.Nil(t, err)
	assert.NotZero(t, page)
	println(page)
}

func TestHttpService_GetCourses(t *testing.T) {
	server := jw_scraper.NewHttpService("https://jw.zjuqsc.com")
	jwbCookie, err := server.Login(
		viper.GetString("TEST_STU_ID"),
		viper.GetString("TEST_PASSWORD"),
		"dDwxNTc0MzA5MTU4Ozs+b5wKASjiu+fSjITNzcKuKXEUyXg=")
	assert.Nil(t, err)
	page, err := server.GetCourses(
		viper.GetString("TEST_STU_ID"),
		jwbCookie,
		"2018-2019",
		"1|秋、冬",
		"dDwtMjQ5Nzk5MzUyO3Q8O2w8aTwwPjs+O2w8dDw7bDxpPDE+O2k8Mz47aTw1PjtpPDg+O2k8MTA+O2k8MTI+O2k8MTQ+O2k8MTY+O2k8MTg+O2k8MjI+O2k8MjY+O2k8Mjg+Oz47bDx0PHQ8OztsPGk8MD47Pj47Oz47dDx0PHA8cDxsPERhdGFUZXh0RmllbGQ7RGF0YVZhbHVlRmllbGQ7PjtsPHhuO3huOz4+Oz47dDxpPDM+O0A8MjAxOC0yMDE5OzIwMTctMjAxODsyMDE2LTIwMTc7PjtAPDIwMTgtMjAxOTsyMDE3LTIwMTg7MjAxNi0yMDE3Oz4+O2w8aTwwPjs+Pjs7Pjt0PHQ8cDxwPGw8RGF0YVRleHRGaWVsZDtEYXRhVmFsdWVGaWVsZDs+O2w8ZHl4cTt4cTE7Pj47Pjt0PGk8Mj47QDzmmKXjgIHlpI8756eL44CB5YasOz47QDwyfOaYpeOAgeWkjzsxfOeni+OAgeWGrDs+PjtsPGk8MD47Pj47Oz47dDxwPHA8bDxUZXh0Oz47bDzlrablj7fvvJozMTYwMTAxMDM0Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzlp5PlkI3vvJrmnY7mmajmm6Y7Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOWtpumZou+8mua1t+a0i+WtpumZojs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w857G7KOS4k+S4minvvJrmtbfmtIvlt6XnqIvkuI7mioDmnK87Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOihjOaUv+ePre+8mua1t+a0i+W3peeoi+S4juaKgOacrzE2MDI7Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFxlOz4+Oz47Oz47dDxAMDxwPHA8bDxWaXNpYmxlO1BhZ2VDb3VudDtfIUl0ZW1Db3VudDtfIURhdGFTb3VyY2VJdGVtQ291bnQ7RGF0YUtleXM7PjtsPG88dD47aTwxPjtpPDg+O2k8OD47bDw+Oz4+Oz47Ozs7Ozs7Ozs7PjtsPGk8MD47PjtsPHQ8O2w8aTwxPjtpPDI+O2k8Mz47aTw0PjtpPDU+O2k8Nj47aTw3PjtpPDg+Oz47bDx0PDtsPGk8MD47aTwxPjtpPDI+O2k8Mz47aTw0PjtpPDU+O2k8Nj47aTw3Pjs+O2w8dDxwPHA8bDxUZXh0Oz47bDxcPEEgaHJlZj0nIycgb25jbGljaz0id2luZG93Lm9wZW4oJ3hzeGpzLmFzcHg/eGtraD1UKDIwMTgtMjAxOS0yKS0wMTFMMDEwMDMxNjAxMDEwMzQnLCdrY2InLCd0b29sYmFyPTAsbG9jYXRpb249MCxkaXJlY3Rvcmllcz0wLHN0YXR1cz0wLG1lbnViYXI9MCxzY3JvbGxiYXJzPTEscmVzaXphYmxlPTEnKSJcPjAxMUwwMTAwXDwvQVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDxcPEEgaHJlZj0nIycgb25jbGljaz0id2luZG93Lm9wZW4oJ3hzeGpzLmFzcHg/eGtraD1UKDIwMTgtMjAxOS0yKS0wMTFMMDEwMDMxNjAxMDEwMzQnLCdrY2InLCd0b29sYmFyPTAsbG9jYXRpb249MCxkaXJlY3Rvcmllcz0wLHN0YXR1cz0wLG1lbnViYXI9MCxzY3JvbGxiYXJzPTEscmVzaXphYmxlPTEnKSJcPuS4reWbvee7j+a1jueDreeCueino+ivu1w8L0FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMDExTDAxMDAzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD7mma/kuYPmnYNcPC9hXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOaYpTs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w85ZGo5LiJ56ysNiw3LDjoioI7Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOe0q+mHkea4r+S4nDFBLTIxNijlpJopOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwyMDE4LTEyLTI3IDA1OjE2OjQxOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwxOz4+Oz47Oz47Pj47dDw7bDxpPDA+O2k8MT47aTwyPjtpPDM+O2k8ND47aTw1PjtpPDY+O2k8Nz47PjtsPHQ8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMDQxSTAxMDAzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD4wNDFJMDEwMFw8L0FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMDQxSTAxMDAzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD7kuJbnlYzlkI3or5fmrKPotY9cPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTA0MUkwMTAwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+5ZC056ybXDwvYVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzmmKU7Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOWRqOS4ieesrDExLDEyLDEz6IqCOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzntKvph5HmuK/opb8yLTIwMyjlpJopIzs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MjAxOC0xMi0yNyAwNToxNTowMjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MTs+Pjs+Ozs+Oz4+O3Q8O2w8aTwwPjtpPDE+O2k8Mj47aTwzPjtpPDQ+O2k8NT47aTw2PjtpPDc+Oz47bDx0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTA0MUkwNDQwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+MDQxSTA0NDBcPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTA0MUkwNDQwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+5Lit5Zu95Y+k5YW45paH5a2m5qyj6LWPXDwvQVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDxcPEEgaHJlZj0nIycgb25jbGljaz0id2luZG93Lm9wZW4oJ3hzeGpzLmFzcHg/eGtraD1UKDIwMTgtMjAxOS0yKS0wNDFJMDQ0MDMxNjAxMDEwMzQnLCdrY2InLCd0b29sYmFyPTAsbG9jYXRpb249MCxkaXJlY3Rvcmllcz0wLHN0YXR1cz0wLG1lbnViYXI9MCxzY3JvbGxiYXJzPTEscmVzaXphYmxlPTEnKSJcPuWtmeaVj+W8ulw8L2FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w85pilOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzlkajkuoznrKwxMSwxMiwxM+iKgjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w857Sr6YeR5riv6KW/MS00MDIo5aSaKTs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MjAxOC0xMi0yNyAwNToxNTo1Nzs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MTs+Pjs+Ozs+Oz4+O3Q8O2w8aTwwPjtpPDE+O2k8Mj47aTwzPjtpPDQ+O2k8NT47aTw2PjtpPDc+Oz47bDx0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTA2MUI5MDkwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+MDYxQjkwOTBcPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTA2MUI5MDkwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+5qaC546H6K665LiO5pWw55CG57uf6K6hXDwvQVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDxcPEEgaHJlZj0nIycgb25jbGljaz0id2luZG93Lm9wZW4oJ3hzeGpzLmFzcHg/eGtraD1UKDIwMTgtMjAxOS0yKS0wNjFCOTA5MDMxNjAxMDEwMzQnLCdrY2InLCd0b29sYmFyPTAsbG9jYXRpb249MCxkaXJlY3Rvcmllcz0wLHN0YXR1cz0wLG1lbnViYXI9MCxzY3JvbGxiYXJzPTEscmVzaXphYmxlPTEnKSJcPuWQtOWbveaholw8L2FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w85pil5aSPOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzlkajkuIDnrKw26IqCXDxiclw+5ZGo5LiA56ysNyw46IqCOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzntKvph5HmuK/opb8xLTIwOCjlpJopOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwyMDE4LTEyLTI3IDA0OjMzOjM1Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwxOz4+Oz47Oz47Pj47dDw7bDxpPDA+O2k8MT47aTwyPjtpPDM+O2k8ND47aTw1PjtpPDY+O2k8Nz47PjtsPHQ8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMDgxQzAxOTEzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD4wODFDMDE5MVw8L0FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMDgxQzAxOTEzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD7mnLrmorDorr7orqHln7rnoYDvvIjnlLLvvIlcPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTA4MUMwMTkxMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+566h5oiQXDwvYVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzmmKXlpI87Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOWRqOS4gOesrDMsNCw16IqCOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzntKvph5HmuK/opb8xLTQwNCjlpJopOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwyMDE4LTEyLTI3IDA1OjE5OjUwOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwxOz4+Oz47Oz47Pj47dDw7bDxpPDA+O2k8MT47aTwyPjtpPDM+O2k8ND47aTw1PjtpPDY+O2k8Nz47PjtsPHQ8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMTAxQzAyNTEzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD4xMDFDMDI1MVw8L0FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8XDxBIGhyZWY9JyMnIG9uY2xpY2s9IndpbmRvdy5vcGVuKCd4c3hqcy5hc3B4P3hra2g9VCgyMDE4LTIwMTktMiktMTAxQzAyNTEzMTYwMTAxMDM0Jywna2NiJywndG9vbGJhcj0wLGxvY2F0aW9uPTAsZGlyZWN0b3JpZXM9MCxzdGF0dXM9MCxtZW51YmFyPTAsc2Nyb2xsYmFycz0xLHJlc2l6YWJsZT0xJykiXD7mlbDlrZfnlLXot6/liIbmnpDkuI7orr7orqFcPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTEwMUMwMjUxMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+5ZGo566tXDwvYVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzmmKXlpI87Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOWRqOWbm+esrDYsNyw4LDnoioJ75Y+M5ZGofVw8YnJcPuWRqOWbm+esrDYsNyw46IqCe+WNleWRqH07Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOe0q+mHkea4r+S4nDMtMzA2XDxiclw+57Sr6YeR5riv6KW/MS0zMTQo5aSaKTs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MjAxOC0xMi0yNyAwNToxMjowMTs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MTs+Pjs+Ozs+Oz4+O3Q8O2w8aTwwPjtpPDE+O2k8Mj47aTwzPjtpPDQ+O2k8NT47aTw2PjtpPDc+Oz47bDx0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTIwMUwwMDMwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+MjAxTDAwMzBcPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTIwMUwwMDMwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+5biC5Zy66JCl6ZSA5qaC6K66XDwvQVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDxcPEEgaHJlZj0nIycgb25jbGljaz0id2luZG93Lm9wZW4oJ3hzeGpzLmFzcHg/eGtraD1UKDIwMTgtMjAxOS0yKS0yMDFMMDAzMDMxNjAxMDEwMzQnLCdrY2InLCd0b29sYmFyPTAsbG9jYXRpb249MCxkaXJlY3Rvcmllcz0wLHN0YXR1cz0wLG1lbnViYXI9MCxzY3JvbGxiYXJzPTEscmVzaXphYmxlPTEnKSJcPuW+kOW/oOa1t1w8L2FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w85aSPOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzlkajkuIDnrKwxMSwxMiwxM+iKgjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w857Sr6YeR5riv6KW/MS0xMDMo5aSaKTs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MjAxOC0xMi0yNyAwNToxODo0Mzs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w8MTs+Pjs+Ozs+Oz4+O3Q8O2w8aTwwPjtpPDE+O2k8Mj47aTwzPjtpPDQ+O2k8NT47aTw2PjtpPDc+Oz47bDx0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTY4MTkwMTMwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+NjgxOTAxMzBcPC9BXD47Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPFw8QSBocmVmPScjJyBvbmNsaWNrPSJ3aW5kb3cub3BlbigneHN4anMuYXNweD94a2toPVQoMjAxOC0yMDE5LTIpLTY4MTkwMTMwMzE2MDEwMTAzNCcsJ2tjYicsJ3Rvb2xiYXI9MCxsb2NhdGlvbj0wLGRpcmVjdG9yaWVzPTAsc3RhdHVzPTAsbWVudWJhcj0wLHNjcm9sbGJhcnM9MSxyZXNpemFibGU9MScpIlw+5L+h5Y+35LiO57O757ufXDwvQVw+Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDxcPEEgaHJlZj0nIycgb25jbGljaz0id2luZG93Lm9wZW4oJ3hzeGpzLmFzcHg/eGtraD1UKDIwMTgtMjAxOS0yKS02ODE5MDEzMDMxNjAxMDEwMzQnLCdrY2InLCd0b29sYmFyPTAsbG9jYXRpb249MCxkaXJlY3Rvcmllcz0wLHN0YXR1cz0wLG1lbnViYXI9MCxzY3JvbGxiYXJzPTEscmVzaXphYmxlPTEnKSJcPui1teWdh1w8L2FcPjs+Pjs+Ozs+O3Q8cDxwPGw8VGV4dDs+O2w85pilOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDzlkajkuoznrKw3LDjoioJcPGJyXD7lkajlm5vnrKwxLDLoioI7Pj47Pjs7Pjt0PHA8cDxsPFRleHQ7PjtsPOe0q+mHkea4r+ilvzEtNTA0KOWkmilcPGJyXD7ntKvph5HmuK/opb8xLTUwNCjlpJopOz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwyMDE5LTAxLTA2IDAwOjUzOjE2Oz4+Oz47Oz47dDxwPHA8bDxUZXh0Oz47bDwxOz4+Oz47Oz47Pj47Pj47Pj47dDxAMDxwPHA8bDxQYWdlQ291bnQ7XyFJdGVtQ291bnQ7XyFEYXRhU291cmNlSXRlbUNvdW50O0RhdGFLZXlzOz47bDxpPDE+O2k8MD47aTwwPjtsPD47Pj47Pjs7Ozs7Ozs7Ozs+Ozs+O3Q8O2w8aTwzPjs+O2w8dDxAMDw7Ozs7Ozs7Ozs7Pjs7Pjs+Pjs+Pjs+Pjs+Alal0Mdy1EQJ38XRwfsObXeLLBw=",
		"xqd",
	)
	assert.Nil(t, err)
	assert.NotZero(t, page)
}
