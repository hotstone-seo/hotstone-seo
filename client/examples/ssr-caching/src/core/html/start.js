import {Helmet} from 'react-helmet';
import config from 'config';
import {normalizeCSS,fontsCSS} from './critical';
import getClientAssets from '../assets';

export default function start(scripts = []) {
  const helmet = Helmet.renderStatic();
  const assets = getClientAssets();
  const preloadedScript = scripts
    .map(script => `<link rel="preload" href="${assets[script].js}" as="script" crossorigin="anonymous" />`)
    .join('\t');

  return `<!DOCTYPE html>
<html ${helmet.htmlAttributes.toString()}>
  <head>
    ${helmet.base.toString()}
    ${helmet.title.toString()}
    <meta charset="UTF-8">
    <meta name="viewport" content="initial-scale=1, minimum-scale=1, maximum-scale=1, user-scalable=no, width=device-width">
    <link rel="manifest" href="/manifest.json">
    <meta name="mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-title" content="${config.app.title}">
    <meta name="theme-color" content="#005ACC">

    <style type="text/css">${normalizeCSS}</style>
    <style type="text/css">${fontsCSS}</style>
    
     ${assets.vendor && assets.vendor.css ? `<link rel="stylesheet" media="screen" href="${assets.vendor.css}" crossorigin="anonymous" />` : ''}
     ${assets.client && assets.client.css ? `<link rel="stylesheet" media="screen" href="${assets.client.css}" crossorigin="anonymous" />` : ''}
     ${preloadedScript}

    ${helmet.meta.toString()}
    ${helmet.link.toString()}
    ${helmet.style.toString()}
    
    <!--[if lt IE 9]>
    <script src="//cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.2/html5shiv.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.2/html5shiv-printshiv.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/es5-shim/3.4.0/es5-shim.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/es5-shim/3.4.0/es5-sham.js"></script>
    <![endif]-->
    
    <script type="text/javascript">(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],j=d.createElement(s),dl=l!=='dataLayer'?'&l='+l:'';j.async=true;j.src='https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);})(window,document,'script','dataLayer','${config.GTM_ID}');</script>
  </head>
  <body ${helmet.bodyAttributes.toString()}>
    <noscript>
      <div>Website memerlukan javascript untuk dapat ditampilkan.</div>
      <iframe src="https://www.googletagmanager.com/ns.html?id=${config.GTM_ID}" height="0" width="0" style="display: none; visibility: hidden;"></iframe>
    </noscript>
    
    <div id="modal-root" ><!--MODAL--></div>
    <div id="app">`;
}