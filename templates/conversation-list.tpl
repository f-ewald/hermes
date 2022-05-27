Conversations (ID: Participant, ...)
====================================
{{ range . -}}
{{ .ID }}: {{ range $i, $e := .Participants }}{{if $i}}, {{end}}{{$e.Number}}{{ end }}
{{ end }}