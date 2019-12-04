import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class Canonical extends Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    const { history } = this.props;
    history.push('/base/CanonicalForm');
  }
  render() {
    return (
      <div className="animated fadeIn">
        <Card>
          <CardHeader>
            Canonical
          </CardHeader>
          <CardBody>
            <div style={{ marginBottom: '.5rem' }}>
              <Button color="primary" onClick={this.handleClick}>Add New</Button>
            </div>
            <Table responsive bordered>
              <thead>
                <tr>
                  <th>Canonical Tag</th>
                  <th>Rule</th>
                  <th>Updated Date</th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>http://tiket.com/asad</td>
                  <td>Airport</td>
                  <td>Nov 15 2019</td>
                  <td>
                    <NavLink href="#">Edit</NavLink>
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
      </div>
    );
  }
}
Canonical.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default Canonical;
