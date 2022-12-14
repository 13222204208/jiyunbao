/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{s as e,G as l,a9 as a,r as t,b as i,o as u,c as s,d as o,e as r,w as d,h as n,t as p,A as c,i as m}from"./index.42b7f490.js";import{f as v,b as y}from"./format.215352f7.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const f=l=>e({url:"/sysDictionaryDetail/createSysDictionaryDetail",method:"post",data:l}),b={class:"gva-search-box"},g={class:"gva-table-box"},D={class:"gva-btn-list"},h=o("p",null,"确定要删除吗？",-1),w={style:{"text-align":"right","margin-top":"8px"}},V={class:"gva-pagination"},_={class:"dialog-footer"},k={name:"SysDictionaryDetail"},z=Object.assign(k,{setup(k){const z=l();a(((e,l)=>{"dictionaryDetail"===e.name&&(q.value.sysDictionaryID=e.params.id,G())}));const C=t({label:null,value:null,status:!0,sort:null}),x=t({label:[{required:!0,message:"请输入展示值",trigger:"blur"}],value:[{required:!0,message:"请输入字典值",trigger:"blur"}],sort:[{required:!0,message:"排序标记",trigger:"blur"}]}),I=t(1),S=t(0),U=t(10),j=t([]),q=t({sysDictionaryID:Number(z.params.id)}),N=()=>{q.value={sysDictionaryID:Number(z.params.id)}},A=()=>{I.value=1,U.value=10,""===q.value.status&&(q.value.status=null),G()},M=e=>{U.value=e,G()},F=e=>{I.value=e,G()},G=async()=>{const l=await(a={page:I.value,pageSize:U.value,...q.value},e({url:"/sysDictionaryDetail/getSysDictionaryDetailList",method:"get",params:a}));var a;0===l.code&&(j.value=l.data.list,S.value=l.data.total,I.value=l.data.page,U.value=l.data.pageSize)};G();const L=t(""),O=t(!1),T=async l=>{const a=await(t={ID:l.ID},e({url:"/sysDictionaryDetail/findSysDictionaryDetail",method:"get",params:t}));var t;L.value="update",0===a.code&&(C.value=a.data.reSysDictionaryDetail,O.value=!0)},B=()=>{O.value=!1,C.value={label:null,value:null,status:!0,sort:null,sysDictionaryID:""}},E=async l=>{l.visible=!1;var a;0===(await(a={ID:l.ID},e({url:"/sysDictionaryDetail/deleteSysDictionaryDetail",method:"delete",data:a}))).code&&(m({type:"success",message:"删除成功"}),1===j.value.length&&I.value>1&&I.value--,G())},H=t(null),J=async()=>{C.value.sysDictionaryID=Number(z.params.id),H.value.validate((async l=>{if(!l)return;let a;switch(L.value){case"create":default:a=await f(C.value);break;case"update":a=await(t=C.value,e({url:"/sysDictionaryDetail/updateSysDictionaryDetail",method:"put",data:t}))}var t;0===a.code&&(m({type:"success",message:"创建/更改成功"}),B(),G())}))},K=()=>{L.value="create",O.value=!0};return(e,l)=>{const a=i("el-input"),t=i("el-form-item"),m=i("el-option"),f=i("el-select"),k=i("el-button"),z=i("el-form"),G=i("el-table-column"),L=i("el-popover"),P=i("el-table"),Q=i("el-pagination"),R=i("el-input-number"),W=i("el-switch"),X=i("el-dialog");return u(),s("div",null,[o("div",b,[r(z,{inline:!0,model:q.value},{default:d((()=>[r(t,{label:"展示值"},{default:d((()=>[r(a,{modelValue:q.value.label,"onUpdate:modelValue":l[0]||(l[0]=e=>q.value.label=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),r(t,{label:"字典值"},{default:d((()=>[r(a,{modelValue:q.value.value,"onUpdate:modelValue":l[1]||(l[1]=e=>q.value.value=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),r(t,{label:"启用状态",prop:"status"},{default:d((()=>[r(f,{modelValue:q.value.status,"onUpdate:modelValue":l[2]||(l[2]=e=>q.value.status=e),placeholder:"请选择"},{default:d((()=>[r(m,{key:"true",label:"是",value:"true"}),r(m,{key:"false",label:"否",value:"false"})])),_:1},8,["modelValue"])])),_:1}),r(t,null,{default:d((()=>[r(k,{size:"small",type:"primary",icon:"search",onClick:A},{default:d((()=>[n("查询")])),_:1}),r(k,{size:"small",icon:"refresh",onClick:N},{default:d((()=>[n("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),o("div",g,[o("div",D,[r(k,{size:"small",type:"primary",icon:"plus",onClick:K},{default:d((()=>[n("新增字典项")])),_:1})]),r(P,{ref:"multipleTable",data:j.value,style:{width:"100%"},"tooltip-effect":"dark","row-key":"ID"},{default:d((()=>[r(G,{type:"selection",width:"55"}),r(G,{align:"left",label:"日期",width:"180"},{default:d((e=>[n(p(c(v)(e.row.CreatedAt)),1)])),_:1}),r(G,{align:"left",label:"展示值",prop:"label",width:"120"}),r(G,{align:"left",label:"字典值",prop:"value",width:"120"}),r(G,{align:"left",label:"启用状态",prop:"status",width:"120"},{default:d((e=>[n(p(c(y)(e.row.status)),1)])),_:1}),r(G,{align:"left",label:"排序标记",prop:"sort",width:"120"}),r(G,{align:"left",label:"按钮组"},{default:d((e=>[r(k,{size:"small",type:"primary",link:"",icon:"edit",onClick:l=>T(e.row)},{default:d((()=>[n("变更")])),_:2},1032,["onClick"]),r(L,{modelValue:e.row.visible,"onUpdate:modelValue":l=>e.row.visible=l,placement:"top",width:"160"},{reference:d((()=>[r(k,{type:"primary",link:"",icon:"delete",size:"small",onClick:l=>e.row.visible=!0},{default:d((()=>[n("删除")])),_:2},1032,["onClick"])])),default:d((()=>[h,o("div",w,[r(k,{size:"small",type:"primary",link:"",onClick:l=>e.row.visible=!1},{default:d((()=>[n("取消")])),_:2},1032,["onClick"]),r(k,{type:"primary",size:"small",onClick:l=>E(e.row)},{default:d((()=>[n("确定")])),_:2},1032,["onClick"])])])),_:2},1032,["modelValue","onUpdate:modelValue"])])),_:1})])),_:1},8,["data"]),o("div",V,[r(Q,{"current-page":I.value,"page-size":U.value,"page-sizes":[10,30,50,100],total:S.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:F,onSizeChange:M},null,8,["current-page","page-size","total"])])]),r(X,{modelValue:O.value,"onUpdate:modelValue":l[7]||(l[7]=e=>O.value=e),"before-close":B,title:"弹窗操作"},{footer:d((()=>[o("div",_,[r(k,{size:"small",onClick:B},{default:d((()=>[n("取 消")])),_:1}),r(k,{size:"small",type:"primary",onClick:J},{default:d((()=>[n("确 定")])),_:1})])])),default:d((()=>[r(z,{ref_key:"dialogForm",ref:H,model:C.value,rules:x.value,size:"medium","label-width":"110px"},{default:d((()=>[r(t,{label:"展示值",prop:"label"},{default:d((()=>[r(a,{modelValue:C.value.label,"onUpdate:modelValue":l[3]||(l[3]=e=>C.value.label=e),placeholder:"请输入展示值",clearable:"",style:{width:"100%"}},null,8,["modelValue"])])),_:1}),r(t,{label:"字典值",prop:"value"},{default:d((()=>[r(R,{modelValue:C.value.value,"onUpdate:modelValue":l[4]||(l[4]=e=>C.value.value=e),modelModifiers:{number:!0},"step-strictly":"",step:1,placeholder:"请输入字典值",clearable:"",style:{width:"100%"}},null,8,["modelValue"])])),_:1}),r(t,{label:"启用状态",prop:"status",required:""},{default:d((()=>[r(W,{modelValue:C.value.status,"onUpdate:modelValue":l[5]||(l[5]=e=>C.value.status=e),"active-text":"开启","inactive-text":"停用"},null,8,["modelValue"])])),_:1}),r(t,{label:"排序标记",prop:"sort"},{default:d((()=>[r(R,{modelValue:C.value.sort,"onUpdate:modelValue":l[6]||(l[6]=e=>C.value.sort=e),modelModifiers:{number:!0},placeholder:"排序标记"},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}}});export{z as default};
