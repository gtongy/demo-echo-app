{{define "content"}}
{{template "header" .}}
<div>
    <form action="/auth" method="post">
        <input type="text" name="user_name">
        <input type="text" name="email">
        <input type="password" name="password">
    </form>
</div>
{{template "footer" .}}
{{end}}