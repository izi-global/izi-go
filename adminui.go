// Copyright 2018 IZI Global. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package izigo

var indexTpl = `
{{define "content"}}
<h1>IZIGo Admin Dashboard</h1>
<p>
For detail usage please check our document:
</p>
<p>
<a target="_blank" href="http://go.izi.asia/docs/module/toolbox.md">Toolbox</a>
</p>
<p>
<a target="_blank" href="http://go.izi.asia/docs/advantage/monitor.md">Live Monitor</a>
</p>
{{.Content}}
{{end}}`

var profillingTpl = `
{{define "content"}}
<h1>{{.Title}}</h1>
<pre id="content">
<div>{{.Content}}</div>
</pre>
{{end}}`

var defaultScriptsTpl = ``

var gcAjaxTpl = `
{{define "scripts"}}
<script type="text/javascript">
	var app = app || {};
(function() {
	app.$el = $('#content');
	app.getGc = function() {
		var that = this;
		$.ajax("/prof?command=gc%20summary&format=json").done(function(data) {
			that.$el.append($('<p>' + data.Content + '</p>'));
		});
	};
	$(document).ready(function() {
		setInterval(function() {
			app.getGc();
		}, 3000);
	});
})();
</script>
{{end}}
`

var qpsTpl = `{{define "content"}}
<h1>Requests statistics</h1>
<table class="table table-striped table-hover ">
	<thead>
	<tr>
	{{range .Content.Fields}}
		<th>
		{{.}}
		</th>
	{{end}}
	</tr>
	</thead>

	<tbody>
	{{range $i, $elem := .Content.Data}}

	<tr>
	    <td>{{index $elem 0}}</td>
	    <td>{{index $elem 1}}</td>
	    <td>{{index $elem 2}}</td>
	    <td data-order="{{index $elem 3}}">{{index $elem 4}}</td>
	    <td data-order="{{index $elem 5}}">{{index $elem 6}}</td>
	    <td data-order="{{index $elem 7}}">{{index $elem 8}}</td>
	    <td data-order="{{index $elem 9}}">{{index $elem 10}}</td>
	</tr>
	{{end}}
	</tbody>

</table>
{{end}}`

var configTpl = `
{{define "content"}}
<h1>Configurations</h1>
<pre>
{{range $index, $elem := .Content}}
{{$index}}={{$elem}}
{{end}}
</pre>
{{end}}
`

var routerAndFilterTpl = `{{define "content"}}


<h1>{{.Title}}</h1>

{{range .Content.Methods}}

<div class="panel panel-default">
<div class="panel-heading lead success"><strong>{{.}}</strong></div>
<div class="panel-body">
<table class="table table-striped table-hover ">
	<thead>
	<tr>
	{{range $.Content.Fields}}
		<th>
		{{.}}
		</th>
	{{end}}
	</tr>
	</thead>

	<tbody>
	{{$slice := index $.Content.Data .}}
	{{range $i, $elem := $slice}}

	<tr>
		{{range $elem}}
			<td>
			{{.}}
			</td>
		{{end}}
	</tr>

	{{end}}
	</tbody>

</table>
</div>
</div>
{{end}}


{{end}}`

var tasksTpl = `{{define "content"}}

<h1>{{.Title}}</h1>

{{if .Message }}
{{ $messageType := index .Message 0}}
<p class="message
{{if eq "error" $messageType}}
bg-danger
{{else if eq "success" $messageType}}
bg-success
{{else}}
bg-warning
{{end}}
">
{{index .Message 1}}
</p>
{{end}}


<table class="table table-striped table-hover ">
<thead>
<tr>
{{range .Content.Fields}}
<th>
{{.}}
</th>
{{end}}
</tr>
</thead>

<tbody>
{{range $i, $slice := .Content.Data}}
<tr>
	{{range $slice}}
	<td>
	{{.}}
	</td>
	{{end}}
	<td>
	<a class="btn btn-primary btn-sm" href="/task?taskname={{index $slice 0}}">Run</a>
	</td>
</tr>
{{end}}
</tbody>
</table>

{{end}}`

