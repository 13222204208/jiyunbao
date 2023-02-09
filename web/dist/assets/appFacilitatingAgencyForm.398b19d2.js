/*! 
 Build based on gin-vue-admin 
 Time : 1675935707000 */
import{f as e,c as a,u as l}from"./appFacilitatingAgency.2640ffde.js";import{g as u}from"./format.a5b70b7c.js";import{C as r,u as t,r as o,a as d,b as i,o as s,c as p,d as n,e as c,w as m,F as v,x as b,m as f,h as g,i as V}from"./index.3b117381.js";import"./date.34b11046.js";import"./dictionary.c64483b9.js";import"./dictionary.2e4ec8e3.js";import"./sysDictionary.3083fcf3.js";const y={class:"gva-form-box"},h={name:"AppFacilitatingAgency"},_=Object.assign(h,{setup(h){const _=r(),w=t(),U=o(""),j=o([]),q=o([]),k=o([]),x=o({phone:"",name:"",way:void 0,principal:"",card:"",mail:"",area:"",address:"",status:void 0,certification:void 0}),F=d({phone:[{required:!0,message:"",trigger:["input","blur"]}],name:[{required:!0,message:"",trigger:["input","blur"]}],way:[{required:!0,message:"",trigger:["input","blur"]}],principal:[{required:!0,message:"",trigger:["input","blur"]}],status:[{required:!0,message:"",trigger:["input","blur"]}],certification:[{required:!0,message:"",trigger:["input","blur"]}]}),A=o();(async()=>{if(_.query.id){const a=await e({ID:_.query.id});0===a.code&&(x.value=a.data.reappFacilitatingAgency,U.value="update")}else U.value="create";j.value=await u("mchType"),q.value=await u("serviceState"),k.value=await u("serviceCertification")})();const C=async()=>{var e;null==(e=A.value)||e.validate((async e=>{if(!e)return;let u;switch(U.value){case"create":default:u=await a(x.value);break;case"update":u=await l(x.value)}0===u.code&&V({type:"success",message:"创建/更改成功"})}))},z=()=>{w.go(-1)};return(e,a)=>{const l=i("el-input"),u=i("el-form-item"),r=i("el-option"),t=i("el-select"),o=i("el-button"),d=i("el-form");return s(),p("div",null,[n("div",y,[c(d,{model:x.value,ref_key:"elFormRef",ref:A,"label-position":"right",rules:F,"label-width":"80px"},{default:m((()=>[c(u,{label:"手机号:",prop:"phone"},{default:m((()=>[c(l,{modelValue:x.value.phone,"onUpdate:modelValue":a[0]||(a[0]=e=>x.value.phone=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"机构名称:",prop:"name"},{default:m((()=>[c(l,{modelValue:x.value.name,"onUpdate:modelValue":a[1]||(a[1]=e=>x.value.name=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"机构类型:",prop:"way"},{default:m((()=>[c(t,{modelValue:x.value.way,"onUpdate:modelValue":a[2]||(a[2]=e=>x.value.way=e),placeholder:"请选择",clearable:!0},{default:m((()=>[(s(!0),p(v,null,b(j.value,((e,a)=>(s(),f(r,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),c(u,{label:"负责人姓名:",prop:"principal"},{default:m((()=>[c(l,{modelValue:x.value.principal,"onUpdate:modelValue":a[3]||(a[3]=e=>x.value.principal=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"身份证号:",prop:"card"},{default:m((()=>[c(l,{modelValue:x.value.card,"onUpdate:modelValue":a[4]||(a[4]=e=>x.value.card=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"邮箱:",prop:"mail"},{default:m((()=>[c(l,{modelValue:x.value.mail,"onUpdate:modelValue":a[5]||(a[5]=e=>x.value.mail=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"所属地区:",prop:"area"},{default:m((()=>[c(l,{modelValue:x.value.area,"onUpdate:modelValue":a[6]||(a[6]=e=>x.value.area=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"详细地区:",prop:"address"},{default:m((()=>[c(l,{modelValue:x.value.address,"onUpdate:modelValue":a[7]||(a[7]=e=>x.value.address=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),c(u,{label:"资料状态:",prop:"status"},{default:m((()=>[c(t,{modelValue:x.value.status,"onUpdate:modelValue":a[8]||(a[8]=e=>x.value.status=e),placeholder:"请选择",clearable:!0},{default:m((()=>[(s(!0),p(v,null,b(q.value,((e,a)=>(s(),f(r,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),c(u,{label:"资质状态:",prop:"certification"},{default:m((()=>[c(t,{modelValue:x.value.certification,"onUpdate:modelValue":a[9]||(a[9]=e=>x.value.certification=e),placeholder:"请选择",clearable:!0},{default:m((()=>[(s(!0),p(v,null,b(k.value,((e,a)=>(s(),f(r,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),c(u,null,{default:m((()=>[c(o,{size:"small",type:"primary",onClick:C},{default:m((()=>[g("保存")])),_:1}),c(o,{size:"small",type:"primary",onClick:z},{default:m((()=>[g("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{_ as default};