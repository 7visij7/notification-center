# Devops Api
URL -http://devops-api.infra.company.local/swagger/index.html

DNS:
    For field "action" valid only two value "a" or "d". ( a - Add record, d - Delete record).
Exp:
    curl -X POST "http://devops-api.infra.company.local/api/v1/dns" -H  "Content-Type: application/json" -H "Authorization:SECRETKEY" -d  '{"name":"devops.api.company.local", "ip":"123.123.123.123" , "action":"d"}'

    curl -X POST "http://devops-api.infra.company.local/api/v1/telegram/devops" -H  "Content-Type: application/json" -H "Authorization:$TOKEN" -d  '{"text":"Hello world"}'