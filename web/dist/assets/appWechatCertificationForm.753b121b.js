/*! 
 Build based on gin-vue-admin 
 Time : 1677739392000 */
import{f as e,h as l,u as a}from"./appWechatCertification.f4eac258.js";import{g as d}from"./format.7121bf90.js";import{C as o,u as t,r as u,a as r,b as m,o as p,c as n,d as c,e as i,w as s,F as b,x as v,m as V,h as f,i as I}from"./index.533d8bbb.js";import"./date.34b11046.js";import"./dictionary.f9434c9f.js";import"./dictionary.a38b2399.js";import"./sysDictionary.dfe03bab.js";const y={class:"gva-form-box"},g={name:"AppWechatCertification"},h=Object.assign(g,{setup(g){const h=o(),_=t(),C=u(""),T=u([]),U=u([]),k=u({reqId:"",mercId:"",mercName:"",idTypeCd:"",idFrontImg:"",idBackImg:"",contIdValidDateBegin:"",contIdValidDateEnd:"",contactCorpId:"",idAddress:"",merContactType:void 0,managementType:void 0,storeName:"",busAutLetterImg:""}),D=r({mercId:[{required:!0,message:"",trigger:["input","blur"]}]}),A=u();(async()=>{if(h.query.id){const l=await e({ID:h.query.id});0===l.code&&(k.value=l.data.reappWechatCertification,C.value="update")}else C.value="create";T.value=await d("merContactType"),U.value=await d("managementType\t")})();const j=async()=>{var e;null==(e=A.value)||e.validate((async e=>{if(!e)return;let d;switch(C.value){case"create":default:d=await l(k.value);break;case"update":d=await a(k.value)}0===d.code&&I({type:"success",message:"创建/更改成功"})}))},w=()=>{_.go(-1)};return(e,l)=>{const a=m("el-input"),d=m("el-form-item"),o=m("el-option"),t=m("el-select"),u=m("el-button"),r=m("el-form");return p(),n("div",null,[c("div",y,[i(r,{model:k.value,ref_key:"elFormRef",ref:A,"label-position":"right",rules:D,"label-width":"80px"},{default:s((()=>[i(d,{label:"流水号:",prop:"reqId"},{default:s((()=>[i(a,{modelValue:k.value.reqId,"onUpdate:modelValue":l[0]||(l[0]=e=>k.value.reqId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"商户号:",prop:"mercId"},{default:s((()=>[i(a,{modelValue:k.value.mercId,"onUpdate:modelValue":l[1]||(l[1]=e=>k.value.mercId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"商户名称:",prop:"mercName"},{default:s((()=>[i(a,{modelValue:k.value.mercName,"onUpdate:modelValue":l[2]||(l[2]=e=>k.value.mercName=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"证件类型:",prop:"idTypeCd"},{default:s((()=>[i(a,{modelValue:k.value.idTypeCd,"onUpdate:modelValue":l[3]||(l[3]=e=>k.value.idTypeCd=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"联系人证件正面照片:",prop:"idFrontImg"},{default:s((()=>[i(a,{modelValue:k.value.idFrontImg,"onUpdate:modelValue":l[4]||(l[4]=e=>k.value.idFrontImg=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"联系人证件反面照片:",prop:"idBackImg"},{default:s((()=>[i(a,{modelValue:k.value.idBackImg,"onUpdate:modelValue":l[5]||(l[5]=e=>k.value.idBackImg=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"联系人证件有效期开始时间:",prop:"contIdValidDateBegin"},{default:s((()=>[i(a,{modelValue:k.value.contIdValidDateBegin,"onUpdate:modelValue":l[6]||(l[6]=e=>k.value.contIdValidDateBegin=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"联系人证件有效期结束时间:",prop:"contIdValidDateEnd"},{default:s((()=>[i(a,{modelValue:k.value.contIdValidDateEnd,"onUpdate:modelValue":l[7]||(l[7]=e=>k.value.contIdValidDateEnd=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"联系人证件号码:",prop:"contactCorpId"},{default:s((()=>[i(a,{modelValue:k.value.contactCorpId,"onUpdate:modelValue":l[8]||(l[8]=e=>k.value.contactCorpId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"法人居住地址:",prop:"idAddress"},{default:s((()=>[i(a,{modelValue:k.value.idAddress,"onUpdate:modelValue":l[9]||(l[9]=e=>k.value.idAddress=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"联系人类型:",prop:"merContactType"},{default:s((()=>[i(t,{modelValue:k.value.merContactType,"onUpdate:modelValue":l[10]||(l[10]=e=>k.value.merContactType=e),placeholder:"请选择",clearable:!0},{default:s((()=>[(p(!0),n(b,null,v(T.value,((e,l)=>(p(),V(o,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),i(d,{label:"小微经营类型:",prop:"managementType"},{default:s((()=>[i(t,{modelValue:k.value.managementType,"onUpdate:modelValue":l[11]||(l[11]=e=>k.value.managementType=e),placeholder:"请选择",clearable:!0},{default:s((()=>[(p(!0),n(b,null,v(U.value,((e,l)=>(p(),V(o,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),i(d,{label:"门店名称:",prop:"storeName"},{default:s((()=>[i(a,{modelValue:k.value.storeName,"onUpdate:modelValue":l[12]||(l[12]=e=>k.value.storeName=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,{label:"业务办理授权函:",prop:"busAutLetterImg"},{default:s((()=>[i(a,{modelValue:k.value.busAutLetterImg,"onUpdate:modelValue":l[13]||(l[13]=e=>k.value.busAutLetterImg=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),i(d,null,{default:s((()=>[i(u,{size:"small",type:"primary",onClick:j},{default:s((()=>[f("保存")])),_:1}),i(u,{size:"small",type:"primary",onClick:w},{default:s((()=>[f("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{h as default};