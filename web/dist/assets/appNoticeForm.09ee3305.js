/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{f as e,c as a,u as l}from"./appNotice.808548d7.js";import{g as t}from"./format.215352f7.js";import{G as s,u as o,r as u,a as r,b as i,o as n,c as d,d as c,e as p,w as m,F as f,x as v,m as y,h as b,i as g}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const w={class:"gva-form-box"},j={name:"AppNotice"},_=Object.assign(j,{setup(j){const _=s(),h=o(),V=u(""),k=u([]),x=u({contents:"",way:void 0}),q=r({contents:[{required:!0,message:"",trigger:["input","blur"]}],way:[{required:!0,message:"",trigger:["input","blur"]}]}),N=u();(async()=>{if(_.query.id){const a=await e({ID:_.query.id});0===a.code&&(x.value=a.data.reappNotice,V.value="update")}else V.value="create";k.value=await t("notice")})();const z=async()=>{var e;null==(e=N.value)||e.validate((async e=>{if(!e)return;let t;switch(V.value){case"create":default:t=await a(x.value);break;case"update":t=await l(x.value)}0===t.code&&g({type:"success",message:"创建/更改成功"})}))},C=()=>{h.go(-1)};return(e,a)=>{const l=i("el-input"),t=i("el-form-item"),s=i("el-option"),o=i("el-select"),u=i("el-button"),r=i("el-form");return n(),d("div",null,[c("div",w,[p(r,{model:x.value,ref_key:"elFormRef",ref:N,"label-position":"right",rules:q,"label-width":"80px"},{default:m((()=>[p(t,{label:"公告内容:",prop:"contents"},{default:m((()=>[p(l,{modelValue:x.value.contents,"onUpdate:modelValue":a[0]||(a[0]=e=>x.value.contents=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),p(t,{label:"公告类型:",prop:"way"},{default:m((()=>[p(o,{modelValue:x.value.way,"onUpdate:modelValue":a[1]||(a[1]=e=>x.value.way=e),placeholder:"请选择",clearable:!0},{default:m((()=>[(n(!0),d(f,null,v(k.value,((e,a)=>(n(),y(s,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),p(t,null,{default:m((()=>[p(u,{size:"small",type:"primary",onClick:z},{default:m((()=>[b("保存")])),_:1}),p(u,{size:"small",type:"primary",onClick:C},{default:m((()=>[b("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{_ as default};
