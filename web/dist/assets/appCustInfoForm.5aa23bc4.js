/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{f as e,h as l,u as a}from"./appCustInfo.6fcc7e99.js";import{g as u}from"./format.215352f7.js";import{G as r,u as o,r as d,a as t,b as c,o as p,c as m,d as n,e as s,w as i,F as b,x as v,m as V,h as g,i as f}from"./index.42b7f490.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const y={class:"gva-form-box"},h={name:"AppCustInfo"},C=Object.assign(h,{setup(h){const C=r(),_=o(),U=d(""),N=d([]),q=d([]),M=d({uid:0,reqId:"",certId:"",agtMercId:"",mercName:"",mercShortName:"",mercType:void 0,mccCd:"",contactMail:"",contactMan:"",contactPhone:"",cusMgrNm:"",notifyUrl:"",crpCertNo:"",crpCertType:"",certBgn:"",certExpire:"",crpNm:"",crpPhone:"",stlAccNo:"",bankSubCode:"",stlAccType:void 0,busProviceCode:"",busCityCode:"",busAddr:""}),A=t({mercName:[{required:!0,message:"",trigger:["input","blur"]}],mercType:[{required:!0,message:"",trigger:["input","blur"]}],mccCd:[{required:!0,message:"",trigger:["input","blur"]}],contactMail:[{required:!0,message:"",trigger:["input","blur"]}],cusMgrNm:[{required:!0,message:"",trigger:["input","blur"]}],notifyUrl:[{required:!0,message:"",trigger:["input","blur"]}],crpCertNo:[{required:!0,message:"",trigger:["input","blur"]}],crpCertType:[{required:!0,message:"",trigger:["input","blur"]}],certBgn:[{required:!0,message:"",trigger:["input","blur"]}],certExpire:[{required:!0,message:"",trigger:["input","blur"]}],crpNm:[{required:!0,message:"",trigger:["input","blur"]}],crpPhone:[{required:!0,message:"",trigger:["input","blur"]}],stlAccNo:[{required:!0,message:"",trigger:["input","blur"]}],bankSubCode:[{required:!0,message:"",trigger:["input","blur"]}],stlAccType:[{required:!0,message:"",trigger:["input","blur"]}],busProviceCode:[{required:!0,message:"",trigger:["input","blur"]}],busCityCode:[{required:!0,message:"",trigger:["input","blur"]}],busAddr:[{required:!0,message:"",trigger:["input","blur"]}]}),I=d();(async()=>{if(C.query.id){const l=await e({ID:C.query.id});0===l.code&&(M.value=l.data.reappCustInfo,U.value="update")}else U.value="create";N.value=await u("mercType"),q.value=await u("stlAccType")})();const T=async()=>{var e;null==(e=I.value)||e.validate((async e=>{if(!e)return;let u;switch(U.value){case"create":default:u=await l(M.value);break;case"update":u=await a(M.value)}0===u.code&&f({type:"success",message:"创建/更改成功"})}))},P=()=>{_.go(-1)};return(e,l)=>{const a=c("el-input"),u=c("el-form-item"),r=c("el-option"),o=c("el-select"),d=c("el-button"),t=c("el-form");return p(),m("div",null,[n("div",y,[s(t,{model:M.value,ref_key:"elFormRef",ref:I,"label-position":"right",rules:A,"label-width":"80px"},{default:i((()=>[s(u,{label:"商户ID:",prop:"uid"},{default:i((()=>[s(a,{modelValue:M.value.uid,"onUpdate:modelValue":l[0]||(l[0]=e=>M.value.uid=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"流水号:",prop:"reqId"},{default:i((()=>[s(a,{modelValue:M.value.reqId,"onUpdate:modelValue":l[1]||(l[1]=e=>M.value.reqId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"发起方商户号:",prop:"certId"},{default:i((()=>[s(a,{modelValue:M.value.certId,"onUpdate:modelValue":l[2]||(l[2]=e=>M.value.certId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"代理商编号:",prop:"agtMercId"},{default:i((()=>[s(a,{modelValue:M.value.agtMercId,"onUpdate:modelValue":l[3]||(l[3]=e=>M.value.agtMercId=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"商户名称:",prop:"mercName"},{default:i((()=>[s(a,{modelValue:M.value.mercName,"onUpdate:modelValue":l[4]||(l[4]=e=>M.value.mercName=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"商户简称:",prop:"mercShortName"},{default:i((()=>[s(a,{modelValue:M.value.mercShortName,"onUpdate:modelValue":l[5]||(l[5]=e=>M.value.mercShortName=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"商户类型:",prop:"mercType"},{default:i((()=>[s(o,{modelValue:M.value.mercType,"onUpdate:modelValue":l[6]||(l[6]=e=>M.value.mercType=e),placeholder:"请选择",clearable:!0},{default:i((()=>[(p(!0),m(b,null,v(N.value,((e,l)=>(p(),V(r,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),s(u,{label:"mcc码:",prop:"mccCd"},{default:i((()=>[s(a,{modelValue:M.value.mccCd,"onUpdate:modelValue":l[7]||(l[7]=e=>M.value.mccCd=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"联系人邮箱:",prop:"contactMail"},{default:i((()=>[s(a,{modelValue:M.value.contactMail,"onUpdate:modelValue":l[8]||(l[8]=e=>M.value.contactMail=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"联系人:",prop:"contactMan"},{default:i((()=>[s(a,{modelValue:M.value.contactMan,"onUpdate:modelValue":l[9]||(l[9]=e=>M.value.contactMan=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"联系人电话:",prop:"contactPhone"},{default:i((()=>[s(a,{modelValue:M.value.contactPhone,"onUpdate:modelValue":l[10]||(l[10]=e=>M.value.contactPhone=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"客户经理:",prop:"cusMgrNm"},{default:i((()=>[s(a,{modelValue:M.value.cusMgrNm,"onUpdate:modelValue":l[11]||(l[11]=e=>M.value.cusMgrNm=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"异步通知地址:",prop:"notifyUrl"},{default:i((()=>[s(a,{modelValue:M.value.notifyUrl,"onUpdate:modelValue":l[12]||(l[12]=e=>M.value.notifyUrl=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"法人证件号:",prop:"crpCertNo"},{default:i((()=>[s(a,{modelValue:M.value.crpCertNo,"onUpdate:modelValue":l[13]||(l[13]=e=>M.value.crpCertNo=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"法人证件类型:",prop:"crpCertType"},{default:i((()=>[s(a,{modelValue:M.value.crpCertType,"onUpdate:modelValue":l[14]||(l[14]=e=>M.value.crpCertType=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"证件开始日期:",prop:"certBgn"},{default:i((()=>[s(a,{modelValue:M.value.certBgn,"onUpdate:modelValue":l[15]||(l[15]=e=>M.value.certBgn=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"证件有效期:",prop:"certExpire"},{default:i((()=>[s(a,{modelValue:M.value.certExpire,"onUpdate:modelValue":l[16]||(l[16]=e=>M.value.certExpire=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"法人姓名:",prop:"crpNm"},{default:i((()=>[s(a,{modelValue:M.value.crpNm,"onUpdate:modelValue":l[17]||(l[17]=e=>M.value.crpNm=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"法人手机号:",prop:"crpPhone"},{default:i((()=>[s(a,{modelValue:M.value.crpPhone,"onUpdate:modelValue":l[18]||(l[18]=e=>M.value.crpPhone=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"结算账号:",prop:"stlAccNo"},{default:i((()=>[s(a,{modelValue:M.value.stlAccNo,"onUpdate:modelValue":l[19]||(l[19]=e=>M.value.stlAccNo=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"开户支行联行号:",prop:"bankSubCode"},{default:i((()=>[s(a,{modelValue:M.value.bankSubCode,"onUpdate:modelValue":l[20]||(l[20]=e=>M.value.bankSubCode=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"结算账户类型:",prop:"stlAccType"},{default:i((()=>[s(o,{modelValue:M.value.stlAccType,"onUpdate:modelValue":l[21]||(l[21]=e=>M.value.stlAccType=e),placeholder:"请选择",clearable:!0},{default:i((()=>[(p(!0),m(b,null,v(q.value,((e,l)=>(p(),V(r,{key:l,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),s(u,{label:"营业归属省代码:",prop:"busProviceCode"},{default:i((()=>[s(a,{modelValue:M.value.busProviceCode,"onUpdate:modelValue":l[22]||(l[22]=e=>M.value.busProviceCode=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"营业归属市代码:",prop:"busCityCode"},{default:i((()=>[s(a,{modelValue:M.value.busCityCode,"onUpdate:modelValue":l[23]||(l[23]=e=>M.value.busCityCode=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,{label:"营业详细地址:",prop:"busAddr"},{default:i((()=>[s(a,{modelValue:M.value.busAddr,"onUpdate:modelValue":l[24]||(l[24]=e=>M.value.busAddr=e),clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),s(u,null,{default:i((()=>[s(d,{size:"small",type:"primary",onClick:T},{default:i((()=>[g("保存")])),_:1}),s(d,{size:"small",type:"primary",onClick:P},{default:i((()=>[g("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{C as default};