import React, { Component } from 'react';
import { Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class Metatag extends Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
    this.handleEdit = this.handleEdit.bind(this);
  }
  handleClick() {
    const { history } = this.props;
    history.push('/metatagForm');
  }
  handleEdit() {
    const { history } = this.props;
    history.push('/metatagEditForm');
  }
  render() {
    return (
      <div className="animated fadeIn">
        <Col xs="12" lg="12">
          <Card>
            <CardHeader>
              Meta-Tag
            </CardHeader>
            <CardBody>
              <div style={{ marginBottom: '.5rem' }}>
                <Button color="primary" onClick={this.handleClick}>Add New</Button>
              </div>
              <Table responsive bordered>
                <thead>
                  <tr>
                    <th>Name</th>
                    <th>Rule</th>
                    <th>Content</th>
                    <th>Updated Date</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>Pompeius René</td>
                    <td>2xxx</td>
                    <td>xx</td>
                    <td>Nov 16 2019</td>
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
                <PaginationItem className="page-item"><PaginationLink tag="button">2</PaginationLink></PaginationItem>
                <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
              </Pagination>
            </CardBody>
          </Card>
        </Col>
      </div>
    );
  }
}
Metatag.propTypes = {
  match: PropTypes.shape({
    path: PropTypes.string,
  }).isRequired,
};

export default Metatag;