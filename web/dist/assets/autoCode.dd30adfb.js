/*! 
 Build based on gin-vue-admin 
 Time : 1675228064000 */
import{s as t}from"./index.6ee4fc76.js";const a=a=>t({url:"/autoCode/preview",method:"post",data:a}),e=a=>t({url:"/autoCode/createTemp",method:"post",data:a,responseType:"blob"}),o=()=>t({url:"/autoCode/getDB",method:"get"}),d=a=>t({url:"/autoCode/getTables",method:"get",params:a}),s=a=>t({url:"/autoCode/getColumn",method:"get",params:a}),u=a=>t({url:"/autoCode/getSysHistory",method:"post",data:a}),r=a=>t({url:"/autoCode/rollback",method:"post",data:a}),l=a=>t({url:"/autoCode/getMeta",method:"post",data:a}),m=a=>t({url:"/autoCode/delSysHistory",method:"post",data:a}),p=a=>t({url:"/autoCode/createPackage",method:"post",data:a}),g=()=>t({url:"/autoCode/getPackage",method:"post"}),h=a=>t({url:"/autoCode/delPackage",method:"post",data:a}),C=a=>t({url:"/autoCode/createPlug",method:"post",data:a});export{s as a,g as b,e as c,o as d,l as e,u as f,d as g,m as h,p as i,h as j,C as k,a as p,r};
