/* eslint-disable */

import React from 'react'
import { createScope, map, transformProxies } from './helpers'
import ContactFormView from './ContactFormView'

const scripts = [
  fetch("https://d3e54v103j8qbb.cloudfront.net/js/jquery-3.5.1.min.dc5e7f18c8.js?site=5f8e5f029f2850ba2966936f").then(body => body.text()),
  fetch("js/webflow.js").then(body => body.text()),
]

let Controller

class IndexView extends React.Component {
  static get Controller() {
    if (Controller) return Controller

    try {
      Controller = require('../controllers/IndexController')
      Controller = Controller.default || Controller

      return Controller
    }
    catch (e) {
      if (e.code == 'MODULE_NOT_FOUND') {
        Controller = IndexView

        return Controller
      }

      throw e
    }
  }
// commented out on  friday 27 at 5pm beacuse i was getting a werid error 
  componentDidMount() {
    // scripts.concat(Promise.resolve()).reduce((loaded, loading) => {
    //   return loaded.then((script) => {
    //     new Function(`
    //       with (this) {
    //         eval(arguments[0])
    //       }
    //     `).call(window, script)

    //     return loading
    //   })
    // })
  }

  render() {
    const proxies = Controller !== IndexView ? transformProxies(this.props.children) : {
      'contact-form': [],
    }

    return (
      <span>
        <style dangerouslySetInnerHTML={{ __html: `
          @import url(/css/normalize.css);
          @import url(/css/webflow.css);
          @import url(/css/deepfakes.webflow.css);
        ` }} />
        <span className="af-view">
          <div className="af-class-body">
            <div data-collapse="medium" data-animation="default" data-duration={400} role="banner" className="af-class-navigation w-nav">
              <div className="af-class-navigation-wrap">
                <a href="index.html" aria-current="page" className="af-class-logo-link w-nav-brand w--current"><img src="images/Deep-Fakes.png" loading="lazy" height={0} alt /></a>
                <div className="af-class-menu">
                  <nav role="navigation" className="af-class-navigation-items w-nav-menu">
                    <a href="index.html" aria-current="page" className="af-class-navigation-item w-nav-link w--current">Home</a>
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
            </div><img src="images/iconfinder_additional_icons-10_2239886.png" loading="eager" width={80} height={80} srcSet="images/iconfinder_additional_icons-10_2239886-p-500.png 500w, images/iconfinder_additional_icons-10_2239886.png 512w" sizes="(max-width: 479px) 100vw, 80px" alt className="af-class-image-3" />
            <div className="af-class-container-4 w-container"><img src="images/Deep-Fakes_1.png" loading="lazy" alt className="af-class-image-2" /></div>
            <div className="af-class-section" />
            <section id="contact-form" className="af-class-contact-form-2">
            {proxies['contact-form'] && <ContactFormView.Controller {...proxies['contact-form']}>{proxies['contact-form'].children}</ContactFormView.Controller>}

            </section>
            <footer id="footer" className="af-class-footer">
              <div className="w-container">
                <div className="af-class-footer-flex-container">
                  <a href="#" className="af-class-footer-logo-link w-inline-block"><img src="images/Deep-Fakes_1.png" alt className="af-class-footer-image" /></a>
                  <div>
                    <h2 className="af-class-footer-heading">Coming Soon</h2>
                    {/* <ul role="list" className="w-list-unstyled">
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                    </ul> */}
                  </div>
                  <div>
                    <h2 className="af-class-footer-heading">Coming Soon</h2>
                    {/* <ul role="list" className="w-list-unstyled">
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                    </ul> */}
                  </div>
                  <div>
                    <h2 className="af-class-footer-heading">Coming Soon</h2>
                    <ul role="list" className="w-list-unstyled">
                      {/* <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li>
                      <li>
                        <a href="#" className="af-class-footer-link">Text Link</a>
                      </li> */}
                    </ul>
                  </div>
                </div>
                <div>Copyright © 2020 DeepFakes. All rights reserved.</div>
              </div>
            </footer>
            {/* [if lte IE 9]><![endif] */}
          </div>
        </span>
      </span>
    )
  }
}

export default IndexView

/* eslint-enable */