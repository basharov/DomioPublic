package templater

import (
    "html/template"
    "bytes"
)

type PageData struct {
    LeftColumnLinks  []Link
    RightColumnLinks []Link
    DomainAddLink    Link
}

func GetTopBar() template.HTML {

    t := template.New("Index")

    output, _ := t.Parse(`
                            <div class="b-top-bar-container">
                                <div class="left-area">
                                    {{range .LeftColumnLinks}}
                                         <a href="{{.Url}}">{{.Label}}</a>
                                    {{end}}
                                </div>
                                <div class="right-area">

                                    {{with .DomainAddLink}}
                                        <a href="{{.Url}}" class="{{.ClassName}}">{{.Label}}</a>
                                    {{end}}

                                    {{range .RightColumnLinks}}
                                         <a href="{{.Url}}">{{.Label}}</a>
                                    {{end}}
                                </div>
                            </div>
                        `)

    pageData := PageData{
        LeftColumnLinks:[]Link{
            Link{Url:"/", Label:"Home"},
            Link{Url:"/domains", Label:"Domains"},
        },
        RightColumnLinks:[]Link{
            Link{Url:"/login", Label:"Login"},
        },
        DomainAddLink:Link{Url:"/domains/add", Label:"Add Domain", ClassName:"b-top-bar-container__domain-add-link"},
    }

    var doc bytes.Buffer
    output.Execute(&doc, pageData)

    return template.HTML(doc.String())
}