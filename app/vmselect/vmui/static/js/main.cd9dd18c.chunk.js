(this.webpackJsonpvmui=this.webpackJsonpvmui||[]).push([[0],{179:function(e,t,n){},194:function(e,t){},199:function(e,t,n){},200:function(e,t,n){},207:function(e,t,n){"use strict";n.r(t);var a=n(0),r=n.n(a),c=n(27),i=n.n(c),o=(n(179),n(14)),s=n(279),l=n(277),u=n(2),j=Object(a.createContext)({showInfoMessage:function(){}}),d=function(){return Object(a.useContext)(j)},b=function(e){var t=e.children,n=Object(a.useState)({}),r=Object(o.a)(n,2),c=r[0],i=r[1],d=Object(a.useState)(!1),b=Object(o.a)(d,2),O=b[0],h=b[1],f=Object(a.useState)(void 0),p=Object(o.a)(f,2),x=p[0],v=p[1];Object(a.useEffect)((function(){x&&(i({message:x,key:(new Date).getTime()}),h(!0))}),[x]);return Object(u.jsxs)(j.Provider,{value:{showInfoMessage:v},children:[Object(u.jsx)(s.a,{open:O,autoHideDuration:4e3,onClose:function(e,t){"clickaway"!==t&&(v(void 0),h(!1))},children:Object(u.jsx)(l.a,{children:c.message})},c.key),t]})},O=n(270),h=n(260),f=n(274),p=n(86),x=n(253),v=n(262),m=n(241),y=n(243),g=n(244),E=n(283),_=n(159),T=n(135),S=n.n(T),A=n(13),C=n(12),w=n(36),I=n.n(w),R=n(131),U=n.n(R),D=n(132),H=n.n(D);I.a.extend(U.a),I.a.extend(H.a);var k,M=window.screen.availWidth/2,L=1e3,N=1578e8,Y="YYYY-MM-DD[T]HH:mm:ss",q=[{long:"days",short:"d",possible:"day"},{long:"weeks",short:"w",possible:"week"},{long:"months",short:"M",possible:"mon"},{long:"years",short:"y",possible:"year"},{long:"hours",short:"h",possible:"hour"},{long:"minutes",short:"m",possible:"min"},{long:"seconds",short:"s",possible:"sec"},{long:"milliseconds",short:"ms",possible:"millisecond"}],B=q.map((function(e){return e.short})),P=function(e){var t=e.match(/\d+/g),n=e.match(/[a-zA-Z]+/g);if(n&&t&&B.includes(n[0]))return Object(A.a)({},n[0],t[0])},Q=function(e,t){var n=(t||new Date).valueOf()/1e3,a=e.trim().split(" ").reduce((function(e,t){var n=P(t);return n?Object(C.a)(Object(C.a)({},e),n):Object(C.a)({},e)}),{}),r=I.a.duration(a).asSeconds();return{start:n-r,end:n,step:Math.ceil(r/M),date:z(t||new Date)}},z=function(e){return I()(e).utc().format(Y)},F=function(e){var t=Math.floor(e%1e3),n=Math.floor(e/1e3%60),a=Math.floor(e/1e3/60%60),r=Math.floor(e/1e3/3600%24),c=Math.floor(e/864e5),i=["d","h","m","s","ms"];return[c,r,a,n,t].map((function(e,t){return e?"".concat(e).concat(i[t]):""})).filter((function(e){return e})).join(" ")},G=function(e){return new Date(1e3*e)},W=function(e,t){t?window.localStorage.setItem(e,JSON.stringify({value:t})):X([e])},V=function(e){var t=window.localStorage.getItem(e);if(null!==t)try{var n;return null===(n=JSON.parse(t))||void 0===n?void 0:n.value}catch(a){return t}},X=function(e){return e.forEach((function(e){return window.localStorage.removeItem(e)}))},J=["BASIC_AUTH_DATA","BEARER_AUTH_DATA"],K=n(133),Z=n.n(K),$=n(117),ee=n.n($),te={query:"g0.expr","time.duration":"g0.range_input","time.period.date":"g0.end_input","time.period.step":"g0.step_input",stacked:"g0.stacked"},ne=function(e){var t=new Map(Object.entries(te)),n=[];t.forEach((function(t,a){encodeURIComponent(t);var r=ee()(e,a,"");if(r){var c=encodeURIComponent(r);n.push("".concat(t,"=").concat(c))}})),function(e){var t=window;if(t){var n="".concat(t.location.protocol,"//").concat(t.location.host).concat(t.location.pathname,"?").concat(e);t.history.pushState({path:n},"",n)}}(n.join("&"))},ae=function(e,t){var n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:window.location.search,a=Z.a.parse(n,{ignoreQueryPrefix:!0});return ee()(a,e,t||"")},re=ae("g0.range_input","1h"),ce=(k=ae("g0.end_input",new Date(I()().utc().format(Y))),I()(k).utcOffset(0,!0).local().format(Y)),ie=ae("g0.expr",V("LAST_QUERY")||"\n"),oe={serverUrl:window.location.href.replace(/\/(?:prometheus\/)?(?:graph|vmui)\/.*/,"/prometheus/"),displayType:"chart",query:ie,queryHistory:{index:0,values:[ie]},time:{duration:re,period:Q(re,new Date(ce))},queryControls:{autoRefresh:!1,autocomplete:V("AUTOCOMPLETE")||!1,nocache:V("NO_CACHE")||!1}};function se(e,t){switch(t.type){case"SET_DISPLAY_TYPE":return Object(C.a)(Object(C.a)({},e),{},{displayType:t.payload});case"SET_SERVER":return Object(C.a)(Object(C.a)({},e),{},{serverUrl:t.payload});case"SET_QUERY":return Object(C.a)(Object(C.a)({},e),{},{query:t.payload});case"SET_QUERY_HISTORY_INDEX":return Object(C.a)(Object(C.a)({},e),{},{queryHistory:Object(C.a)(Object(C.a)({},e.queryHistory),{},{index:t.payload})});case"SET_QUERY_HISTORY_VALUES":return Object(C.a)(Object(C.a)({},e),{},{queryHistory:Object(C.a)(Object(C.a)({},e.queryHistory),{},{values:t.payload})});case"SET_DURATION":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{duration:t.payload,period:Q(t.payload,G(e.time.period.end))})});case"SET_UNTIL":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{period:Q(e.time.duration,t.payload)})});case"SET_PERIOD":var n=function(e){var t=e.to.valueOf()-e.from.valueOf();return F(t)}(t.payload);return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{autoRefresh:!1}),time:Object(C.a)(Object(C.a)({},e.time),{},{duration:n,period:Q(n,t.payload.to)})});case"TOGGLE_AUTOREFRESH":return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{autoRefresh:!e.queryControls.autoRefresh})});case"TOGGLE_AUTOCOMPLETE":return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{autocomplete:!e.queryControls.autocomplete})});case"NO_CACHE":return Object(C.a)(Object(C.a)({},e),{},{queryControls:Object(C.a)(Object(C.a)({},e.queryControls),{},{nocache:!e.queryControls.nocache})});case"RUN_QUERY":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{period:Q(e.time.duration,G(e.time.period.end))})});case"RUN_QUERY_TO_NOW":return Object(C.a)(Object(C.a)({},e),{},{time:Object(C.a)(Object(C.a)({},e.time),{},{period:Q(e.time.duration)})});default:throw new Error}}var le=Object(a.createContext)({}),ue=function(){return Object(a.useContext)(le).state},je=function(){return Object(a.useContext)(le).dispatch},de=Object.entries(oe).reduce((function(e,t){var n=Object(o.a)(t,2),a=n[0],r=n[1];return Object(C.a)(Object(C.a)({},e),{},Object(A.a)({},a,ae(a)||r))}),{}),be=function(e){var t=e.children,n=Object(a.useReducer)(se,de),r=Object(o.a)(n,2),c=r[0],i=r[1];Object(a.useEffect)((function(){ne(c)}),[c]);var s=Object(a.useMemo)((function(){return{state:c,dispatch:i}}),[c,i]);return Object(u.jsx)(le.Provider,{value:s,children:t})},Oe=function(e){return Object(u.jsxs)(f.a,{position:"relative",display:"inline-flex",children:[Object(u.jsx)(m.a,Object(C.a)({variant:"determinate"},e)),Object(u.jsx)(f.a,{top:0,left:0,bottom:0,right:0,position:"absolute",display:"flex",alignItems:"center",justifyContent:"center",children:Object(u.jsx)(p.a,{variant:"caption",component:"div",children:"".concat(e.label,"s")})})]})},he=n(242),fe=Object(he.a)({colorizing:{color:"white"}}),pe=function(){var e=fe(),t=je(),n=ue().queryControls.autoRefresh,c=Object(a.useState)(5),i=Object(o.a)(c,2),s=i[0],l=i[1],j=Object(a.useState)(),d=Object(o.a)(j,2),b=d[0],O=d[1],h=r.a.useState(100),p=Object(o.a)(h,2),x=p[0],v=p[1];Object(a.useEffect)((function(){var e;return n&&(O((new Date).valueOf()),e=setInterval((function(){O((new Date).valueOf()),t({type:"RUN_QUERY_TO_NOW"})}),1e3*s)),function(){e&&clearInterval(e)}}),[s,n]),Object(a.useEffect)((function(){var e=setInterval((function(){if(n&&b){var e=((new Date).valueOf()-b)/1e3,t=Math.floor(e/s*100);v(t)}}),16);return function(){clearInterval(e)}}),[n,b,s]);var m=function(){l((function(e){switch(e){case 1:return 2;default:return 5;case 5:return 1}}))};return Object(u.jsxs)(f.a,{display:"flex",alignItems:"center",children:[Object(u.jsx)(y.a,{control:Object(u.jsx)(g.a,{size:"small",className:e.colorizing,checked:n,onChange:function(){t({type:"TOGGLE_AUTOREFRESH"})}}),label:"Auto-refresh"}),n&&Object(u.jsxs)(u.Fragment,{children:[Object(u.jsx)(Oe,{className:e.colorizing,label:s,value:x,onClick:function(){m()}}),Object(u.jsx)(E.a,{title:"Change delay refresh",children:Object(u.jsx)(f.a,{ml:1,children:Object(u.jsx)(_.a,{onClick:function(){m()},children:Object(u.jsx)(S.a,{style:{color:"white"}})})})})]})]})},xe=n(138),ve=n.n(xe),me=n(136),ye=n.n(me),ge=n(137),Ee=n.n(ge),_e=n(246),Te=n(286),Se=n(10),Ae=Object(Se.a)({root:{padding:6,color:"white","&.Mui-selected":{color:"white"}}})(_e.a),Ce=function(){var e=ue().displayType,t=je();return Object(u.jsxs)(Te.a,{value:e,exclusive:!0,onChange:function(n,a){return t({type:"SET_DISPLAY_TYPE",payload:null!==a&&void 0!==a?a:e})},children:[Object(u.jsxs)(Ae,{value:"chart","aria-label":"display as chart",children:[Object(u.jsx)(ye.a,{}),"\xa0Query Range as Chart"]}),Object(u.jsxs)(Ae,{value:"code","aria-label":"display as code",children:[Object(u.jsx)(Ee.a,{}),"\xa0Instant Query as JSON"]}),Object(u.jsxs)(Ae,{value:"table","aria-label":"display as table",children:[Object(u.jsx)(ve.a,{}),"\xa0Instant Query as Table"]})]})},we=n(139),Ie=n.n(we),Re=(n(198),n(140)),Ue=n.n(Re),De=(n(199),n(200),{yaxis:{limits:{enable:!1,range:[0,0]}}});function He(e,t){switch(t.type){case"TOGGLE_ENABLE_YAXIS_LIMITS":return Object(C.a)(Object(C.a)({},e),{},{yaxis:Object(C.a)(Object(C.a)({},e.yaxis),{},{limits:Object(C.a)(Object(C.a)({},e.yaxis.limits),{},{enable:!e.yaxis.limits.enable})})});case"SET_YAXIS_LIMITS":return Object(C.a)(Object(C.a)({},e),{},{yaxis:Object(C.a)(Object(C.a)({},e.yaxis),{},{limits:Object(C.a)(Object(C.a)({},e.yaxis.limits),{},{range:t.payload})})});default:throw new Error}}var ke=Object(a.createContext)({}),Me=function(){return Object(a.useContext)(ke).state},Le=function(){return Object(a.useContext)(ke).dispatch},Ne=function(e){var t=e.children,n=Object(a.useReducer)(He,De),r=Object(o.a)(n,2),c=r[0],i=r[1],s=Object(a.useMemo)((function(){return{state:c,dispatch:i}}),[c,i]);return Object(u.jsx)(ke.Provider,{value:s,children:t})},Ye=n(26),qe=function(e){for(var t=0,n=0;n<e.length;n++)t=e.charCodeAt(n)+((t<<5)-t);for(var a="#",r=0;r<3;r++){a+=("00"+(t>>8*r&255).toString(16)).substr(-2)}return a},Be=n(106),Pe=["__name__"],Qe=function(e){if(0===Object.keys(e.metric).length)return"Query result";var t=e.metric,n=t.__name__,a=Object(Be.a)(t,Pe);return"".concat(n||""," {").concat(Object.entries(a).map((function(e){return"".concat(e[0],": ").concat(e[1])})).join(", "),"}")},ze=function(e){var t,n=e.u,a=e.tooltipIdx,r=e.data,c=e.series,i=e.tooltip,o=e.tooltipOffset,s=a.seriesIdx,l=a.dataIdx,u=n.data[s][l],j=n.data[0][l],d=(null===(t=r[s-1])||void 0===t?void 0:t.metric)||{},b=qe(c[s].label||""),O=n.over.getBoundingClientRect(),h=O.width,f=O.height,p=n.valToPos(u||0,"y"),x=n.valToPos(j,"x"),v=i.getBoundingClientRect(),m=v.width,y=v.height,g=x+m>=h,E=p+y>=f;i.style.display="grid",i.style.top="".concat(o.top+p+10-(E?y+10:0),"px"),i.style.left="".concat(o.left+x+10-(g?m+20:0),"px");var _=I()(new Date(1e3*j)).format("YYYY-MM-DD HH:mm:ss:SSS (Z)"),T=Object.keys(d).filter((function(e){return"__name__"!==e})).map((function(e){return"<div><b>".concat(e,"</b>: ").concat(d[e],"</div>")})).join(""),S='<div class="u-tooltip__marker" style="background: '.concat(b,'"></div>');i.innerHTML="<div>".concat(_,'</div>\n                       <div class="u-tooltip-data">\n                         ').concat(S).concat(d.__name__||"",": <b>").concat(u,'</b>\n                       </div>\n                       <div class="u-tooltip__info">').concat(T,"</div>")},Fe=function(e){var t=e.data,n=void 0===t?[]:t,r=je(),c=ue().time.period,i=Object(a.useState)({min:c.start,max:c.end}),s=Object(o.a)(i,2),l=s[0],j=s[1],d=Object(a.useRef)(null),b=Object(a.useState)(!1),O=Object(o.a)(b,2),h=O[0],f=O[1],p=Object(a.useState)(0),x=Object(o.a)(p,2),v=x[0],m=x[1],y={seriesIdx:1,dataIdx:0},g={left:0,top:0},E=Me().yaxis,_=Le(),T=Object(a.useMemo)((function(){for(var e=function(e){var t=e.map((function(e){return e.values.map((function(e){return e[0]}))})).flat().sort((function(e,t){return e-t}));return[t[0],t[t.length-1]]}(n),t=Object(o.a)(e,2),a=t[0],r=t[1],i=[],s=a;s<r;s+=c.step||1)i.push(s);return i}),[n]),S=Object(a.useMemo)((function(){return function(e){return[{}].concat(Object(Ye.a)(e.map((function(e){return{label:Qe(e),width:1.5,stroke:qe(Qe(e))}}))))}(n)}),[n]),A=Object(a.useMemo)((function(){return function(e,t){return[t].concat(Object(Ye.a)(e.map((function(e){return t.map((function(t){var n=e.values.find((function(e){return e[0]===t}));return n?+n[1]:null}))}))))}(n,T)}),[n]),C=document.createElement("div");C.className="u-tooltip";Object(a.useEffect)((function(){var e;e=function(e){var t=e.map((function(e){return e.values.map((function(e){return+e[1]}))})).flat().sort((function(e,t){return e-t}));return[t[0],t[t.length-1]]}(n),E.limits.enable&&!E.limits.range.every((function(e){return!e}))||_({type:"SET_YAXIS_LIMITS",payload:e})}),[n]),Object(a.useEffect)((function(){j({min:c.start,max:c.end})}),[c]),Object(a.useEffect)((function(){var e=(c.end-c.start)/3,t=e/(l.max-l.min);(l.max>c.end+e||l.min<c.start-e||t>=.7)&&r({type:"SET_PERIOD",payload:{from:new Date(1e3*l.min),to:new Date(1e3*l.max)}})}),[l]);var w={width:d.current?d.current.offsetWidth:400,height:500,series:S,plugins:[{hooks:{ready:function(e){var t;g.left=parseFloat(e.over.style.left),g.top=parseFloat(e.over.style.top),null===(t=e.root.querySelector(".u-wrap"))||void 0===t||t.appendChild(C),e.over.addEventListener("mousedown",(function(t){if(0===t.button){f(!0),t.preventDefault();var n=t.clientX,a=function(t){t.preventDefault();var a=(e.posToVal(1,"x")-e.posToVal(0,"x"))*(t.clientX-n),r=(e.scales.x.min||1)-a,c=(e.scales.x.max||1)-a;e.setScale("x",{min:r,max:c}),j({min:r,max:c})};document.addEventListener("mousemove",a),document.addEventListener("mouseup",(function e(){f(!1),document.removeEventListener("mousemove",a),document.removeEventListener("mouseup",e)}))}})),e.over.addEventListener("wheel",(function(t){if(t.ctrlKey||t.metaKey){t.preventDefault();var n=e.over.getBoundingClientRect().width;e.cursor.left&&e.cursor.left>0&&m(e.cursor.left);var a=e.posToVal(v,"x"),r=(e.scales.x.max||0)-(e.scales.x.min||0),c=t.deltaY<0?.85*r:r/.85,i=a-v/n*c,o=i+c;e.batch((function(){e.setScale("x",{min:i,max:o}),j({min:i,max:o})}))}}))},setCursor:function(e){y.dataIdx!==e.cursor.idx&&(y.dataIdx=e.cursor.idx||0,y.seriesIdx&&y.dataIdx&&ze({u:e,tooltipIdx:y,data:n,series:S,tooltip:C,tooltipOffset:g}))},setSeries:function(e,t){y.seriesIdx!==t&&(y.seriesIdx=t||0,t&&y.dataIdx?ze({u:e,tooltipIdx:y,data:n,series:S,tooltip:C,tooltipOffset:g}):C.style.display="none")}}}],cursor:{drag:{x:!1,y:!1},focus:{prox:30}},axes:[{space:80},{show:!0,font:"10px Arial",values:function(e,t){return t.map((function(e){return e>1e3?Ue()(e).format("0.0a"):e}))}}],scales:{x:{range:function(){return[l.min,l.max]}},y:{range:function(e,t,n){return E.limits.enable?E.limits.range:[t,n]}}}};return Object(u.jsx)("div",{ref:d,style:{pointerEvents:h?"none":"auto"},children:A&&Object(u.jsx)(Ie.a,{options:w,data:A})})},Ge=function(e){var t=e.data,n=void 0===t?[]:t;return Object(u.jsx)(u.Fragment,{children:n.length>0?Object(u.jsx)(Fe,{data:n}):Object(u.jsx)("div",{style:{textAlign:"center"},children:"No data to show"})})},We=n(247),Ve=n(157),Xe=n(248),Je=n(249),Ke=n(250),Ze=n(251),$e=n(252),et=Object(he.a)({deemphasized:{opacity:.4}}),tt=function(e){var t=e.data,n=et(),r=function(e){return Object(a.useMemo)((function(){var t={};return e.forEach((function(e){return Object.entries(e.metric).forEach((function(e){return t[e[0]]?t[e[0]].options.add(e[1]):t[e[0]]={options:new Set([e[1]])}}))})),Object.entries(t).map((function(e){return{key:e[0],variations:e[1].options.size}})).sort((function(e,t){return e.variations-t.variations}))}),[e])}(t),c=Object(a.useMemo)((function(){return null===t||void 0===t?void 0:t.map((function(e){return{metadata:r.map((function(t){return e.metric[t.key]||"-"})),value:e.value[1]}}))}),[r,t]);return Object(u.jsx)(u.Fragment,{children:c.length>0?Object(u.jsx)(We.a,{component:Ve.a,children:Object(u.jsxs)(Xe.a,{"aria-label":"simple table",children:[Object(u.jsx)(Je.a,{children:Object(u.jsxs)(Ke.a,{children:[r.map((function(e,t){return Object(u.jsx)(Ze.a,{style:{textTransform:"capitalize"},children:e.key},t)})),Object(u.jsx)(Ze.a,{align:"right",children:"Value"})]})}),Object(u.jsx)($e.a,{children:c.map((function(e,t){return Object(u.jsxs)(Ke.a,{children:[e.metadata.map((function(e,a){var r=c[t-1]&&c[t-1].metadata[a];return Object(u.jsx)(Ze.a,{className:r===e?n.deemphasized:void 0,children:e},a)})),Object(u.jsx)(Ze.a,{align:"right",children:e.value})]},t)}))})]})}):Object(u.jsx)("div",{style:{textAlign:"center"},children:"No data to show"})})},nt=n(280),at=n(268),rt=n(269),ct=n(265),it=n(275),ot=n(4),st=n(9),lt=n(103),ut=n(152),jt=n(154),dt=function(e){var t=e.query,n=e.queryHistory,r=e.setHistoryIndex,c=e.setQuery,i=e.runQuery,s=e.server,l=e.oneLiner,j=void 0!==l&&l,d=e.autocomplete,b=Object(a.useRef)(null),O=Object(a.useState)(),h=Object(o.a)(O,2),f=h[0],p=h[1];Object(a.useEffect)((function(){return b.current&&p(new st.d({parent:b.current})),function(){return null===f||void 0===f?void 0:f.destroy()}}),[]),Object(a.useEffect)((function(){var e=new ut.a;e.activateCompletion(d),e.setComplete({remote:{url:s}});var n=st.d.updateListener.of((function(e){e.docChanged&&c(e.state.doc.toJSON().map((function(e){return e.trim()})).join(""))}));null===f||void 0===f||f.setState(ot.f.create({doc:t,extensions:[jt.a,st.k.of(lt.a),n,e.asExtension()]}))}),[s,f,d,n]);return Object(u.jsx)(u.Fragment,{children:Object(u.jsx)("div",{ref:b,className:j?"one-line-scroll":"multi-line-scroll",onKeyUp:function(e){var t=e.key,n=e.ctrlKey,a=e.metaKey,c=n||a;"Enter"===t&&c?i():"ArrowUp"===t&&c?r(-1):"ArrowDown"===t&&c&&r(1)}})})},bt=n(256),Ot=n(273),ht=function(){return Object(u.jsx)(We.a,{component:Ve.a,children:Object(u.jsxs)(Xe.a,{"aria-label":"simple table",size:"small",children:[Object(u.jsx)(Je.a,{children:Object(u.jsxs)(Ke.a,{children:[Object(u.jsx)(Ze.a,{children:"Long"}),Object(u.jsx)(Ze.a,{children:"Short"})]})}),Object(u.jsx)($e.a,{children:q.map((function(e,t){return Object(u.jsxs)(Ke.a,{children:[Object(u.jsx)(Ze.a,{component:"th",scope:"row",children:e.long}),Object(u.jsx)(Ze.a,{children:e.short})]},t)}))})]})})},ft=Object(he.a)({inlineBtn:{"&:hover":{cursor:"pointer"}}}),pt=function(e){var t=e.handler,n=e.text,a=ft();return Object(u.jsx)(x.a,{component:"span",className:a.inlineBtn,onClick:t,children:n})},xt=function(e){var t=e.setDuration,n=Object(a.useState)(!1),c=Object(o.a)(n,2),i=c[0],s=c[1],l=r.a.useState(null),j=Object(o.a)(l,2),d=j[0],b=j[1],O=Object(a.useState)(),h=Object(o.a)(O,2),x=h[0],v=h[1],m=ue().time,y=m.period.end,g=m.duration,E=je(),_=Object(a.useState)(g),T=Object(o.a)(_,2),S=T[0],A=T[1];Object(a.useEffect)((function(){A(g)}),[g]),Object(a.useEffect)((function(){v(function(e){return I()(e).format(Y)}(G(y)))}),[y]),Object(a.useEffect)((function(){if(!i){var e=function(e){var t=e.trim().split(" ").reduce((function(e,t){var n=P(t);return n?Object(C.a)(Object(C.a)({},e),n):Object(C.a)({},e)}),{}),n=I.a.duration(t).asMilliseconds();return n<L?F(L):n>N?F(N):e}(S);A(e),t(e)}}),[S,i]);var w=function(){b(null)},R=Boolean(d);return Object(u.jsxs)(f.a,{m:1,flexDirection:"row",display:"flex",children:[Object(u.jsxs)(f.a,{px:1,children:[Object(u.jsx)(f.a,{children:Object(u.jsx)(it.a,{label:"Duration",value:S,onChange:function(e){A(e.target.value)},fullWidth:!0,onBlur:function(){s(!1)},onFocus:function(){s(!0)}})}),Object(u.jsx)(f.a,{my:2,children:Object(u.jsxs)(p.a,{variant:"body2",children:["Possible options",Object(u.jsx)("span",{"aria-owns":R?"mouse-over-popover":void 0,"aria-haspopup":"true",style:{cursor:"pointer"},onMouseEnter:function(e){b(e.currentTarget)},onMouseLeave:w,children:":\xa0"}),Object(u.jsx)(bt.a,{open:R,anchorEl:d,anchorOrigin:{vertical:"bottom",horizontal:"left"},transformOrigin:{vertical:"top",horizontal:"left"},style:{pointerEvents:"none"},onClose:w,disableRestoreFocus:!0,children:Object(u.jsx)(ht,{})}),Object(u.jsx)(pt,{handler:function(){return A("5m")},text:"5m"}),",\xa0",Object(u.jsx)(pt,{handler:function(){return A("1h")},text:"1h"}),",\xa0",Object(u.jsx)(pt,{handler:function(){return A("1h 30m")},text:"1h 30m"})]})})]}),Object(u.jsxs)(f.a,{px:1,children:[Object(u.jsx)(f.a,{children:Object(u.jsx)(Ot.a,{variant:"inline",ampm:!1,label:"Until",value:x,onChange:function(e){return E({type:"SET_UNTIL",payload:e})},onError:console.log,format:"DD/MM/YYYY HH:mm:ss"})}),Object(u.jsx)(f.a,{my:2,children:Object(u.jsxs)(p.a,{variant:"body2",children:["Will be changed to current time for auto-refresh mode.\xa0",Object(u.jsx)(pt,{handler:function(){return E({type:"RUN_QUERY_TO_NOW"})},text:"Switch to now"})]})})]})]})},vt=n(144),mt=n.n(vt),yt=n(146),gt=n.n(yt),Et=n(266),_t=n(285),Tt=n(281),St=n(263),At=n(267),Ct=n(276),wt=n(258),It=n(255),Rt=n(284),Ut=n(254),Dt=n(261),Ht=n(278),kt=n(257),Mt=n(264),Lt=n(259),Nt=["children","value","index"],Yt=function(e){var t=e.children,n=e.value,a=e.index,r=Object(Be.a)(e,Nt);return Object(u.jsx)("div",Object(C.a)(Object(C.a)({role:"tabpanel",hidden:n!==a,id:"auth-config-tabpanel-".concat(a),"aria-labelledby":"auth-config-tab-".concat(a)},r),{},{children:n===a&&Object(u.jsx)(f.a,{py:2,children:t})}))},qt=n(142),Bt=n.n(qt),Pt=n(143),Qt=n.n(Pt),zt={authMethod:"NO_AUTH",saveAuthLocally:!1},Ft=V("AUTH_TYPE"),Gt=V("BASIC_AUTH_DATA"),Wt=V("BEARER_AUTH_DATA"),Vt=Object(C.a)(Object(C.a)({},zt),{},{authMethod:Ft||zt.authMethod,basicData:Gt,bearerData:Wt,saveAuthLocally:!(!Gt&&!Wt)}),Xt=function(){X(J)};function Jt(e,t){switch(t.type){case"SET_BASIC_AUTH":return t.payload.checkbox?W("BASIC_AUTH_DATA",t.payload.value):Xt(),W("AUTH_TYPE","BASIC_AUTH"),Object(C.a)(Object(C.a)({},e),{},{authMethod:"BASIC_AUTH",basicData:t.payload.value});case"SET_BEARER_AUTH":return t.payload.checkbox?W("BEARER_AUTH_DATA",t.payload.value):Xt(),W("AUTH_TYPE","BEARER_AUTH"),Object(C.a)(Object(C.a)({},e),{},{authMethod:"BEARER_AUTH",bearerData:t.payload.value});case"SET_NO_AUTH":return!t.payload.checkbox&&Xt(),W("AUTH_TYPE","NO_AUTH"),Object(C.a)(Object(C.a)({},e),{},{authMethod:"NO_AUTH"});default:throw new Error}}var Kt=Object(a.createContext)({}),Zt=function(){return Object(a.useContext)(Kt).state},$t=function(e){var t=e.children,n=Object(a.useReducer)(Jt,Vt),r=Object(o.a)(n,2),c=r[0],i=r[1],s=Object(a.useMemo)((function(){return{state:c,dispatch:i}}),[c,i]);return Object(u.jsx)(Kt.Provider,{value:s,children:t})},en=Object(he.a)((function(){return Object(Tt.a)({tabsContent:{height:"200px"}})})),tn="Bearer ",nn=[{title:"No auth",id:"NO_AUTH"},{title:"Basic Auth",id:"BASIC_AUTH"},{title:"Bearer Token",id:"BEARER_AUTH"}],an=function(e){var t=en(),n=e.onClose,r=e.open,c=Zt(),i=c.saveAuthLocally,s=c.basicData,l=c.bearerData,j=c.authMethod,d=Object(a.useContext)(Kt).dispatch,b=Object(a.useState)(i),O=Object(o.a)(b,2),h=O[0],x=O[1],v=Object(a.useState)(s||{password:"",login:""}),m=Object(o.a)(v,2),g=m[0],E=m[1],_=Object(a.useState)((null===l||void 0===l?void 0:l.token)||tn),T=Object(o.a)(_,2),S=T[0],A=T[1],w=Object(a.useState)(nn.findIndex((function(e){return e.id===j}))||0),I=Object(o.a)(w,2),R=I[0],U=I[1],D=function(){n()};return Object(u.jsxs)(_t.a,{onClose:D,"aria-labelledby":"simple-dialog-title",open:r,children:[Object(u.jsx)(Et.a,{id:"simple-dialog-title",children:"Request Auth Settings"}),Object(u.jsxs)(St.a,{children:[Object(u.jsx)(At.a,{children:"This affects Authorization header sent to the server you specify. Not shown in URL and can be optionally stored on a client side"}),Object(u.jsx)(Ct.a,{value:R,onChange:function(e,t){U(t)},indicatorColor:"primary",textColor:"primary",children:nn.map((function(e){return Object(u.jsx)(wt.a,{label:e.title},e.id)}))}),Object(u.jsxs)(f.a,{p:0,display:"flex",flexDirection:"column",className:t.tabsContent,children:[Object(u.jsxs)(f.a,{flexGrow:1,children:[Object(u.jsx)(Yt,{value:R,index:0,children:Object(u.jsx)(p.a,{style:{fontStyle:"italic"},children:"No Authorization Header"})}),Object(u.jsxs)(Yt,{value:R,index:1,children:[Object(u.jsxs)(It.a,{margin:"dense",fullWidth:!0,children:[Object(u.jsx)(Rt.a,{htmlFor:"basic-login",children:"User"}),Object(u.jsx)(Ut.a,{id:"basic-login",startAdornment:Object(u.jsx)(Dt.a,{position:"start",children:Object(u.jsx)(Bt.a,{})}),required:!0,onChange:function(e){return E((function(t){return Object(C.a)(Object(C.a)({},t),{},{login:e.target.value||""})}))},value:(null===g||void 0===g?void 0:g.login)||""})]}),Object(u.jsxs)(It.a,{margin:"dense",fullWidth:!0,children:[Object(u.jsx)(Rt.a,{htmlFor:"basic-pass",children:"Password"}),Object(u.jsx)(Ut.a,{id:"basic-pass",startAdornment:Object(u.jsx)(Dt.a,{position:"start",children:Object(u.jsx)(Qt.a,{})}),onChange:function(e){return E((function(t){return Object(C.a)(Object(C.a)({},t),{},{password:e.target.value||""})}))},value:(null===g||void 0===g?void 0:g.password)||""})]})]}),Object(u.jsx)(Yt,{value:R,index:2,children:Object(u.jsx)(it.a,{id:"bearer-auth",label:"Bearer token",multiline:!0,fullWidth:!0,value:S,onChange:function(e){var t=e.target.value;t.startsWith(tn)?A(t):A(tn)},InputProps:{onPaste:function(e){var t=e.clipboardData.getData("text/plain");t.startsWith(tn)?A(t):A(tn+t),e.preventDefault()}},maxRows:6})})]}),Object(u.jsxs)(It.a,{children:[Object(u.jsx)(y.a,{control:Object(u.jsx)(Ht.a,{checked:h,onChange:function(){return x((function(e){return!e}))},name:"checkedB",color:"primary"}),label:"Persist Auth Data Locally"}),Object(u.jsx)(kt.a,{children:h?"Auth Data and the Selected method will be saved to LocalStorage":"Auth Data won't be saved. All previously saved Auth Data will be removed"})]})]})]}),Object(u.jsx)(Mt.a,{children:Object(u.jsx)(Lt.a,{onClick:function(){switch(R){case 0:d({type:"SET_NO_AUTH",payload:{checkbox:h}});break;case 1:d({type:"SET_BASIC_AUTH",payload:{checkbox:h,value:g}});break;case 2:d({type:"SET_BEARER_AUTH",payload:{checkbox:h,value:{token:S}}})}D()},color:"primary",children:"Apply"})})]})},rn=n(145),cn=n.n(rn),on=n(245),sn=n(119),ln=n.n(sn),un=function(){var e=ue(),t=e.serverUrl,n=e.query,r=e.queryHistory,c=e.time.duration,i=e.queryControls,s=i.autocomplete,l=i.nocache,j=je(),d=Me().yaxis,b=Le(),O=Object(a.useState)(!1),h=Object(o.a)(O,2),x=h[0],v=h[1],m=Object(a.useState)(!0),T=Object(o.a)(m,2),S=T[0],A=T[1],C=Object(a.useRef)(null),w=function(){var e=r.values;n!==e[e.length-1]&&(j({type:"RUN_QUERY"}),j({type:"SET_QUERY_HISTORY_INDEX",payload:e.length}),j({type:"SET_QUERY_HISTORY_VALUES",payload:[].concat(Object(Ye.a)(e),[n])}))},I=function(e){n!==e&&j({type:"SET_QUERY",payload:e})};return Object(u.jsxs)(u.Fragment,{children:[Object(u.jsxs)(nt.a,{expanded:S,onChange:function(){return A((function(e){return!e}))},children:[Object(u.jsxs)(at.a,{expandIcon:Object(u.jsx)(mt.a,{}),"aria-controls":"panel1a-content",id:"panel1a-header",children:[Object(u.jsx)(f.a,{mr:2,children:Object(u.jsx)(p.a,{variant:"h6",component:"h2",children:"Query Configuration"})}),Object(u.jsx)(f.a,{flexGrow:1,onClick:function(e){return e.stopPropagation()},onFocusCapture:function(e){return e.stopPropagation()},children:Object(u.jsx)(on.a,{disablePortal:!S,container:C.current,children:Object(u.jsxs)(f.a,{display:"flex",alignItems:"center",children:[Object(u.jsx)(f.a,{width:"100%",children:Object(u.jsx)(dt,{server:t,query:n,oneLiner:!S,autocomplete:s,queryHistory:r,setHistoryIndex:function(e){var t=r.index+e;t<-1||t>r.values.length||(j({type:"SET_QUERY_HISTORY_INDEX",payload:t}),I(r.values[t]||""))},runQuery:w,setQuery:I})}),Object(u.jsx)(E.a,{title:"Execute Query",children:Object(u.jsx)(_.a,{onClick:w,children:Object(u.jsx)(cn.a,{})})})]})})})]}),Object(u.jsx)(rt.a,{children:Object(u.jsxs)(ct.a,{container:!0,spacing:2,children:[Object(u.jsx)(ct.a,{item:!0,xs:12,md:6,children:Object(u.jsxs)(f.a,{display:"grid",gridGap:16,children:[Object(u.jsxs)(f.a,{display:"flex",alignItems:"center",children:[Object(u.jsx)(it.a,{variant:"outlined",fullWidth:!0,label:"Server URL",value:t,inputProps:{style:{fontFamily:"Monospace"}},onChange:function(e){var t=e.target.value;j({type:"SET_SERVER",payload:t})}}),Object(u.jsx)(f.a,{children:Object(u.jsx)(E.a,{title:"Request Auth Settings",children:Object(u.jsx)(_.a,{onClick:function(){return v(!0)},children:Object(u.jsx)(gt.a,{})})})})]}),Object(u.jsx)(f.a,{flexGrow:1,children:Object(u.jsx)("div",{ref:C})})]})}),Object(u.jsx)(ct.a,{item:!0,xs:8,md:6,children:Object(u.jsx)(f.a,{style:{borderRadius:"4px",borderColor:"#b9b9b9",borderStyle:"solid",borderWidth:"1px",height:"100%"},children:Object(u.jsx)(xt,{setDuration:function(e){return j({type:"SET_DURATION",payload:e})},duration:c})})}),Object(u.jsx)(ct.a,{item:!0,xs:12,children:Object(u.jsxs)(f.a,{px:1,display:"flex",alignItems:"center",minHeight:52,children:[Object(u.jsx)(f.a,{children:Object(u.jsx)(y.a,{control:Object(u.jsx)(g.a,{size:"small",checked:s,onChange:function(){j({type:"TOGGLE_AUTOCOMPLETE"}),W("AUTOCOMPLETE",!s)}}),label:"Enable autocomplete"})}),Object(u.jsx)(f.a,{ml:4,children:Object(u.jsx)(y.a,{control:Object(u.jsx)(g.a,{size:"small",checked:!l,onChange:function(){j({type:"NO_CACHE"}),W("NO_CACHE",!l)}}),label:"Enable cache"})}),Object(u.jsxs)(f.a,{ml:4,display:"flex",alignItems:"center",children:[Object(u.jsx)(y.a,{control:Object(u.jsx)(g.a,{size:"small",checked:d.limits.enable,onChange:function(){b({type:"TOGGLE_ENABLE_YAXIS_LIMITS"})}}),label:"fix the limits for y-axis"}),d.limits.enable&&Object(u.jsxs)(f.a,{display:"grid",gridTemplateColumns:"120px 120px",gridGap:10,children:[Object(u.jsx)(it.a,{label:"Min",type:"number",size:"small",variant:"outlined",defaultValue:d.limits.range[0],onChange:ln()((function(e){var t=e.target.value;b({type:"SET_YAXIS_LIMITS",payload:[+t,d.limits.range[1]]})}),750)}),Object(u.jsx)(it.a,{label:"Max",type:"number",size:"small",variant:"outlined",defaultValue:d.limits.range[1],onChange:ln()((function(e){var t=e.target.value;b({type:"SET_YAXIS_LIMITS",payload:[d.limits.range[0],+t]})}),750)})]})]})]})})]})})]}),Object(u.jsx)(an,{open:x,onClose:function(){return v(!1)}})]})},jn=n(101),dn=n.n(jn),bn=n(147),On=function(){var e=ue(),t=e.query,n=e.displayType,r=e.serverUrl,c=e.time.period,i=e.queryControls.nocache,s=Zt(),l=s.basicData,u=s.bearerData,j=s.authMethod,d=Object(a.useState)(!1),b=Object(o.a)(d,2),O=b[0],h=b[1],f=Object(a.useState)(),p=Object(o.a)(f,2),x=p[0],v=p[1],m=Object(a.useState)(),y=Object(o.a)(m,2),g=y[0],E=y[1],_=Object(a.useState)(),T=Object(o.a)(_,2),S=T[0],A=T[1];Object(a.useEffect)((function(){S&&(v(void 0),E(void 0))}),[S]);var w=Object(a.useMemo)((function(){if(c){if(!r)return void A("Please enter Server URL");if(!t.trim())return void A("Please enter a valid Query and execute it");if(function(e){var t;try{t=new URL(e)}catch(n){return!1}return"http:"===t.protocol||"https:"===t.protocol}(r)){var e=(c.end-c.start)/2,a=Object(C.a)(Object(C.a)({},c),{},{start:c.start-e,end:c.end+e});return"chart"===n?function(e,t,n,a){return"".concat(e,"/api/v1/query_range?query=").concat(encodeURIComponent(t),"&start=").concat(n.start,"&end=").concat(n.end,"&step=").concat(n.step).concat(a?"&nocache=1":"")}(r,t,a,i):function(e,t,n){return"".concat(e,"/api/v1/query?query=").concat(encodeURIComponent(t),"&start=").concat(n.start,"&end=").concat(n.end,"&step=").concat(n.step)}(r,t,c)}A("Please provide a valid URL")}}),[r,c,n]);return Object(a.useEffect)((function(){Object(bn.a)(dn.a.mark((function e(){var a,r,c,i;return dn.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!w){e.next=38;break}return a=new Headers,"BASIC_AUTH"===j&&a.set("Authorization","Basic "+btoa("".concat((null===l||void 0===l?void 0:l.login)||"",":").concat((null===l||void 0===l?void 0:l.password)||""))),"BEARER_AUTH"===j&&a.set("Authorization",(null===u||void 0===u?void 0:u.token)||""),h(!0),e.prev=5,e.next=8,fetch(w,{headers:a});case 8:if(!(r=e.sent).ok){e.next=18;break}return W("LAST_QUERY",t),e.next=13,r.json();case 13:c=e.sent,A(void 0),"chart"===n?v(c.data.result):E(c.data.result),e.next=32;break;case 18:return e.t0=A,e.next=21,r.json();case 21:if(e.t2=i=e.sent,e.t1=null===e.t2,e.t1){e.next=25;break}e.t1=void 0===i;case 25:if(!e.t1){e.next=29;break}e.t3=void 0,e.next=30;break;case 29:e.t3=i.error;case 30:e.t4=e.t3,(0,e.t0)(e.t4);case 32:e.next=37;break;case 34:e.prev=34,e.t5=e.catch(5),e.t5 instanceof Error&&A(e.t5.message);case 37:h(!1);case 38:case"end":return e.stop()}}),e,null,[[5,34]])})))()}),[w,r,n]),{fetchUrl:w,isLoading:O,graphData:x,liveData:g,error:S}},hn=function(e){var t=e.data,n=d().showInfoMessage,r=Object(a.useMemo)((function(){return JSON.stringify(t,null,2)}),[t]);return Object(u.jsxs)(f.a,{position:"relative",children:[Object(u.jsx)(f.a,{flexDirection:"column",justifyContent:"flex-end",display:"flex",style:{position:"fixed",right:"16px"},children:Object(u.jsx)(Lt.a,{variant:"outlined",onClick:function(e){navigator.clipboard.writeText(r),n("Formatted JSON has been copied"),e.preventDefault()},children:"Copy JSON"})}),Object(u.jsx)("pre",{children:r})]})},fn=n(148),pn=n.n(fn),xn=function(e){var t=e.url,n=d().showInfoMessage;return Object(u.jsx)(f.a,{pl:2,py:1,flexShrink:0,display:"flex",children:Object(u.jsx)(E.a,{title:"Copy Query URL",children:Object(u.jsx)(_.a,{size:"small",onClick:function(e){t&&(navigator.clipboard.writeText(t),n("Value has been copied"),e.preventDefault())},children:Object(u.jsx)(pn.a,{style:{color:"white"}})})})})},vn=function(){var e=ue(),t=e.displayType,n=e.time.period,a=On(),r=a.fetchUrl,c=a.isLoading,i=a.liveData,o=a.graphData,s=a.error;return Object(u.jsxs)(u.Fragment,{children:[Object(u.jsx)(O.a,{position:"static",children:Object(u.jsxs)(h.a,{children:[Object(u.jsxs)(f.a,{display:"flex",children:[Object(u.jsxs)(p.a,{variant:"h5",children:[Object(u.jsx)("span",{style:{fontWeight:"bolder"},children:"VM"}),Object(u.jsx)("span",{style:{fontWeight:"lighter"},children:"UI"})]}),Object(u.jsx)("div",{style:{fontSize:"10px",marginTop:"-2px"},children:Object(u.jsx)("div",{children:"BETA"})})]}),Object(u.jsx)("div",{style:{fontSize:"10px",position:"absolute",top:"40px",opacity:".4"},children:Object(u.jsx)(x.a,{color:"inherit",href:"https://github.com/VictoriaMetrics/VictoriaMetrics/issues/new",target:"_blank",children:"Create an issue"})}),Object(u.jsx)(f.a,{ml:4,flexGrow:1,children:Object(u.jsx)(pe,{})}),Object(u.jsx)(Ce,{}),Object(u.jsx)(xn,{url:r})]})}),Object(u.jsxs)(f.a,{p:2,display:"grid",gridTemplateRows:"auto 1fr",gridGap:"20px",style:{minHeight:"calc(100vh - 64px)"},children:[Object(u.jsx)(f.a,{children:Object(u.jsx)(un,{})}),Object(u.jsxs)(f.a,{height:"100%",children:[c&&Object(u.jsx)(v.a,{in:c,style:{transitionDelay:c?"300ms":"0ms"},children:Object(u.jsx)(f.a,{alignItems:"center",justifyContent:"center",flexDirection:"column",display:"flex",style:{width:"100%",maxWidth:"calc(100vw - 32px)",position:"absolute",height:"50%",background:"linear-gradient(rgba(255,255,255,.7), rgba(255,255,255,.7), rgba(255,255,255,0))"},children:Object(u.jsx)(m.a,{})})}),Object(u.jsxs)(f.a,{height:"100%",p:3,bgcolor:"#fff",children:[s&&Object(u.jsx)(l.a,{color:"error",severity:"error",style:{fontSize:"14px"},children:s}),o&&n&&"chart"===t&&Object(u.jsx)(Ge,{data:o}),i&&"code"===t&&Object(u.jsx)(hn,{data:i}),i&&"table"===t&&Object(u.jsx)(tt,{data:i})]})]})]})]})},mn=n(272),yn=n(153),gn=n(271),En=n(33),_n=n(149),Tn=function(){var e=Object(yn.a)({typography:{fontSize:10}});return Object(u.jsxs)(u.Fragment,{children:[Object(u.jsx)(gn.a,{})," ",Object(u.jsxs)(En.a,{utils:_n.a,children:[" ",Object(u.jsxs)(mn.a,{theme:e,children:["  ",Object(u.jsxs)(be,{children:[" ",Object(u.jsxs)($t,{children:[" ",Object(u.jsxs)(Ne,{children:[" ",Object(u.jsxs)(b,{children:[" ",Object(u.jsx)(vn,{})]})]})]})]})]})]})]})},Sn=function(e){e&&e instanceof Function&&n.e(3).then(n.bind(null,288)).then((function(t){var n=t.getCLS,a=t.getFID,r=t.getFCP,c=t.getLCP,i=t.getTTFB;n(e),a(e),r(e),c(e),i(e)}))};i.a.render(Object(u.jsx)(r.a.StrictMode,{children:Object(u.jsx)(Tn,{})}),document.getElementById("root")),Sn()}},[[207,1,2]]]);