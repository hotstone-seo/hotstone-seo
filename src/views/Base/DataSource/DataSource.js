import React, { Component } from 'react';
import {  Card, CardBody, CardHeader, Col, Pagination, PaginationItem, PaginationLink, Row, Table, Button, NavLink } from 'reactstrap';
import PropTypes from 'prop-types';

class DataSource extends Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }
  handleClick() {
    const { history } = this.props;
    history.push('/base/DataSourceForm');
  }
  render() {
    return (

      <div className="animated fadeIn">       
        <Row>
        <Col xs="6" lg="2">
        <Button block color="primary" onClick={this.handleClick} >Add New</Button>
        
            </Col>
            
          <Col xs="12" lg="12">
            <Card>
              <CardHeader>
                Data Source
              </CardHeader>
              <CardBody>
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
                    <th>Nov 15 2019</th>
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
         </Col>
        </Row>
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
