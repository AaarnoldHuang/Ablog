package views

const head  = `
<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>Ablog</title>
	<style>
	body {margin: 0;}
		ull {
			list-style-type: none;
			margin: 0;
			padding: 0;
			overflow: hidden;
			background-color: #4FC3F7;
			position: fixed;
			top: 0;
			width: 100%;
		}
		
		lii {
			float: left;
		}
		
		lii a {
			display: block;
			color: white;
			text-align: center;
			padding: 33px 30px;
			text-decoration: none;
		}
		
		lii a:hover{
			background-color: #01579b;
		}
	
	</style>
</head>

<body>
	<ull>
	<lii><div style = "margin-left:30px;color:#FFFFFF"><h1>ABlog</h1></div></lii>
	<lii><a class="active" href="home">主页</a></lii>
	<lii><a href="blogs">博客</a></lii>
	<lii><a href="about">关于</a></lii>
	</ull>
<div style="margin-left:30px;margin-top:100px;height:1500px;">

`
const tail  = `
	</div>
</body>
</html>
`

type webstyle struct {
	head string
	tail string
}


func GetHead() string {
	return head
}

func GetTail() string {
	return tail
}

func Getweb() (string, string) {
	return head, tail
}