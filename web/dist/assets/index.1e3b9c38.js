/*! 
 Build based on gin-vue-admin 
 Time : 1673519692000 */
import{_ as e,U as a,g as l,e as t}from"./common.c8e60ca0.js";import{W as s}from"./warningBar.e39232e5.js";import{r as n,b as u,o,m as r,w as i,e as m,d as c,h as d,c as p,F as v,x as g,t as y,A as f,i as x}from"./index.29e958a1.js";const b={class:"gva-btn-list"},h={class:"media"},C={class:"header-img-box-list"},_={class:"header-img-box-list"},k=c("picture",null,null,-1),w=["onClick"],z={__name:"index",props:{target:{type:Object,default:null},targetKey:{type:String,default:""}},emits:["enterImg"],setup(z,{expose:U,emit:S}){const V=n(""),j=n(""),B=n({}),O=n(1),F=n(0),I=n(20),K=e=>{I.value=e,P()},T=e=>{O.value=e,P()},A=n(!1),E=n([]),M=n("/api/"),P=async()=>{const e=await l({page:O.value,pageSize:I.value,...B.value});0===e.code&&(E.value=e.data.list,F.value=e.data.total,O.value=e.data.page,I.value=e.data.pageSize,A.value=!0)};return U({open:P}),(l,n)=>{const U=u("el-input"),W=u("el-form-item"),q=u("el-button"),D=u("el-form"),G=u("el-icon"),H=u("el-image"),J=u("el-pagination"),L=u("el-drawer");return o(),r(L,{modelValue:A.value,"onUpdate:modelValue":n[3]||(n[3]=e=>A.value=e),title:"媒体库",size:"650px"},{default:i((()=>[m(s,{title:"点击“文件名/备注”可以编辑文件名或者备注内容。"}),c("div",b,[m(e,{imageCommon:j.value,"onUpdate:imageCommon":n[0]||(n[0]=e=>j.value=e),class:"upload-btn-media-library",onOnSuccess:P},null,8,["imageCommon"]),m(a,{imageUrl:V.value,"onUpdate:imageUrl":n[1]||(n[1]=e=>V.value=e),"file-size":512,"max-w-h":1080,class:"upload-btn-media-library",onOnSuccess:P},null,8,["imageUrl"]),m(D,{ref:"searchForm",inline:!0,model:B.value},{default:i((()=>[m(W,{label:""},{default:i((()=>[m(U,{modelValue:B.value.keyword,"onUpdate:modelValue":n[2]||(n[2]=e=>B.value.keyword=e),class:"keyword",placeholder:"请输入文件名或备注"},null,8,["modelValue"])])),_:1}),m(W,null,{default:i((()=>[m(q,{size:"small",type:"primary",icon:"search",onClick:P},{default:i((()=>[d("查询")])),_:1})])),_:1})])),_:1},8,["model"])]),c("div",h,[(o(!0),p(v,null,g(E.value,((e,a)=>(o(),p("div",{key:a,class:"media-box"},[c("div",C,[(o(),r(H,{key:a,src:e.url&&"http"!==e.url.slice(0,4)?M.value+e.url:e.url,onClick:a=>{return l=e.url,t=z.target,s=z.targetKey,t&&s&&(t[s]=l),S("enterImg",l),void(A.value=!1);var l,t,s}},{error:i((()=>[c("div",_,[m(G,null,{default:i((()=>[k])),_:1})])])),_:2},1032,["src","onClick"]))]),c("div",{class:"img-title",onClick:a=>(async e=>{f.prompt("请输入文件名或者备注","编辑",{confirmButtonText:"确定",cancelButtonText:"取消",inputPattern:/\S/,inputErrorMessage:"不能为空",inputValue:e.name}).then((async({value:a})=>{e.name=a,0===(await t(e)).code&&(x({type:"success",message:"编辑成功!"}),P())})).catch((()=>{x({type:"info",message:"取消修改"})}))})(e)},y(e.name),9,w)])))),128))]),m(J,{"current-page":O.value,"page-size":I.value,total:F.value,style:{"justify-content":"center"},layout:"total, prev, pager, next, jumper",onCurrentChange:T,onSizeChange:K},null,8,["current-page","page-size","total"])])),_:1},8,["modelValue"])}}};export{z as _};