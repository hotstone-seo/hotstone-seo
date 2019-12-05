import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink,  Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class DataSource extends Component {
  constructor(props) {
    super(props);
    this.handleAdd = this.handleAdd.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
  }
  handleAdd() {
    const { history } = this.props;
    history.push('/dataSourceForm');
  }
  handleEdit() {
    const { history } = this.props;
    history.push('/dataSourceEditForm');
  }
  render() {
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
