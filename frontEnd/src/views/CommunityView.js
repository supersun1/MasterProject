/* eslint-disable */

import React from 'react'
import { createScope, map, transformProxies } from './helpers'

const scripts = [
  fetch("https://d3e54v103j8qbb.cloudfront.net/js/jquery-3.5.1.min.dc5e7f18c8.js?site=5f8e5f029f2850ba2966936f").then(body => body.text()),
  fetch("js/webflow.js").then(body => body.text()),
]

let Controller

class CommunityView extends React.Component {
  static get Controller() {
    if (Controller) return Controller

    try {
      Controller = require('../controllers/CommunityController')
      Controller = Controller.default || Controller

      return Controller
    }
    catch (e) {
      if (e.code == 'MODULE_NOT_FOUND') {
        Controller = CommunityView

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
    const proxies = Controller !== CommunityView ? transformProxies(this.props.children) : {

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
                    <a href="community.html" aria-current="page" className="af-class-navigation-item w-nav-link w--current">Community</a>
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
                  <div className="af-class-heading-jumbo">Community<br /></div>
                </div>
              </div>
            </div>
            <div className="af-class-section">
              <div className="af-class-container">
                <div className="af-class-section-heading-wrap">
                  <div className="af-class-label af-class-cc-light">Our Team<br /></div>
                  <h2>Meet the creators<br /></h2>
                </div>
                <div className="w-layout-grid af-class-team-members">
                  <div id="w-node-6bd9fce9803c-e3b757c9">
                    <div className="af-class-team-pic"><img src="images/IMG_2036.jpg" loading="lazy" srcSet="images/IMG_2036-p-500.jpeg 500w, images/IMG_2036-p-800.jpeg 800w, images/IMG_2036-p-1080.jpeg 1080w, images/IMG_2036.jpg 1242w" sizes="(max-width: 479px) 94vw, (max-width: 767px) 96vw, 94vw" alt className="af-class-image-6" /></div>
                    <div className="af-class-team-member-title-wrap">
                      <div className="af-class-team-member-name">Lucero Davalos</div>
                      <div className="af-class-paragraph-light"></div>
                    </div>
                  </div>
                  <div id="w-node-6bd9fce98043-e3b757c9">
                    <div className="af-class-team-pic"><img src="images/IMG_2038.jpg" loading="lazy" srcSet="images/IMG_2038-p-500.jpeg 500w, images/IMG_2038-p-800.jpeg 800w, images/IMG_2038-p-1080.jpeg 1080w, images/IMG_2038.jpg 1242w" sizes="(max-width: 479px) 94vw, (max-width: 767px) 96vw, 94vw" alt className="af-class-image-5" /></div>
                    <div className="af-class-team-member-title-wrap">
                      <div className="af-class-team-member-name">Konstantin Pavlenko</div>
                      <div className="af-class-paragraph-light"></div>
                    </div>
                  </div>
                  <div id="w-node-6bd9fce9804a-e3b757c9">
                    <div className="af-class-team-pic"><img src="images/IMG_2037.jpg" loading="lazy" srcSet="images/IMG_2037-p-500.jpeg 500w, images/IMG_2037-p-800.jpeg 800w, images/IMG_2037-p-1080.jpeg 1080w, images/IMG_2037.jpg 1242w" sizes="(max-width: 479px) 94vw, (max-width: 767px) 96vw, 94vw" alt className="af-class-image-4" /></div>
                    <div className="af-class-team-member-title-wrap">
                      <div className="af-class-team-member-name">William Forsythe</div>
                      <div className="af-class-paragraph-light"></div>
                    </div>
                  </div>
                  <div id="w-node-6bd9fce98051-e3b757c9">
                    <div className="af-class-team-pic"><img src="images/IMG_2039.jpg" loading="lazy" srcSet="images/IMG_2039-p-500.jpeg 500w, images/IMG_2039-p-800.jpeg 800w, images/IMG_2039-p-1080.jpeg 1080w, images/IMG_2039.jpg 1242w" sizes="(max-width: 479px) 94vw, (max-width: 767px) 96vw, 94vw" alt className="af-class-image-8" /></div>
                    <div className="af-class-team-member-title-wrap">
                      <div className="af-class-team-member-name">Dylan Sun</div>
                      <div className="af-class-paragraph-light"></div>
                    </div>
                  </div>
                  <div id="w-node-6bd9fce98058-e3b757c9">
                    <div className="af-class-team-pic"><img src="images/IMG_2040.jpg" loading="lazy" srcSet="images/IMG_2040-p-500.jpeg 500w, images/IMG_2040-p-800.jpeg 800w, images/IMG_2040-p-1080.jpeg 1080w, images/IMG_2040.jpg 1242w" sizes="(max-width: 479px) 94vw, (max-width: 767px) 96vw, 94vw" alt className="af-class-image-7" /></div>
                    <div className="af-class-team-member-title-wrap">
                      <div className="af-class-team-member-name">Tommy Yao</div>
                      <div className="af-class-paragraph-light"></div>
                    </div>
                  </div>
                  <div id="w-node-6bd9fce9805f-e3b757c9">
                    <div className="af-class-team-member-title-wrap" />
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

export default CommunityView

/* eslint-enable */