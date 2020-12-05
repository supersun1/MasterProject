/* eslint-disable */

import React from 'react'
import { createScope, map, transformProxies } from './helpers'

const scripts = [
  fetch("https://d3e54v103j8qbb.cloudfront.net/js/jquery-3.5.1.min.dc5e7f18c8.js?site=5f8e5f029f2850ba2966936f").then(body => body.text()),
  fetch("js/webflow.js").then(body => body.text()),
]

let Controller

class AboutView extends React.Component {
  static get Controller() {
    if (Controller) return Controller

    try {
      Controller = require('../controllers/AboutController')
      Controller = Controller.default || Controller

      return Controller
    }
    catch (e) {
      if (e.code == 'MODULE_NOT_FOUND') {
        Controller = AboutView

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
    const proxies = Controller !== AboutView ? transformProxies(this.props.children) : {

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
                    <a href="about.html" aria-current="page" className="af-class-navigation-item w-nav-link w--current">About</a>
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
            <div className="af-class-section af-class-cc-home-wrap">
              <div className="af-class-intro-header af-class-cc-subpage">
                <div className="af-class-intro-content">
                  <div className="af-class-heading-jumbo">About Us<br /></div>
                </div>
              </div>
              <div className="af-class-container">
                <div className="af-class-motto-wrap">
                  <div className="af-class-heading-jumbo-small">We’re a group of truth-seekers who want to share science-based knowledge with the world!</div>
                </div>
                <div className="af-class-divider" />
                <div className="af-class-about-story-wrap">
                  <p className="af-class-paragraph-light"><strong>What we do<br />‍</strong><br />We collect information from reliable sources, science-based and peer-reviewed research.<br />‍<br />We have built a state of the art Machine Learning Algorithm that is able to provide you with a percentage of data accuracy.<br /><br /><strong>How it works </strong><br />‍<br />Just submit the text you would like to analyze and click submit, then let our Algorithm do the work and provide you with an accuracy percentage in a matter of seconds.<br />‍<br /><strong>Why we do it<br /><br /></strong>We believe that reliable sources of information has become a real problem in this era, and misinformation can lead to many problems. We saw an issue that doesn't currently appear to have an adequate solution, so we decided to see if we could challenge ourselves to come up with a new innovative solution. <br /><br /><br /><br />‍</p>
                </div>
                <div className="af-class-divider" />
              </div>
            </div>
            <div className="af-class-section">
              <div className="af-class-container" />
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

export default AboutView

/* eslint-enable */