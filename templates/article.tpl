{{ define "head" }}
<title>{{ .Title }}</title>
{{ end }}
{{ define "body" }}
<div class="masthead">
	<h1><a href="/">Essocony</a></h1>
</div>
<div class="articles">
	<div class="article">
		<h1 class="article-title"><a href="/{{ .Slug }}">{{ .Title }}</a></h1>
		<h2 class="article-info">{{ .Author }}&mdash;{{ .Time }}</h2>
		<div class="article contents">
			{{ .HtmlContents }}
		</div>
	</div>
</div>
{{ end }}
