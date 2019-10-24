package utils

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/IcanFun/utils/utils/log"

	"github.com/fsnotify/fsnotify"
	"github.com/nicksnyder/go-i18n/i18n"
)

var htmlTemplates *template.Template

type HTMLTemplate struct {
	TemplateName string
	Props        map[string]interface{}
	Html         map[string]template.HTML
	Locale       string
}

func InitHTML() {
	InitHTMLWithDir("templates")
}

func InitHTMLWithDir(dir string) {

	if htmlTemplates != nil {
		return
	}

	templatesDir, _ := FindDir(dir)
	log.Debug(T("api.api.init.parsing_templates.debug"), templatesDir)
	var err error
	if htmlTemplates, err = template.ParseGlob(templatesDir + "*.html"); err != nil {
		log.Error(T("api.api.init.parsing_templates.error"), err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(T("web.create_dir.error"), err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Info(T("web.reparse_templates.info"), event.Name)
					if htmlTemplates, err = template.ParseGlob(templatesDir + "*.html"); err != nil {
						log.Error(T("web.parsing_templates.error"), err)
					}
				}
			case err := <-watcher.Errors:
				log.Error(T("web.dir_fail.error"), err)
			}
		}
	}()

	err = watcher.Add(templatesDir)
	if err != nil {
		log.Error(T("web.watcher_fail.error"), err)
	}
}

func NewHTMLTemplate(templateName string, locale string) *HTMLTemplate {
	return &HTMLTemplate{
		TemplateName: templateName,
		Props:        make(map[string]interface{}),
		Html:         make(map[string]template.HTML),
		Locale:       locale,
	}
}

func (t *HTMLTemplate) addDefaultProps(feedbackOrganization, supportEmail, siteName string) {
	var localT i18n.TranslateFunc
	if len(t.Locale) > 0 {
		localT = GetUserTranslations(t.Locale)
	} else {
		localT = T
	}

	t.Props["Footer"] = localT("api.templates.email_footer")

	if feedbackOrganization != "" {
		t.Props["Organization"] = localT("api.templates.email_organization") + feedbackOrganization
	} else {
		t.Props["Organization"] = ""
	}

	t.Html["EmailInfo"] = template.HTML(localT("api.templates.email_info",
		map[string]interface{}{"SupportEmail": supportEmail, "SiteName": siteName}))
}

func (t *HTMLTemplate) Render(feedbackOrganization, supportEmail, siteName string) string {
	t.addDefaultProps(feedbackOrganization, supportEmail, siteName)

	var text bytes.Buffer

	if err := htmlTemplates.ExecuteTemplate(&text, t.TemplateName, t); err != nil {
		log.Error(T("api.api.render.error"), t.TemplateName, err)
	}

	return text.String()
}

func (t *HTMLTemplate) RenderToWriter(w http.ResponseWriter, feedbackOrganization, supportEmail, siteName string) error {
	t.addDefaultProps(feedbackOrganization, supportEmail, siteName)

	if err := htmlTemplates.ExecuteTemplate(w, t.TemplateName, t); err != nil {
		log.Error(T("api.api.render.error"), t.TemplateName, err)
		return err
	}
	return nil
}
