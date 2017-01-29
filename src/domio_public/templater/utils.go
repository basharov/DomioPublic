package templater

import (
    "html/template"
    "net/http"
    "log"
    "github.com/tdewolff/minify"
    "github.com/tdewolff/minify/html"
    "encoding/json"
    "io"
    "os"
)

type Link struct {
    Url       string
    Label     string
    ClassName string
}

type BaseTemplateData struct {
    PageName string
}

type FullPageData struct {
    BaseTemplateData BaseTemplateData
    TopBarData       TopBarData
    PageData         interface{}
}

type TemplateAdditionals func(*template.Template)

func BuildTemplate(addTemplatesToBaseTemplate TemplateAdditionals) *template.Template {

    parsedTemplate, parseErr := template.New("base_template").Parse(getBaseTemplateContent())
    parsedTemplate.New("top_bar_template").Parse(GetTopBarTemplate())

    if (parseErr != nil) {
        log.Fatalln(parseErr)
    }

    addTemplatesToBaseTemplate(parsedTemplate)
    return parsedTemplate
}

func WriteTemplate(w http.ResponseWriter, tmpl *template.Template, pageName string, data interface{}) {
    w.Header().Set("Content-Type", "text/html")

    fullData := FullPageData{
        BaseTemplateData:BaseTemplateData{PageName: pageName},
        TopBarData:GetTopBarData(),
        PageData:data,
    }
    result, _ := json.Marshal(fullData)
    io.Writer(os.Stdout).Write(result)

    execErr := tmpl.Execute(w, fullData)

    if (execErr != nil) {
        log.Print(execErr)
    }
}

func InitMinifier(w http.ResponseWriter, req *http.Request) {
    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    mw := m.ResponseWriter(w, req)
    defer mw.Close()

    w = mw
}
