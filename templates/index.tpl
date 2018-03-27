{{define "content"}}
{{template "header" .}}
<div>
    {{ if .user }}
        <p>top page session has user</p>
        <div class="alert alert-success" role="alert">
            <div>ID: {{ .user.ID }}</div>
            <div>Email: {{ .user.Email }}</div>
            <div>Password: {{ .user.Password }}</div>
        </div>
    {{ else }}
        <p>top page</p>
    {{end}}
</div>
{{template "footer" .}}
{{end}}