{{define "form"}} 
{{template "header" .}}
<div class="container">
    <div class="row top-buffer">
        <div class="col-md-12">
            <div class="col-md-6 offset-md-3">
                    {{ if .error }}
                        <div class="alert alert-warning" role="alert">
                            {{ .error }}
                        </div>
                    {{end}}
                <div class="modal-content">
                    {{ if .new }}
                    <form action="/user/create" method="post" class="form">
                    {{ else }}
                    <form action="/auth" method="post" class="form">
                    {{end}}
                        <input type="hidden" name="csrf" value="{{ .csrfToken }}">
                        <div class="modal-header">
                            {{ if .new }}
                            <h5 class="modal-title">Register</h5>
                            {{ else }}
                            <h5 class="modal-title">Login</h5>
                            {{end}}
                        </div>
                        <div class="modal-body">
                            <div class="form-group">
                                <label>Email address</label>
                                <input type="email" name="email" class="form-control" aria-describedby="emailHelp" placeholder="Enter email">
                            </div>
                            <div class="form-group">
                                <label>Password</label>
                                <input type="password" name="password" class="form-control" placeholder="Password">
                            </div>
                        </div>
                        <div class="modal-footer">
                            {{ if .new }}
                            <button type="submit" class="btn btn-primary">Register</button>
                            {{ else }}
                            <button type="submit" class="btn btn-primary">Login</button>
                            {{end}}
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>
{{template "footer" .}} {{end}}