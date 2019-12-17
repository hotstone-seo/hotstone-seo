import React, { Component } from 'react';
//import { Link } from 'react-router-dom';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink, Popconfirm, message } from 'reactstrap';
import PropTypes from 'prop-types';
import axios from 'axios';

class RuleList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      rules: [],
    };   
    this.handleClick = this.handleClick.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
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

  handleClick() {
    const { history } = this.props;
    history.push('/ruleForm');
  }

  handleEdit() {
    const { history } = this.props;
    history.push('/ruleEditForm');
  }

  handleDelete(id) {
    axios.delete(`http://localhost:8089/rules/${id}`)
      .then(() => {
        const { rules } = this.state;
        this.setState({ rules: rules.filter((env) => env.id !== id) });
      })
      .catch((error) => {
        alert(error.message)
      });
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
                <Button color="primary" onClick={this.handleClick}>Add New</Button>
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
                          <button className="button muted-button">Edit</button>
                          <button className="button muted-button" onClick={() => this.handleDelete(rule.id)}>Delete</button>
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
