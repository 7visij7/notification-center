package types

type DnsRecord struct {
    Name string
    Ip  string
    Action  string
}

type EvalMatche struct {
    Value float64
    Metric string
}

type GrafanaAlert struct {
    DashboardId int
    EvalMatches []EvalMatche
    RuleName string
    RuleUrl string 
    State string 
    Title string
}

type Options struct {
    Value int `json:"value"`
    Column string `json:"column"`
    Op string  `json:"op"`
}

type Alert struct {
    Description string `json:"Description"`
    Updated_at string `json:"updated_at"`
    Id int `json:"id"`
    Last_triggered_at string `json:"last_triggered_at"`
    User_id int `json:"user_id"`
    Name string `json:"name"`
    Rearm int `json:"rearm"`
    Title string `json:"yitle"`
    Created_at string `json:"created_at"`
    State string `json:"state"`
    Query_id int `json:"query_id"`
    Option Options `json:"options"`
}

type Redash struct {
    Url_base string `json:"url_base"`
    Event string `json:"event"`
    RedashAlerts Alert `json:"alert"`
}

type Message struct {
    Text string `json:"text"`
}