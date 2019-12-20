import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';
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
      },
      URL_API: process.env.REACT_APP_API_URL + 'datasources'
    }
    this.handleAdd = this.handleAdd.bind(this);
    this.handleEdit = this.handleEdit.bind(this);

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
    const { datasourcesFormValues, rules, actionForm, record } = this.state;
    const isUpdate = actionForm !== "Add";

    datasourcesFormValues.id = record.id;

    if (isUpdate) {
      axios.put(this.state.URL_API, datasourcesFormValues)
        .then(() => {
          const index = rules.findIndex((rul) => rul.id === record.id);
          if (index > -1) {
            rules[index] = datasourcesFormValues;
            this.setState({ rules });
          }
        })
        .then(() => {
          this.getDataSourceList();
        })
        .catch((error) => {
          console.log(error.message)
        });
    }
    else {
      axios.post(this.state.URL_API, datasourcesFormValues)
        .then((response) => {
          this.setState({ rules: [...datasources, datasourcesFormValues] });
        })
        .then(() => {
          this.getDataSourceList();
        })
        .catch((error) => {

        });
      this.setState({ datasourcesFormValuess: {} });
    }
    this.setState({ formVisible: false });
  }

  handleOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { datasourcesFormValues } = this.state;

    this.setState({
      datasourcesFormValues: {
        ...datasourcesFormValuess,
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
                <Button color="primary" onClick={this.handleAdd}>Add New</Button>
              </div>
              <Table responsive bordered>
                <thead>
                  <tr>
                    <th>Data Source Name</th>
                    <th>Webhook</th>
                    <th>Fields</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Airport</td>
                    <td>http://fligh-service/airport</td>
                    <td>Id, name, address, province</td>
                    <td>Nov 15 2019</td>
                    <td>
                      <NavLink href="#" onClick={this.handleEdit}>Edit</NavLink>
                    </td>
                  </tr>

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
        </Col>
      </div>
    );
  }
}
DataSource.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default DataSource;
