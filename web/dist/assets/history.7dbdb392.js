/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{G as e,u as a,r as s,j as l,Y as t,J as n,R as u,a0 as o,b as r,o as i,c as v,e as m,w as c,F as d,x as y,m as p,d as g,H as f,A as h,h as b,t as S,a1 as I,a2 as q,S as O,U as w}from"./index.42b7f490.js";const N={class:"router-history"},k=["tab"],x={name:"HistoryComponent"},C=Object.assign(x,{setup(x){const C=e(),J=a(),j=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),V=s([]),E=s(""),A=s(!1),P=l(),R=e=>e.name+JSON.stringify(e.query)+JSON.stringify(e.params),T=s(0),L=s(0),H=s(!1),U=s(!1),Y=s(""),_=t((()=>P.userInfo.authority.defaultRouter)),F=()=>{V.value=[{name:_.value,meta:{title:"首页"},query:{},params:{}}],J.push({name:_.value}),A.value=!1,sessionStorage.setItem("historys",JSON.stringify(V.value))},G=()=>{let e;const a=V.value.findIndex((a=>(j(a)===Y.value&&(e=a),j(a)===Y.value))),s=V.value.findIndex((e=>j(e)===E.value));V.value.splice(0,a),a>s&&J.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},X=()=>{let e;const a=V.value.findIndex((a=>(j(a)===Y.value&&(e=a),j(a)===Y.value))),s=V.value.findIndex((e=>j(e)===E.value));V.value.splice(a+1,V.value.length),a<s&&J.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},$=()=>{let e;V.value=V.value.filter((a=>(j(a)===Y.value&&(e=a),j(a)===Y.value))),J.push(e),sessionStorage.setItem("historys",JSON.stringify(V.value))},z=e=>{if(!V.value.some((a=>((e,a)=>{if(e.name!==a.name)return!1;if(Object.keys(e.query).length!==Object.keys(a.query).length||Object.keys(e.params).length!==Object.keys(a.params).length)return!1;for(const s in e.query)if(e.query[s]!==a.query[s])return!1;for(const s in e.params)if(e.params[s]!==a.params[s])return!1;return!0})(a,e)))){const a={};a.name=e.name,a.meta={...e.meta},delete a.meta.matched,a.query=e.query,a.params=e.params,V.value.push(a)}window.sessionStorage.setItem("activeValue",j(e))},B=s({});n((()=>V.value),(()=>{B.value={},V.value.forEach((e=>{B.value[j(e)]=e}))}));const D=e=>{const a=B.value[e];J.push({name:a.name,query:a.query,params:a.params})},K=e=>{const a=V.value.findIndex((a=>j(a)===e));j(C)===e&&(1===V.value.length?J.push({name:_.value}):a<V.value.length-1?J.push({name:V.value[a+1].name,query:V.value[a+1].query,params:V.value[a+1].params}):J.push({name:V.value[a-1].name,query:V.value[a-1].query,params:V.value[a-1].params})),V.value.splice(a,1)};n((()=>A.value),(()=>{A.value?document.body.addEventListener("click",(()=>{A.value=!1})):document.body.removeEventListener("click",(()=>{A.value=!1}))})),n((()=>C),((e,a)=>{"Login"!==e.name&&"Reload"!==e.name&&(V.value=V.value.filter((e=>!e.meta.closeTab)),z(e),sessionStorage.setItem("historys",JSON.stringify(V.value)),E.value=window.sessionStorage.getItem("activeValue"))}),{deep:!0}),n((()=>V.value),(()=>{sessionStorage.setItem("historys",JSON.stringify(V.value))}),{deep:!0});return(()=>{o.on("closeThisPage",(()=>{K(R(C))})),o.on("closeAllPage",(()=>{F()})),o.on("mobile",(e=>{U.value=e})),o.on("collapse",(e=>{H.value=e}));const e=[{name:_.value,meta:{title:"首页"},query:{},params:{}}];V.value=JSON.parse(sessionStorage.getItem("historys"))||e,window.sessionStorage.getItem("activeValue")?E.value=window.sessionStorage.getItem("activeValue"):E.value=j(C),z(C),"true"===window.sessionStorage.getItem("needCloseAll")&&(F(),window.sessionStorage.removeItem("needCloseAll"))})(),u((()=>{o.off("collapse"),o.off("mobile")})),(e,a)=>{const s=r("el-tab-pane"),l=r("el-tabs");return i(),v("div",N,[m(l,{modelValue:E.value,"onUpdate:modelValue":a[0]||(a[0]=e=>E.value=e),closable:!(1===V.value.length&&e.$route.name===h(_)),type:"card",onContextmenu:a[1]||(a[1]=q((e=>(e=>{if(1===V.value.length&&C.name===_.value)return!1;let a="";if(a="SPAN"===e.srcElement.nodeName?e.srcElement.offsetParent.id:e.srcElement.id,a){let s;A.value=!0,s=H.value?54:220,U.value&&(s=0),T.value=e.clientX-s,L.value=e.clientY+10,Y.value=a.substring(4)}})(e)),["prevent"])),onTabChange:D,onTabRemove:K},{default:c((()=>[(i(!0),v(d,null,y(V.value,(e=>(i(),p(s,{key:R(e),label:e.meta.title,name:R(e),tab:e,class:"gva-tab"},{label:c((()=>[g("span",{tab:e,style:f({color:E.value===R(e)?h(P).activeColor:"#333"})},[g("i",{class:"dot",style:f({backgroundColor:E.value===R(e)?h(P).activeColor:"#ddd"})},null,4),b(" "+S(h(I)(e.meta.title,e)),1)],12,k)])),_:2},1032,["label","name","tab"])))),128))])),_:1},8,["modelValue","closable"]),O(g("ul",{style:f({left:T.value+"px",top:L.value+"px"}),class:"contextmenu"},[g("li",{onClick:F},"关闭所有"),g("li",{onClick:G},"关闭左侧"),g("li",{onClick:X},"关闭右侧"),g("li",{onClick:$},"关闭其他")],4),[[w,A.value]])])}}});export{C as default};
