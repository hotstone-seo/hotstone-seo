import React, { Component } from "react";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Form,
  FormGroup,
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  ModalFooter,
  Pagination,
  PaginationItem,
  PaginationLink,
  Row,
  Table
} from "reactstrap";

import CanonicalForm from "../Canonical/CanonicalForm";
import MetaTagForm from "../Metatag/MetatagForm";
import ScriptTagForm from "../Scripttag/ScripttagForm";
import TitleTagForm from "../Titletag/TitletagForm";

import axios from "axios";

export const parseQuery = subject => {
  const results = {};
  const parser = /[^&?]+/g;
  let match = parser.exec(subject);
  while (match !== null) {
    const parts = match[0].split("=");
    results[parts[0]] = parts[1];
    match = parser.exec(subject);
  }
  return results;
};

class RuleDetail extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300,
      URL_API: process.env.REACT_APP_API_URL + "rules",
      URL_TAG_API: process.env.REACT_APP_API_URL + "tags",
      URL_ADDMETA_API: process.env.REACT_APP_API_URL + "center/addMetaTag",
      URL_ADDTITLE_API: process.env.REACT_APP_API_URL + "center/addTitleTag",
      URL_ADDCANONICAL_API:
        process.env.REACT_APP_API_URL + "center/addCanonicalTag",
      URL_ADDSCRIPT_API: process.env.REACT_APP_API_URL + "center/addScriptTag",

      canonicalFormValues: {
        id: null,
        name: null,
        rule_id: null
      },
      metaTagFormValues: {
        id: null,
        name: null,
        content: null,
        rule_id: null
      },
      scriptTagFormValues: {
        id: null,
        name: null,
        rule_id: null
      },
      titleTagFormValues: {
        id: null,
        name: null,
        rule_id: null
      },
      canonicalFormVisible: false,
      actionCanonicalForm: "",
      metaTagFormVisible: false,
      actionMetaTagForm: "",
      scriptTagFormVisible: false,
      actionScriptTagForm: "",
      titleTagFormVisible: false,
      actionTitleTagForm: "",
      ruleIdParam: 0,
      rules: [],
      warning: false,
      warningAPI: false,
      tags: []
    };
    this.handleEditCanonical = this.handleEditCanonical.bind(this);
    this.handleCancelAddCanonical = this.handleCancelAddCanonical.bind(this);

    this.handleCancelAddScriptTag = this.handleCancelAddScriptTag.bind(this);
    this.handleCancelAddTitleTag = this.handleCancelAddTitleTag.bind(this);

    this.handleDelete = this.handleDelete.bind(this);
  }
  componentDidMount() {
    const query = parseQuery((window.location || {}).search || "");
    const { ruleId } = query || {};
    const { rules } = this.state;
    axios
      .get(this.state.URL_API + `/${ruleId}`)
      .then(res => {
        const { data: resData } = res || {};
        const rulesdata = resData;
        this.setState({ rules: [...rules, rulesdata] });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });

    this.getTagList();
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState(prevState => {
      return { fadeIn: !prevState };
    });
  }

  handleEditCanonical() {
    const { history } = this.props;
    history.push("/canonicalEditForm");
  }

  showForm(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionForm: "Edit" });
    } else {
      this.setState({ record: {} });
      this.setState({ actionForm: "Add" });
    }
    this.setState({ canonicalFormVisible: true });
  }
  showFormMetaTag(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionMetaTagForm: "Edit" });
    } else {
      this.setState({ record: {} });
      this.setState({ actionMetaTagForm: "Add" });
    }
    this.setState({ metaTagFormVisible: true });
  }

  showFormScriptTag(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionForm: "Edit" });
    } else {
      this.setState({ record: {} });
      this.setState({ actionForm: "Add" });
    }
    this.setState({ scriptTagFormVisible: true });
  }

  showFormTitleTag(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionForm: "Edit" });
    } else {
      this.setState({ record: {} });
      this.setState({ actionForm: "Add" });
    }
    this.setState({ titleTagFormVisible: true });
  }

  handleOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { canonicalFormValues } = this.state;

    this.setState({
      canonicalFormValues: {
        ...canonicalFormValues,
        [type]: value
      }
    });
  }
  handleMetaTagOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { metaTagFormValues } = this.state;

    this.setState({
      metaTagFormValues: {
        ...metaTagFormValues,
        [type]: value
      }
    });
  }
  handleCancelAddCanonical() {
    this.setState({ canonicalFormVisible: false });
  }
  handleCancelAddMetaTag() {
    this.setState({ metaTagFormVisible: false });
  }
  handleCancelAddScriptTag() {
    this.setState({ scriptTagFormVisible: false });
  }
  handleCancelAddTitleTag() {
    this.setState({ titleTagFormVisible: false });
  }
  toggleWarningAPI(errmsg) {
    this.setState({
      warningAPI: !this.state.warningAPI,
      errorMessage: errmsg
    });
  }
  toggleWarning() {
    //this.setState({
    //  warning: !this.state.warning
    //});
  }

  handleDelete(id) {
    axios
      .delete(this.state.URL_TAG_API + `/${id}`)
      .then(() => {
        const { tags } = this.state;
        this.setState({ tags: tags.filter(tag => tag.id !== id) });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
    this.toggleWarning();
  }

  handleSaveMetaTag() {
    const { metaTagFormValues, tags, actionMetaTagForm, record } = this.state;
    const isUpdate = actionMetaTagForm !== "Add";

    metaTagFormValues.id = record.id;

    if (isUpdate) {
      axios
        .put(this.state.URL_ADDMETA_API, metaTagFormValues)
        .then(() => {
          const index = tags.findIndex(tg => tg.id === record.id);
          if (index > -1) {
            tags[index] = metaTagFormValues;
            this.setState({ tags });
          }
        })
        .then(() => {
          //this.getRuleList(); TO DO:
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      axios
        .post(this.state.URL_ADDMETA_API, metaTagFormValues)
        .then(response => {
          this.setState({ rules: [...tags, metaTagFormValues] });
        })
        .then(() => {
          //this.getRuleList();  TO DO :
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
      this.setState({ metaTagFormValues: {} });
    }
    this.setState({ formMetaTagVisible: false });
  }

  getTagList() {
    axios
      .get(this.state.URL_TAG_API)
      .then(res => {
        const tags = res.data;
        this.setState({ tags });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }

  render() {
    const { rules, tags } = this.state;
    return (
      <div className="animated fadeIn">
        {rules.map((rule, index) => (
          <Row key={index}>
            <Col xs="12" md="9" lg="6">
              <Card>
                <CardHeader>
                  <strong>Detail Rule ID {rule.id}</strong>
                </CardHeader>
                <CardBody>
                  <Form
                    action=""
                    method="post"
                    encType="multipart/form-data"
                    className="form-horizontal"
                  >
                    <FormGroup row>
                      <Col md="3">
                        <Label htmlFor="text-input">Name</Label>
                      </Col>
                      <Col xs="12" md="9">
                        {rule.name}
                      </Col>
                    </FormGroup>
                    <FormGroup row>
                      <Col md="3">
                        <Label htmlFor="text-input">URL Pattern</Label>
                      </Col>
                      <Col xs="12" md="9">
                        {rule.url_pattern}
                      </Col>
                    </FormGroup>

                    <FormGroup row>
                      <Col md="3">
                        <Label htmlFor="text-input">Data Source</Label>
                      </Col>
                      <Col xs="12" md="9">
                        Airport
                      </Col>
                    </FormGroup>
                  </Form>
                </CardBody>
              </Card>
            </Col>
          </Row>
        ))}

        <Row>
          <Col>
            <Card>
              <CardHeader>
                <i className="fa fa-align-justify"></i>
              </CardHeader>
              <CardBody>
                <div style={{ marginBottom: ".5rem" }}>
                  <Button
                    color="primary"
                    onClick={() => this.showForm()}
                    style={{ marginRight: "0.4em" }}
                  >
                    Add New Canonical
                  </Button>
                  <Button
                    color="primary"
                    onClick={() => this.showFormMetaTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    Add New Meta-Tag
                  </Button>
                  <Button
                    color="primary"
                    onClick={() => this.showFormScriptTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    Add New Script Tag
                  </Button>
                  <Button
                    color="primary"
                    onClick={() => this.showFormTitleTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    Add New Title-Tag
                  </Button>
                </div>
                <Table responsive bordered>
                  <thead>
                    <tr>
                      <th>Type</th>
                      <th>Attribute</th>
                      <th>Value</th>
                      <td>Language</td>
                      <th>Action</th>
                    </tr>
                  </thead>

                  <tbody>
                    {tags.length > 0 ? (
                      tags.map((tag, index) => (
                        <tr key={index}>
                          <td>{tag.type}</td>
                          <td></td>
                          <td>{tag.value}</td>
                          <td>{tag.locale_id}</td>
                          <td>
                            <button
                              className="button muted-button"
                              onClick={() => this.showFormMetaTag(tag)}
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
                              className={
                                "modal-warning " + this.props.className
                              }
                            >
                              <ModalHeader toggle={this.toggleWarning}>
                                Delete Confirmation
                              </ModalHeader>
                              <ModalBody>
                                Are you sure want to delete this row ?
                              </ModalBody>
                              <ModalFooter>
                                <Button
                                  color="warning"
                                  onClick={() => this.handleDelete(tag.id)}
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
                        <td colSpan={5}>No Tag</td>
                      </tr>
                    )}
                  </tbody>
                </Table>
                <nav>
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
                </nav>
              </CardBody>
            </Card>
            <CanonicalForm
              visible={this.state.canonicalFormVisible}
              onCancel={this.handleCancelAddCanonical}
              onSave={this.handleSave}
              canonical={this.state.record}
              action={this.state.actionForm}
              onChange={this.handleOnChange.bind(this)}
            />
            <MetaTagForm
              visible={this.state.metaTagFormVisible}
              onCancel={this.handleCancelAddMetaTag.bind(this)}
              onSave={this.handleSaveMetaTag.bind(this)}
              metatag={this.state.record}
              action={this.state.actionForm}
              onChange={this.handleMetaTagOnChange.bind(this)}
            />
            <ScriptTagForm
              visible={this.state.scriptTagFormVisible}
              onCancel={this.handleCancelAddScriptTag}
              onSave={this.handleSave}
              scripttag={this.state.record}
              action={this.state.actionForm}
              onChange={this.handleOnChange.bind(this)}
            />
            <TitleTagForm
              visible={this.state.titleTagFormVisible}
              onCancel={this.handleCancelAddTitleTag}
              onSave={this.handleSave}
              titletag={this.state.record}
              action={this.state.actionForm}
              onChange={this.handleOnChange.bind(this)}
            />
            <Modal
              isOpen={this.state.warningAPI}
              toggle={this.toggleWarningAPI}
              className={"modal-warning " + this.props.className}
            >
              <ModalHeader toggle={this.toggleWarningAPI}>
                Information
              </ModalHeader>
              <ModalBody>
                <span>{this.state.errorMessage}</span>
                <br></br>
                <span>
                  Sorry, failed to connect API. API currently not available/API
                  in problem
                </span>
              </ModalBody>
            </Modal>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RuleDetail;
