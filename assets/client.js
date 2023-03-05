(()=>{function l(e,t){let n;t?n=t.querySelectorAll(e):n=document.querySelectorAll(e);let i=[];return n.forEach(s=>{i.push(s)}),i}function u(e,t){let n=l(e,t);switch(n.length){case 0:return;case 1:return n[0];default:console.warn(`found [${n.length}] elements with selector [${e}], wanted zero or one`)}}function o(e,t){let n=u(e,t);if(!n)throw new Error(`no element found for selector [${e}]`);return n}function ge(e,t){return typeof e=="string"&&(e=o(e)),e.innerHTML=t,e}function P(e,t,n="block"){return typeof e=="string"&&(e=o(e)),e.style.display=t?n:"none",e}function r(e,t,...n){let i=document.createElement(e);for(let s in t)if(s==="for"&&(s="for"),s==="className"&&(s="class"),s&&t.hasOwnProperty(s)){let a=t[s];s==="dangerouslySetInnerHTML"?ge(i,a.__html):a===!0?i.setAttribute(s,s):a!==!1&&a!==null&&a!==void 0&&i.setAttribute(s,a.toString())}for(let s of n)if(Array.isArray(s))s.forEach(a=>{if(s==null)throw new Error(`child array for tag [${e}] is ${s}
${i.outerHTML}`);if(a==null)throw new Error(`child for tag [${e}] is ${a}
${i.outerHTML}`);typeof a=="string"&&(a=document.createTextNode(a)),i.appendChild(a)});else{if(s==null)throw new Error(`child for tag [${e}] is ${s}
${i.outerHTML}`);s.nodeType||(s=document.createTextNode(s.toString())),i.appendChild(s)}return i}function he(){for(let e of l(".menu-container .final"))e.scrollIntoView({block:"nearest"})}var ne="mode-light",oe="mode-dark";function ve(){for(let e of l(".mode-input"))e.onclick=()=>{switch(e.value){case"":document.body.classList.remove(ne),document.body.classList.remove(oe);break;case"light":document.body.classList.add(ne),document.body.classList.remove(oe);break;case"dark":document.body.classList.remove(ne),document.body.classList.add(oe);break}}}function be(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function E(e,t,n){let i=document.getElementById("flash-container");i===null&&(i=document.createElement("div"),i.id="flash-container",document.body.insertAdjacentElement("afterbegin",i));let s=document.createElement("div");s.className="flash";let a=document.createElement("input");a.type="radio",a.style.display="none",a.id="hide-flash-"+e,s.appendChild(a);let m=document.createElement("label");m.htmlFor="hide-flash-"+e;let c=document.createElement("span");c.innerHTML="\xD7",m.appendChild(c),s.appendChild(m);let f=document.createElement("div");f.className="content flash-"+t,f.innerText=n,s.appendChild(f),i.appendChild(s),be(s)}function ye(){let e=document.getElementById("flash-container");if(e===null)return E;let t=e.querySelectorAll(".flash");if(t.length>0)for(let n of t)be(n);return E}function Te(){for(let e of l(".link-confirm"))e.onclick=()=>{let t=e.dataset.message;return t&&t.length===0&&(t="Are you sure?"),confirm(t)}}function Me(e){let t=Date.UTC(e.getUTCFullYear(),e.getUTCMonth(),e.getUTCDate(),e.getUTCHours(),e.getUTCMinutes(),e.getUTCSeconds());return new Date(t).toISOString().substring(0,19).replace("T"," ")}function V(e,t){let n=(e||"").replace(/-/gu,"/").replace(/[TZ]/gu," ")+" UTC",i=new Date(n),s=(new Date().getTime()-i.getTime())/1e3,a=Math.floor(s/86400),m=i.getFullYear(),c=i.getMonth()+1,f=i.getDate();if(isNaN(a)||a<0||a>=31)return m.toString()+"-"+(c<10?"0"+c.toString():c.toString())+"-"+(f<10?"0"+f.toString():f.toString());let d,b;return a===0?s<5?(b=1,d="just now"):s<60?(b=1,d=Math.floor(s)+" seconds ago"):s<120?(b=10,d="1 minute ago"):s<3600?(b=30,d=Math.floor(s/60)+" minutes ago"):s<7200?(b=60,d="1 hour ago"):(b=60,d=Math.floor(s/3600)+" hours ago"):a===1?(b=600,d="yesterday"):a<7?(b=600,d=a+" days ago"):(b=6e3,d=Math.ceil(a/7)+" weeks ago"),t&&(t.innerText=d,setTimeout(()=>V(e,t),b*1e3)),d}function Ee(){return l(".reltime").forEach(e=>{V(e.dataset.time||"",e)}),V}function Ft(e,t){let n=0;return(...i)=>{n!==0&&window.clearTimeout(n),n=window.setTimeout(()=>{e(null,...i)},t)}}function At(e,t,n,i,s){if(!e)return;let a=e.id+"-list",m=document.createElement("datalist"),c=document.createElement("option");c.value="",c.innerText="Loading...",m.appendChild(c),m.id=a,e.parentElement?.prepend(m),e.setAttribute("autocomplete","off"),e.setAttribute("list",a);let f={},d="";function b(h){let M=t;return M.includes("?")?M+"&t=json&"+n+"="+encodeURIComponent(h):M+"?t=json&"+n+"="+encodeURIComponent(h)}function B(h){let M=f[h];!M||!M.frag||(d=h,m.replaceChildren(M.frag.cloneNode(!0)))}function W(){let h=e.value;if(h.length===0)return;let M=b(h),ee=!h||!d;if(!ee){let I=f[d];I&&(ee=!I.data.find(N=>h===s(N)))}if(!!ee){if(f[h]&&f[h].url===M){B(h);return}fetch(M,{credentials:"include"}).then(I=>I.json()).then(I=>{if(!I)return;let N=Array.isArray(I)?I:[I],ue=document.createDocumentFragment(),pe=10;for(let G=0;G<N.length&&pe>0;G++){let fe=s(N[G]),Dt=i(N[G]);if(fe){let te=document.createElement("option");te.value=fe,te.innerText=Dt,ue.appendChild(te),pe--}}f[h]={url:M,data:N,frag:ue,complete:!1},B(h)})}}e.oninput=Ft(W,250),console.log("managing ["+e.id+"] autocomplete")}function ke(){return At}function Le(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function g(e,t,n){return t||(t=18),n==null&&(n="icon"),`<svg class="${n}" style="width: ${t}px; height: ${t+"px"};"><use xlink:href="#svg-${e}"></use></svg>`}function Y(e,t,n){return{__html:g(e,t,n)}}function xe(e){e||(e="");let t=g("right",15,"expand"),n=g("down",15,"collapse");return{__html:t+n+e}}function v(e){return setTimeout(()=>{e.setSelectionRange(e.value.length,e.value.length),e.focus()},100),!0}function Rt(e,t){return e.parentElement!==t.parentElement?null:e===t?0:e.compareDocumentPosition(t)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var K;function re(e){let t=[],n=l(".item .value",e);for(let s of n)t.push(s.innerText);let i=o("input.result",e);i.value=t.join(", ")}function Ut(e){let t=e.parentElement?.parentElement;e.remove(),t&&re(t)}function He(e){let t=o(".value",e),n=o(".editor",e);n.value=t.innerText;let i=()=>{if(n.value===""){e.remove();return}t.innerText=n.value,P(t,!0),P(n,!1);let s=e.parentElement?.parentElement;s&&re(s)};n.onblur=i,n.onkeydown=s=>{if(s.code==="Enter")return s.preventDefault(),i(),!1},P(t,!1),P(n,!0),n.focus()}function Ie(e,t){let n=document.createElement("div");n.className="item",n.draggable=!0,n.ondragstart=m=>{m.dataTransfer?.setDragImage(document.createElement("div"),0,0),n.classList.add("dragging"),K=n},n.ondragover=()=>{let m=Rt(n,K);if(!m)return;let c=m===-1?n:n.nextSibling;K.parentElement?.insertBefore(K,c),re(t)},n.ondrop=m=>{m.preventDefault()},n.ondragend=m=>{n.classList.remove("dragging"),m.preventDefault()};let i=document.createElement("div");i.innerText=e,i.className="value",i.onclick=()=>{He(n)},n.appendChild(i);let s=document.createElement("input");s.className="editor",n.appendChild(s);let a=document.createElement("div");return a.innerHTML=g("times",13),a.className="close",a.onclick=()=>{Ut(n)},n.appendChild(a),n}function $t(e,t){let n=Ie("",t);e.appendChild(n),He(n)}function O(e){let t=o("input.result",e),n=o(".tags",e),i=t.value.split(",").map(a=>a.trim()).filter(a=>a!=="");P(t,!1),n.innerHTML="";for(let a of i)n.appendChild(Ie(a,e));u(".add-item",e)?.remove();let s=document.createElement("div");s.className="add-item",s.innerHTML=g("plus",22),s.onclick=()=>{$t(n,e)},e.insertBefore(s,o(".clear",e))}function we(){for(let e of l(".tag-editor"))O(e);return O}var Ce="--selected";function Nt(e){let t=e.parentElement?.parentElement?.querySelector("input");if(!t)throw new Error("no associated input found");t.value="\u2205"}function ie(e){e.onreset=()=>ie(e);let t={},n={};for(let i of e.elements){let s=i;if(s.name.length>0)if(s.name.endsWith(Ce))n[s.name]=s;else{(s.type!=="radio"||s.checked)&&(t[s.name]=s.value);let a=()=>{let m=n[s.name+Ce];m&&(m.checked=t[s.name]!==s.value)};s.onchange=a,s.onkeyup=a}}}function De(){for(let e of l("form.editor"))ie(e);return[Nt,ie]}var Pt=[];function Fe(e,t,n){let i=e.querySelectorAll(t);if(i.length===0)throw new Error("empty query selector ["+t+"]");i.forEach(s=>n(s))}function _(e,t,n){Fe(e,t,i=>{i.style.backgroundColor=n})}function w(e,t,n){Fe(e,t,i=>{i.style.color=n})}function Ot(e,t,n){let i=document.querySelector("#mockup-"+e);if(!i){console.error("can't find mockup for mode ["+e+"]");return}switch(t){case"color-foreground":w(i,".mock-main",n);break;case"color-background":_(i,".mock-main",n);break;case"color-foreground-muted":w(i,".mock-main .mock-muted",n);break;case"color-background-muted":_(i,".mock-main .mock-muted",n);break;case"color-link-foreground":w(i,".mock-main .mock-link",n);break;case"color-link-visited-foreground":w(i,".mock-main .mock-link-visited",n);break;case"color-nav-foreground":w(i,".mock-nav",n),w(i,".mock-nav .mock-link",n);break;case"color-nav-background":_(i,".mock-nav",n);break;case"color-menu-foreground":w(i,".mock-menu",n),w(i,".mock-menu .mock-link",n);break;case"color-menu-background":_(i,".mock-menu",n);break;case"color-menu-selected-foreground":w(i,".mock-menu .mock-link-selected",n);break;case"color-menu-selected-background":_(i,".mock-menu .mock-link-selected",n);break;default:console.error("invalid key ["+t+"]")}}function Ae(){for(let e of l(".color-var")){let t=e.dataset.var,n=e.dataset.mode;Pt.push(t),!(!t||t.length===0)&&(e.oninput=()=>{Ot(n,t,e.value)})}}var Re=!1;function Ue(e){if(e||(e="/connect"),e.indexOf("ws")===0)return e;let t=document.location,n="ws";return t.protocol==="https:"&&(n="wss"),e.indexOf("/")!==0&&(e="/"+e),n+`://${t.host}${e}`}var Q=class{constructor(t,n,i,s,a){this.debug=t,this.open=n,this.recv=i,this.err=s,this.url=Ue(a),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){window.onbeforeunload=()=>{Re=!0},this.connectTime=Date.now(),this.sock=new WebSocket(Ue(this.url));let t=this;this.sock.onopen=()=>{t.connected=!0,t.pendingMessages.forEach(t.send),t.pendingMessages=[],t.debug&&console.log("WebSocket connected"),t.open()},this.sock.onmessage=n=>{let i=JSON.parse(n.data);t.debug&&console.debug("[socket]: receive",i),t.recv(i)},this.sock.onerror=n=>()=>{t.err("socket",n.type)},this.sock.onclose=()=>{if(Re)return;t.connected=!1;let n=t.connectTime?Date.now()-t.connectTime:0;n>0&&n<2e3?(t.pauseSeconds=t.pauseSeconds*2,t.debug&&console.debug(`socket closed immediately, reconnecting in ${t.pauseSeconds} seconds`),setTimeout(()=>{t.connect()},t.pauseSeconds*1e3)):(console.debug("socket closed after ["+n+"ms]"),setTimeout(()=>{t.connect()},t.pauseSeconds*500))}}disconnect(){}send(t){if(this.debug&&console.debug("out",t),!this.sock)throw new Error("socket not initialized");if(this.connected){let n=JSON.stringify(t,null,2);this.sock.send(n)}else this.pendingMessages.push(t)}};function $e(){return Q}function C(e,t,n){return e?`<img class="${n}" style="width: ${t+"px"}; height: ${t+"px"};" src="${e}" />`:g("profile",t,n)}function Ne(e,t,n,i){return r("tr",{id:"member-"+e,class:"member","data-id":e},r("td",null,r("a",{href:"#modal-member-"+e},r("div",{class:"left prs member-picture",dangerouslySetInnerHTML:{__html:C(i,18,"")}}),r("span",{class:"member-name member-"+e+"-name"},t))),r("td",{class:"shrink text-align-right"},r("em",{class:"member-status"},n)),r("td",{class:"shrink online-status",title:"offline",dangerouslySetInnerHTML:Y("circle",18,"right")}))}function Pe(e,t,n,i){return r("div",{id:"modal-member-"+e,"data-id":e,class:"modal modal-member",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,r("span",{class:"member-picture",dangerouslySetInnerHTML:{__html:C(i,18,"")}}),r("span",{class:"member-name member-"+e+"-name"},t))),r("div",{class:"modal-body"},r("em",null,"Role"),r("br",null),r("span",{class:"member-role"},n))))}function Oe(e,t,n,i){let s="return confirm('Are you sure you wish to remove this user?');",a=[["owner","Owner"],["member","Member"],["observer","Observer"]];return r("div",{id:"modal-member-"+e,"data-id":e,class:"modal modal-member",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,r("span",{className:"member-picture",dangerouslySetInnerHTML:{__html:C(i,18,"")}}),r("span",{className:"member-name member-"+e+"-name"},t))),r("div",{class:"modal-body"},r("form",{action:document.location.pathname,method:"post",class:"expanded"},r("input",{type:"hidden",name:"userID",value:e}),r("em",null,"Role"),r("br",null),r("select",{class:"input-role",name:"role"},a.map(m=>m[0]===n?r("option",{selected:"selected",value:m[0]},m[1]):r("option",{value:m[0]},m[1]))),r("hr",null),r("div",{class:"right"},r("button",{class:"member-update",type:"submit",name:"action",value:"member-update"},"Save")),r("button",{type:"submit",class:"member-remove",name:"action",value:"member-remove",onClick:s},"Remove")))))}function T(e){if(!e)return"System";let t=qe(e);return t||"Unknown User"}function y(){return o("#self-id").innerText}function qt(){return o("#self-role").innerText==="owner"}function jt(){let e=o("#modal-member-"+y()),t=u("form",e);t&&(t.onsubmit=()=>{let n=o('input[name="name"]',t),i=o('input[name="choice"]:checked',t),s=u('input[name="picture"]:checked',t),a={name:n.value,choice:i.value};return s&&(a.picture=s.value),p("self",a),l(".member-"+y()+"-name").forEach(m=>{m.innerText=n.value}),o("#self-picture").innerHTML=C(s?s.value:"",20,"icon"),document.location.hash="",!1})}function qe(e){return e===y()?o("#self-name").innerText:o("#member-"+e+" .member-name").innerText}function je(e){let t=u("form",e);if(!t)return;let n=i=>{let s=o('input[name="userID"]',t).value,a=o('select[name="role"]',t).value;p(i,{userID:s,role:a});let m=o("#member-"+s);return i==="member-update"?o(".member-role",m).innerText=a:i==="member-remove"&&m.remove(),document.location.hash="",!1};o(".member-update",t).onclick=()=>n("member-update"),o(".member-remove",t).onclick=()=>confirm("Are you sure you wish to remove this user?")?n("member-remove"):!1}function Vt(){let e=l(".modal-member");for(let t of e)je(t)}function se(e){if(l(".member-"+e.userID+"-name").forEach(t=>{t.innerText=e.name}),e.userID===y())o("#self-role").innerText=e.role,o("#self-picture").innerHTML=C(e.picture?e.picture:"",20,"icon");else{let t=o("#member-"+e.userID);o(".member-role",t).innerText=e.role,o(".member-picture",t).innerHTML=C(e.picture?e.picture:"",18,"");let n=o("#modal-member-"+e.userID),i=u(".member-role",n);i&&(i.innerText=e.role);let s=u('select[name="role"]',n);s&&(s.value=e.role),o(".member-picture",n).innerHTML=C(e.picture?e.picture:"",18,"")}if(qe(e.userID)!==e.name){let t=o("#panel-members table tbody"),i=[...t.children];i.sort((s,a)=>{let m=o(".member-name",s).innerText,c=o(".member-name",a).innerText;return m.localeCompare(c,void 0,{sensitivity:"accent"})}),t.replaceChildren(...i)}}function Ve(e){if(u("#member-"+e.userID)||e.userID===y())return se(e);let n=o("#panel-members table tbody"),i=-1;for(let c=0;c<n.children.length;c++){let f=n.children.item(c);if(o(".member-name",f).innerText.localeCompare(e.name,void 0,{sensitivity:"accent"})>0){i=c;break}}let s=Ne(e.userID,e.name,e.role,e.picture?e.picture:"");i===-1?n.appendChild(s):n.insertBefore(s,n.children[i]);let a=o("#member-modals"),m;qt()?m=Oe(e.userID,e.name,e.role,e.picture?e.picture:""):m=Pe(e.userID,e.name,e.role,e.picture?e.picture:""),a.appendChild(m),o(".member-picture",m).innerHTML=C(e.picture?e.picture:"",24,"icon"),je(m)}function _e(e){o("#member-"+e).remove()}function Be(e){if(e.userID===y())return;let t=u("#member-"+e.userID+" .online-status");if(!t)throw new Error("missing panel #member-"+e.userID);t.title=e.connected?"online":"offline";let n=e.connected?"check-circle":"circle";t.innerHTML=g(n,18,"right")}function We(){o("#self-edit-link").onclick=()=>{v(o("#self-name-input"))},jt(),Vt()}function Ge(e,t){let n=r("li",null),i=Me(new Date),s=r("span",{class:"nowrap reltime","data-time":i},"just now");V(i,s);let a=r("div",{class:"right"});return a.appendChild(s),n.appendChild(a),n.appendChild(r("div",null,e.content)),n.appendChild(r("div",null,r("em",{class:"member-"+e.userID+"-name"},t))),n}function A(e,t){let n="comment-link-"+e+"-"+t,i="#modal-"+e+"-"+t+"-comments",s=r("a",{id:n,class:"comment-link","data-key":e+"-"+t,href:i,title:"0 comments",dangerouslySetInnerHTML:Y("comment-alt",18,"member-icon")});return ae(s),s}function R(e,t,n){return r("div",{id:"modal-"+e+"-"+t+"-comments",class:"modal comments",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,n," Comments")),r("div",{class:"modal-body"},r("ul",{id:"comment-list-"+e+"-"+t,class:"comment-list"}),r("form",{action:"#",method:"post",class:"expanded"},r("input",{type:"hidden",name:"action",value:"comment"}),r("input",{type:"hidden",name:"svc",value:e}),r("input",{type:"hidden",name:"modelID",value:t}),r("div",null,r("textarea",{name:"content",placeholder:"Add a comment, Markdown and HTML supported"})),r("div",{class:"mt right"},r("button",{type:"submit"},"Add Comment"))))))}function ae(e){e.onclick=()=>v(o("#modal-"+e.dataset.key+"-comments form textarea"))}function le(e){let t=o("#comment-list-"+e.svc+"-"+e.modelID),n=T(e.userID),i=Ge(e,n);t.appendChild(i);let s=t.children.length,a=o("#comment-link-"+e.svc+"-"+e.modelID);a.title=s+(s===1?" comment":" comments"),a.innerHTML.indexOf("comment-dots")===-1&&(a.innerHTML=g("comment-dots",18,"right"))}function D(e){let t=o("form",e);t.onsubmit=()=>{let n={svc:o('input[name="svc"]',t).value,modelID:o('input[name="modelID"]',t).value},i=o("textarea",t);return n.content=i.value,n.userID=y(),p("comment",n),le(n),i.value="",!1}}function Ye(){for(let e of l(".comment-link"))ae(e);for(let e of l(".modal.comments"))D(e)}function Z(e,t,n){if(n){let i=u(`select[name="${e}"] option[value="${n}"]`,t);if(i){let s=i.innerText;return`<a href="/${e}/${n}">${s}</a> `}return`<a href="/${e}/${n}">${n}</a> `}return""}function _t(e,t,n,i){let s="";return s+=Z("sprint",t,i),s+=e==="retro"?"retrospective":e,n&&(s+=" in "+Z("team",t,n)),s}function z(e){let t=o(`#${e.type}-list tbody`),n=u(".empty",t);n&&n.remove();let i=document.createElement("tr");i.id=`${e.type}-list-${e.id}`;let s=document.createElement("td"),a=document.createElement("div");a.classList.add("right"),a.appendChild(A(e.type,e.id)),a.appendChild(R(e.type,e.id,e.title)),s.appendChild(a);let m=document.createElement("a");m.href=e.path;let c=document.createElement("span");c.classList.add("model-span-icon"),c.innerHTML=g(e.icon,16,"icon"),m.appendChild(c);let f=document.createElement("span");f.classList.add("model-span-title"),f.innerText=e.title,m.appendChild(f),s.appendChild(m),i.appendChild(s),t.appendChild(i),D(o(".modal",a))}function S(e){let t=o(`#${e.type}-list-${e.model.id}`);o(".model-span-icon",t).innerHTML=g(e.model.icon),o(".model-span-title",t).innerText=e.model.title}function J(e){o(`#${e.type}-list-${e.id}`).remove();let t=o(`#${e.type}-list tbody`);if(t.children.length===0){let n=document.createElement("tr");n.classList.add("empty");let i=document.createElement("em");i.innerText="no "+e.type+"s",n.appendChild(i),t.appendChild(n)}}function k(e,t,n,i,s,a){let m=u('input[name="title"]',t);m&&(m.value=s);for(let d of l(".config-panel-team",t))d.innerHTML=Z("team",t,n);for(let d of l(".config-panel-sprint",t))d.innerHTML=Z("sprint",t,i);for(let d of l(".view-icon",t))d.innerHTML=g(a,24,"icon");for(let d of l(".view-title",t))d.innerText=s;for(let d of l('input[name="icon"]',t))d.checked=a===d.value;let c=u('select[name="team"]',t);c!=null&&(c.value=n||"");let f=u('select[name="sprint"]',t);f!=null&&(f.value=i||""),o("#model-title").innerText=s,o("#model-icon").innerHTML=g(a,20),o("#model-banner").innerHTML=_t(e,t,n||"",i||"")}function L(e){o(`#modal-${e}-config-link`).onclick=()=>{let t=u(`#modal-${e}-config form input[name="title"]`);t&&v(t)}}function x(e){console.log("permissionsUpdate",e)}function F(e){o(".permission-config-team").style.display=e?"block":"none"}function $(e){o(".permission-config-sprint").style.display=e?"block":"none"}function H(e){let t={};for(let n of l(".perm-option",e))n.checked&&(t[n.name]=n.value);return t}function U(e,t){F(e.value!==""),t&&$(t.value!=="")}function Ke(){L("team");let e=u("#modal-team-config form");e&&(e.onsubmit=()=>{let t=o('input[name="title"]',e).value,n=o('input[name="icon"]:checked',e).value;return p("update",{title:t,icon:n,...H(e)}),document.location.hash="",!1})}function Bt(e){let t=o("#modal-team-config");k("team",t,null,null,e.title,e.icon)}function Qe(e){switch(e.cmd){case"update":return Bt(e.param);case"child-add":return z(e.param);case"child-update":return S(e.param);case"child-remove":return J(e.param);case"permissions":return x(e.param);default:throw new Error("invalid team command ["+e.cmd+"]")}}function Ze(){L("sprint");let e=u("#modal-sprint-config form");if(e){let t=o('select[name="team"]',e);t.onchange=()=>{F(t.value!=="")},U(t),e.onsubmit=()=>{let n=o('input[name="title"]',e).value,i=o('input[name="icon"]:checked',e).value,s=o('input[name="startDate"]',e).value,a=o('input[name="endDate"]',e).value;return p("update",{title:n,icon:i,startDate:s,endDate:a,team:t.value,...H(e)}),document.location.hash="",!1}}}function q(e){return`${e}`.split("T")[0]}function Wt(e){let t="";return e.startDate&&(t+="starts ",t+=q(e.startDate),e.endDate&&(t+=", ")),e.endDate&&(t+="ends ",t+=q(e.endDate)),t}function Gt(e){let t=o("#modal-sprint-config"),n=u("form",t);n?(o('input[name="startDate"]',n).value=q(e.startDate),o('input[name="endDate"]',n).value=q(e.endDate)):(o(".config-panel-startDate",t).innerText=q(e.startDate),o(".config-panel-endDate",t).innerText=q(e.endDate)),o("#model-summary").innerText=Wt(e),k("sprint",t,e.teamID,null,e.title,e.icon)}function ze(e){switch(e.cmd){case"update":return Gt(e.param);case"child-add":return z(e.param);case"child-update":return S(e.param);case"child-remove":return J(e.param);case"permissions":return x(e.param);default:throw new Error("invalid sprint command ["+e.cmd+"]")}}function Se(e){return r("tr",{class:"story-row",id:"story-row-"+e.id,"data-idx":"s.Idx"},r("td",null,r("a",{href:"#modal-story-"+e.id},r("div",{class:"story-title"},e.title))),r("td",{class:"story-status"},e.status),r("td",{class:"story-final-vote"},e.finalVote?e.finalVote:"-"),r("td",null,A("story",e.id),R("story",e.id,e.title)))}function Yt(e,t){t.onclick=()=>{v(o("#modal-story-"+e+'-edit form input[name="title"]'))}}function Kt(e,t){t.onsubmit=()=>(p("child-status",{storyID:e,status:"new"}),!1)}function Je(e,t){t.onsubmit=()=>(p("child-status",{storyID:e,status:"active"}),!1)}function Qt(e,t){t.onsubmit=()=>(p("child-status",{storyID:e,status:"complete"}),!1)}function Zt(e,t){l(".vote-option",t).forEach(n=>{n.onclick=()=>(o('input[name="vote"]',n).checked=!0,p("vote",{storyID:e,vote:n.dataset.choice}),l("#modal-story-"+e+" .story-members .member").forEach(i=>{i.dataset.member==y()&&(o(".choice",i).innerHTML=g("check",18,""))}),!1)})}function me(e,t){t.onsubmit=()=>(confirm("Are you sure you want to delete this story?")&&(p("child-remove",{storyID:e}),document.location.hash=""),!1)}function ce(e){let t=e.dataset.id;if(!t){console.warn("no id in dataset",e);return}l(".vote-submit-button",e).forEach(n=>{n.style.display="none"}),l(".link-edit",e).forEach(n=>Yt(t,n)),l(".status-new-form-delete",e).forEach(n=>me(t,n)),l(".status-new-form-next",e).forEach(n=>Je(t,n)),l(".status-active-form-prev",e).forEach(n=>Kt(t,n)),l(".status-active-form-next",e).forEach(n=>Qt(t,n)),l(".story-vote-options",e).forEach(n=>Zt(t,n)),l(".status-complete-form-prev",e).forEach(n=>Je(t,n))}function Xe(e){let t=o(".story-delete-button",e);t.onclick=()=>(confirm("Are you sure you want to delete this story?")&&(p("child-remove",{storyID:t.dataset.id}),document.location.hash=""),!1),e.onsubmit=()=>{let n=o('input[name="storyID"]',e).value,i=o('input[name="title"]',e).value;return p("child-update",{storyID:n,title:i}),document.location.hash="",!1}}function et(){l(".add-story-link").forEach(n=>{n.onclick=()=>v(o("#story-add-title"))});let e=o("#modal-story--add"),t=o("form",e);t.onsubmit=()=>{let n=o('input[name="title"]',t),i=n.value;return n.value="",p("child-add",{title:i}),document.location.hash="",!1};for(let n of l(".modal-story-edit"))Xe(o("form",n));l("#story-modals .modal-story").forEach(ce)}function tt(e){let t=o("#panel-detail table tbody"),n=-1;for(let d=0;d<t.children.length;d++){let b=t.children.item(d),B=o(".story-title",b).innerText,W=b.dataset.index;if(W){if(parseInt(W,10)>=e.idx){n=d;break}else if(B.localeCompare(e.title,void 0,{sensitivity:"accent"})>=0){n=d;break}}}let i=Se(e);n===-1?t.appendChild(i):t.insertBefore(i,t.children[n]),D(o(".modal",i));let a=o(".modal-story-edit-new").cloneNode(!0);a.id="modal-story-"+e.id+"-edit",a.classList.remove("modal-story-edit-new"),l('input[name="storyID"]',a).forEach(d=>{d.value=e.id}),o('input[name="title"]',a).value=e.title,o(".story-delete-button",a).dataset.id=e.id,Xe(o("form",a)),l(".edit-form-delete",a).forEach(d=>me(e.id,d)),o("#story-modals").appendChild(a);let c=o("#modal-story-new").cloneNode(!0);c.id="modal-story-"+e.id,c.dataset.id=e.id,c.dataset.status=e.status,c.classList.add("modal-story");let f=o(".link-edit",c);f.href="#modal-story-"+e.id+"-edit",f.dataset.id=e.id,o("#story-modals").appendChild(c),ce(c),de(e),(document.location.hash==="modal-story--add"||document.location.hash==="")&&(document.location.hash="modal-story-"+e.id)}function de(e){let t=o("#story-row-"+e.id);o(".story-status",t).innerText=e.status,o(".story-title",t).innerText=e.title,o(".story-final-vote",t).innerText=e.finalVote;let n=o("#modal-story-"+e.id+"-edit");o('form input[name="title"]',n).innerText=e.title;let i=o("#modal-story-"+e.id);o("h2.billboard",i).innerText=e.title}function nt(e){let t=o("#story-row-"+e.id);o(".story-status",t).innerText=e.status;let n=o("#modal-story-"+e.id);o(".status-new",n).style.display=e.status==="new"?"block":"none",o(".status-active",n).style.display=e.status==="active"?"block":"none",o(".status-complete",n).style.display=e.status==="complete"?"block":"none"}function ot(e){let t=o("#story-row-"+e),n=o(".story-title",t).innerText;E(e+"-removed","success",`story [${n}] has been removed`),t.remove(),o("#modal-story-"+e).remove()}function rt(e){l("#modal-story-"+e.storyID+" .story-members .member").forEach(t=>{t.dataset.member==e.userID&&(o(".choice",t).innerHTML=g("check",18,""))}),console.log("VOTE",e)}function it(){L("estimate");let e=u("#modal-estimate-config form");if(e){let t=o('select[name="team"]',e);t.onchange=()=>{F(t.value!=="")};let n=o('select[name="sprint"]',e);n.onchange=()=>{$(n.value!=="")},U(t,n),e.onsubmit=()=>{let i=o('input[name="title"]',e).value,s=o('input[name="icon"]:checked',e).value,a=o('input[name="choices"]',e).value;return p("update",{title:i,icon:s,choices:a,team:t.value,sprint:n.value,...H(e)}),document.location.hash="",!1}}et()}function zt(e){let t=o("#modal-estimate-config"),n=u("form",t);if(n){let i=o('input[name="choices"]',n);i.value=e.choices.join(","),i.parentElement&&O(i.parentElement)}else o(".config-panel-choices",t).innerText=e.choices.join(", ");k("estimate",t,e.teamID,e.sprintID,e.title,e.icon)}function st(e){switch(e.cmd){case"update":return zt(e.param);case"child-add":return tt(e.param);case"child-update":return de(e.param);case"child-status":return nt(e.param);case"child-remove":return ot(e.param);case"vote":return rt(e.param);case"permissions":return x(e.param);default:throw new Error("invalid estimate command ["+e.cmd+"]")}}function at(e){return r("li",{id:"report-group-"+e},r("input",{id:"accordion-"+e,type:"checkbox",hidden:!0,checked:"checked"}),r("label",{for:"accordion-"+e,dangerouslySetInnerHTML:xe(e)}),r("div",{class:"bd"}))}function lt(e){return r("div",{class:"report",id:"report-"+e.id},r("div",null,r("div",{class:"right"},A("report",e.id)),R("report",e.id,e.id),r("a",{href:"#modal-report-"+e.id,"data-id":e.id,class:"clean"},r("h4",{class:"username member-"+e.userID+"-name"},T(e.userID)),r("div",{class:"pt"},e.html))))}function mt(e){return r("div",{id:"modal-report-"+e.id,class:"modal modal-report-view",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,e.day,": ",r("span",{class:"member-"+e.userID+"-name"},T(e.userID)))),r("div",{class:"modal-body",dangerouslySetInnerHTML:e.html})))}function ct(e){return r("div",{id:"modal-report-"+e.id,class:"modal modal-report-edit",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,e.day,": ",r("span",{className:"member-"+e.userID+"-name"},T(e.userID)))),r("div",{class:"modal-body"},r("form",{action:"#",method:"post"},r("input",{type:"hidden",name:"reportID",value:e.id}),r("div",{class:"mb expanded"},r("label",{for:"input-day-"+e.id},r("em",{class:"title"},"Day")),r("div",null,r("input",{type:"date",id:"input-day-"+e.id,name:"day",value:e.day}))),r("div",{class:"mb expanded"},r("label",{for:"input-content-"+e.id},r("em",{class:"title"},"Content")),r("div",null,r("textarea",{id:"input-content-"+e.id,name:"content"},e.content))),r("div",{class:"right"},r("button",{class:"report-edit-save",type:"submit",name:"action",value:"child-update"},"Save Changes")),r("button",{class:"report-edit-delete",type:"submit",name:"action",value:"child-remove",onclick:"return confirm('Are you sure you want to delete this report?');"},"Delete")))))}function dt(e){let t=o("form",e),n=o('input[name="reportID"]',t).value;o(".report-edit-delete",t).onclick=()=>(confirm("Are you sure you want to delete this report?")&&(p("child-remove",{reportID:n}),document.location.hash=""),!1),t.onsubmit=()=>{let i=o('input[name="day"]',t).value,a=o('textarea[name="content"]',t).value;return p("child-update",{reportID:n,day:i,content:a}),document.location.hash="",!1}}function ut(){l(".add-report-link").forEach(n=>{n.onclick=()=>v(o("#report-add-content"))}),l(".modal-report-edit-link").forEach(n=>{n.onclick=()=>v(o("#input-content-"+n.dataset.id))});let e=o("#modal-report--add"),t=o("form",e);t.onsubmit=()=>{let n=o('input[name="day"]',t).value,i=o('textarea[name="content"]',t),s=i.value;return i.value="",p("child-add",{day:n,content:s}),document.location.hash="",!1},l(".modal-report-edit").forEach(dt)}function pt(e){e.day.length>10&&(e.day=e.day.substring(0,10));let t=u("#report-group-"+e.day+" .bd");if(!t){let a=at(e.day);o("#report-groups").appendChild(a),t=o(".bd",a)}let n=-1,i=T(e.userID);for(let a=0;a<t.children.length;a++){let m=t.children.item(a);if(o(".username",m).innerText.localeCompare(i,void 0,{sensitivity:"accent"})>=0){n=a;break}}let s=lt(e);if(n===-1?t.appendChild(s):t.insertBefore(s,t.children[n]),y()===e.userID){let a=ct(e);dt(a),o("#report-modals").appendChild(a)}else o("#report-modals").appendChild(mt(e));D(o(".modal",s))}function ft(e){o("#report-"+e).remove(),E(e+"-removed","success","report has been removed"),o("#modal-report-"+e).remove()}function gt(){L("standup");let e=u("#modal-standup-config form");if(e){let t=o('select[name="team"]',e);t.onchange=()=>{F(t.value!=="")};let n=o('select[name="sprint"]',e);n.onchange=()=>{$(n.value!=="")},U(t,n),e.onsubmit=()=>{let i=o('input[name="title"]',e).value,s=o('input[name="icon"]:checked',e).value;return p("update",{title:i,icon:s,team:t.value,sprint:n.value,...H(e)}),document.location.hash="",!1}}ut()}function St(e){let t=o("#modal-standup-config");k("standup",t,e.teamID,e.sprintID,e.title,e.icon)}function ht(e){switch(e.cmd){case"update":return St(e.param);case"child-add":return pt(e.param);case"child-update":return console.log("TODO: child-update");case"child-remove":return ft(e.param);case"permissions":return x(e.param);default:throw new Error("invalid standup command ["+e.cmd+"]")}}function vt(e){return r("div",{id:"category-"+e,"data-category":e,class:"category"},r("div",{class:"right"},r("a",{class:"add-feedback-link","data-category":e,href:"#modal-feedback--add-"+e},r("button",null,"New"))),r("h4",null,r("a",{href:"#modal-feedback--add-"+e},e)),r("div",{class:"clear"}),r("div",{class:"feedback-list"}),r("div",{id:"modal-feedback--add-"+e,class:"modal modal-feedback-add","data-category":e,style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,"New Feedback")),r("div",{class:"modal-body"},r("form",{action:"#",method:"post"},r("input",{type:"hidden",name:"action",value:"child-add"}),r("input",{type:"hidden",name:"action",value:"child-add"}),r("div",{class:"mb expanded"},r("label",{for:"category-"+e+"-select"},r("em",{class:"title"},"Category")),r("div",null,r("select",{name:"category",id:"category-"+e+"-select"},r("option",{selected:"selected",value:"TODO"},"TODO")))),r("div",{class:"mb expanded"},r("label",{for:"feedback-add-content-"+e},r("em",{class:"title"},"Content")),r("div",null,r("textarea",{rows:"8",id:"feedback-add-content-"+e,name:"content",placeholder:"HTML and Markdown supported"}))),r("button",{type:"submit"},"Add Feedback"))))))}function bt(e){return r("div",{id:"feedback-"+e.id,class:"feedback mt clear"},r("div",{class:"right"},A("feedback",e.id)),R("feedback",e.id,e.id),r("a",{href:"#modal-feedback-"+e.id,class:"clean modal-feedback-edit-link","data-id":e.id},r("div",null,T(e.userID)),r("div",{class:"pt feedback-content"},e.html)))}function yt(e){return r("div",{id:"modal-feedback-"+e.id,class:"modal modal-feedback-view",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,e.category+" :: "+T(e.userID))),r("div",{class:"modal-body",dangerouslySetInnerHTML:e.html})))}var Jt="return confirm('Are you sure you want to delete this feedback?');";function Xt(e){return l(".category").map(t=>t.dataset.category).map(t=>t===e?r("option",{value:t,selected:"selected"},t):r("option",{value:t},t))}function Tt(e){return r("div",{id:"modal-feedback-"+e.id,class:"modal modal-feedback-edit",style:"display: none;"},r("a",{class:"backdrop",href:"#"}),r("div",{class:"modal-content"},r("div",{class:"modal-header"},r("a",{href:"#",class:"modal-close"},"\xD7"),r("h2",null,e.category+" :: "+T(e.userID))),r("div",{class:"modal-body"},r("form",{action:"#",method:"post"},r("input",{type:"hidden",name:"feedbackID",value:e.id}),r("div",{class:"mb expanded"},r("label",{for:"input-category-"+e.id},r("em",{class:"title"},"Category")),r("div",null,r("select",{id:"input-category-"+e.id,name:"category"},Xt(e.category)))),r("div",{class:"mb expanded"},r("label",{for:"input-content-"+e.id},r("em",{class:"title"},"Content")),r("div",null,r("textarea",{id:"input-content-"+e.id,name:"content"},e.content))),r("div",{class:"right"},r("button",{class:"feedback-edit-save",type:"submit",name:"action",value:"child-update"},"Save Changes")),r("button",{class:"feedback-edit-delete",type:"submit",name:"action",value:"child-remove",onclick:Jt},"Delete")))))}function en(e){let t=o("form",e);t.onsubmit=()=>{let n=o('select[name="category"]',t).value,i=o('textarea[name="content"]',t),s=i.value;return i.value="",p("child-add",{category:n,content:s}),document.location.hash="",!1}}function Mt(e){let t=o("form",e),n=o('input[name="feedbackID"]',t).value;o(".feedback-edit-delete",t).onclick=()=>(confirm("Are you sure you want to delete this feedback?")&&(p("child-remove",{feedbackID:n}),document.location.hash=""),!1),t.onsubmit=()=>{let i=o('select[name="category"]',t).value,a=o('textarea[name="content"]',t).value;return p("child-update",{feedbackID:n,category:i,content:a}),document.location.hash="",!1}}function Et(){l(".add-feedback-link").forEach(e=>{e.onclick=()=>v(o("#feedback-add-content-"+e.dataset.category))}),l(".modal-feedback-edit-link").forEach(e=>{e.onclick=()=>v(o("#input-content-"+e.dataset.id))});for(let e of l(".modal-feedback-add"))en(e);l(".modal-feedback-edit").forEach(Mt)}function kt(e){let t=u("#category-"+e.category+" .feedback-list");if(!t){let s=vt(e.category);o("#category-list").appendChild(s),t=o(".feedback-list",s)}let n=-1;for(let s=0;s<t.children.length;s++){let a=t.children.item(s),m=o(".feedback-content",a).innerText,c=a.dataset.index;if(c){if(parseInt(c,10)>=e.idx){n=s;break}else if(m.localeCompare(e.content,void 0,{sensitivity:"accent"})>=0){n=s;break}}}let i=bt(e);if(n===-1?t.appendChild(i):t.insertBefore(i,t.children[n]),y()===e.userID){let s=Tt(e);Mt(s),o("#feedback-modals").appendChild(s)}else o("#feedback-modals").appendChild(yt(e));D(o(".modal",i))}function Lt(e){o("#feedback-"+e).remove(),E(e+"-removed","success","feedback has been removed"),o("#modal-feedback-"+e).remove()}function xt(){L("retro");let e=u("#modal-retro-config form");if(e){let t=o('select[name="team"]',e);t.onchange=()=>{F(t.value!=="")};let n=o('select[name="sprint"]',e);n.onchange=()=>{$(n.value!=="")},U(t,n),e.onsubmit=()=>{let i=o('input[name="title"]',e).value,s=o('input[name="icon"]:checked',e).value,a=o('input[name="categories"]',e).value;return p("update",{title:i,icon:s,categories:a,team:t.value,sprint:n.value,...H(e)}),document.location.hash="",!1}}Et()}function tn(e){let t=o("#modal-retro-config"),n=u("form",t);if(n){let i=o('input[name="categories"]',n);i.value=e.categories.join(","),i.parentElement&&O(i.parentElement)}else o(".config-panel-categories",t).innerText=e.categories.join(", ");k("retro",t,e.teamID,e.sprintID,e.title,e.icon)}function Ht(e){switch(e.cmd){case"update":return tn(e.param);case"child-add":return kt(e.param);case"child-update":return console.log("TODO: child-update");case"child-remove":return Lt(e.param);case"permissions":return x(e.param);default:throw new Error("invalid retro command ["+e.cmd+"]")}}function nn(e){E("adhoc-message",e.level,e.message)}function It(e,t){switch(t.cmd){case"reset":document.location.href="/";return;case"message":return nn(t.param);case"comment":return le(t.param);case"online-update":return Be(t.param);case"member-add":return Ve(t.param);case"member-update":return se(t.param);case"member-remove":return _e(t.param)}switch(e){case"team":return Qe(t);case"sprint":return ze(t);case"estimate":return st(t);case"standup":return ht(t);case"retro":return Ht(t);default:throw new Error("invalid service ["+e+"]")}}var wt,j,X;function on(){console.log("[socket]: open")}function rn(e){let t=document.getElementById("socket-list");if(t){let n=document.createElement("pre");n.innerText=JSON.stringify(e,null,2),t.append(n)}It(j,e)}function sn(e){console.log(`[socket error]: ${e}`)}function an(e,t){switch(j=e,X=t,Ye(),We(),j){case"team":Ke();break;case"sprint":Ze();break;case"estimate":it();break;case"standup":gt();break;case"retro":xt();break}wt=new Q(!0,on,rn,sn,"/"+j+"/"+X+"/connect"),console.log("loaded ["+j+"] workspace ["+X+"]")}function p(e,t){wt.send({channel:j+":"+X,cmd:e,param:t})}function Ct(){window.initWorkspace=an}function ln(){let[e,t]=De();window.rituals={relativeTime:Ee(),autocomplete:ke(),setSiblingToNull:e,initForm:t,flash:ye(),tags:we(),Socket:$e()},he(),ve(),Te(),Le(),Ae(),window.JSX=r,Ct()}document.addEventListener("DOMContentLoaded",ln);})();
//# sourceMappingURL=client.js.map
