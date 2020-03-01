{{ range $i, $_ := .persons }}{{ if gt $i 0 }}
{{ end }}Hello, {{ .name }}!{{ end }}
