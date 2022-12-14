/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{f as e,c as a,u as l}from"./appUser.c38357a7.js";import{g as o}from"./format.215352f7.js";import{G as u,u as d,r,a as t,b as p,o as s,c as m,d as n,e as c,w as i,F as v,x as f,m as b,h,i as V}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const y={class:"gva-form-box"},_={name:"AppUser"},j=Object.assign(_,{setup(_){const j=u(),w=d(),U=r(""),g=r([]),k=r({nickname:"",phone:"",serviceId:0,joinCode:"",mchType:void 0,agreement:"",avatar:"",wechat:""}),C=t({}),x=r();(async()=>{if(j.query.id){const a=await e({ID:j.query.id});0===a.code&&(k.value=a.data.reappUser,U.value="update")}else U.value="create";g.value=await o("mchType")})();const I=async()=>{var e;null==(e=x.value)||e.validate((async e=>{if(!e)return;let o;switch(U.value){case"create":default:o=await a(k.value);break;case"update":o=await l(k.value)}0===o.code&&V({type:"success",message:"创建/更改成功"})}))},T=()=>{w.go(-1)};return(e,a)=>{const l=p("el-input"),o=p("el-form-item"),u=p("el-option"),d=p("el-select"),r=p("el-button"),t=p("el-form");return s(),m("div",null,[n("div",y,[c(t,{model:k.value,ref_key:"elFormRef",ref:x,"label-position":"right",rules:C,"label-width":"80px"},{default:i((()=>[c(o,{label:"昵称:",prop:"nickname"},{default:i((()=>[c(l,{modelValue:k.value.nickname,"onUpdate:modelValue":a[0]||(a[0]=e=>k.value.nickname=e),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,{label:"手机号:",prop:"phone"},{default:i((()=>[c(l,{modelValue:k.value.phone,"onUpdate:modelValue":a[1]||(a[1]=e=>k.value.phone=e),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,{label:"服务机构:",prop:"serviceId"},{default:i((()=>[c(l,{modelValue:k.value.serviceId,"onUpdate:modelValue":a[2]||(a[2]=e=>k.value.serviceId=e),modelModifiers:{number:!0},clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,{label:"联保代码:",prop:"joinCode"},{default:i((()=>[c(l,{modelValue:k.value.joinCode,"onUpdate:modelValue":a[3]||(a[3]=e=>k.value.joinCode=e),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,{label:"商户类型:",prop:"mchType"},{default:i((()=>[c(d,{modelValue:k.value.mchType,"onUpdate:modelValue":a[4]||(a[4]=e=>k.value.mchType=e),placeholder:"请选择",clearable:!1},{default:i((()=>[(s(!0),m(v,null,f(g.value,((e,a)=>(s(),b(u,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),c(o,{label:"商户入驻协议:",prop:"agreement"},{default:i((()=>[c(l,{modelValue:k.value.agreement,"onUpdate:modelValue":a[5]||(a[5]=e=>k.value.agreement=e),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,{label:"头像:",prop:"avatar"},{default:i((()=>[c(l,{modelValue:k.value.avatar,"onUpdate:modelValue":a[6]||(a[6]=e=>k.value.avatar=e),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,{label:"微信:",prop:"wechat"},{default:i((()=>[c(l,{modelValue:k.value.wechat,"onUpdate:modelValue":a[7]||(a[7]=e=>k.value.wechat=e),clearable:!1,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(o,null,{default:i((()=>[c(r,{size:"small",type:"primary",onClick:I},{default:i((()=>[h("保存")])),_:1}),c(r,{size:"small",type:"primary",onClick:T},{default:i((()=>[h("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{j as default};
