/*! 
 Build based on gin-vue-admin 
 Time : 1677739392000 */
import{g as e,i as a,d as l,f as t,a as s,c as i,u as o}from"./appUploadImg.6bf3a25e.js";import{f as u,a as r,g as n}from"./format.7121bf90.js";import{r as p,j as d,C as c,u as v,a as g,b as m,o as f,c as y,d as b,e as w,w as h,h as I,t as k,z as _,F as z,x as C,m as x,H as U,i as j,A as T}from"./index.533d8bbb.js";import"./date.34b11046.js";import"./dictionary.f9434c9f.js";import"./dictionary.a38b2399.js";import"./sysDictionary.dfe03bab.js";const q={class:"gva-table-box"},D={class:"gva-btn-list"},V=b("p",null,"确定要删除吗？",-1),F={style:{"text-align":"right","margin-top":"8px"}},A={class:"gva-pagination"},S={class:"dialog-footer"},B={name:"AppUploadImg"},R=Object.assign(B,{setup(B){const R=p("/api"),$=d(),H=p(""),K=c(),O=v();console.log("获取的参数",K.query.id),O.isReady().then((()=>{P.value.custInfoId=K.query.id}));const E=()=>{window.open("https://jiyunbao.vvv5g.com/api/uploads/explain.jpg","_blank")},G=p(!1),J=e=>{G.value=!0;const a="image/jpeg"===e.type,l="image/png"===e.type,t=e.size/1024/1024<.5;return a||l||(j.error("上传图片只能是 jpg或png 格式!"),G.value=!1),t||(j.error("未压缩未见上传图片大小不能超过 500KB，请使用压缩上传"),G.value=!1),(l||a)&&t},L=e=>{if(G.value=!1,console.log("图片",e),0===e.code){var a=e.data.file.url;H.value="/api/"+a,console.log("图片地址",H.value),P.value.imgUrl=a,j({type:"success",message:"上传成功"}),0===e.code&&se()}else j({type:"warning",message:e.msg})},M=()=>{j({type:"error",message:"上传失败"}),G.value=!1};p(null);const N=p([]),P=p({custInfoId:K.query.id,reqId:"",picType:void 0,imgUrl:"",sysFlowId:""}),Q=g({picType:[{required:!0,message:"",trigger:["input","blur"]}],imgUrl:[{required:!0,message:"",trigger:["input","blur"]}]}),W=p(),X=p(1),Y=p(0),Z=p(10),ee=p([]),ae=p({}),le=e=>{Z.value=e,se()},te=e=>{X.value=e,se()},se=async()=>{ae.value.custInfoId=K.query.id;const a=await e({page:X.value,pageSize:Z.value,...ae.value});0===a.code&&(ee.value=a.data.list,Y.value=a.data.total,X.value=a.data.page,Z.value=a.data.pageSize)};se();(async()=>{N.value=await n("uploadImgType")})();const ie=p([]),oe=e=>{ie.value=e},ue=p(!1),re=async()=>{const e=[];if(0===ie.value.length)return void j({type:"warning",message:"请选择要删除的数据"});ie.value&&ie.value.map((a=>{e.push(a.ID)}));0===(await l({ids:e})).code&&(j({type:"success",message:"删除成功"}),ee.value.length===e.length&&X.value>1&&X.value--,ue.value=!1,se())},ne=p(""),pe=async e=>{0===(await s({ID:e.ID})).code&&(j({type:"success",message:"删除成功"}),1===ee.value.length&&X.value>1&&X.value--,se())},de=p(!1),ce=()=>{ne.value="create",de.value=!0},ve=()=>{de.value=!1,P.value={custInfoId:K.query.id,reqId:"",picType:void 0,imgUrl:"",sysFlowId:""}},ge=async()=>{var e;null==(e=W.value)||e.validate((async e=>{if(!e)return;let a;switch(P.value.custInfoId=parseInt(P.value.custInfoId),console.log("接收的参数",K.query.id),ne.value){case"create":default:a=await i(P.value);break;case"update":a=await o(P.value)}0===a.code&&(j({type:"success",message:"创建/更改成功"}),ve(),se())}))};return(e,l)=>{const s=m("el-button"),i=m("el-popover"),o=m("el-table-column"),n=m("el-image"),p=m("el-table"),d=m("el-pagination"),c=m("el-option"),v=m("el-select"),g=m("el-form-item"),B=m("el-upload"),K=m("el-form"),O=m("el-dialog");return f(),y("div",null,[b("div",q,[b("div",D,[w(s,{size:"small",type:"primary",icon:"plus",onClick:ce},{default:h((()=>[I("新增")])),_:1}),w(i,{visible:ue.value,"onUpdate:visible":l[1]||(l[1]=e=>ue.value=e),placement:"top",width:"160"},{reference:h((()=>[w(s,{type:"danger",onClick:E},{default:h((()=>[I("图片上传说明")])),_:1})])),default:h((()=>[V,b("div",F,[w(s,{size:"small",type:"primary",link:"",onClick:l[0]||(l[0]=e=>ue.value=!1)},{default:h((()=>[I("取消")])),_:1}),w(s,{size:"small",type:"primary",onClick:re},{default:h((()=>[I("确定")])),_:1})])])),_:1},8,["visible"])]),w(p,{ref:"multipleTable",style:{width:"100%"},"tooltip-effect":"dark",data:ee.value,"row-key":"ID",onSelectionChange:oe},{default:h((()=>[w(o,{align:"left",label:"日期",width:"180"},{default:h((e=>[I(k(_(u)(e.row.CreatedAt)),1)])),_:1}),w(o,{align:"left",label:"商户进件编号",prop:"custInfoId",width:"150"}),w(o,{align:"left",label:"流水号",prop:"reqId",width:"150"}),w(o,{align:"left",label:"图片类别",prop:"picType",width:"120"},{default:h((e=>[I(k(_(r)(e.row.picType,N.value)),1)])),_:1}),w(o,{align:"left",label:"图片地址",prop:"imgUrl",width:"150"},{default:h((e=>[w(n,{style:{width:"100px",height:"100px"},src:R.value+"/"+e.row.imgUrl,fit:"cover"},null,8,["src"])])),_:1}),w(o,{align:"left",label:"入网申请流水号",prop:"sysFlowId",width:"180"}),w(o,{align:"left",label:"按钮组"},{default:h((e=>[w(s,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:l=>(async e=>{const l=await a(e);console.log("结果确认",l.code),0===l.code&&j({type:"success",message:l.msg})})(e.row)},{default:h((()=>[I("图片上送")])),_:2},1032,["onClick"]),w(s,{type:"primary",link:"",icon:"edit",size:"small",class:"table-button",onClick:a=>(async e=>{const a=await t({ID:e.ID});ne.value="update",0===a.code&&(P.value=a.data.reappUploadImg,console.log("数据详情",a.data),H.value="/api/"+a.data.reappUploadImg.imgUrl,de.value=!0)})(e.row)},{default:h((()=>[I("变更图片")])),_:2},1032,["onClick"]),w(s,{type:"primary",link:"",icon:"delete",size:"small",onClick:a=>{return l=e.row,void T.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((()=>{pe(l)}));var l}},{default:h((()=>[I("删除")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),b("div",A,[w(d,{layout:"total, sizes, prev, pager, next, jumper","current-page":X.value,"page-size":Z.value,"page-sizes":[10,30,50,100],total:Y.value,onCurrentChange:te,onSizeChange:le},null,8,["current-page","page-size","total"])])]),w(O,{modelValue:de.value,"onUpdate:modelValue":l[3]||(l[3]=e=>de.value=e),"before-close":ve,title:"弹窗操作"},{footer:h((()=>[b("div",S,[w(s,{size:"small",onClick:ve},{default:h((()=>[I("取 消")])),_:1}),w(s,{size:"small",type:"primary",onClick:ge},{default:h((()=>[I("确 定")])),_:1})])])),default:h((()=>[w(K,{model:P.value,"label-position":"right",ref_key:"elFormRef",ref:W,rules:Q,"label-width":"120px"},{default:h((()=>[w(g,{label:"图片类别:",prop:"picType"},{default:h((()=>[w(v,{modelValue:P.value.picType,"onUpdate:modelValue":l[2]||(l[2]=e=>P.value.picType=e),placeholder:"请选择",style:{width:"100%"},clearable:!0},{default:h((()=>[(f(!0),y(z,null,C(N.value,((e,a)=>(f(),x(c,{key:a,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),w(g,{label:"图片:"},{default:h((()=>[b("div",{class:"user-headpic-update",style:U({"background-image":`url(${H.value})`,"background-repeat":"no-repeat","background-size":"cover",width:"200px",height:"200px"})},null,4),w(B,{action:`${R.value}/fileUploadAndDownload/upload`,"before-upload":J,headers:{"x-token":_($).token},"on-error":M,"on-success":L,"show-file-list":!1,class:"upload-btn"},{default:h((()=>[w(s,{size:"small",type:"primary"},{default:h((()=>[I("上传图片")])),_:1})])),_:1},8,["action","headers"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{R as default};