{{define "form"}}
{{template "header" .}}
    {{ if .new }}
    <form action="/user/create" method="post">
    {{ else }}
    <form action="/auth" method="post">
    {{end}}
        <input type="text" name="email">
        <input type="password" name="password">
    </form>
{{template "footer" .}}
{{end}}