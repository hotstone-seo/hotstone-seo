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
import DataSourceForm from "./DataSourceForm";

class DataSource extends Component {
  constructor(props) {
    super(props);
    this.state = {
      datasources: [],
      record: {},
      modal: false,
      warning: false,
      formVisible: false,
      actionForm: "",
      datasourceFormValues: {
        id: null,
        name: null,
        url: null
      },
      URL_API: process.env.REACT_APP_API_URL + "data_sources",
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

  getDataSourceList() {
    axios
      .get(this.state.URL_API)
      .then(res => {
        const datasources = res.data;
        this.setState({ datasources });
      })
      .catch(error => {
        this.toggleWarningAPI(error.message);
      });
  }
  componentDidMount() {
    this.getDataSourceList();
  }
  handleDelete(id) {
    axios
      .delete(this.state.URL_API + `/${id}`)
      .then(() => {
        const { datasources } = this.state;
        this.setState({
          datasources: datasources.filter(rul => rul.id !== id)
        });
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
    const {
      datasourcesFormValues,
      datasources,
      actionForm,
      record
    } = this.state;

    if (actionForm !== "Add") {
      datasourcesFormValues.id = record.id;
      axios
        .put(this.state.URL_API, datasourcesFormValues)
        .then(() => {
          const index = datasources.findIndex(ds => ds.id === record.id);
          if (index > -1) {
            datasources[index] = datasourcesFormValues;
            this.setState({ datasources });
          }
        })
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
    } else {
      let lastid = this.getLastID();
      axios
        .post(this.state.URL_API, datasourcesFormValues)
        .then(response => {
          this.getDataSourceList();
          //datasourcesFormValues.id = lastid + 1;
          //this.setState({
          //  datasources: [...datasources, datasourcesFormValues]
          //});
        })
        //.then(() => {
        //  this.getDataSourceList();
        //})
        .catch(error => {
          this.toggleWarningAPI(error.message);
        });
      this.setState({ datasourcesFormValues: {} });
    }
    this.setState({ formVisible: false });
  }

  handleOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { datasourcesFormValues } = this.state;

    this.setState({
      datasourcesFormValues: {
        ...datasourcesFormValues,
        [type]: value
      }
    });
  }
  handleCloseWarningAPI() {
    this.setState({ warningAPI: false });
  }
  getLastID() {
    const { datasources } = this.state;
    let lastid = 0;
    if (datasources.length > 0) {
      lastid = datasources[datasources.length - 1].id;
    }
    return lastid;
  }
  render() {
    const { datasources } = this.state;
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>Data Source</CardHeader>
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
                    <th>URL</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  {datasources.length > 0 ? (
                    datasources.map((datasource, index) => (
                      <tr key={index}>
                        <td>{datasource.id}</td>
                        <td>{datasource.name}</td>
                        <td>{datasource.url}</td>
                        <td>
                          <button
                            className="button muted-button"
                            onClick={() => this.showForm(datasource)}
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
                              Are you sure want to delete data source name{" "}
                              {datasource.name} ?
                            </ModalBody>
                            <ModalFooter>
                              <Button
                                color="warning"
                                onClick={() => this.handleDelete(datasource.id)}
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
                      <td colSpan={5}>No Data Source</td>
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
          <DataSourceForm
            wrappedComponentRef={this.saveFormRef}
            visible={this.state.formVisible}
            onCancel={this.handleCancel}
            onSave={this.handleSave}
            datasource={this.state.record}
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

export default DataSource;
