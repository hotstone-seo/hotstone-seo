import axios from 'axios';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { Button, Card, CardBody, CardHeader, Col, Modal, ModalBody, ModalFooter, ModalHeader, Table } from 'reactstrap';
import RuleForm from './RuleForm';

class RuleList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      rules: [],
      record: {}, 
      modal: false,
      warning: false,
      formVisible: false,
      actionForm: ""
    };   
    
    this.handleEdit = this.handleEdit.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
    this.toggleWarning = this.toggleWarning.bind(this);
    this.handleCancel = this.handleCancel.bind(this);

    this.showForm = this.showForm.bind(this);
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
  componentDidMount() {
    axios.get('http://localhost:8089/rules')
      .then((res) => {
        const rules = res.data;
        this.setState({ rules });
      }).catch((error) => {
        alert(error.message)
      });
  }

  handleEdit(record) {
    const { history } = this.props;
    if (record !== undefined) {
      history.push({
        pathname: '/ruleEditForm',
        data: record
      })
    }
  }

  handleDelete(id) {
      axios.delete(`http://localhost:8089/rules/${id}`)
       .then(() => {
        const { rules } = this.state;
        this.setState({ rules: rules.filter((rul) => rul.id !== id) });
      })
      .catch((error) => {
        alert(error.message)
      });
      this.toggleWarning()
  }
  showForm(record) {
    if (record !== undefined) {
      this.setState({ record: record });
      this.setState({ actionForm: "Edit" });
    }
    else
    {
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
  render() {
    const { rules } = this.state;
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>
              Rule
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
                    <th>URL Pattern</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  {rules.length > 0 ? (
                    rules.map(rule => (
                      <tr key={rule.id}>
                        <td>{rule.id}</td>
                        <td>{rule.name}</td>
                        <td>{rule.url_pattern}</td>
                        <td>{rule.updated_at}</td>
                        <td>
                          <button className="button muted-button" onClick={() => this.showForm(rule)}>Edit</button>
                          <button className="button muted-button" onClick={this.toggleWarning}>Delete</button>
                          <Modal isOpen={this.state.warning} toggle={this.toggleWarning}
                              className={'modal-warning ' + this.props.className}>
                          <ModalHeader toggle={this.toggleWarning}>Delete Confirmation</ModalHeader>
                          <ModalBody>
                            Are you sure want to delete {rule.name} ?
                          </ModalBody>
                          <ModalFooter>
                            <Button color="warning" onClick={() => this.handleDelete(rule.id)}>YES</Button>{' '}
                            <Button color="secondary" onClick={this.toggleWarning}>NO</Button>
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
        />
        </Col>
      </div>
    );
  }
}
RuleList.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default RuleList;
