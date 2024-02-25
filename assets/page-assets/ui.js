"use strict"
import * as _ from "https://cdn.jsdelivr.net/npm/markdown-it@14.0.0/dist/markdown-it.min.js";

// globals
const TEXT_CHANGE_EVENT = "textChanged";
const LOCATION_QUERY_PARAM_HEADER = "l"
const NAV_ID_PREFIX = "nav-location-"
const DEFAULT_NAV_ITEM = "home"
const MARKDOWN_FILE_MAP = {
    "home": "./README.md",
    "docs": "./docs/man.md",
    "changelogs": "./docs/CHANGELOG.md",
    "credits": "./docs/credits.md",
}

const md = window.markdownit({
    html: true,
    xhtmlOut: true,
    breaks: false,
    linkify: false, // Autoconvert URL-like text to links
    typographer: true,
    highlight: function (str, lang) {
        var esc = md.utils.escapeHtml;
        return '<pre class="big-code-bg"><code>' + esc(str) + '</code></pre>';

    },
})

// doing every content body changes via this one 
// (so that the loader placeholder can handle indepedently)
function _changeContent(htmlContent = "", domElem = document.getElementById("content-body")) {
    domElem.innerHTML = htmlContent;
    domElem.dispatchEvent(new Event(TEXT_CHANGE_EVENT))
}


// load markdown as html
function loadHtmlFromMarkdown(markdownFile) {
    if (!markdownFile) markdownFile = "./README.md"
    fetch(markdownFile)
        .then(response => response.text())
        .then(text => {
            const html = md.render(text)
            _changeContent(html)
        })
        .catch(err => {
            console.log("error : ", err)
        })
}

// bind listsners to automatically hide and show loader place holder in content body
function bindLoadingPlaceHolder() {
    const loaderPlaceHolder = document.getElementById("loader-placeholder")
    const contentBody = document.getElementById("content-body")

    contentBody.addEventListener(TEXT_CHANGE_EVENT, () => {
        if (contentBody.textContent.trim()) {
            loaderPlaceHolder.classList.add("d-none")
        } else {
            loaderPlaceHolder.classList.remove("d-none")
        }
    })
}

// triger navigation
function initiateNavigation() {
    // triger initiation
    const params = new URLSearchParams(window.location.search)
    try {
        document.getElementById(NAV_ID_PREFIX + params.get(LOCATION_QUERY_PARAM_HEADER)).click()
    } catch {
        // handle not found case
        document.getElementById(NAV_ID_PREFIX + DEFAULT_NAV_ITEM).click()
    }
}

// bind listeners to the nav buttons :)
function renderNavItems() {
    const navItems = document.querySelectorAll(".navitems")
    document.querySelectorAll(".navitems").forEach(elemnt => elemnt.addEventListener("click", (e) => {
        if (e.target.classList.contains("enabled") || e.target.classList.contains("heading"))
            return

        let navItem = e.target.id.slice(NAV_ID_PREFIX.length)
        if (!(navItem in MARKDOWN_FILE_MAP))
            navItem = DEFAULT_NAV_ITEM

        navItems.forEach((dom) => {
            if (dom.classList.contains("enabled"))
                dom.classList.remove("enabled")
        })
        e.target.classList.add("enabled")
        let queryParams = new URLSearchParams(window.location.search)
        queryParams.set(LOCATION_QUERY_PARAM_HEADER, navItem)
        window.history.pushState(null, null, "?" + queryParams.toString())
        loadHtmlFromMarkdown(MARKDOWN_FILE_MAP[navItem])
    }))
}

// append footer link for each page for seo
function appendFooterLinks() {
    const footer = document.getElementsByTagName("footer-links")[0]
    for (let key in MARKDOWN_FILE_MAP) {
        const elm = document.createElement("a");
        elm.innerText = key

        let loc = window.location.pathname
        if (loc[loc.length-1] == "/") {
            loc = loc.slice(0, -1)
        }  
        elm.href = `${window.location.origin}${loc}?${LOCATION_QUERY_PARAM_HEADER}=${key}`
        footer.appendChild(elm);
    } 
}

// starts here
window.addEventListener("popstate", initiateNavigation)
window.onload = () => {
    bindLoadingPlaceHolder();
    renderNavItems();
    initiateNavigation();
    appendFooterLinks();
}