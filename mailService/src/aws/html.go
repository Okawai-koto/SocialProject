package aws

func getHtmlUpside(verify string) string {
	HtmlContent :=
		`
<html>
	<body> 
		<h1>Bu bir test e-postasıdır.</h1>
		<h1>Okawaii KOTTO!!</h1>
		<p>Merhaba, bu e-posta bir test e-postasıdır.</p>
		<br>
	</body>
</html>
` + verify

	return HtmlContent
}
