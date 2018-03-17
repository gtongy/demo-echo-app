{{define "content"}}
    {{template "header" .}}
            <div>
                <input type="text"> {{index . "name"}}
            </div>
    {{template "footer" .}}
{{end}}