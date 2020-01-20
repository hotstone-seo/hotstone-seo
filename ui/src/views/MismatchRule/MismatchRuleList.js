import axios from "axios";
import PropTypes from "prop-types";
import React, { Component } from "react";
import {
  Card,
  CardBody,
  CardHeader,
  Col,
  Modal,
  ModalBody,
  ModalHeader,
  Table
} from "reactstrap";
import { format, formatDistance } from "date-fns";

class MismatchRuleList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      mismatchRules: [],
      record: {},
      warning: false,
      URL_API: process.env.REACT_APP_API_URL + "metrics/mismatched",
      errorMessage: ""
    };

    this.toggleWarning = this.toggleWarning.bind(this);

    this.handleCloseWarningAPI = this.handleCloseWarningAPI.bind(this);
    this.toggleWarningAPI = this.toggleWarningAPI.bind(this);
  }
  toggle() {
    this.setState({
      modal: !this.state.modal
    });
  }
  toggleWarning() {
    this.setState({
      warning: !this.state.warning
    });
  }
  toggleWarningAPI(errmsg) {
    this.setState({
      warningAPI: !this.state.warningAPI,
      errorMessage: errmsg
    });
  }
  getList() {
    axios
      .get(this.state.URL_API)
      .then(res => {
        const mismatchRules = res.data;
        this.setState({ mismatchRules });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }
  componentDidMount() {
    this.getList();
  }

  handleCancel() {
    this.setState({ formVisible: false });
  }

  handleCloseWarningAPI() {
    this.setState({ warningAPI: false });
  }

  formatSince(since) {
    const sinceDate = new Date(since);

    const full = format(sinceDate, "yyyy-MM-dd hh:mm");
    const relative = formatDistance(sinceDate, new Date());

    return `${full} (${relative} ago)`;
  }

  render() {
    const { mismatchRules } = this.state;
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>Mismatch Rule</CardHeader>
            <CardBody>
              <Table responsive bordered>
                <thead>
                  <tr>
                    <th>No</th>
                    <th>URL</th>
                    <th>Since</th>
                    <th>Count</th>
                  </tr>
                </thead>
                <tbody>
                  {mismatchRules.length > 0 ? (
                    mismatchRules.map((mismatchRule, index) => (
                      <tr key={index}>
                        <td>{++index}</td>
                        <td>{mismatchRule.request_path}</td>
                        <td>{this.formatSince(mismatchRule.since)}</td>
                        <td>{mismatchRule.count}</td>
                      </tr>
                    ))
                  ) : (
                    <tr>
                      <td colSpan={5}>No Rule</td>
                    </tr>
                  )}
                </tbody>
              </Table>
            </CardBody>
          </Card>

          <Modal
            isOpen={this.state.warningAPI}
            toggle={this.toggleWarningAPI}
            className={"modal-warning " + this.props.className}
          >
            <ModalHeader toggle={() => this.toggleWarningAPI("")}>
              Information
            </ModalHeader>
            <ModalBody>
              <span>{this.state.errorMessage}</span>
              <br></br>
              <span>
                Sorry, failed to connect API. API currently not available/API in
                problem
              </span>
            </ModalBody>
          </Modal>
        </Col>
      </div>
    );
  }
}
MismatchRuleList.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string
  }).isRequired
};

export default MismatchRuleList;
