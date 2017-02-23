package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData}}
                    <input class="subscription-id-input" type="hidden" value="{{.Subscription.Id}}"></input>

                    <div class="b-subscription-config-mainview">

                        <div>

                            {{range .Records}}

                                <div>
                                    <label><strong>{{.Type}}:</strong></label>

                                    <ul>

                                        {{range .ResourceRecords}}

                                            <li>
                                                <label>Value:</label>
                                                <label>{{.Value}}</label>

                                                <button class="delete-record-button" value="{{.Value}}">Delete</button>

                                            </li>

                                        {{end}}

                                    </ul>
                                <hr>
                                </div>

                            {{end}}

                        </div>

                        <form class="b-subscription-record-form">
                            <select class="value-type-select">
                                <option value="a">A – IPv4 address</option>
                                <option value="cname">CNAME – Canonical name</option>
                                <option value="mx">MX – Mail exchange</option>
                                <option value="aaaa">AAAA – IPv6 address</option>
                                <option value="txt">TXT – Text</option>
                                <option value="ptr">PTR – Pointer</option>
                                <option value="srv">SRV – Service locator</option>
                                <option value="spf">SPF – Sender Policy Framework</option>
                                <option value="naptr">NAPTR – Name Authority Pointer</option>
                                <option value="ns">NS – Name server</option>
                                <option value="soa" disabled="disabled" title="Only one SOA record can exist in a zone, you can select it to edit its properties.">SOA – Start of authority</option>
                            </select>

                            <input class="value-input"></input>
                            <input type="submit" class="update-subscription-button" type="submit" value="Add"></input>
                        </form>

                    </div>

                {{end}}

            {{end}}
        `
}