{{ define "head" }}
<title>Hello World</title>
{{ end }}
{{ define "body" }}
<div class="masthead">
	<h1><a href="/">Hello World</a></h1>
</div>
<div class="articles">
	{{ range . }}
	<div class="article">
		<h1 class="article-title"><a href="/{{ .Slug }}">{{ .Title }}</a></h1>
		<h2 class="article-info">{{ .Author }}&mdash;{{ .Time }}</h2>
		<div class="article contents">
			{{ .HtmlContents }}
		</div>
	</div>
</div>
{{ end }}
{{ end }}
