/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import e from"./menuItem.7c07b863.js";import t from"./asyncSubmenu.9b612524.js";import{Y as o,b as n,o as r,m as s,w as u,c as l,F as a,x as i,f,p as m,A as c}from"./index.42b7f490.js";const p={name:"AsideComponent"},d=Object.assign(p,{props:{routerInfo:{type:Object,default:()=>null},isCollapse:{default:function(){return!1},type:Boolean},theme:{default:function(){return{}},type:Object}},setup(p){const d=p,h=o((()=>d.routerInfo.children&&d.routerInfo.children.filter((e=>!e.hidden)).length?t:e));return(e,t)=>{const o=n("AsideComponent");return p.routerInfo.hidden?f("",!0):(r(),s(m(c(h)),{key:0,"is-collapse":p.isCollapse,theme:p.theme,"router-info":p.routerInfo},{default:u((()=>[p.routerInfo.children&&p.routerInfo.children.length?(r(!0),l(a,{key:0},i(p.routerInfo.children,(e=>(r(),s(o,{key:e.name,"is-collapse":!1,"router-info":e,theme:p.theme},null,8,["router-info","theme"])))),128)):f("",!0)])),_:1},8,["is-collapse","theme","router-info"]))}}});export{d as default};
