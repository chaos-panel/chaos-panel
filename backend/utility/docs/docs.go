package docs

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	DOCS_HTML = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Docs</title>
			<style>
				body{font-family:Arial,sans-serif;background:#282c34;color:#fff;margin:0;padding:0;display:flex;flex-direction:column;height:100vh;overflow:hidden}
				header{background:#282c34;padding:10px;display:flex;align-items:center;justify-content:space-between;border-bottom:1px solid #333}
				.title{font-size:2.5em;color:#61dafb}
				.nav-menu{display:flex;gap:20px}
				.nav-button{background:#61dafb;color:#282c34;border:none;border-radius:5px;padding:10px 20px;font-size:1.2em;cursor:pointer;transition:background .3s}
				.nav-button.selected{background:#21a1f1}
				.nav-button:hover{background:#21a1f1}
				.timestamp{font-size:1em;color:#888}
				.content{flex:1;display:flex;overflow:hidden}
				iframe{border:none;width:100%;height:100%;display:none;background:#fff}
				iframe.active{display:block}
			</style>
		</head>
		<body>
			<header>
				<div class="title">Docs</div>
				<div class="nav-menu">
					<button class="nav-button selected" id="docsify-btn" onclick="showIframe('docsify-frame')">Docsify</button>
					<div stuyle="width:32px"></div>
					<button class="nav-button" id="swagger-btn" onclick="showIframe('swagger-frame')">Swagger</button>
					<button class="nav-button" id="openapi-btn" onclick="showIframe('openapi-frame')">OpenApi</button>
					<button class="nav-button" id="redoc-btn" onclick="showIframe('redoc-frame')">Redoc</button>
				</div>
				<div class="timestamp" id="timestamp"></div>
			</header>
			<div class="content">
				<iframe id="docsify-frame" src="/help/docsify-ui.html" class="active"></iframe>
				<iframe id="swagger-frame" src="/docs/swagger-ui.html"></iframe>
				<iframe id="openapi-frame" src="/docs/openapi-ui.html" ></iframe>
				<iframe id="redoc-frame" src="/docs/redoc-ui.html"></iframe>
			</div>
			<script>
				function showIframe(id) {
					document.querySelectorAll('iframe').forEach(function(iframe) {
						iframe.classList.remove('active');
					});
					document.getElementById(id).classList.add('active');
					document.querySelectorAll('.nav-button').forEach(function(button) {
						button.classList.remove('selected');
					});
					document.getElementById(id.split('-')[0] + '-btn').classList.add('selected');
				}
				function updateTime() {
					var now = new Date();
					var formattedTime = now.toLocaleString();
					document.getElementById('timestamp').textContent = formattedTime;
				}
				setInterval(updateTime, 1000); // Update every second
				updateTime(); // Initial call to set time immediately
			</script>
		</body>
		</html>
		`

	SwaggerUI = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1" />
			<meta name="description" content="SwaggerUI"/>
			<title>SwaggerUI</title>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.10.5/swagger-ui.min.css" />
		</head>
		<body>
			<div id="swagger-ui"></div>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.10.5/swagger-ui-bundle.js" crossorigin></script>
			<script>
				window.onload = () => {
					window.ui = SwaggerUIBundle({
						url: '/api.json',
						dom_id: '#swagger-ui',
					});
				};
			</script>
		</body>
		</html>
		`
	OpenApiUI = `
		<!doctype html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<title>OpenAPI UI</title>
		</head>
		<body>
			<div id="openapi-ui-container" spec-url="/api.json" theme="light"></div>
			<script src="https://cdn.jsdelivr.net/npm/openapi-ui-dist@latest/lib/openapi-ui.umd.js"></script>
		</body>
		</html>
		`

	RedocUI = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>API Reference</title>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<style>
				body {
					margin: 0;
					padding: 0;
				}
			</style>
		</head>
		<body>
			<redoc spec-url="/api.json" show-object-schema-examples="true"></redoc>
    		<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
		</body>
		</html>
	    `
)

func InitSwagger(ctx context.Context, s *ghttp.Server) {

	docsUi, _ := g.View().ParseContent(ctx, DOCS_HTML)

	s.SetSwaggerPath("/docs")
	s.SetSwaggerUITemplate(docsUi)

	s.Group("/docs", func(group *ghttp.RouterGroup) {
		group.GET("/swagger-ui.html", func(r *ghttp.Request) {
			r.Response.Write(SwaggerUI)
		})
		group.GET("/openapi-ui.html", func(r *ghttp.Request) {
			r.Response.Write(OpenApiUI)
		})
		group.GET("/redoc-ui.html", func(r *ghttp.Request) {
			r.Response.Write(RedocUI)
		})
	})

}
