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
      actionForm: "",
      ruleFormValues: {
        name: null,
        urlPattern: null
      }
    };

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
  componentDidMount() {
    axios.get('http://localhost:8089/rules')
      .then((res) => {
        const rules = res.data;
        this.setState({ rules });
      }).catch((error) => {
        
      });
  }

  handleDelete(id) {
    axios.delete(`http://localhost:8089/rules/${id}`)
      .then(() => {
        const { rules } = this.state;
        this.setState({ rules: rules.filter((rul) => rul.id !== id) });
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
    const { ruleFormValues } = this.state;
    axios.post('http://localhost:8089/rules', ruleFormValues)
      .then((response) => {
        this.setState({ ruleFormValues: [...ruleFormValues, response.data] });
        //message.success('Environment created');
      })
      .catch((error) => {

      });
    this.setState({ formVisible: false });
  }

  handleOnChange(type, e) {
    const { target } = e || {};
    const { value } = target || {};
    const { ruleFormValues } = this.state;

    this.setState({
      ruleFormValues: {
        ...ruleFormValues,
        [type]: value
      }
    });
  }

  render() {
    const { rules, ruleFormValues } = this.state;

    console.log(ruleFormValues, 'pasrender')
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
            onChange={this.handleOnChange.bind(this)}
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
