import React, { Component } from "react";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Form,
  FormGroup,
  Input,
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
import UniquePageCounterCard from "../Analytic/UniquePageCounterCard";
import HitCounterCard from "../Analytic/HitCounterCard";

import axios from "axios";

import HotstoneAPI from "../../api/hotstone";

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
        canonical: null,
        rule_id: null,
        locale: null
      },
      metaTagFormValues: {
        id: null,
        name: null,
        content: null,
        rule_id: null,
        locale: null
      },
      scriptTagFormValues: {
        id: null,
        type: null,
        rule_id: null,
        locale: null,
        datasource_id: null
      },
      titleTagFormValues: {
        id: null,
        title: null,
        rule_id: null,
        locale: null
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
      tags: [],
      tag_update: {
        id: null,
        type: null,
        attributes: null,
        value: null,
        locale: null,
        rule_id: null
      },
      metatag_attr: {
        name: null,
        content: null
      },
      languages: [],
      localeTag: "id-ID",
      metaTagPreviewValue: "",
      titleTagPreviewVal: "",
      scriptTagPreviewVal: ""
    };

    this.handleEditCanonical = this.handleEditCanonical.bind(this);
    this.handleDelete = this.handleDelete.bind(this);

    this.toggleWarning = this.toggleWarning.bind(this);
    this.toggleWarningAPI = this.toggleWarningAPI.bind(this);
    this.refreshTag = this.refreshTag.bind(this);
  }
  componentDidMount() {
    const query = parseQuery((window.location || {}).search || "");
    const { id } = query || {};

    const { rules } = this.state;
    axios
      .get(this.state.URL_API + `/${id}`)
      .then(res => {
        const { data: resData } = res || {};
        const rulesdata = resData;
        this.setState({ rules: [...rules, rulesdata] });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
    this.setState({ ruleId: id });
    this.getTagList(parseInt(id));

    const languages = HotstoneAPI.getLocales();
    this.setState({ languages });
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
      const { metaTagFormValues } = this.state;
      metaTagFormValues.id = record.id;
      metaTagFormValues.name = record.attributes.name;
      metaTagFormValues.locale = record.locale;
      metaTagFormValues.content = record.attributes.content;

      this.setState({ metaTagFormValues: metaTagFormValues });
      this.setState({ actionMetaTagForm: "Edit" });
    } else {
      this.resetMetaForm();
      this.setState({ record: {} });
      this.setState({ actionMetaTagForm: "Add" });
    }
    this.setState({ metaTagFormVisible: true });
  }

  showFormScriptTag(record) {
    if (record !== undefined) {
      const { scriptTagFormValues } = this.state;
      scriptTagFormValues.id = record.id;
      scriptTagFormValues.type = record.value;
      scriptTagFormValues.locale = record.locale;

      this.setState({ scriptTagFormValues: scriptTagFormValues });
      this.setState({ actionScriptTagForm: "Edit" });
    } else {
      this.resetScriptForm();
      this.setState({ record: {} });
      this.setState({ actionScriptTagForm: "Add" });
    }
    this.setState({ scriptTagFormVisible: true });
  }

  showFormTitleTag(record) {
    if (record !== undefined) {
      const { titleTagFormValues } = this.state;
      titleTagFormValues.id = record.id;
      titleTagFormValues.name = record.value;
      titleTagFormValues.locale = record.locale;

      this.setState({ titleTagFormValues: titleTagFormValues });
      this.setState({ actionTitleTagForm: "Edit" });
    } else {
      this.resetTitleForm();
      this.setState({ record: {} });
      this.setState({ actionTitleTagForm: "Add" });
    }
    this.setState({ titleTagFormVisible: true });
  }

  showFormCanonicalTag(record) {
    if (record !== undefined) {
      const { canonicalFormValues } = this.state;
      canonicalFormValues.id = record.id;
      canonicalFormValues.canonical = record.value;
      canonicalFormValues.locale = record.locale;

      this.setState({ canonicalFormValues: canonicalFormValues });
      this.setState({ actionCanonicalForm: "Edit" });
    } else {
      this.resetCanonicalForm();
      this.setState({ record: {} });
      this.setState({ actionCanonicalForm: "Add" });
    }
    this.setState({ canonicalFormVisible: true });
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

    let name = "";
    let content = "";
    if (target.name === "name") {
      name = target.value;
      content = metaTagFormValues.content;
    } else if (target.name === "content") {
      name = metaTagFormValues.name;
      content = target.value;
    }
    this.generatePreviewMetaTag(name, content);
  }

  handleTitleTagOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { titleTagFormValues } = this.state;

    this.setState({
      titleTagFormValues: {
        ...titleTagFormValues,
        [type]: value
      }
    });
    this.generatePreviewTitleTag(target.value);
  }
  handleScriptTagOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { scriptTagFormValues } = this.state;

    this.setState({
      scriptTagFormValues: {
        ...scriptTagFormValues,
        [type]: value
      }
    });

    if (target.name === "types") {
      this.generatePreviewScriptTag(target.value);
    }
  }

  handleCanonicalTagOnChange(type, e) {
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
    this.setState({
      warning: !this.state.warning
    });
  }

  handleDelete(id) {
    const { ruleId } = this.state;
    axios
      .delete(this.state.URL_TAG_API + `/${id}`)
      .then(() => {
        this.getTagList(parseInt(ruleId));
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
    const { tags } = this.state;
    this.setState({ tags: tags.filter(tag => tag.id !== id) });
    this.toggleWarning();
  }

  handleSaveMetaTag() {
    const {
      metaTagFormValues,
      actionMetaTagForm,
      ruleId,
      tag_update,
      metatag_attr,
      localeTag
    } = this.state;
    const isUpdate = actionMetaTagForm !== "Add";

    if (isUpdate) {
      metatag_attr.name = metaTagFormValues.name;
      metatag_attr.content = metaTagFormValues.content;

      tag_update.id = metaTagFormValues.id;
      tag_update.type = "meta";
      tag_update.attributes = metatag_attr;
      tag_update.locale = metaTagFormValues.locale;
      tag_update.value = metaTagFormValues.name;
      tag_update.rule_id = parseInt(ruleId);

      axios
        .put(this.state.URL_TAG_API, tag_update)
        .then(() => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      metaTagFormValues.rule_id = parseInt(ruleId);
      metaTagFormValues.locale = localeTag;
      axios
        .post(this.state.URL_ADDMETA_API, metaTagFormValues)
        .then(response => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    }
    this.resetMetaForm();
    this.setState({ metaTagFormVisible: false });
  }
  handleSaveTitleTag() {
    const {
      titleTagFormValues,
      actionTitleTagForm,
      ruleId,
      tag_update,
      localeTag
    } = this.state;
    const isUpdate = actionTitleTagForm !== "Add";

    if (isUpdate) {
      tag_update.id = titleTagFormValues.id;
      tag_update.type = "title";
      tag_update.attributes = "{}";
      tag_update.locale = titleTagFormValues.locale;
      tag_update.value = titleTagFormValues.title;
      tag_update.rule_id = parseInt(ruleId);

      axios
        .put(this.state.URL_TAG_API, tag_update)
        .then(() => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      titleTagFormValues.rule_id = parseInt(ruleId);
      titleTagFormValues.locale = localeTag;
      axios
        .post(this.state.URL_ADDTITLE_API, titleTagFormValues)
        .then(response => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    }
    this.resetTitleForm();
    this.setState({ titleTagFormVisible: false });
  }
  handleSaveScriptTag() {
    const {
      scriptTagFormValues,
      actionScriptTagForm,
      ruleId,
      tag_update
    } = this.state;
    const isUpdate = actionScriptTagForm !== "Add";

    if (isUpdate) {
      tag_update.id = scriptTagFormValues.id;
      tag_update.type = "script";
      tag_update.attributes = "{}";
      tag_update.locale = scriptTagFormValues.locale;
      tag_update.value = scriptTagFormValues.type;
      tag_update.rule_id = parseInt(ruleId);

      axios
        .put(this.state.URL_TAG_API, tag_update)
        .then(() => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      scriptTagFormValues.rule_id = parseInt(ruleId);
      scriptTagFormValues.datasource_id = 1;
      scriptTagFormValues.locale = this.state.localeTag;
      axios
        .post(this.state.URL_ADDSCRIPT_API, scriptTagFormValues)
        .then(response => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    }
    this.resetScriptForm();
    this.setState({ scriptTagFormVisible: false });
  }
  handleSaveCanonicalTag() {
    const {
      canonicalFormValues,
      actionCanonicalForm,
      ruleId,
      tag_update,
      localeTag
    } = this.state;
    const isUpdate = actionCanonicalForm !== "Add";

    if (isUpdate) {
      tag_update.id = canonicalFormValues.id;
      tag_update.type = "canonical";
      tag_update.attributes = "{}";
      tag_update.locale = canonicalFormValues.locale;
      tag_update.value = canonicalFormValues.canonical;
      tag_update.rule_id = parseInt(ruleId);

      axios
        .put(this.state.URL_TAG_API, tag_update)
        .then(() => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      canonicalFormValues.rule_id = parseInt(ruleId);
      canonicalFormValues.locale = localeTag;

      axios
        .post(this.state.URL_ADDCANONICAL_API, canonicalFormValues)
        .then(response => {
          this.getTagList(parseInt(ruleId));
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    }
    this.resetCanonicalForm();
    this.setState({ canonicalFormVisible: false });
  }
  getTagList(rule_id) {
    const { localeTag } = this.state;
    axios
      .get(this.state.URL_TAG_API + "?rule_id=" + rule_id)
      .then(res => {
        const tags = res.data;
        this.setState({ tags });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }
  handleEdit(record) {
    if (record !== undefined) {
      const typeTag = record.type;
      if (typeTag === "canonical") this.showFormCanonicalTag(record);
      else if (typeTag === "meta") this.showFormMetaTag(record);
      else if (typeTag === "script") this.showFormScriptTag(record);
      else if (typeTag === "title") this.showFormTitleTag(record);
    }
  }

  refreshTag(e) {
    let localeSelected = e.target.value;
    const { ruleId } = this.state;
    //localeSelected = localeSelected.toUpperCase();

    this.setState({ localeTag: localeSelected });

    // TO DO : next below code will be merged to function getTagList
    axios
      .get(
        this.state.URL_TAG_API +
          "?locale=" +
          localeSelected +
          "&rule_id=" +
          ruleId
      )
      .then(res => {
        const tags = res.data;
        this.setState({ tags });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }
  generatePreviewMetaTag(nm, ct) {
    let metaTagPreviewValue = "";
    metaTagPreviewValue =
      '<meta name="' + nm + '" content="' + (ct !== null ? ct : "") + '">';
    this.setState({ metaTagPreviewValue });
  }
  generatePreviewTitleTag(ttl) {
    let titleTagPreviewValue = "";
    titleTagPreviewValue = "<title>" + (ttl !== null ? ttl : "") + "</title>";
    this.setState({ titleTagPreviewValue });
  }
  generatePreviewScriptTag(scp) {
    let scriptTagPreviewValue = "";
    scriptTagPreviewValue =
      '<script type="' + (scp !== null ? scp : "") + '"></script>';
    this.setState({ scriptTagPreviewValue });
  }
  resetCanonicalForm() {
    this.setState({
      canonicalFormValues: {
        id: null,
        canonical: null,
        rule_id: null,
        locale: null
      }
    });
  }
  resetTitleForm() {
    this.setState({
      titleTagFormValues: {
        id: null,
        title: null,
        rule_id: null,
        locale: null
      }
    });
  }
  resetMetaForm() {
    this.setState({
      canonicalFormValues: {
        id: null,
        canonical: null,
        rule_id: null,
        locale: null
      }
    });
  }
  resetScriptForm() {
    this.setState({
      scriptTagFormValues: {
        id: null,
        type: null,
        rule_id: null,
        locale: null,
        datasource_id: null
      }
    });
  }
  render() {
    const { rules, tags, languages } = this.state;

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
            <Col>
              <HitCounterCard ruleID={rule.id} />
            </Col>
            <Col>
              <UniquePageCounterCard ruleID={rule.id} />
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
                    onClick={() => this.showFormCanonicalTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    <i className="fa fa-plus"></i>&nbsp;New Canonical
                  </Button>
                  <Button
                    color="primary"
                    onClick={() => this.showFormMetaTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    <i className="fa fa-plus"></i>&nbsp;New Meta Tag
                  </Button>
                  <Button
                    color="primary"
                    onClick={() => this.showFormScriptTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    <i className="fa fa-plus"></i>&nbsp; New Script Tag
                  </Button>
                  <Button
                    color="primary"
                    onClick={() => this.showFormTitleTag()}
                    style={{ marginRight: "0.4em" }}
                  >
                    <i className="fa fa-plus"></i>&nbsp; New Title Tag
                  </Button>
                </div>
                <div>
                  <FormGroup row>
                    <Col md="1">
                      <Label htmlFor="text-input">Locale:</Label>
                    </Col>
                    <Col xs="6" md="3">
                      <Input
                        type="select"
                        name="lang_code"
                        id="lang_code"
                        onChange={this.refreshTag}
                        value={this.state.localeTag}
                      >
                        {languages.map(ds => (
                          <option
                            key={ds.lang_code}
                            value={ds.lang_code + "-" + ds.country_code}
                          >
                            {ds.lang_code + "-" + ds.country_code}
                          </option>
                        ))}
                      </Input>
                    </Col>
                  </FormGroup>
                </div>
                <Table responsive bordered>
                  <thead>
                    <tr>
                      <th>Type</th>
                      <th>Attribute</th>
                      <th>Value</th>
                      <td>Locale</td>
                      <th>Action</th>
                    </tr>
                  </thead>

                  <tbody>
                    {tags.length > 0 ? (
                      tags.map((tag, index) => (
                        <tr key={index}>
                          <td>{tag.type}</td>
                          <td>
                            {tag.type === "meta" &&
                            tag.attributes.name !== undefined
                              ? "Name : " + tag.attributes.name
                              : ""}
                            {tag.type === "meta" &&
                            tag.attributes.name !== undefined
                              ? " Content :" + tag.attributes.content
                              : ""}
                          </td>
                          <td>{tag.value}</td>
                          <td>{tag.locale}</td>
                          <td>
                            <Button
                              color="secondary"
                              onClick={() => this.handleEdit(tag)}
                            >
                              <i className="fa fa-pencil"></i>&nbsp; Edit
                            </Button>
                            {"  "}
                            <Button color="danger" onClick={this.toggleWarning}>
                              <i className="fa fa-trash"></i>&nbsp;Delete
                            </Button>
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
              onCancel={this.handleCancelAddCanonical.bind(this)}
              onSave={this.handleSaveCanonicalTag.bind(this)}
              canonical={this.state.canonicalFormValues}
              action={this.state.actionCanonicalForm}
              onChange={this.handleCanonicalTagOnChange.bind(this)}
              languages={this.state.languages}
              languageDefault={this.state.localeTag}
            />
            <MetaTagForm
              visible={this.state.metaTagFormVisible}
              onCancel={this.handleCancelAddMetaTag.bind(this)}
              onSave={this.handleSaveMetaTag.bind(this)}
              metatag={this.state.metaTagFormValues}
              action={this.state.actionMetaTagForm}
              onChange={this.handleMetaTagOnChange.bind(this)}
              languages={this.state.languages}
              languageDefault={this.state.localeTag}
              metaTagPreviewValue={this.state.metaTagPreviewValue}
            />
            <ScriptTagForm
              visible={this.state.scriptTagFormVisible}
              onCancel={this.handleCancelAddScriptTag.bind(this)}
              onSave={this.handleSaveScriptTag.bind(this)}
              scripttag={this.state.scriptTagFormValues}
              action={this.state.actionScriptTagForm}
              onChange={this.handleScriptTagOnChange.bind(this)}
              languages={this.state.languages}
              languageDefault={this.state.localeTag}
              scriptTagPreviewValue={this.state.scriptTagPreviewValue}
            />
            <TitleTagForm
              visible={this.state.titleTagFormVisible}
              onCancel={this.handleCancelAddTitleTag.bind(this)}
              onSave={this.handleSaveTitleTag.bind(this)}
              titletag={this.state.titleTagFormValues}
              action={this.state.actionTitleTagForm}
              onChange={this.handleTitleTagOnChange.bind(this)}
              languages={this.state.languages}
              languageDefault={this.state.localeTag}
              titleTagPreviewValue={this.state.titleTagPreviewValue}
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
