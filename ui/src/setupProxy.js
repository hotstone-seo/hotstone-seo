const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    createProxyMiddleware(["/api/**", "/auth/**", "/p/**"], {
      target: "http://localhost:8089",
      changeOrigin: true,
    })
  );
};
