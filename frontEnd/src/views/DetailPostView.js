/* eslint-disable */

import React from 'react'
import { createScope, map, transformProxies } from './helpers'

const scripts = [
  fetch("https://d3e54v103j8qbb.cloudfront.net/js/jquery-3.5.1.min.dc5e7f18c8.js?site=5f8e5f029f2850ba2966936f").then(body => body.text()),
  fetch("js/webflow.js").then(body => body.text()),
]

let Controller

class DetailPostView extends React.Component {
  static get Controller() {
    if (Controller) return Controller

    try {
      Controller = require('../controllers/DetailPostController')
      Controller = Controller.default || Controller

      return Controller
    }
    catch (e) {
      if (e.code == 'MODULE_NOT_FOUND') {
        Controller = DetailPostView

        return Controller
      }

      throw e
    }
  }

  componentDidMount() {
    scripts.concat(Promise.resolve()).reduce((loaded, loading) => {
      return loaded.then((script) => {
        new Function(`
          with (this) {
            eval(arguments[0])
          }
        `).call(window, script)

        return loading
      })
    })
  }

  render() {
    const proxies = Controller !== DetailPostView ? transformProxies(this.props.children) : {

    }

    return (
      <span>
        <style dangerouslySetInnerHTML={{ __html: `
          @import url(/css/normalize.css);
          @import url(/css/webflow.css);
          @import url(/css/deepfakes.webflow.css);
        ` }} />
        <span className="af-view">
          <div>
            <div data-collapse="medium" data-animation="default" data-duration={400} role="banner" className="af-class-navigation w-nav">
              <div className="af-class-navigation-wrap">
                <a href="index.html" className="af-class-logo-link w-nav-brand"><img src="images/Deep-Fakes.png" loading="lazy" height={0} alt /></a>
                <div className="af-class-menu">
                  <nav role="navigation" className="af-class-navigation-items w-nav-menu">
                    <a href="index.html" className="af-class-navigation-item w-nav-link">Home</a>
                    <a href="about.html" className="af-class-navigation-item w-nav-link">About</a>
                    <a href="community.html" className="af-class-navigation-item w-nav-link">Community</a>
                    <a href="api.html" className="af-class-navigation-item w-nav-link">api</a>
                  </nav>
                  <div className="af-class-menu-button w-nav-button"><img src="images/menu-icon_1menu-icon.png" width={22} alt className="af-class-menu-icon" /></div>
                </div>
                <a href="mailto:lucerodavi96@gmail.com?subject=DeepFakes Inquiry!" className="af-class-button af-class-cc-contact-us w-inline-block">
                  <div>join us</div>
                </a>
              </div>
            </div>
            <div className="af-class-section">
              <div className="af-class-container af-class-cc-blog-detail">
                <div className="af-class-blog-detail-header-wrap">
                  <h1 />
                  <div className="af-class-label af-class-cc-blog-date" />
                </div>
                <div className="af-class-detail-header-image" />
                <div className="af-class-rich-text w-richtext" />
              </div>
            </div>
            <div className="af-class-section af-class-cc-cta">
              <div className="af-class-container">
                <div className="af-class-cta-wrap">
                  <div>
                    <div className="af-class-cta-text">
                      <div className="af-class-heading-jumbo-small">Grow your business.<br /></div>
                      <div className="af-class-paragraph-bigger af-class-cc-bigger-light">Today is the day to build the business of your dreams. Share your mission with the world — and blow your customers away.<br /></div>
                    </div>
                    <a href="api.html" className="af-class-button af-class-cc-jumbo-button w-inline-block">
                      <div>Start Now</div>
                    </a>
                  </div>
                </div>
              </div>
            </div>
            <div className="af-class-section">
              <div className="af-class-container">
                <div className="af-class-footer-wrap">
                  <a href="https://webflow.com/" target="_blank" className="af-class-webflow-link w-inline-block" />
                </div>
              </div>
            </div>
            {/* [if lte IE 9]><![endif] */}
          </div>
        </span>
      </span>
    )
  }
}

export default DetailPostView

/* eslint-enable */