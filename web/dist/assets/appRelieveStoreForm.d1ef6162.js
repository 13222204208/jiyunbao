/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{f as e,c as a,u as l}from"./appRelieveStore.64335204.js";import{g as t}from"./format.215352f7.js";import{G as o,u as s,r as u,a as r,b as d,o as i,c as n,d as p,e as m,w as c,F as v,x as f,m as b,h as y,i as V}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const _={class:"gva-form-box"},g={name:"AppRelieveStore"},h=Object.assign(g,{setup(g){const h=o(),j=s(),w=u(""),k=u([]),x=u({uid:0,storeId:0,contents:"",status:void 0}),I=r({contents:[{required:!0,message:"",trigger:["input","blur"]}]}),R=u();(async()=>{if(h.query.id){const a=await e({ID:h.query.id});0===a.code&&(x.value=a.data.reappRelieveStore,w.value="update")}else w.value="create";k.value=await t("relieveStore")})();const S=async()=>{var e;null==(e=R.value)||e.validate((async e=>{if(!e)return;let t;switch(w.value){case"create":default:t=await a(x.value);break;case"update":t=await l(x.value)}0===t.code&&V({type:"success",message:"创建/更改成功"})}))},U=()=>{j.go(-1)};return(e,a)=>{const l=d("el-input"),t=d("el-form-item"),o=d("el-option"),s=d("el-select"),u=d("el-button"),r=d("el-form");return i(),n("div",null,[p("div",_,[m(r,{model:x.value,ref_key:"elFormRef",ref:R,"label-position":"right",rules:I,"label-width":"80px"},{default:c((()=>[m(t,{label:"用户:",prop:"uid"},{default:c((()=>[m(l,{modelValue:x.value.uid,"onUpdate:modelValue":a[0]||(a[0]=e=>x.value.uid=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(t,{label:"门店:",prop:"storeId"},{default:c((()=>[m(l,{modelValue:x.value.storeId,"onUpdate:modelValue":a[1]||(a[1]=e=>x.value.storeId=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(t,{label:"解除原因:",prop:"contents"},{default:c((()=>[m(l,{modelValue:x.value.contents,"onUpdate:modelValue":a[2]||(a[2]=e=>x.value.contents=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),m(t,{label:"状态:",prop:"status"},{default:c((()=>[m(s,{modelValue:x.value.status,"onUpdate:modelValue":a[3]||(a[3]=e=>x.value.status=e),placeholder:"请选择",clearable:!0},{default:c((()=>[(i(!0),n(v,null,f(k.value,((e,a)=>(i(),b(o,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),m(t,null,{default:c((()=>[m(u,{size:"small",type:"primary",onClick:S},{default:c((()=>[y("保存")])),_:1}),m(u,{size:"small",type:"primary",onClick:U},{default:c((()=>[y("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{h as default};