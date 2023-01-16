/*! 
 Build based on gin-vue-admin 
 Time : 1673519692000 */
import{f as a}from"./sysDictionary.77e8fb14.js";import{B as t,r as e}from"./index.29e958a1.js";const i=t("dictionary",(()=>{const t=e({}),i=a=>{t.value={...t.value,...a}};return{dictionaryMap:t,setDictionaryMap:i,getDictionary:async e=>{if(t.value[e]&&t.value[e].length)return t.value[e];{const s=await a({type:e});if(0===s.code){const a={},r=[];return s.data.resysDictionary.sysDictionaryDetails&&s.data.resysDictionary.sysDictionaryDetails.forEach((a=>{r.push({label:a.label,value:a.value})})),a[s.data.resysDictionary.type]=r,i(a),t.value[e]}}}}}));export{i as u};
