/*! 
 Build based on gin-vue-admin 
 Time : 1677739392000 */
System.register(["./index-legacy.3496f7e7.js","./index-legacy.f6c9b804.js","./menuItem-legacy.514f55f8.js","./asyncSubmenu-legacy.f7b8fde8.js"],(function(e,n){"use strict";var t,a,u,l,o,r,c,i,f,d,s,v,m,p,h,g,b,x,y,k,T,_=document.createElement("style");return _.textContent=".el-sub-menu__title:hover,.el-menu-item:hover{background:transparent}.el-scrollbar .el-scrollbar__view{height:100%}.menu-info .menu-contorl{line-height:52px;font-size:20px;display:table-cell;vertical-align:middle}\n",document.head.appendChild(_),{setters:[function(e){t=e.C,a=e.u,u=e.j,l=e.I,o=e.r,r=e.J,c=e.R,i=e.a0,f=e.b,d=e.o,s=e.c,v=e.e,m=e.w,p=e.T,h=e.F,g=e.x,b=e.z,x=e.m,y=e.f,k=e.H},function(e){T=e.default},function(){},function(){}],execute:function(){var n={name:"Aside"};e("default",Object.assign(n,{setup:function(e){var n=t(),_=a(),j=u(),w=l(),F=o({}),M=function(){switch(j.sideMode){case"#fff":F.value={background:"#fff",activeBackground:"#4D70FF",activeText:"#fff",normalText:"#333",hoverBackground:"rgba(64, 158, 255, 0.08)",hoverText:"#333"};break;case"#191a23":F.value={background:"#191a23",activeBackground:"#4D70FF",activeText:"#fff",normalText:"#fff",hoverBackground:"rgba(64, 158, 255, 0.08)",hoverText:"#fff"}}};M();var B=o("");r((function(){return n}),(function(){console.log(n.meta.activeName),B.value=n.meta.activeName||n.name}),{deep:!0}),r((function(){return j.sideMode}),(function(){M()}));var q=o(!1);B.value=n.meta.activeName||n.name,document.body.clientWidth<1e3&&(q.value=!q.value),i.on("collapse",(function(e){q.value=e})),c((function(){i.off("collapse")}));var C=function(e,t,a,u){var l,o,r={},c={};(null===(l=w.routeMap[e])||void 0===l?void 0:l.parameters)&&(null===(o=w.routeMap[e])||void 0===o||o.parameters.forEach((function(e){"query"===e.type?r[e.key]=e.value:c[e.key]=e.value}))),e!==n.name&&(e.indexOf("http://")>-1||e.indexOf("https://")>-1?window.open(e):_.push({name:e,query:r,params:c}))};return function(e,n){var t=f("el-menu"),a=f("el-scrollbar");return d(),s("div",{style:k({background:b(j).sideMode})},[v(a,{style:{height:"calc(100vh - 60px)"}},{default:m((function(){return[v(p,{duration:{enter:800,leave:100},mode:"out-in",name:"el-fade-in-linear"},{default:m((function(){return[v(t,{collapse:q.value,"collapse-transition":!1,"default-active":B.value,"background-color":F.value.background,"active-text-color":F.value.active,class:"el-menu-vertical","unique-opened":"",onSelect:C},{default:m((function(){return[(d(!0),s(h,null,g(b(w).asyncRouters[0].children,(function(e){return d(),s(h,null,[e.hidden?y("",!0):(d(),x(T,{key:e.name,"is-collapse":q.value,"router-info":e,theme:F.value},null,8,["is-collapse","router-info","theme"]))],64)})),256))]})),_:1},8,["collapse","default-active","background-color","active-text-color"])]})),_:1})]})),_:1})],4)}}}))}}}));