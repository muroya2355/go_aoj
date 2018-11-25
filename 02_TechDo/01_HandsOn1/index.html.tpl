<html>
<body>
	<h1>Hello, golang!</h1>
	{{ range $name, $profile := . }}
		<div><a href="Profile/{{$profile.Name}}">{{ $profile.Name }}</a></div>
	{{ end }}
</body>
</html>