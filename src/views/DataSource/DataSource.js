import axios from 'axios';
import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Modal, ModalBody, ModalFooter, ModalHeader, Table, Button } from 'reactstrap';
import DataSourceForm from './DataSourceForm';

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
        url: null,
      },
      URL_API: process.env.REACT_APP_API_URL + 'data_sources'
    }
    this.handleDelete = this.handleDelete.bind(this);
    this.handleSave = this.handleSave.bind(this);
    this.handleCancel = this.handleCancel.bind(this);

    this.toggleWarning = this.toggleWarning.bind(this);
    this.saveFormRef = this.saveFormRef.bind(this);

  }
  toggle() {
    this.setState({
      modal: !this.state.modal,
    });
  }
  toggleWarning() {
    this.setState({
      warning: !this.state.warning,
    });
  }
  getDataSourceList() {
    axios.get(this.state.URL_API)
      .then((res) => {
        const datasources = res.data;
        this.setState({ datasources });
      }).catch((error) => {
        // TODO: show error in dialog box
      });
  }
  componentDidMount() {
    this.getDataSourceList();
  }
  handleDelete(id) {
    axios.delete(this.state.URL_API + `/${id}`)
      .then(() => {
        const { datasources } = this.state;
        this.setState({ datasources: datasources.filter((rul) => rul.id !== id) });
      })
      .catch((error) => {
        // TODO: show error in dialog box
      });
    this.toggleWarning()
  }
  showForm(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionForm: "Edit" });
    }
    else {
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
    const { datasourcesFormValues, datasources, actionForm, record } = this.state;
    const isUpdate = actionForm !== "Add";

    datasourcesFormValues.id = record.id;

    if (isUpdate) {
      axios.put(this.state.URL_API, datasourcesFormValues)
        .then(() => {
          const index = datasources.findIndex((ds) => ds.id === record.id);
          if (index > -1) {
            datasources[index] = datasourcesFormValues;
            this.setState({ datasources });
          }
        })
        .catch((error) => {
          console.log(error.message)
        });
    }
    else {
      axios.post(this.state.URL_API, datasourcesFormValues)
        .then((response) => {
          this.setState({ datasources: [...datasources, datasourcesFormValues] });
        })
        .then(() => {
          this.getDataSourceList();
        })
        .catch((error) => {
          // TODO: show error in dialog box
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

  render() {
    const { datasources } = this.state;
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>
              Data Source
            </CardHeader>
            <CardBody>
              <div style={{ marginBottom: '.5rem' }}>
                <Button color="primary" onClick={() => this.showForm()}>Add New</Button>
              </div>
              <Table responsive bordered>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  {datasources.length > 0 ? (
                    datasources.map((datasource,index) => (
                      <tr key={index}>
                        <td>{datasource.id}</td>
                        <td>{datasource.name}</td>
                        <td></td>
                        <td>
                          <button className="button muted-button" onClick={() => this.showForm(datasource)}>Edit</button>
                          <button className="button muted-button" onClick={this.toggleWarning}>Delete</button>
                          <Modal isOpen={this.state.warning} toggle={this.toggleWarning}
                            className={'modal-warning ' + this.props.className}>
                            <ModalHeader toggle={this.toggleWarning}>Delete Confirmation</ModalHeader>
                            <ModalBody>
                              Are you sure want to delete data source name {datasource.name} ?
                          </ModalBody>
                            <ModalFooter>
                              <Button color="warning" onClick={() => this.handleDelete(datasource.id)}>YES</Button>{' '}
                              <Button color="secondary" onClick={this.toggleWarning}>NO</Button>
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
                <PaginationItem><PaginationLink previous tag="button">Prev</PaginationLink></PaginationItem>
                <PaginationItem active>
                  <PaginationLink tag="button">1</PaginationLink>
                </PaginationItem>
                <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
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
        </Col>
      </div>
    );
  }
}

export default DataSource;
