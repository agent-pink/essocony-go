{{ define "head" }}
<title>Hello World</title>
{{ end }}
{{ define "body" }}
<h1>Hello World</h1>
<div class="articles">
	{{ range . }}
	<div class="article">
		<h1 class="article-title">{{ .Title }}</h1>
		<h2 class="article-info">{{ .Author }}&mdash;{{ .Time }}</h2>
		<div class="article contents">
			{{ .HtmlContents }}
		</div>
	</div>
</div>
{{ end }}
{{ end }}
