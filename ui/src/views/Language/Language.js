import axios from "axios";
import React, { Component } from "react";
import {
  Card,
  CardBody,
  CardHeader,
  Col,
  Pagination,
  PaginationItem,
  PaginationLink,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Table,
  Button
} from "reactstrap";
import PropTypes from "prop-types";
import LanguageForm from "./LanguageForm";

class Language extends Component {
  constructor(props) {
    super(props);
    this.state = {
      languages: [],
      record: {},
      modal: false,
      warning: false,
      formVisible: false,
      actionForm: "",
      languageFormValues: {
        id: null,
        lang_code: null,
        country_code: null
      },
      URL_API: process.env.REACT_APP_API_URL + "locales",
      warningAPI: false,
      errorMessage: ""
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
  getLanguageList() {
    axios
      .get(this.state.URL_API)
      .then(res => {
        const languages = res.data;
        this.setState({ languages });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }
  componentDidMount() {
    this.getLanguageList();
  }
  handleDelete(id) {
    axios
      .delete(this.state.URL_API + `/${id}`)
      .then(() => {
        const { languages } = this.state;
        this.setState({ languages: languages.filter(rul => rul.id !== id) });
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
    const { languageFormValues, languages, actionForm, record } = this.state;

    if (actionForm !== "Add") {
      languageFormValues.id = record.id;
      axios
        .put(this.state.URL_API, languageFormValues)
        .then(() => {
          const index = languages.findIndex(lg => lg.id === record.id);
          if (index > -1) {
            languages[index] = languageFormValues;
            this.setState({ languages });
          }
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      let lastid = this.getLastID();
      axios
        .post(this.state.URL_API, languageFormValues)
        .then(response => {
          languageFormValues.id = lastid + 1;
          this.setState({ languages: [...languages, languageFormValues] });
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
      this.setState({
        languageFormValues: { id: null, lang_code: null, country_code: null }
      });
    }
    this.setState({ formVisible: false });
  }

  handleOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { languageFormValues } = this.state;

    this.setState({
      languageFormValues: {
        ...languageFormValues,
        [type]: value
      }
    });
  }
  handleCloseWarningAPI() {
    this.setState({ warningAPI: false });
  }
  getLastID() {
    const { languages } = this.state;
    let lastid = 0;
    if (languages.length > 0) {
      lastid = languages[languages.length - 1].id;
    }
    return lastid;
  }
  render() {
    const { languages } = this.state;
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>Language</CardHeader>
            <CardBody>
              <div style={{ marginBottom: ".5rem" }}>
                <Button color="primary" onClick={() => this.showForm()}>
                  <i class="fa fa-plus" />
                  &nbsp;New Language
                </Button>
              </div>
              <Table responsive bordered>
                <thead>
                  <tr>
                    <th>Language Code</th>
                    <th>Country Code</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  {languages.length > 0 ? (
                    languages.map((language, index) => (
                      <tr key={index}>
                        <td>{language.lang_code}</td>
                        <td>{language.country_code}</td>
                        <td>
                          <Button
                            color="secondary"
                            onClick={() => this.showForm(language)}
                          >
                            <i class="fa fa-pencil" />
                            &nbsp; Edit
                          </Button>
                          {"  "}
                          <Button color="danger" onClick={this.toggleWarning}>
                            <i class="fa fa-trash" />
                            &nbsp; Delete
                          </Button>
                          <Modal
                            isOpen={this.state.warning}
                            toggle={this.toggleWarning}
                            className={"modal-warning " + this.props.className}
                          >
                            <ModalHeader toggle={this.toggleWarning}>
                              Delete Confirmation
                            </ModalHeader>
                            <ModalBody>
                              Are you sure want to delete language{" "}
                              {language.lang_code} ?
                            </ModalBody>
                            <ModalFooter>
                              <Button
                                color="warning"
                                onClick={() => this.handleDelete(language.id)}
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
                      <td colSpan={5}>No Data Language</td>
                    </tr>
                  )}
                </tbody>
              </Table>
              <Pagination>
                <PaginationItem>
                  <PaginationLink previous tag="button">
                    Prev
                  </PaginationLink>
                </PaginationItem>
                <PaginationItem active>
                  <PaginationLink tag="button">1</PaginationLink>
                </PaginationItem>
                <PaginationItem>
                  <PaginationLink next tag="button">
                    Next
                  </PaginationLink>
                </PaginationItem>
              </Pagination>
            </CardBody>
          </Card>

          <LanguageForm
            wrappedComponentRef={this.saveFormRef}
            visible={this.state.formVisible}
            onCancel={this.handleCancel}
            onSave={this.handleSave}
            language={this.state.record}
            action={this.state.actionForm}
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
Language.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string
  }).isRequired
};

export default Language;
