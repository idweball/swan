{{- if and (exist "vm.s1.nginx.domain.git.server_name") (exist "vm.s1.nginx.domain.git.port") (exist "vm.s1.nginx.domain.git.root") -}}
    server {
    {{- with $domain := getv "vm.s1.nginx.domain.git.server_name" }}
        server_name {{ $domain }};
    {{- end -}}
    {{- with $port := getv "vm.s1.nginx.domain.git.port" }}
        listen {{ $port }};
    {{- end -}}
    {{- with $root := getv "vm.s1.nginx.domain.git.root" }}
        root {{ $root }};
    {{- end }}
    index index.html index.php;
    }
{{- end -}}