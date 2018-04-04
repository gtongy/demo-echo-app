{{define "content"}}
{{template "header" .}}
<div>
    {{ if .user }}
        <p>top page session has user</p>
        <div class="alert alert-success" role="alert">
            <div>ID: {{ .user.ID }}</div>
            <div>Email: {{ .user.Email }}</div>
            <div>Token: {{ .user.AccessToken }}</div>
        </div>
        <a href="/logout">logout</a>
    {{ else }}
        <p>top page</p>
        <a href="/login">login</a>
    {{end}}
</div>
{{template "footer" .}}
{{end}}