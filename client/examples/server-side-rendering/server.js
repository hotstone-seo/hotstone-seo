!function(e){var t={};function r(n){if(t[n])return t[n].exports;var o=t[n]={i:n,l:!1,exports:{}};return e[n].call(o.exports,o,o.exports,r),o.l=!0,o.exports}r.m=e,r.c=t,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)r.d(n,o,function(t){return e[t]}.bind(null,o));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="",r(r.s=18)}([function(e,t){e.exports=require("react")},function(e,t){e.exports=require("react-helmet")},function(e,t){e.exports=require("@babel/runtime/regenerator")},function(e,t){e.exports=require("express")},function(e,t){e.exports=require("core-js/modules/es6.object.to-string")},function(e,t){e.exports=require("@babel/runtime/helpers/asyncToGenerator")},function(e,t){e.exports=require("react-dom/server")},function(e,t){e.exports=require("hotstone-client")},function(e,t){e.exports=require("serialize-javascript")},function(e,t){e.exports=require("core-js/modules/es6.regexp.match")},function(e,t){e.exports=require("regenerator-runtime/runtime")},function(e,t){e.exports=require("core-js/modules/es6.string.link")},function(e,t){e.exports=require("core-js/modules/es6.regexp.to-string")},function(e,t){e.exports=require("core-js/modules/es6.date.to-string")},function(e,t){e.exports=require("core-js/modules/web.dom.iterable")},function(e,t){e.exports=require("core-js/modules/es6.array.iterator")},function(e,t){e.exports=require("core-js/modules/es6.object.keys")},function(e,t){e.exports=require("core-js/modules/es6.array.map")},function(e,t,r){"use strict";r.r(t);var n=r(2),o=r.n(n),u=(r(9),r(10),r(5)),a=r.n(u),c=(r(11),r(12),r(13),r(4),r(3)),i=r.n(c),l=r(0),s=r.n(l),p=r(6),d=r(1),f=r(7),m=r.n(f),b=r(8),x=r.n(b);r(14),r(15),r(16),r(17);function y(e){var t=e.data,r=(t.rule,t.tags),n=void 0===r?[]:r;return console.log("tags:",n),s.a.createElement("div",null,s.a.createElement(d.Helmet,null,function(e){return e.map((function(e){var t=e.type,r=e.attributes,n=e.value;return s.a.createElement(t,r,n)}))}(n)),s.a.createElement("h1",null,"Sample Application using HotStone"),s.a.createElement("table",null,s.a.createElement("thead",null,s.a.createElement("tr",null,s.a.createElement("th",null,"Type"),s.a.createElement("th",null,"Attributes"),s.a.createElement("th",null,"Value"))),s.a.createElement("tbody",null,function(e){return e.map((function(e,t){var r=e.type,n=e.attributes,o=e.value,u=Object.keys(n).map((function(e,t){return s.a.createElement("li",{key:t},e,": ",n[e])}));return s.a.createElement("tr",{key:t},s.a.createElement("td",null,r),s.a.createElement("td",null,s.a.createElement("ul",null,u)),s.a.createElement("td",null,o))}))}(n))))}var g=i()(),v=new m.a("http://localhost:8089"),h=function(e,t){var r=e.body,n=e.head;return"\n    <!DOCTYPE html>\n    <html ".concat(n.htmlAttributes.toString(),">\n      <head>\n        ").concat(n.title.toString(),"\n        ").concat(n.meta.toString(),"\n        ").concat(n.link.toString(),"\n      </head>\n      <body ").concat(n.bodyAttributes.toString(),'>\n        <div id="root">').concat(r,"</div>\n        <script>window.__INITIAL_DATA__ = ").concat(x()(t),'<\/script>\n      </body>\n      <script src="/public/bundle.js"><\/script>\n    </html>\n  ')};g.use(i.a.static("public")),g.get("*",(function(e,t,r){a()(o.a.mark((function n(){var u,a,c,i,l;return o.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return n.prev=0,n.next=3,v.match(e.path);case 3:return u=n.sent,n.next=6,v.tags(u,"en-US");case 6:a=n.sent,c={rule:u,tags:a},i=Object(p.renderToString)(s.a.createElement(y,{data:c})),l=d.Helmet.renderStatic(),t.send(h({body:i,head:l},c)),n.next=16;break;case 13:n.prev=13,n.t0=n.catch(0),r(n.t0);case 16:case"end":return n.stop()}}),n,null,[[0,13]])})))()})),g.listen(4e3)}]);