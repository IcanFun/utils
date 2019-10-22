package utils

import (
	"net/http"
	"net/url"
)

func RenderWebError(message, details, siteName string, statusCode int, w http.ResponseWriter, r *http.Request) {
	T, _ := GetTranslationsAndLocale(r)

	title := T("api.templates.error.title", map[string]interface{}{"SiteName": siteName})
	link := "/"
	linkMessage := T("api.templates.error.link")

	status := http.StatusTemporaryRedirect
	if statusCode != http.StatusInternalServerError {
		status = statusCode
	}

	http.Redirect(
		w,
		r,
		"/error?title="+url.QueryEscape(title)+
			"&message="+url.QueryEscape(message)+
			"&details="+url.QueryEscape(details)+
			"&link="+url.QueryEscape(link)+
			"&linkmessage="+url.QueryEscape(linkMessage),
		status,
	)
}
