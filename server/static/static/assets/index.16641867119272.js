var C=Object.defineProperty;var y=Object.getOwnPropertySymbols;var D=Object.prototype.hasOwnProperty,T=Object.prototype.propertyIsEnumerable;var b=(a,e,i)=>e in a?C(a,e,{enumerable:!0,configurable:!0,writable:!0,value:i}):a[e]=i,_=(a,e)=>{for(var i in e||(e={}))D.call(e,i)&&b(a,i,e[i]);if(y)for(var i of y(e))T.call(e,i)&&b(a,i,e[i]);return a};import{_ as I,u as P,a as z,r as B,c as w,f as j,o as S,t as U,b as E,d as f,e as N,g as F,w as g,h as m,i as A,F as M,j as q,k as L,n as O,l as k,m as G}from"./index.1664186711927.js";import{A as R}from"./Api.1664186711927.js";var x=globalThis&&globalThis.__assign||function(){return(x=Object.assign||function(a){for(var e,i=1,r=arguments.length;i<r;i++)for(var t in e=arguments[i])Object.prototype.hasOwnProperty.call(e,t)&&(a[t]=e[t]);return a}).apply(this,arguments)},c=function(){function a(e,i,r){var t=this;this.target=e,this.endVal=i,this.options=r,this.version="2.0.8",this.defaults={startVal:0,decimalPlaces:0,duration:2,useEasing:!0,useGrouping:!0,smartEasingThreshold:999,smartEasingAmount:333,separator:",",decimal:".",prefix:"",suffix:""},this.finalEndVal=null,this.useEasing=!0,this.countDown=!1,this.error="",this.startVal=0,this.paused=!0,this.count=function(h){t.startTime||(t.startTime=h);var n=h-t.startTime;t.remaining=t.duration-n,t.useEasing?t.countDown?t.frameVal=t.startVal-t.easingFn(n,0,t.startVal-t.endVal,t.duration):t.frameVal=t.easingFn(n,t.startVal,t.endVal-t.startVal,t.duration):t.countDown?t.frameVal=t.startVal-(t.startVal-t.endVal)*(n/t.duration):t.frameVal=t.startVal+(t.endVal-t.startVal)*(n/t.duration),t.countDown?t.frameVal=t.frameVal<t.endVal?t.endVal:t.frameVal:t.frameVal=t.frameVal>t.endVal?t.endVal:t.frameVal,t.frameVal=Number(t.frameVal.toFixed(t.options.decimalPlaces)),t.printValue(t.frameVal),n<t.duration?t.rAF=requestAnimationFrame(t.count):t.finalEndVal!==null?t.update(t.finalEndVal):t.callback&&t.callback()},this.formatNumber=function(h){var n,o,s,l,V=h<0?"-":"";n=Math.abs(h).toFixed(t.options.decimalPlaces);var u=(n+="").split(".");if(o=u[0],s=u.length>1?t.options.decimal+u[1]:"",t.options.useGrouping){l="";for(var d=0,v=o.length;d<v;++d)d!==0&&d%3==0&&(l=t.options.separator+l),l=o[v-d-1]+l;o=l}return t.options.numerals&&t.options.numerals.length&&(o=o.replace(/[0-9]/g,function(p){return t.options.numerals[+p]}),s=s.replace(/[0-9]/g,function(p){return t.options.numerals[+p]})),V+t.options.prefix+o+s+t.options.suffix},this.easeOutExpo=function(h,n,o,s){return o*(1-Math.pow(2,-10*h/s))*1024/1023+n},this.options=x(x({},this.defaults),r),this.formattingFn=this.options.formattingFn?this.options.formattingFn:this.formatNumber,this.easingFn=this.options.easingFn?this.options.easingFn:this.easeOutExpo,this.startVal=this.validateValue(this.options.startVal),this.frameVal=this.startVal,this.endVal=this.validateValue(i),this.options.decimalPlaces=Math.max(this.options.decimalPlaces),this.resetDuration(),this.options.separator=String(this.options.separator),this.useEasing=this.options.useEasing,this.options.separator===""&&(this.options.useGrouping=!1),this.el=typeof e=="string"?document.getElementById(e):e,this.el?this.printValue(this.startVal):this.error="[CountUp] target is null or undefined"}return a.prototype.determineDirectionAndSmartEasing=function(){var e=this.finalEndVal?this.finalEndVal:this.endVal;this.countDown=this.startVal>e;var i=e-this.startVal;if(Math.abs(i)>this.options.smartEasingThreshold){this.finalEndVal=e;var r=this.countDown?1:-1;this.endVal=e+r*this.options.smartEasingAmount,this.duration=this.duration/2}else this.endVal=e,this.finalEndVal=null;this.finalEndVal?this.useEasing=!1:this.useEasing=this.options.useEasing},a.prototype.start=function(e){this.error||(this.callback=e,this.duration>0?(this.determineDirectionAndSmartEasing(),this.paused=!1,this.rAF=requestAnimationFrame(this.count)):this.printValue(this.endVal))},a.prototype.pauseResume=function(){this.paused?(this.startTime=null,this.duration=this.remaining,this.startVal=this.frameVal,this.determineDirectionAndSmartEasing(),this.rAF=requestAnimationFrame(this.count)):cancelAnimationFrame(this.rAF),this.paused=!this.paused},a.prototype.reset=function(){cancelAnimationFrame(this.rAF),this.paused=!0,this.resetDuration(),this.startVal=this.validateValue(this.options.startVal),this.frameVal=this.startVal,this.printValue(this.startVal)},a.prototype.update=function(e){cancelAnimationFrame(this.rAF),this.startTime=null,this.endVal=this.validateValue(e),this.endVal!==this.frameVal&&(this.startVal=this.frameVal,this.finalEndVal||this.resetDuration(),this.finalEndVal=null,this.determineDirectionAndSmartEasing(),this.rAF=requestAnimationFrame(this.count))},a.prototype.printValue=function(e){var i=this.formattingFn(e);this.el.tagName==="INPUT"?this.el.value=i:this.el.tagName==="text"||this.el.tagName==="tspan"?this.el.textContent=i:this.el.innerHTML=i},a.prototype.ensureNumber=function(e){return typeof e=="number"&&!isNaN(e)},a.prototype.validateValue=function(e){var i=Number(e);return this.ensureNumber(i)?i:(this.error="[CountUp] invalid start or end value: "+e,null)},a.prototype.resetDuration=function(){this.startTime=null,this.duration=1e3*Number(this.options.duration),this.remaining=this.duration},a}();const $={getIndexCount:R.create("/common/index/count","get")};const H={name:"HomePage",setup(){const a=P(),e=z(),i=B({topCardItemList:[{title:"\u9879\u76EE\u6570",id:"projectNum",color:"#FEBB50"},{title:"Linux\u673A\u5668\u6570",id:"machineNum",color:"#F95959"},{title:"\u6570\u636E\u5E93\u603B\u6570",id:"dbNum",color:"#8595F4"},{title:"redis\u603B\u6570",id:"redisNum",color:"#1abc9c"}]}),r=w(()=>j(new Date)),t=async()=>{const o=await $.getIndexCount.request();G(()=>{new c("projectNum",o.projectNum).start(),new c("machineNum",o.machineNum).start(),new c("dbNum",o.dbNum).start(),new c("redisNum",o.redisNum).start()})},h=o=>{switch(o.id){case"personal":{a.push("/personal");break}case"projectNum":{a.push("/ops/projects");break}case"machineNum":{a.push("/ops/machines");break}case"dbNum":{a.push("/ops/dbms/dbs");break}case"redisNum":{a.push("/ops/redis/manage");break}}};S(()=>{t()});const n=w(()=>e.state.userInfos.userInfos);return _({getUserInfos:n,currentTime:r,toPage:h},U(i))}},J={class:"home-container"},K={class:"flex-margin flex"},Q=["src"],W={class:"home-card-first-right ml15"},X={class:"flex-margin"},Y={class:"home-card-first-right-title"},Z=["onClick"],tt={class:"home-card-item-flex"},et={class:"home-card-item-title pb3"},at=["id"];function it(a,e,i,r,t,h){const n=E("el-col"),o=E("el-row");return f(),N("div",J,[F(o,{gutter:15},{default:g(()=>[F(n,{sm:6,class:"mb15"},{default:g(()=>[m("div",{onClick:e[0]||(e[0]=s=>r.toPage({id:"personal"})),class:"home-card-item home-card-first"},[m("div",K,[m("img",{src:r.getUserInfos.photo},null,8,Q),m("div",W,[m("div",X,[m("div",Y,A(`${r.currentTime}, ${r.getUserInfos.username}`),1)])])])])]),_:1}),(f(!0),N(M,null,q(a.topCardItemList,(s,l)=>(f(),L(n,{sm:3,class:"mb15",key:l},{default:g(()=>[m("div",{onClick:V=>r.toPage(s),class:"home-card-item home-card-item-box",style:k({background:s.color})},[m("div",tt,[m("div",et,A(s.title),1),m("div",{class:"home-card-item-title-num pb6",id:s.id},null,8,at)]),m("i",{class:O(s.icon),style:k({color:s.iconColor})},null,6)],12,Z)]),_:2},1024))),128))]),_:1})])}var st=I(H,[["render",it],["__scopeId","data-v-77501f64"]]);export{st as default};
