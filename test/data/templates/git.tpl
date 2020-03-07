{{- if exist "vm.s1.site.git.welcome" -}}
    {{- with $msg := getv "vm.s1.site.git.welcome" -}}
        Welcome {{ $msg }}
    {{- end -}}
{{- else -}}
    Welcome!
{{- end -}}