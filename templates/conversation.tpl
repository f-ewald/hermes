Conversation
============
Participants:
{{ range .Participants -}}
* {{ .ID }}: {{ .Number }}
{{- end }}

Messages:
{{ range .Messages -}}
{{printf "%05d" .SenderID}} ({{.Date}}): {{ .Text }}
{{ end }}