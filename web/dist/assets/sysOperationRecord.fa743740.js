/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{s as e,r as l,b as a,o as t,c as o,d as i,e as s,w as r,h as u,t as d,A as n,m as p,i as c}from"./index.42b7f490.js";import{f as v}from"./format.215352f7.js";import"./date.34b11046.js";import"./dictionary.528e6e48.js";import"./dictionary.3462a5b2.js";import"./sysDictionary.e268ed4c.js";const m={class:"gva-search-box"},f={class:"gva-table-box"},g={class:"gva-btn-list"},y=i("p",null,"确定要删除吗？",-1),h={style:{"text-align":"right","margin-top":"8px"}},b={class:"popover-box"},w={key:1},_={class:"popover-box"},k={key:1},C=i("p",null,"确定要删除吗？",-1),V={style:{"text-align":"right","margin-top":"8px"}},z={class:"gva-pagination"},x={name:"SysOperationRecord"},O=Object.assign(x,{setup(x){const O=l(1),S=l(0),j=l(10),R=l([]),I=l({}),U=()=>{I.value={}},D=()=>{O.value=1,j.value=10,""===I.value.status&&(I.value.status=null),B()},N=e=>{j.value=e,B()},A=e=>{O.value=e,B()},B=async()=>{const l=await(a={page:O.value,pageSize:j.value,...I.value},e({url:"/sysOperationRecord/getSysOperationRecordList",method:"get",params:a}));var a;0===l.code&&(R.value=l.data.list,S.value=l.data.total,O.value=l.data.page,j.value=l.data.pageSize)};B();const E=l(!1),J=l([]),L=e=>{J.value=e},P=async()=>{const l=[];J.value&&J.value.forEach((e=>{l.push(e.ID)}));var a;0===(await(a={ids:l},e({url:"/sysOperationRecord/deleteSysOperationRecordByIds",method:"delete",data:a}))).code&&(c({type:"success",message:"删除成功"}),R.value.length===l.length&&O.value>1&&O.value--,E.value=!1,B())},T=async l=>{l.visible=!1;var a;0===(await(a={ID:l.ID},e({url:"/sysOperationRecord/deleteSysOperationRecord",method:"delete",data:a}))).code&&(c({type:"success",message:"删除成功"}),1===R.value.length&&O.value>1&&O.value--,B())},q=e=>{try{return JSON.parse(e)}catch(l){return e}};return(e,l)=>{const c=a("el-input"),x=a("el-form-item"),B=a("el-button"),F=a("el-form"),G=a("el-popover"),H=a("el-table-column"),K=a("el-tag"),M=a("warning"),Q=a("el-icon"),W=a("el-table"),X=a("el-pagination");return t(),o("div",null,[i("div",m,[s(F,{inline:!0,model:I.value},{default:r((()=>[s(x,{label:"请求方法"},{default:r((()=>[s(c,{modelValue:I.value.method,"onUpdate:modelValue":l[0]||(l[0]=e=>I.value.method=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),s(x,{label:"请求路径"},{default:r((()=>[s(c,{modelValue:I.value.path,"onUpdate:modelValue":l[1]||(l[1]=e=>I.value.path=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),s(x,{label:"结果状态码"},{default:r((()=>[s(c,{modelValue:I.value.status,"onUpdate:modelValue":l[2]||(l[2]=e=>I.value.status=e),placeholder:"搜索条件"},null,8,["modelValue"])])),_:1}),s(x,null,{default:r((()=>[s(B,{size:"small",type:"primary",icon:"search",onClick:D},{default:r((()=>[u("查询")])),_:1}),s(B,{size:"small",icon:"refresh",onClick:U},{default:r((()=>[u("重置")])),_:1})])),_:1})])),_:1},8,["model"])]),i("div",f,[i("div",g,[s(G,{modelValue:E.value,"onUpdate:modelValue":l[5]||(l[5]=e=>E.value=e),placement:"top",width:"160"},{reference:r((()=>[s(B,{icon:"delete",size:"small",style:{"margin-left":"10px"},disabled:!J.value.length,onClick:l[4]||(l[4]=e=>E.value=!0)},{default:r((()=>[u("删除")])),_:1},8,["disabled"])])),default:r((()=>[y,i("div",h,[s(B,{size:"small",type:"primary",link:"",onClick:l[3]||(l[3]=e=>E.value=!1)},{default:r((()=>[u("取消")])),_:1}),s(B,{size:"small",type:"primary",onClick:P},{default:r((()=>[u("确定")])),_:1})])])),_:1},8,["modelValue"])]),s(W,{ref:"multipleTable",data:R.value,style:{width:"100%"},"tooltip-effect":"dark","row-key":"ID",onSelectionChange:L},{default:r((()=>[s(H,{align:"left",type:"selection",width:"55"}),s(H,{align:"left",label:"操作人",width:"140"},{default:r((e=>[i("div",null,d(e.row.user.userName)+"("+d(e.row.user.nickName)+")",1)])),_:1}),s(H,{align:"left",label:"日期",width:"180"},{default:r((e=>[u(d(n(v)(e.row.CreatedAt)),1)])),_:1}),s(H,{align:"left",label:"状态码",prop:"status",width:"120"},{default:r((e=>[i("div",null,[s(K,{type:"success"},{default:r((()=>[u(d(e.row.status),1)])),_:2},1024)])])),_:1}),s(H,{align:"left",label:"请求IP",prop:"ip",width:"120"}),s(H,{align:"left",label:"请求方法",prop:"method",width:"120"}),s(H,{align:"left",label:"请求路径",prop:"path",width:"240"}),s(H,{align:"left",label:"请求",prop:"path",width:"80"},{default:r((e=>[i("div",null,[e.row.body?(t(),p(G,{key:0,placement:"left-start",trigger:"click"},{reference:r((()=>[s(Q,{style:{cursor:"pointer"}},{default:r((()=>[s(M)])),_:1})])),default:r((()=>[i("div",b,[i("pre",null,d(q(e.row.body)),1)])])),_:2},1024)):(t(),o("span",w,"无"))])])),_:1}),s(H,{align:"left",label:"响应",prop:"path",width:"80"},{default:r((e=>[i("div",null,[e.row.resp?(t(),p(G,{key:0,placement:"left-start",trigger:"click"},{reference:r((()=>[s(Q,{style:{cursor:"pointer"}},{default:r((()=>[s(M)])),_:1})])),default:r((()=>[i("div",_,[i("pre",null,d(q(e.row.resp)),1)])])),_:2},1024)):(t(),o("span",k,"无"))])])),_:1}),s(H,{align:"left",label:"按钮组"},{default:r((e=>[s(G,{modelValue:e.row.visible,"onUpdate:modelValue":l=>e.row.visible=l,placement:"top",width:"160"},{reference:r((()=>[s(B,{icon:"delete",size:"small",type:"primary",link:"",onClick:l=>e.row.visible=!0},{default:r((()=>[u("删除")])),_:2},1032,["onClick"])])),default:r((()=>[C,i("div",V,[s(B,{size:"small",type:"primary",link:"",onClick:l=>e.row.visible=!1},{default:r((()=>[u("取消")])),_:2},1032,["onClick"]),s(B,{size:"small",type:"primary",onClick:l=>T(e.row)},{default:r((()=>[u("确定")])),_:2},1032,["onClick"])])])),_:2},1032,["modelValue","onUpdate:modelValue"])])),_:1})])),_:1},8,["data"]),i("div",z,[s(X,{"current-page":O.value,"page-size":j.value,"page-sizes":[10,30,50,100],total:S.value,layout:"total, sizes, prev, pager, next, jumper",onCurrentChange:A,onSizeChange:N},null,8,["current-page","page-size","total"])])])])}}});export{O as default};
