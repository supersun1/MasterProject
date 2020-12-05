/* eslint-disable */

import React from 'react'
import { createScope, map, transformProxies } from './helpers'

const scripts = [
  fetch("https://d3e54v103j8qbb.cloudfront.net/js/jquery-3.5.1.min.dc5e7f18c8.js?site=5f8e5f029f2850ba2966936f").then(body => body.text()),
  fetch("js/webflow.js").then(body => body.text()),
]

let Controller

class StyleguideView extends React.Component {
  static get Controller() {
    if (Controller) return Controller

    try {
      Controller = require('../controllers/StyleguideController')
      Controller = Controller.default || Controller

      return Controller
    }
    catch (e) {
      if (e.code == 'MODULE_NOT_FOUND') {
        Controller = StyleguideView

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
    const proxies = Controller !== StyleguideView ? transformProxies(this.props.children) : {

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
              <div className="af-class-styleguide-header-wrap">
                <div className="af-class-heading-jumbo">Styleguide</div>
                <div className="af-class-paragraph-bigger af-class-cc-bigger-light">Source of truth of this template.</div>
              </div>
              <div className="af-class-container">
                <div className="af-class-styleguide-content-wrap">
                  <div className="af-class-styleguide-block">
                    <div className="af-class-label af-class-cc-styleguide-label">Headings</div>
                    <div className="af-class-heading-jumbo">Heading Jumbo</div>
                    <div className="af-class-heading-jumbo-small">Heading Jumbo</div>
                    <div className="af-class-heading-jumbo-tiny">Heading Jumbo</div>
                    <h1>Heading 1</h1>
                    <h2>Heading 2</h2>
                    <h3>Heading 3</h3>
                    <h4>Heading 4</h4>
                    <h5>Heading 5</h5>
                    <h6>Heading 6</h6>
                  </div>
                  <div className="af-class-divider" />
                  <div className="af-class-styleguide-block">
                    <div className="af-class-label af-class-cc-styleguide-label">Paragraphs &amp; Links</div>
                    <p className="af-class-paragraph-bigger">Paragraph-bigger</p>
                    <p>Paragraph</p>
                    <p className="af-class-paragraph-light">Paragraph-de-emphasized</p>
                    <p className="af-class-paragraph-small">Paragraph-smaller</p>
                    <p className="af-class-paragraph-tiny">Paragraph-tiny</p>
                    <div className="af-class-label">Label</div>
                    <a href="#">A text link</a>
                  </div>
                  <div className="af-class-divider" />
                  <div className="af-class-styleguide-block">
                    <div className="af-class-label af-class-cc-styleguide-label">Buttons &amp; Form&nbsp;Elements</div>
                    <div className="af-class-styleguide-button-wrap">
                      <a href="#" className="af-class-button w-inline-block">
                        <div>Explore</div>
                      </a>
                    </div>
                    <div className="af-class-styleguide-button-wrap">
                      <a href="#" className="af-class-button af-class-cc-jumbo-button w-inline-block">
                        <div>Explore</div>
                      </a>
                    </div>
                  </div>
                  <div className="af-class-divider" />
                  <div className="af-class-styleguide-block">
                    <div className="af-class-label af-class-cc-styleguide-label">Rich Text Element</div>
                    <div className="af-class-rich-text w-richtext">
                      <h2>What’s a Rich Text element?</h2>
                      <p>The rich text element allows you to create and format headings, paragraphs, blockquotes, images, and video all in one place instead of having to add and format them individually. Just double-click and easily create content. </p>
                      <blockquote>The rich text element allows you to create and format headings, paragraphs, blockquotes, images, and video all in one place instead of having to add and format them individually. Just double-click and easily create content.</blockquote>
                      <h4>Static and dynamic content editing</h4>
                      <p>A rich text element can be used with static or dynamic content. For static content, just drop it into any page and begin editing. For dynamic content, add a rich text field to any collection and then connect a rich text element to that field in the settings panel. Voila!</p>
                      <figure style={{maxWidth: 1306}} id="w-node-e50305e3e80a-35b757c5" className="w-richtext-align-fullwidth w-richtext-figure-type-image">
                        <div><img src="images/placeholder-1.svg" alt /></div>
                      </figure>
                      <h4>How to customize formatting for each rich text</h4>
                      <p>Headings, paragraphs, blockquotes, figures, images, and figure captions can all be styled after a class is added to the rich text element using the "When inside of" nested selector system.</p>
                      <ul role="list">
                        <li>The rich text element allows you to create and format headings</li>
                        <li>Headings, paragraphs, blockquotes, figures, images, and figure captions</li>
                        <li>A rich text element can be used with static or dynamic content.</li>
                      </ul>
                    </div>
                  </div>
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

export default StyleguideView

/* eslint-enable */