import axios from "axios";
import PropTypes from "prop-types";
import React, { Component } from "react";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Table,
  NavLink
} from "reactstrap";
import RuleForm from "./RuleForm";

class RuleList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      rules: [],
      record: {},
      modal: false,
      warning: false,
      formVisible: false,
      actionForm: "",
      ruleFormValues: {
        id: null,
        name: null,
        url_pattern: null,
        data_source_id: null
      },
      URL_API: process.env.REACT_APP_API_URL + "rules",
      URL_API_DATASOURCES: process.env.REACT_APP_API_URL + "data_sources",
      warningAPI: false,
      errorMessage: "",
      dataSources: []
    };

    this.handleDelete = this.handleDelete.bind(this);
    this.handleSave = this.handleSave.bind(this);
    this.handleCancel = this.handleCancel.bind(this);

    this.toggleWarning = this.toggleWarning.bind(this);
    this.saveFormRef = this.saveFormRef.bind(this);

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
  getRuleList() {
    axios
      .get(this.state.URL_API)
      .then(res => {
        const rules = res.data;
        this.setState({ rules });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }
  componentDidMount() {
    this.getRuleList();
    axios
      .get(this.state.URL_API_DATASOURCES)
      .then(res => {
        const dataSources = res.data;
        this.setState({ dataSources });
      })
      .catch(error => {});
  }

  handleDelete(id) {
    axios
      .delete(this.state.URL_API + `/${id}`)
      .then(() => {
        const { rules } = this.state;
        this.setState({ rules: rules.filter(rul => rul.id !== id) });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
    this.toggleWarning();
  }
  showForm(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionForm: "Edit" });
    } else {
      this.setState({ record: {} });
      this.setState({ actionForm: "Add" });
    }
    this.setState({ formVisible: true });
  }

  saveFormRef(formRef) {
    this.formRef = formRef;
  }
  handleCancel() {
    this.setState({ formVisible: false });
  }

  handleSave() {
    const { ruleFormValues, rules, actionForm, record } = this.state;

    ruleFormValues.data_source_id = parseInt(ruleFormValues.data_source_id);

    if (actionForm !== "Add") {
      ruleFormValues.id = record.id;
      axios
        .put(this.state.URL_API, ruleFormValues)
        .then(() => {
          const index = rules.findIndex(rul => rul.id === record.id);
          if (index > -1) {
            rules[index] = ruleFormValues;
            this.setState({ rules });
          }
        })
        .then(() => {
          this.getRuleList();
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      let lastid = this.getLastID();
      axios
        .post(this.state.URL_API, ruleFormValues)
        .then(response => {
          ruleFormValues.id = lastid + 1;
          this.setState({ rules: [...rules, ruleFormValues] });
        })
        //.then(() => {
        //  this.getRuleList();
        //})
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
      this.setState({
        ruleFormValues: {
          id: null,
          name: null,
          url_pattern: null,
          data_source_id: null
        }
      });
    }
    this.setState({ formVisible: false });
  }

  handleOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { ruleFormValues } = this.state;
    console.log(ruleFormValues, "onchange");
    this.setState({
      ruleFormValues: {
        ...ruleFormValues,
        [type]: value
      }
    });
  }

  handleDetail(record) {
    const { history } = this.props;
    /*history.push({
      pathname: "/ruleDetail?ruleId=" + record.id,
      data: record
    });*/
    history.push({
      pathname: "/ruleDetail/?ruleId=" + record.id,
      data: record
    });
  }

  handleCloseWarningAPI() {
    this.setState({ warningAPI: false });
  }
  getLastID() {
    const { rules } = this.state;
    let lastid = 0;
    if (rules.length > 0) {
      lastid = rules[rules.length - 1].id;
    }
    return lastid;
  }
  render() {
    const { rules } = this.state;
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>Rule</CardHeader>
            <CardBody>
              <div style={{ marginBottom: ".5rem" }}>
                <Button color="primary" onClick={() => this.showForm()}>
                  Add New
                </Button>
              </div>

              <Table responsive bordered>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>URL Pattern</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  {rules.length > 0 ? (
                    rules.map((rule, index) => (
                      <tr key={index}>
                        <td>{rule.id}</td>
                        <td>
                          <NavLink
                            href="#"
                            onClick={() => this.handleDetail(rule)}
                          >
                            {rule.name}
                          </NavLink>
                        </td>
                        <td>{rule.url_pattern}</td>
                        <td>{rule.updated_at}</td>
                        <td>
                          <button
                            className="button muted-button"
                            onClick={() => this.showForm(rule)}
                          >
                            Edit
                          </button>
                          {"  "}
                          <button
                            className="button muted-button"
                            onClick={this.toggleWarning}
                          >
                            Delete
                          </button>
                          <Modal
                            isOpen={this.state.warning}
                            toggle={this.toggleWarning}
                            className={"modal-warning " + this.props.className}
                          >
                            <ModalHeader toggle={this.toggleWarning}>
                              Delete Confirmation
                            </ModalHeader>
                            <ModalBody>
                              Are you sure want to delete {rule.name} ?
                            </ModalBody>
                            <ModalFooter>
                              <Button
                                color="warning"
                                onClick={() => this.handleDelete(rule.id)}
                              >
                                YES
                              </Button>{" "}
                              <Button
                                color="secondary"
                                onClick={this.toggleWarning}
                              >
                                NO
                              </Button>
                            </ModalFooter>
                          </Modal>
                        </td>
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

          <RuleForm
            wrappedComponentRef={this.saveFormRef}
            visible={this.state.formVisible}
            onCancel={this.handleCancel}
            onSave={this.handleSave}
            rule={this.state.record}
            action={this.state.actionForm}
            dataSources={this.state.dataSources}
            onChange={this.handleOnChange.bind(this)}
          />

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
RuleList.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string
  }).isRequired
};

export default RuleList;
