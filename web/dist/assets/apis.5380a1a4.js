/*! 
 Build based on gin-vue-admin 
 Time : 1670837959000 */
import{e}from"./api.ac6c4cd2.js";import{s as a,_ as t,r as s,J as o,b as l,o as r,c,d as n,e as d,w as i,h as p,i as u}from"./index.42b7f490.js";const h={class:"clearfix sticky-button"},f={class:"tree-content"},y={name:"Apis"},m=t(Object.assign(y,{props:{row:{default:function(){return{}},type:Object}},setup(t,{expose:y}){const m=t,v=s({children:"children",label:"description"}),b=s(""),k=s([]),I=s([]),w=s("");(async()=>{const t=(await e()).data.apis;k.value=j(t);const s=await(o={authorityId:m.row.authorityId},a({url:"/casbin/getPolicyPathByAuthorityId",method:"post",data:o}));var o;w.value=m.row.authorityId,I.value=[],s.data.paths&&s.data.paths.forEach((e=>{I.value.push("p:"+e.path+"m:"+e.method)}))})();const g=s(!1),x=()=>{g.value=!0},j=e=>{const a={};e&&e.forEach((e=>{e.onlyId="p:"+e.path+"m:"+e.method,Object.prototype.hasOwnProperty.call(a,e.apiGroup)?a[e.apiGroup].push(e):Object.assign(a,{[e.apiGroup]:[e]})}));const t=[];for(const s in a){const e={ID:s,description:s+"组",children:a[s]};t.push(e)}return t},O=s(null),C=async()=>{const e=O.value.getCheckedNodes(!0);var t=[];e&&e.forEach((e=>{var a={path:e.path,method:e.method};t.push(a)}));var s;0===(await(s={authorityId:w.value,casbinInfos:t},a({url:"/casbin/updateCasbin",method:"post",data:s}))).code&&u({type:"success",message:"api设置成功"})};y({needConfirm:g,enterAndNext:()=>{C()}});const _=(e,a)=>!e||-1!==a.description.indexOf(e);return o(b,(e=>{O.value.filter(e)})),(e,a)=>{const t=l("el-input"),s=l("el-button"),o=l("el-tree");return r(),c("div",null,[n("div",h,[d(t,{modelValue:b.value,"onUpdate:modelValue":a[0]||(a[0]=e=>b.value=e),class:"fitler",placeholder:"筛选"},null,8,["modelValue"]),d(s,{class:"fl-right",size:"small",type:"primary",onClick:C},{default:i((()=>[p("确 定")])),_:1})]),n("div",f,[d(o,{ref_key:"apiTree",ref:O,data:k.value,"default-checked-keys":I.value,props:v.value,"default-expand-all":"","highlight-current":"","node-key":"onlyId","show-checkbox":"","filter-node-method":_,onCheck:x},null,8,["data","default-checked-keys","props"])])])}}}),[["__scopeId","data-v-c620bda6"]]);export{m as default};
