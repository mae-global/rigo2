package main

const html_template = `
{{define "docheader"}}
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>RenderMan Asset Browser</title>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=10">
		<meta name="viewport" content="user-scalable=yes">

		<style>
			body { font-size: 9px; }
			section,article { display: block;  }
			article {  clear:both; width:100%; }
			section {  float: left; width: 100%; margin: 10px 0; }
			section.navigation { margin-right:1px; margin: 1px 0; }
			section.navigation:last-of-type {margin-bottom: -8px; }
			nav { padding: 0; list-style: none; width: 100%; float:right; display: block;  }
			nav li { padding: 9px 15px; text-align:left;  background-color: #888; cursor: pointer;  }
			nav li:first-of-type { margin-left: 0; }
			nav li:last-of-type { margin-right: 0; }
	    nav li.selected { background-color: pink; }
			nav li.library { cursor: unset; background-color:red; }
			nav a { text-decoration: none; color: #000; margin:0 1px; float: left;  }
			nav a:visited { color: #000; }
			figure { padding: 10px; margin:0; width: 110px; float: left; }
			figure figcaption { text-align: center; margin-top: 5px; }
		</style>

	</head>
	<body>
	
{{end}}

{{define "docfooter"}}
</body>
</html>
{{end}}

{{define "index"}}
{{template "docheader" .}}
<section class="navigation">
	<nav>
		<a href="/library/RenderMan/"><li>RenderMan</li></a>
	</nav>
</section>
{{template "docfooter" .}}
{{end}}

{{define "library"}}
{{template "docheader" .}}
	<section class="navigation">
		<article>
			<nav>
				<a href="/library/RenderMan"><li class="library">RenderMan</li></a>
				<a href="/library/custom"><li>Custom</li></a>
			</nav>
		</article>
	</section>
	{{range $section := .Body}}
		<section class="navigation">
			<article>
			<nav>		
			{{range $part := $section.Navigation}}
		
					<a href="{{$part.Href}}"><li {{if $part.Selected}}class="selected"{{end}}>{{$part.Label}}</li></a>
			{{end}}
			</nav>
			</article>
		</section>
	 {{end}}
	 {{range $section := .Body}}
		{{if $section.Assets}}
		 <section>
		 <article class="swatches">
			{{range $part := $section.Assets}}
				<figure>
					<img src="{{$part.Href}}">
					<figcaption>{{$part.Label}}</figcaption>
				</figure>
			{{end}}
			</article>
		</section>
	{{end}}
	{{end}}
{{template "docfooter" .}}
{{end}}
`