var healthCheckTpl = `
{{define "content"}}

<h1>{{.Title}}</h1>
<table class="table table-striped table-hover ">
<thead>
<tr>
{{range .Content.Fields}}
	<th>
	{{.}}
	</th>
{{end}}
</tr>
</thead>
<tbody>
{{range $i, $slice := .Content.Data}}
	{{ $header := index $slice 0}}
	{{ if eq "success" $header}}
	<tr class="success">
	{{else if eq "error" $header}}
	<tr class="danger">
	{{else}}
	<tr>
	{{end}}
		{{range $j, $elem := $slice}}
		{{if ne $j 0}}
		<td>
		{{$elem}}
		</td>
		{{end}}
		{{end}}
		<td>
		{{$header}}
		</td>
	</tr>
{{end}}

</tbody>
</table>
{{end}}`

// The base dashboardTpl
var dashboardTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
<!-- Meta, title, CSS, favicons, etc. -->
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">

<title>

Welcome to IZIGo Admin Dashboard

</title>

<link href="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">
<link href="//cdn.datatables.net/plug-ins/725b2a2115b/integration/bootstrap/3/dataTables.bootstrap.css" rel="stylesheet">

<style type="text/css">
ul.nav li.dropdown:hover > ul.dropdown-menu {
	display: block;    
}
#logo {
	width: 102px;
	height: 32px;
	margin-top: 5px;
}
.message {
	padding: 15px;
}
</style>

</head>
<body>

<header class="navbar navbar-default navbar-static-top bs-docs-nav" id="top" role="banner">
<div class="container">
<div class="navbar-header">
<button class="navbar-toggle" type="button" data-toggle="collapse" data-target=".bs-navbar-collapse">
<span class="sr-only">Toggle navigation</span>
<span class="icon-bar"></span>
<span class="icon-bar"></span>
<span class="icon-bar"></span>
</button>

