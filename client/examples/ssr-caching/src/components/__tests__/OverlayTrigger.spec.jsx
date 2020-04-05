/* eslint-disable react/no-find-dom-node,react/no-render-return-value */
import PropTypes from 'prop-types';
import React, { cloneElement } from 'react';
import ReactDOM from 'react-dom';
import sinon from 'sinon';
import ReactTestUtils from 'react-dom/test-utils';

import OverlayTrigger from '../OverlayTrigger';
import Tooltip from '../Tooltip';

function customRender(element, mountPoint) {
  let mount = mountPoint || document.createElement('div');
  let instance = ReactDOM.render(element, mount);

  if (instance && !instance.renderWithProps) {
    instance.renderWithProps = newProps => customRender(cloneElement(element, newProps), mount);
  }

  return instance;
}

describe('<OverlayTrigger>', () => {
  // Swallow extra props.
  const Div = ({ className, children }) => <div className={className}>{children}</div>;

  Div.propTypes = {
    className: PropTypes.string,
    children: PropTypes.node
  };

  it('Should match snapshot', () => {
    const tree = render(
      <OverlayTrigger overlay={<Div>test</Div>}>
        <button>button</button>
      </OverlayTrigger>
    );

    expect(tree).toMatchSnapshot();
  });

  it('Should create OverlayTrigger element', () => {
    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger overlay={<Div>test</Div>}>
        <button>button</button>
      </OverlayTrigger>
    );
    const overlayTrigger = ReactDOM.findDOMNode(instance);

    expect(overlayTrigger.nodeName).toEqual('BUTTON');
  });

  it('Should pass OverlayTrigger onClick prop to child', () => {
    const callback = sinon.spy();
    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger overlay={<Div>test</Div>} onClick={callback}>
        <button>button</button>
      </OverlayTrigger>
    );
    const overlayTrigger = ReactDOM.findDOMNode(instance);

    ReactTestUtils.Simulate.click(overlayTrigger);
    expect(callback.calledOnce).toEqual(true);
  });

  it('Should show after click trigger', () => {
    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger trigger="click" overlay={<Div>test</Div>}>
        <button>button</button>
      </OverlayTrigger>
    );
    const overlayTrigger = ReactDOM.findDOMNode(instance);

    ReactTestUtils.Simulate.click(overlayTrigger);
    expect(instance.state.show).toBe(true);
  });

  it('Should not set aria-describedby if the state is not show', () => {
    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger trigger="click" overlay={<Div>test</Div>}>
        <button>button</button>
      </OverlayTrigger>
    );
    const overlayTrigger = ReactDOM.findDOMNode(instance);

    expect(overlayTrigger.getAttribute('aria-describedby')).toEqual(null);
  });

  it('Should set aria-describedby if the state is show', () => {
    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger trigger="click" overlay={<Div id="overlayid">test</Div>}>
        <button>button</button>
      </OverlayTrigger>
    );
    const overlayTrigger = ReactDOM.findDOMNode(instance);

    ReactTestUtils.Simulate.click(overlayTrigger);

    expect(overlayTrigger.getAttribute('aria-describedby')).toBeTruthy();
  });

  describe('trigger handlers', () => {
    let mountPoint;

    beforeEach(() => {
      mountPoint = document.createElement('div');
      document.body.appendChild(mountPoint);
    });

    afterEach(() => {
      ReactDOM.unmountComponentAtNode(mountPoint);
      document.body.removeChild(mountPoint);
    });

    it('Should keep trigger handlers', done => {
      const instance = customRender(
        <div>
          <OverlayTrigger trigger="focus" overlay={<Div>test</Div>}>
            <button onBlur={() => done()}>button</button>
          </OverlayTrigger>
          <input id="target" />
        </div>,
        mountPoint
      );

      const overlayTrigger = instance.firstChild;

      ReactTestUtils.Simulate.blur(overlayTrigger);
    });
  });

  it('Should maintain overlay classname', () => {
    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger trigger="click" overlay={<Div className="test-overlay">test</Div>}>
        <button>button</button>
      </OverlayTrigger>
    );

    const overlayTrigger = ReactDOM.findDOMNode(instance);

    ReactTestUtils.Simulate.click(overlayTrigger);

    expect(document.getElementsByClassName('test-overlay').length).toEqual(1);
  });

  it('Should pass transition callbacks to Transition', done => {
    let count = 0;
    const increment = () => count++;

    let overlayTrigger;

    const instance = ReactTestUtils.renderIntoDocument(
      <OverlayTrigger
        trigger="click"
        overlay={<Div>test</Div>}
        onExit={increment}
        onExiting={increment}
        onExited={() => {
          increment();
          expect(count).toEqual(6);
          done();
        }}
        onEnter={increment}
        onEntering={increment}
        onEntered={() => {
          increment();
          ReactTestUtils.Simulate.click(overlayTrigger);
        }}
      >
        <button>button</button>
      </OverlayTrigger>
    );

    overlayTrigger = ReactDOM.findDOMNode(instance);
    ReactTestUtils.Simulate.click(overlayTrigger);
  });

  it('Should forward requested context', () => {
    const contextTypes = {
      key: PropTypes.string
    };

    const contextSpy = sinon.spy();

    class ContextReader extends React.Component {
      render() {
        contextSpy(this.context.key);

        return <div />;
      }
    }

    ContextReader.contextTypes = contextTypes;

    class ContextHolder extends React.Component {
      getChildContext() {
        return { key: 'value' };
      }

      render() {
        return (
          <OverlayTrigger trigger="click" overlay={<ContextReader />}>
            <button>button</button>
          </OverlayTrigger>
        );
      }
    }
    ContextHolder.childContextTypes = contextTypes;

    const instance = ReactTestUtils.renderIntoDocument(<ContextHolder />);
    const overlayTrigger = ReactDOM.findDOMNode(instance);

    ReactTestUtils.Simulate.click(overlayTrigger);

    expect(contextSpy.calledWith('value')).toBe(true);
  });

  describe('overlay types', () => {
    [
      {
        name: 'Tooltip',
        overlay: <Tooltip id="test-tooltip">test</Tooltip>
      }
    ].forEach(testCase => {
      describe(testCase.name, () => {
        let instance;
        let overlayTrigger;

        beforeEach(() => {
          instance = ReactTestUtils.renderIntoDocument(
            <OverlayTrigger trigger="click" overlay={testCase.overlay}>
              <button>button</button>
            </OverlayTrigger>
          );
          overlayTrigger = ReactDOM.findDOMNode(instance);
        });

        it('Should handle trigger without warnings', () => {
          ReactTestUtils.Simulate.click(overlayTrigger);
        });
      });
    });
  });

  describe('rootClose', () => {
    [
      {
        label: 'true',
        rootClose: true,
        shownAfterClick: false
      },
      {
        label: 'default (false)',
        rootClose: null,
        shownAfterClick: true
      }
    ].forEach(testCase => {
      describe(testCase.label, () => {
        let instance;

        beforeEach(() => {
          instance = ReactTestUtils.renderIntoDocument(
            <OverlayTrigger
              overlay={<Div>test</Div>}
              trigger="click"
              rootClose={testCase.rootClose}
            >
              <button>button</button>
            </OverlayTrigger>
          );
          const overlayTrigger = ReactDOM.findDOMNode(instance);

          ReactTestUtils.Simulate.click(overlayTrigger);
        });

        it('Should have correct show state', () => {
          // Need to click this way for it to propagate to document element.
          document.documentElement.click();

          expect(instance.state.show).toEqual(testCase.shownAfterClick);
        });
      });
    });

    describe('clicking on trigger to hide', () => {
      let mountNode;

      beforeEach(() => {
        mountNode = document.createElement('div');
        document.body.appendChild(mountNode);
      });

      afterEach(() => {
        ReactDOM.unmountComponentAtNode(mountNode);
        document.body.removeChild(mountNode);
      });

      it('should hide after clicking on trigger', () => {
        const instance = ReactDOM.render(
          <OverlayTrigger overlay={<Div>test</Div>} trigger="click" rootClose>
            <button>button</button>
          </OverlayTrigger>,
          mountNode
        );

        const node = ReactDOM.findDOMNode(instance);

        expect(instance.state.show).toBe(false);

        node.click();
        expect(instance.state.show).toBe(true);

        // Need to click this way for it to propagate to document element.
        node.click();
        expect(instance.state.show).toBe(false);
      });
    });

    describe('replaced overlay', () => {
      let instance;

      beforeEach(() => {
        class ReplacedOverlay extends React.Component {
          constructor(props) {
            super(props);

            this.handleClick = this.handleClick.bind(this);
            this.state = { replaced: false };
          }

          handleClick() {
            this.setState({ replaced: true });
          }

          render() {
            if (this.state.replaced) {
              return <div>replaced</div>;
            }

            return (
              <div>
                <a id="replace-overlay" onClick={this.handleClick}>
                  original
                </a>
              </div>
            );
          }
        }

        instance = ReactTestUtils.renderIntoDocument(
          <OverlayTrigger overlay={<ReplacedOverlay />} trigger="click" rootClose>
            <button>button</button>
          </OverlayTrigger>
        );
        const overlayTrigger = ReactDOM.findDOMNode(instance);

        ReactTestUtils.Simulate.click(overlayTrigger);
      });

      it('Should still be shown', () => {
        // Need to click this way for it to propagate to document element.
        const replaceOverlay = document.getElementById('replace-overlay');

        replaceOverlay.click();

        expect(instance.state.show).toBe(true);
      });
    });
  });
});
