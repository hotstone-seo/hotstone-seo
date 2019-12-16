import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
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
  }

  componentDidMount() {
    axios.get('http://localhost:4000/environments')
      .then((res) => {
        const rules = res.data;
        this.setState({ rules });
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

  render() {
    const ruleLink = '/ruleDetail/100'
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
                    <th>Name</th>
                    <th>URL Pattern</th>
                    <th>Data Source</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td><Link to={ruleLink}>Airport Detail</Link></td>
                    <td>http://xxx</td>
                    <td>DataSource1</td>
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
RuleList.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default RuleList;