<a href="/">
<img id="logo" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGwAAAAhCAYAAAA8oE5/AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAA6cAAAOnAEHlFPdAAAAB3RJTUUH4ggUBCI6HuSsbgAAGitJREFUaN61W2eYFVW2XftUurHv7W6goSWIBMGAAUQEgdEZRH0KiGHUpyNGxICYZsY8D4Rx1BHEBCpilscoCJgQFVQUFRUElYxIQ9P55nvrVtU5+/240tJ0qw349vedP3tV7VV11jl10i6KP/bnIVCqMxlmkJVkgBjSiUG3PoGT2xU44+8wux2P1EsTIOu29WE33wvMCvtrBJBmMnRjFeeSW7V23VE05gkAQOKZKyOcjQ8GiXYkNAMgsHTyYFUhisqWcy6Zj1z94m9SJJ66FJyuh1beW8iaLR2hvENBohyABWbFysuTbgFgsOeASDNBJADOQejV0IxNIlhcIRsqpPBH0BrOAu8YqGQNyAr5IJ3uINELJEpBRFAyD91UZPgEO1kF1yEQmSBSAOoh9E1k+jepVF1eRNohcsXsFjl0eM5RMHwLQFqG7aQnwqUWS7cznNwZ5A+/nJg+OgMAKlUXgtD7ioD/YwjN3W/BAHA+42fPGUyBaM759t1djf5MrC9YuWT53mInZ5NuEumWn/OZQSpZfaJW2mUZAPlLce0v5sL+fC5IMwRrRi9ZtfFIAArM28DyW+hWXASjeW/7Nyq7/HNwHvANOhxG9wEaZ+MWlCqGkgdByX4qUTWADP9a0q3v409cyNFxL/+yUE+OAdtJgNkkK3QsWPUEiQYI7TvSjF3kC2fCY2a4tVdHYfQcBGftYoQveJDcH77S2bVDAMrBshvbqaPJsNYB9E1ixkVe0dgXQERNBWPPqSTDt51MP8hOgXQfZHZnCvnURKRqh5i9hl4LvJxUmViIND3OzJvJX3QAajEABpxcH5VumK2173k9sGUTAMDLmyzlRvJHqsjwAYYFSC+mUnU7IN3ZKhN7AMCslsKmX5+I7HuPQ0TK2qpkzVAALsAf612O2eVu/YJJMyBCJdA79YH/j9ei5J6jCiI/dC9k9SbpVa7LcjaeZeXtNI854wt39dsdWNoDFMsO6TmvvZecPRZFl85sxhv7939B1v8IESptz052GJSqIcP3hoxXNohgMeALQ2vfE/aKl9BuZgLAWwCA7AXEorSzq+q2xdjJxVSi6jtR0qkUyhvAmVg3soJLEg+fFWvewwCKXvPKXg9xBqlsLALpHaWS1REASbg2oIUYmo7IlS1319Za/ZShIFYuu/Zw6OYoAA/slhMEsXf8hntPFCydQyjPA1sSLD3/H8i+/SDMw/7YA55zIggr4pNeXx+9YwRUqg7FNxcqKf/tEjhr3oK7eQXij58PSA/2ipdAgSi0docgNPIuAEDyxUMAVruEv2ihSlaPC474UzH5wvPQQu8m3QA8vQtYnQTSl8uazZv1Tn1g9hiE0Dn3Fp7vtbvhVm1AYubFYOmCrBBkrBJ6xyPgP3cKACCzeBrczZ/WZ9+b92bg5FFHsJcfBdO3MP7IOfXR619tIlhL3QAAMcASqjBciUiZy54TIjMg6u86VhWu+Q0TOlTVRrSdmWjiNjoeCVmzJcJODiBq0zqZSbU0djIz4g8Mh3nEsEMhvf4AFrFSddE7RyI6YQEAIP74+VCJaqReuB6hc6eI/JfzQlDKD1YEkE26lUs9f1e+YeIJoHAbsJ1E9MaFiD92vmQ3n4SXf9TbsXYHgBV7csenjwZAB0HTB4HVm5BurXXc2QiNvBvMjMSMbZC1W7Hx7InoPntsEKyCYBZgZWvlvVLZtx+Q8WmjoJV1R3D4hEJjeXYcVLzyWygl4eYvoGB0buzB02qKb3n71wRrobp84Tin6xtk3bYJAJfjNxVjkBApo89pM4A5NXsiQ/7yKD6YPPhQsAKYJVpr1NyVnHEhyF/UHso7AUJbBFb1Zu+TETztxkKlThsFZ9UbMLoP6CwiZadl335wCJTsDiDKzAJARsV21lhHD1xLpv8TIvF5zRWzdzbceyJEsBicS3jKybaDkz1hT8EST14CKBngfOZkMv0fgkRt6Ox7oZf3KjSi+4dB73y0oeK7Bh1874lneRVrjgFzO4A1gNJy1/rNFGqzBJq+wN28ojrzzlQET70RRWOeQOLJS6CSNetA1I1jlbeSFfyf1Jxb0+HzH2i9YOw5UmvX7R2VrvdD6HcDKtxiDe5WCyygmc+TYaX2BBJPXAhlpw9WiaoRYABK7mi1YHuZs/kz5JbOMOHYJ4PER5zP1GsH920UK/bQGYBhBfWDj71UZePjwOowMAMkJISoJdLiLB2XhX4QpOzNOW8M2enq0oknvAGi578e88S3fSYPMcEMMvzt9+T2dq2HFi3vD6If2XO2cy4JvbxXoZFMPRMw/aXu1s/vZCc3BkpGQQSQaCCh17PyFEE7kp3sABV3LqBA5H8CwycsS72yA+EL/o3IVc8hPm0kKFj8rqz7cRRnE1dYnY95GAC3WrDIlbORnne3IjOwkAz/cCgv8oudjIQGof2RhD6L7VRut9tetQjOd0t0bqi4DtLrDU2rJSv42f4KlnntTlAg0gfghmHj529dumgyQmfeURBr6pkQvqJSWbdtCrv2pWBlQOhVZFoLyfC9QVZoPQUicVn7gy2K2pnsZIPwnE6svAHwnJPhuc/1uXfQEnhOLzCDpWvv5k3Ougwql4rAczpCM18T/iJEbnqjgM2+GirbUKTqK+5nN38pwES6tZp8oWehm0tFqLRaZWI54Y9YKhsrYcfuD+mdk3j03BhYfbObIzphAeLTRjpk+F7iXGKaveLFxckXx68T+1JBodETASgX0vmcpfMuS7eF4r1L/kglmYEPjF5D1kfHz/u5gufdA7lz3VB2coUXMQP/q7XvsXZ/BaNwGwtK9QJpn7838y+NYqXfuh9k+CxZt+0OdnNXAmyQ4XtHhIpHBUbcOY5dexHAmyBELUApkKgHaDtL7xNV9+O/9c5Hn02+0DWACDOrEyAIAG9pJNZMADgCurkd0stF9pjyG4f0g0rWjmHP+QsITFbwZVHUbpT7w8qHIb01IFENEkkIUQvmDWwnX6BAdCKI+pEVCu/5ftEJCyDCbb8Ac5xzqYv1sh7U6h6224pvfvsXscTMv4A9uxjSOwaaMd/bvqaxCyafvw4qVdtG1W+/E0qWkG6tE8GSaXDz+7WmyyyeBnfDR11BFA+dPTGmH3R4I+asfgsAn86ufRWYiXzh17SSTtfLum278itfRcmdy38xbnbpTLibVtgsvc/9wyestD98+gWw6k+ByOLGi5gEQAeTbi4BaU3uz69Z3BFOdixY6WT63xQlnSZwNl4bPPteBIdd34zP/noBch/MrCF/aCt7zpGJWVd8Grn86UZ8+djnMyfce+JH7OXPzK9dPH2fetivWWrePSArSFByKEis5mwiERn7PABAKQm9Ux9SiarL2csPBgmbfOH74w/M39LS2qY1pnc8AgD1ALAhs+ifTTARaReGa18DVkEy/CtFpP1NKtOwK3zRdOxZGS1Z4KSxiFz1LEpuXwbn23cVu/ZnsnbrdM7Eqxr1shNRMIOscH3k2jk/i73sabCdHMrSOxSaUU/BkimcjdcW3/J2i2IBgO/YkbCOPAUkjHWA6kJWsIkmp/y4CtCMlWA+mPOZ/r+bYLJiDVSqtheYDTKstUavoY1Y8uGzkP9qfl92steDWSPTv0BEO8xt8/Bd+81nr3jFBLgIQqvcc+tI5bPgdENfVvIECC1HgcgD3g8rt0dvfgtG1777xBG5dCZKbluKNvdvQvFf3230s5LtQIhxPtFkhquX9yJ2cgPBSiNNf0c/6LCVvsGX/CZPYPgEQNProZTOmYbAnpjR5RiQbv4AwIN0huzzJ7ElSz47DpxPB9m1B0CYb7LnyNCZtxWwF8aDs7GQqtl6O6R3EDSzgoLF/1KZ+uzuHrg/xvl0EQCP/JHsnv4vrQC6Tx58IpQMkuH7UETL3/EPvQJEhPjUEVCZenx35yd46Dfinwzgwn8cBy1aDlYmim/6zx7k3Aag+sBpfwPw8/hlr3jJgnR6gIihGYtVJub6jz+/Ve/jG3Sxm3v/8Szns2EA6T0xMqwGEFJgHPu7CKZ3Ox7O6kWDANro7fimps196xqxugnTEb3lpD+za58OIo+swMOBU29ebfYaCtxE+0+qZBGY05xLNpmqHv7WA1r+y3mHAwRoxtvpObNSkSt+2hwhaCLS4fg+My7q8GwrKNjJOiBaTcKryH3yAvyDLt6NhED0o9HlqKbXZ+IhMHcAiRTp1neRy55q9evIyvWA0G0QhfbGKFRqU7I2DeV1PWDBEo9fAOebN7uA0ZZMa5lv4EUACrO1+PTRgJI9ZaLqVrCyyPS/J4raPeOsfYet3n84MGLpBaAbdvS6uU3c7qZPDVZeGYhc0o2v2jyyFOh9UgE0A+Vw7C4qW/0+S+dXTxxI6CBfqBjMp8D0v5n+z22NYxiUMlk6mb3vUek6k5UMgkQD+cI12AcLDLsO8WmjctANX7NnIc0BqywDBx/QGJae/w+AhAmlBkOIj8HshEYUxMouexIi3NZUmYZb4TmHQtPrKRCdolJ1sfB59x2YWADI8vvgOXYzvy9sABQEURKasdPaLRYAuHaEWe5iJ1vDdqruV0s+U8esNqlsnFTDjmf1Tn067Q7Dbg4k9GY7NKKkc4A0PQigASRS2EdjVpLt5re5FWs89pwskQgeUA/LfTgLRrcBx4NQHbl27nYAwE9rkszrE6F36HU6O7nzQcRk+J80uvb7yOg+EPj7+wcsGDMTe06zSlOZBgVWHkjYZAYze90EANDKugO0V1slAgkdRZc92ehKz5mC/HdzcyAxnGTgVABP7Xl9M+5UjcPSc0g3PVFc3vott58fEFRY4zUxo2s/0/txVZCVR/vdw+KPnA2jx8AyELqTGfgk9dx1jWc3yWeuhHFwv3K2U7dByRDp1leiqN1jKl4lrT6nHrBYhdqRDvnD/r3dorjcg9AyRAJkWE1BzUiSZnRgzykBqwiYfyoqAuYgqOleW+j829G4/cYobdRKt4hdu3ndMTsgckCCyPDt0wBtfzkPZPj8MH1es7C5pAHmACCwX4KlF0wC6T4N0h0K5hWcz2SLxjzWSCyK2mkqG7uOPfc4CC1N/vA/86vf2Lln6/0dBMvCc3zJZ69u4ja6DXCIxHZmqatcwtgTI92oBKsqlaz+l6zauFRWrf+pbFgqd62f6VWuC/4SHe/ungCg6S6ZgeaNJdwmT6Sl2XNI1m7dp9fZ0m80wOyH4vTemMolgqxkmIi8/RLM3fAR2MsfBSJbRMo2+Ab/vRHLvf8E3K0rT+R89kqAiQz/K6K0yxvhi6b/fmIBIN1KgUSQfOEmLVnVbWMIfRUYfnhuyV617oF5KZiXsOcY7OYDjUW67eG5RivZsxCi2WxORDpkQFQJIAyQr3WxCtbxqTEaWPlIN5LNQM8pA7gY4IZ9Fizx5CWgYHEUrPpA6B9xLsm+Y/oDAJKzr4IItylmO3U7lNeGdHOjCJU+yKlaJ3DSVb+rYND0OFiZsmZrk49+aPREkBVYDqIse/meiRn/3YhFrnkF0QkLoB90+Gtk+v8A8GAAgwEMFuG2z0E3tVZxk4gDoiQ5e1wTd+jcKQ6E9j2AKDy71cfyyTm3gPOZIgAEw9ekh6UXTAJY9QJzCKCd+yRYZtEUkC9MUHIISKzhbDweGftCI271HU0yUXUpe87JIJGHFXwgcsP8jdEbXv99xQIgSjvnQcIGq7L0gklNsXCbdUT0HjxnqCgqazaWcD4jAdQDqN1dSDczrSIGAEItlFdMobZN3O7WLwDdXA7Ax/lsx9yKl1oVTtX+AEjvEBBVcS7ZZLJidBtA7Nr9wUzQ9PX7JJi75TOoRNWhYPaR4fvG6HZ8Ixabeiay7z7ch/OZG8BKJ8P3phYtn5N58/7fXSwA4HQ9A9gKcM/8yteaYm7egeF7BEAvr/L7HrEHm050IlfORsmdy1E6aTVKJ61C6aRVAHOrJwkktDooacnq9U22kcxux0P4I58Q0Q/s5o/LvjP1N2PFHz4L5AtbENrhEPpas+fgRszZvAK5D59u81Pag4LQP2u1YKk5t4KscABEJ0AzP2QvLwvHLYUEGBFqE2A7eRuk2xmaUUnB4vtUNp4OnfG3/xfBisbMAIS+CYwOWvueocTMixux6PWvwjx0yEoYvufZsUeDND3xxIVIL5zcYqzM2//eJ27z6LPy0PQEgC7JZ5p+6qM3zK+B4ZsBln8Q0Q7BvSdFe1pi5sVgJwvOJYcAqAj86drawPAbfn6u1+6CSlQPZun1Amm1ZAU+afU6zBp0CbLz7x4AElvc75dUt328vhHTOx4OZ/2yc9mxR4HII8M3lbPxlVpZDySeurTlgF4erCSi1/7vfovG2ViGrNAPkE4/mapdphKVEJHyQvhdG1iESp7lbOJ00oxRYF5ed90dVbGpZ4J0X2Hxa/gB6cD+9EVopZ31XzlBb/roO9YASn0LIfqRP7wee5zkJp+7BipZ/YJK1R8EL38qSLwef3iUhNAhwm3BygOn6gAlAen5yAoOgdBC5Au/kf96YSNHau7foeq3h2TNlsvByiLd/FQUd/y+1T0su3BSKViVQ2hf+E+7pQlmf/Jie0hvLDTdpUB0m97xyHq9a78RFIj8crFCp4BEm9hD/wWvvmK/BKNwW5A//BWATiJY0ikxc0wjFr1uLuA5ntH7pEUQYj072ZOit4weTbo1AEQ9SeidQNQVQj9WK+08AoBOvmC2NbzhcyZCRMq2A9Bl3Y8Hxx/7eYO36JLHwU4uK0Il/yIzqFSi6jxo5gkg6spevhyeUw4husGwhrKXPxcgh8zAIrh5J3zBgwCAzDtTseW8+6ASVeez9P4IEg4M3yuycl2u9WOYkt0A2gZme/dOfGPFWYGUVtrpBr2sxzCt7SEXQdO/B6tqqF8ozNUsXQknewoZfl9i+uj9Eiw67mWw5+Qg9GVgeTKZgeLY1DORfa+wJozeuAjetq+Ypfste85cKLUSygOk2xWG2RfKOwzgMIT2HZmBuXLn9y0Ixi3mQRR2WWg5wENI0wOJWVc0YsU3vwVILxMZP28+gGVgpUF6R7CdHsRObiCU6gVWOTL8b7GdXAbpupFxhQlK+tU7kPvwaXR58NTjlZ2+HawsMnwfau0OWRI8777WJ+FAumGw2mX2Gwmg6YYrK5kBYyXAhd0O8VuzY8LGK2ej25ShT0C6f9Dadr0JWJnGflhhDEhViFDpF+zaZ5JmLLY/e6U6/vgFiF7zCnanKMQfPVfCzVewUhUAoGo2Q0Q6AIYFIg3QLYQvewp4aFhj7MTTl0FWfm+y1/xQPHL50/gSqOj+6Lkb2HPOEMHihfFpI23z8GEIDLsO0RsXItX+VoCxC8rdxV7+p8+gg0JSjw8wBfQufRH+831gZiSfGoPssidhdOl7LGdij0O6XaHp1eQLT5bVW+JWz0EFwWIPjUDxTQt/vWKUBJFQ3g9fN2/p18zBvhozI/bPk4rYc0ZwNv4ogDX7I1jxhIVIvTgB7rav1pHplyzdsWT637eOOn1Fw5RKpZV1h1bSCaGz/tHC3U3TSVjNQvKF66FilVCpWhRdPguxyYMPA9wWR7eeL46Hu2H551rbg3VVt+1CCH1R7uPZtYmZF0PvehyCp4xvcr1bsRYggtHxiEZf7tOXkHx2HOL/+hP0Q/obRqejzlCZhsmQbm+QSJEZuNs8cvjH5pHDgdsPgw4ijXMxNEweAjIscD4DESqRKlnT5BlJt7IglLkbl6+vv3vfTm6bGQGfEqH3pIEWiPwUiLQr+IUAMTVMGlRogWCoXBIgwfiVGUH4ommIP3IO5K71G0VJp1fZy/81t+yps0nTXyHdWmt/8oLdMPEEkC8EaCbYTkJEOoDtFMgMgJ0M2E6j4Z6+SM5ajeLbzo6QbvaOTR48gt3cGIDBSsb35i26aDpi9w/j8NjnlienjjSZeQoFoosB+iA1a3xDw6SBIH8RCj9eALmlM8HZeCH9DgzOpZB+9Xb4+p/ng24c42786DJ28+dBySIIvYFM3z0i1OYZWbVBGacWUvd00ow2sEJ/hpI2WQGwkrqs3XYYWPUAoTFFjQzrO/bckSIQvYo9p/CXywFY78mDw+zY/QF22MkVjkl0aydJZzBM/9Ew/A7AIKXC8PKngtn4tdTV6PWvIvfZHOSWPPK9iJT9TWVi49nJPeX9uKpKRDt8CKF9DVAFwDEA9k9JrAArP0BhaHo78kd6Rq4eeLSsr+gLJQ8DqwhAIN1YR4UFcTMr/usSJMv+Cnfb1x8Y3Y532MndLmu3TDCPHLycDOtzkNgMcAMzZwucCmDWQRQlw+qotT3kWHfTpyez5wyA8kpABNLNlWSFJln9zn4zv2qBKhrz8zpTJyvwvMol7mI7NR6ZmFFIX2YCA6SZq8gXygMASy8pQiWveqmaieza48HKxIEaM6AZWyDdHQCgdzx8rVe5PsfJ2pc4mzgGYFFIbWaASMGwGn4tnH/A+XB/XI3MvLuro7cuvjMxbcSrnM9cDM+9mN38rZzPegAnwZyR9dsZrAASIbAKgbkIYH8hS50YJOpIM5ZCN94nX9H8osuf+R63lrXIW3TJE0g8fRm8ijXLtXbdLlHxXedBuv+tcvkrf5q0JAGkOF3PUJJZSROgYrCK7sGZhW5+QWbgPyJU+opXsXant301im9d3IRLh9BTZPimQqkdUN6hZAbasnRzYN5BVvBd7aDD64BFiF43F8nnrslpJR2nqPiubzmfLT+Q38QAAEQQvvAGrbz3DmAxOBvnyLiXN8UfOGU8hHYaiDqRboXZtetJNzdRIPLqb4U0uhwNAEi2uYXZtVfNu+PjVedMHz2Nc8kB7DknQrpHAdwZzMWslAArJs3wANoBISoh9A2km6tg+L4SoZIN+dVvJY0uxyCzcNKv8kaueAbZj56Bs/rNWhXb+ZjRbcAcla7rz05uKJTXD0odwtItAeAjoRusvBTpZh2ItkEzviahL6NAZIW7/sNq6jEQ/iGXInjG35vx/B/+3CojyEFe8QAAABZ0RVh0Q3JlYXRpb24gVGltZQAwOC8yMC8xOEAyQxQAAAAldEVYdGRhdGU6Y3JlYXRlADIwMTgtMDgtMjBUMDQ6MzQ6NTgtMDQ6MDBsI2wpAAAAJXRFWHRkYXRlOm1vZGlmeQAyMDE4LTA4LTIwVDA0OjM0OjU4LTA0OjAwHX7UlQAAABx0RVh0U29mdHdhcmUAQWRvYmUgRmlyZXdvcmtzIENTNui8sowAAAAASUVORK5CYII="/>
</a>

