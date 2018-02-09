{{define "navbar"}}
<a class="navbar-brand" href="/" style="color:blue">我的博客</a>
<div>
	<ul class="nav navbar-nav" >
		{{if .IsHome}}
		<li class="active"><a href="/" style="color:red" >首页</a></li>	
		{{else}}
		<li><a href="/">首页</a></li>		
		{{end}}
		{{if .IsCategory}}
		<li class="active"><a href="/category" style="color:green">分类</a></li>
		{{else}}
			<li><a href="/category">分类</a></li>		
		{{end}}
		{{if .IsTopic}}
		<li class="active"><a href="/topic" style="color:blue">文章</a></li>
		{{else}}
			<li><a href="/topic">文章</a></li>		
		{{end}}	</ul>
</div>
<div class="pull-right">
	<ul class="nav navbar-nav">
		<a href="/"><img src="/static/img/avatar.jpg" class="img-circle" width="auto" height="45px"> </a>
		{{if .IsLogin}}
		<li><a href="/login?exit=true" style="color:green">退出登录</a></li>
		{{else}}
		<li><a href="/login" style="color:red">管理员登录</a></li>
		{{end}}
	</ul>
</div>
<div class="pull-right">
</div>
</div>

{{end}}