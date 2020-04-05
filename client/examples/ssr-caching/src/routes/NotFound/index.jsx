import React, { Component } from 'react';
import { canUseDOM } from 'fbjs/lib/ExecutionEnvironment';

class NotFound extends Component {
  /* istanbul ignore next */
  goBack () {
    if (canUseDOM) {
      window.history.go(-1);
    }
  }

  render() {
    return (
      <div className="container">
        <div className="row" style={{marginTop: 80}}>
          <div className="col-12 text-center">
            <h1>Ooopss... Are you lost ?</h1>
            <div style={{marginTop: 20}}>
              <a className="btn primary" onClick={this.goBack}>Go back</a>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default NotFound;