</div>
<nav class="collapse navbar-collapse bs-navbar-collapse" role="navigation">
<ul class="nav navbar-nav">
<li>
<a href="/qps">
Requests statistics
</a>
</li>
<li>

<li class="dropdown">
<a href="#" class="dropdown-toggle disabled" data-toggle="dropdown">Performance profiling<span class="caret"></span></a>
<ul class="dropdown-menu" role="menu">

<li><a href="/prof?command=lookup goroutine">lookup goroutine</a></li>
<li><a href="/prof?command=lookup heap">lookup heap</a></li>
<li><a href="/prof?command=lookup threadcreate">lookup threadcreate</a></li>
<li><a href="/prof?command=lookup block">lookup block</a></li>
<li><a href="/prof?command=get cpuprof">get cpuprof</a></li>
<li><a href="/prof?command=get memprof">get memprof</a></li>
<li><a href="/prof?command=gc summary">gc summary</a></li>

</ul>
</li>

<li>
<a href="/healthcheck">
Healthcheck
</a>
</li>

<li>
<a href="/task" class="dropdown-toggle disabled" data-toggle="dropdown">Tasks</a>
</li>

<li class="dropdown">
<a href="#" class="dropdown-toggle disabled" data-toggle="dropdown">Config Status<span class="caret"></span></a>
<ul class="dropdown-menu" role="menu">
<li><a href="/listconf?command=conf">Configs</a></li>
<li><a href="/listconf?command=router">Routers</a></li>
<li><a href="/listconf?command=filter">Filters</a></li>
</ul>
</li>
</ul>
</nav>
</div>
</header>

<div class="container">
{{template "content" .}}
</div>

<script src="//code.jquery.com/jquery-1.11.1.min.js"></script>
<script src="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
<script src="//cdn.datatables.net/1.10.2/js/jquery.dataTables.min.js"></script>
<script src="//cdn.datatables.net/plug-ins/725b2a2115b/integration/bootstrap/3/dataTables.bootstrap.js"></script>

<script type="text/javascript">
$(document).ready(function() {
    $('.table').dataTable();
});
</script>
{{template "scripts" .}}
</body>
</html>
`
